// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package release

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"mime"
	"os"
	"path/filepath"
	"strings"
)

// PutReleaseOptions are options to for the PutRelease function.
type PutReleaseOptions struct {
	// BucketName is the bucket to upload the files to.
	BucketName string
	// NoCache is true if we should set the NoCache option to S3.
	NoCache bool
	// Platform is the platform of the release.
	Platform Platform
	// VersionStr is the version (SHA/branch name) of the release.
	VersionStr string

	// Files are all the files to be included in the archive.
	Files      []ArchiveFile
	ExtraFiles []ArchiveFile
}

// PutNonReleaseOptions are options to pass into PutNonRelease.
type PutNonReleaseOptions struct {
	// Branch is the branch from which the release is being uploaded from.
	Branch string
	// BucketName is the bucket to upload the files to.
	BucketName string

	// Files are all the files to be uploaded into S3.
	Files []NonReleaseFile
}

// PutRelease uploads a compressed archive containing the release
// files and a checksum file of the archive to S3.
func PutRelease(svc ObjectPutGetter, o PutReleaseOptions) {
	keys := makeArchiveKeys(o.Platform, o.VersionStr, "cockroach")
	var body bytes.Buffer

	if strings.HasSuffix(keys.archive, ".zip") {
		if err := createZip(o.Files, &body, keys.base); err != nil {
			log.Fatalf("cannot create zip %s: %s", keys.archive, err)
		}
	} else {
		if err := createTarball(o.Files, &body, keys.base); err != nil {
			log.Fatalf("cannot create tarball %s: %s", keys.archive, err)
		}
	}

	log.Printf("Uploading to s3://%s/%s", o.BucketName, keys.archive)
	putObjectInput := PutObjectInput{
		Bucket: &o.BucketName,
		Key:    &keys.archive,
		Body:   bytes.NewReader(body.Bytes()),
	}
	if o.NoCache {
		putObjectInput.CacheControl = &NoCache
	}
	if err := svc.PutObject(&putObjectInput); err != nil {
		log.Fatalf("s3 upload %s: %s", keys.archive, err)
	}
	// Generate a SHA256 checksum file with a single entry.
	checksumContents := fmt.Sprintf("%x %s\n", sha256.Sum256(body.Bytes()),
		filepath.Base(keys.archive))
	targetChecksum := keys.archive + ChecksumSuffix
	log.Printf("Uploading to s3://%s/%s", o.BucketName, targetChecksum)
	putObjectInputChecksum := PutObjectInput{
		Bucket: &o.BucketName,
		Key:    &targetChecksum,
		Body:   strings.NewReader(checksumContents),
	}
	if o.NoCache {
		putObjectInputChecksum.CacheControl = &NoCache
	}
	if err := svc.PutObject(&putObjectInputChecksum); err != nil {
		log.Fatalf("s3 upload %s: %s", targetChecksum, err)
	}
	for _, f := range o.ExtraFiles {
		keyBase, hasExe := TrimDotExe(f.ArchiveFilePath)
		targetKeys := makeArchiveKeys(o.Platform, o.VersionStr, keyBase)
		targetKey := targetKeys.base
		if hasExe {
			targetKey += ".exe"
		}
		log.Printf("Uploading to s3://%s/%s", o.BucketName, targetKey)
		handle, err := os.Open(f.LocalAbsolutePath)
		if err != nil {
			log.Fatalf("failed to open %s: %s", f.LocalAbsolutePath, err)
		}
		putObjectInput := PutObjectInput{
			Bucket: &o.BucketName,
			Key:    &targetKey,
			Body:   handle,
		}
		if o.NoCache {
			putObjectInput.CacheControl = &NoCache
		}
		if err := svc.PutObject(&putObjectInput); err != nil {
			log.Fatalf("s3 upload %s: %s", targetKey, err)
		}
	}
}

func createZip(files []ArchiveFile, body *bytes.Buffer, prefix string) error {
	zw := zip.NewWriter(body)
	for _, f := range files {
		file, err := os.Open(f.LocalAbsolutePath)
		if err != nil {
			return fmt.Errorf("failed to open file: %s", f.LocalAbsolutePath)
		}
		defer func() { _ = file.Close() }()

		stat, err := file.Stat()
		if err != nil {
			return fmt.Errorf("failed to stat: %s", f.LocalAbsolutePath)
		}

		zipHeader, err := zip.FileInfoHeader(stat)
		if err != nil {
			return err
		}
		zipHeader.Name = filepath.Join(prefix, f.ArchiveFilePath)
		zipHeader.Method = zip.Deflate

		zfw, err := zw.CreateHeader(zipHeader)
		if err != nil {
			return err
		}
		if _, err := io.Copy(zfw, file); err != nil {
			return err
		}
	}
	if err := zw.Close(); err != nil {
		return err
	}
	return nil
}

func createTarball(files []ArchiveFile, body *bytes.Buffer, prefix string) error {
	gzw := gzip.NewWriter(body)
	tw := tar.NewWriter(gzw)
	for _, f := range files {
		file, err := os.Open(f.LocalAbsolutePath)
		if err != nil {
			return fmt.Errorf("failed to open file: %s", f.LocalAbsolutePath)
		}
		defer func() { _ = file.Close() }()

		stat, err := file.Stat()
		if err != nil {
			return fmt.Errorf("failed to stat: %s", f.LocalAbsolutePath)
		}

		// Set the tar header from the file info. Overwrite name.
		tarHeader, err := tar.FileInfoHeader(stat, "")
		if err != nil {
			return err
		}
		tarHeader.Name = filepath.Join(prefix, f.ArchiveFilePath)
		if err := tw.WriteHeader(tarHeader); err != nil {
			return err
		}

		if _, err := io.Copy(tw, file); err != nil {
			return err
		}
	}
	if err := tw.Close(); err != nil {
		return err
	}
	if err := gzw.Close(); err != nil {
		return err
	}
	return nil
}

// PutNonRelease uploads non-release related files to S3.
// Files are uploaded to /cockroach/<S3FilePath> for each non release file.
// A `latest` key is then put at cockroach/<S3RedirectPrefix>.<BranchName> that redirects
// to the above file.
func PutNonRelease(svc ObjectPutGetter, o PutNonReleaseOptions) {
	const repoName = "cockroach"
	for _, f := range o.Files {
		disposition := mime.FormatMediaType("attachment", map[string]string{
			"filename": f.S3FileName,
		})

		fileToUpload, err := os.Open(f.LocalAbsolutePath)
		if err != nil {
			log.Fatalf("failed to open %s: %s", f.LocalAbsolutePath, err)
		}
		defer func() {
			_ = fileToUpload.Close()
		}()

		// NB: The leading slash is required to make redirects work
		// correctly since we reuse this key as the redirect location.
		versionKey := fmt.Sprintf("/%s/%s", repoName, f.S3FilePath)
		log.Printf("Uploading to s3://%s%s", o.BucketName, versionKey)
		if err := svc.PutObject(&PutObjectInput{
			Bucket:             &o.BucketName,
			ContentDisposition: &disposition,
			Key:                &versionKey,
			Body:               fileToUpload,
		}); err != nil {
			log.Fatalf("s3 upload %s: %s", versionKey, err)
		}

		latestSuffix := o.Branch
		if latestSuffix == "master" {
			latestSuffix = "LATEST"
		}
		latestKey := fmt.Sprintf("%s/%s.%s", repoName, f.S3RedirectPathPrefix, latestSuffix)
		if err := svc.PutObject(&PutObjectInput{
			Bucket:                  &o.BucketName,
			CacheControl:            &NoCache,
			Key:                     &latestKey,
			WebsiteRedirectLocation: &versionKey,
		}); err != nil {
			log.Fatalf("s3 redirect to %s: %s", versionKey, err)
		}
	}
}

type archiveKeys struct {
	base    string
	archive string
}

// makeArchiveKeys extracts the target archive base and archive
// name for the given parameters.
func makeArchiveKeys(platform Platform, versionStr string, binaryPrefix string) archiveKeys {
	suffix := SuffixFromPlatform(platform)
	targetSuffix, hasExe := TrimDotExe(suffix)
	if platform == PlatformLinux {
		targetSuffix = strings.Replace(targetSuffix, "gnu-", "", -1)
		targetSuffix = osVersionRe.ReplaceAllLiteralString(targetSuffix, "")
	}
	archiveBase := fmt.Sprintf("%s-%s", binaryPrefix, versionStr)
	targetArchiveBase := archiveBase + targetSuffix
	keys := archiveKeys{
		base: targetArchiveBase,
	}
	if hasExe {
		keys.archive = targetArchiveBase + ".zip"
	} else {
		keys.archive = targetArchiveBase + ".tgz"
	}
	return keys
}

const latestStr = "latest"

// LatestOpts are parameters passed to MarkLatestReleaseWithSuffix
type LatestOpts struct {
	Platform   Platform
	VersionStr string
	BucketName string
}

// MarkLatestReleaseWithSuffix adds redirects to release files using "latest" instead of the version
func MarkLatestReleaseWithSuffix(svc ObjectPutGetter, o LatestOpts, suffix string) {
	keys := makeArchiveKeys(o.Platform, o.VersionStr, "cockroach")
	versionedKey := keys.archive + suffix
	oLatest := o
	oLatest.VersionStr = latestStr
	latestKeys := makeArchiveKeys(oLatest.Platform, oLatest.VersionStr, "cockroach")
	latestKey := latestKeys.archive + suffix
	log.Printf("Adding redirect to s3://%s/%s", o.BucketName, latestKey)
	if err := svc.PutObject(&PutObjectInput{
		Bucket:                  &o.BucketName,
		CacheControl:            &NoCache,
		Key:                     &latestKey,
		WebsiteRedirectLocation: &versionedKey,
	}); err != nil {
		log.Fatalf("s3 redirect to %s: %s", versionedKey, err)
	}
}

// GetObjectInput specifies input parameters for GetOject
type GetObjectInput struct {
	Bucket *string
	Key    *string
}

// GetObjectOutput specifies output parameters for GetOject
type GetObjectOutput struct {
	Body io.ReadCloser
}

// PutObjectInput specifies input parameters for PutOject
type PutObjectInput struct {
	Bucket                  *string
	Key                     *string
	Body                    io.ReadSeeker
	CacheControl            *string
	ContentDisposition      *string
	WebsiteRedirectLocation *string
}

// ObjectPutGetter specifies a minimal interface for cloud storage providers
type ObjectPutGetter interface {
	GetObject(*GetObjectInput) (*GetObjectOutput, error)
	PutObject(*PutObjectInput) error
}
