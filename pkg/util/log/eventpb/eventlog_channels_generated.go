// Code generated by gen.go. DO NOT EDIT.

package eventpb

import "github.com/cockroachdb/cockroach/pkg/util/log/logpb"

// LoggingChannel implements the EventPayload interface.
func (m *CertsReload) LoggingChannel() logpb.Channel { return logpb.Channel_OPS }

// LoggingChannel implements the EventPayload interface.
func (m *NodeDecommissioned) LoggingChannel() logpb.Channel { return logpb.Channel_OPS }

// LoggingChannel implements the EventPayload interface.
func (m *NodeDecommissioning) LoggingChannel() logpb.Channel { return logpb.Channel_OPS }

// LoggingChannel implements the EventPayload interface.
func (m *NodeJoin) LoggingChannel() logpb.Channel { return logpb.Channel_OPS }

// LoggingChannel implements the EventPayload interface.
func (m *NodeRecommissioned) LoggingChannel() logpb.Channel { return logpb.Channel_OPS }

// LoggingChannel implements the EventPayload interface.
func (m *NodeRestart) LoggingChannel() logpb.Channel { return logpb.Channel_OPS }

// LoggingChannel implements the EventPayload interface.
func (m *RuntimeStats) LoggingChannel() logpb.Channel { return logpb.Channel_HEALTH }

// LoggingChannel implements the EventPayload interface.
func (m *Import) LoggingChannel() logpb.Channel { return logpb.Channel_OPS }

// LoggingChannel implements the EventPayload interface.
func (m *Restore) LoggingChannel() logpb.Channel { return logpb.Channel_OPS }

// LoggingChannel implements the EventPayload interface.
func (m *SetClusterSetting) LoggingChannel() logpb.Channel { return logpb.Channel_DEV }

// LoggingChannel implements the EventPayload interface.
func (m *AdminQuery) LoggingChannel() logpb.Channel { return logpb.Channel_SENSITIVE_ACCESS }

// LoggingChannel implements the EventPayload interface.
func (m *SensitiveTableAccess) LoggingChannel() logpb.Channel { return logpb.Channel_SENSITIVE_ACCESS }

// LoggingChannel implements the EventPayload interface.
func (m *QueryExecute) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_EXEC }

// LoggingChannel implements the EventPayload interface.
func (m *AlterDatabaseAddRegion) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *AlterDatabaseDropRegion) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *AlterDatabasePlacement) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *AlterDatabasePrimaryRegion) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *AlterDatabaseSurvivalGoal) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *AlterIndex) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *AlterSequence) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *AlterTable) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *AlterType) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *CommentOnColumn) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *CommentOnDatabase) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *CommentOnIndex) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *CommentOnTable) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *ConvertToSchema) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *CreateDatabase) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *CreateIndex) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *CreateSchema) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *CreateSequence) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *CreateStatistics) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *CreateTable) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *CreateType) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *CreateView) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *DropDatabase) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *DropIndex) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *DropSchema) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *DropSequence) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *DropTable) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *DropType) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *DropView) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *FinishSchemaChange) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *FinishSchemaChangeRollback) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *ForceDeleteTableDataEntry) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *RenameDatabase) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *RenameSchema) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *RenameTable) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *RenameType) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *ReverseSchemaChange) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *SetSchema) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *TruncateTable) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *UnsafeDeleteDescriptor) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *UnsafeDeleteNamespaceEntry) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *UnsafeUpsertDescriptor) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *UnsafeUpsertNamespaceEntry) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_SCHEMA }

// LoggingChannel implements the EventPayload interface.
func (m *AlterDatabaseOwner) LoggingChannel() logpb.Channel { return logpb.Channel_PRIVILEGES }

// LoggingChannel implements the EventPayload interface.
func (m *AlterDefaultPrivileges) LoggingChannel() logpb.Channel { return logpb.Channel_PRIVILEGES }

// LoggingChannel implements the EventPayload interface.
func (m *AlterSchemaOwner) LoggingChannel() logpb.Channel { return logpb.Channel_PRIVILEGES }

// LoggingChannel implements the EventPayload interface.
func (m *AlterTableOwner) LoggingChannel() logpb.Channel { return logpb.Channel_PRIVILEGES }

// LoggingChannel implements the EventPayload interface.
func (m *AlterTypeOwner) LoggingChannel() logpb.Channel { return logpb.Channel_PRIVILEGES }

// LoggingChannel implements the EventPayload interface.
func (m *ChangeDatabasePrivilege) LoggingChannel() logpb.Channel { return logpb.Channel_PRIVILEGES }

// LoggingChannel implements the EventPayload interface.
func (m *ChangeSchemaPrivilege) LoggingChannel() logpb.Channel { return logpb.Channel_PRIVILEGES }

// LoggingChannel implements the EventPayload interface.
func (m *ChangeTablePrivilege) LoggingChannel() logpb.Channel { return logpb.Channel_PRIVILEGES }

// LoggingChannel implements the EventPayload interface.
func (m *ChangeTypePrivilege) LoggingChannel() logpb.Channel { return logpb.Channel_PRIVILEGES }

// LoggingChannel implements the EventPayload interface.
func (m *ClientAuthenticationFailed) LoggingChannel() logpb.Channel { return logpb.Channel_SESSIONS }

// LoggingChannel implements the EventPayload interface.
func (m *ClientAuthenticationInfo) LoggingChannel() logpb.Channel { return logpb.Channel_SESSIONS }

// LoggingChannel implements the EventPayload interface.
func (m *ClientAuthenticationOk) LoggingChannel() logpb.Channel { return logpb.Channel_SESSIONS }

// LoggingChannel implements the EventPayload interface.
func (m *ClientConnectionEnd) LoggingChannel() logpb.Channel { return logpb.Channel_SESSIONS }

// LoggingChannel implements the EventPayload interface.
func (m *ClientConnectionStart) LoggingChannel() logpb.Channel { return logpb.Channel_SESSIONS }

// LoggingChannel implements the EventPayload interface.
func (m *ClientSessionEnd) LoggingChannel() logpb.Channel { return logpb.Channel_SESSIONS }

// LoggingChannel implements the EventPayload interface.
func (m *LargeRow) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_PERF }

// LoggingChannel implements the EventPayload interface.
func (m *SlowQuery) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_PERF }

// LoggingChannel implements the EventPayload interface.
func (m *TxnRowsReadLimit) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_PERF }

// LoggingChannel implements the EventPayload interface.
func (m *TxnRowsWrittenLimit) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_PERF }

// LoggingChannel implements the EventPayload interface.
func (m *LargeRowInternal) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_INTERNAL_PERF }

// LoggingChannel implements the EventPayload interface.
func (m *SlowQueryInternal) LoggingChannel() logpb.Channel { return logpb.Channel_SQL_INTERNAL_PERF }

// LoggingChannel implements the EventPayload interface.
func (m *TxnRowsReadLimitInternal) LoggingChannel() logpb.Channel {
	return logpb.Channel_SQL_INTERNAL_PERF
}

// LoggingChannel implements the EventPayload interface.
func (m *TxnRowsWrittenLimitInternal) LoggingChannel() logpb.Channel {
	return logpb.Channel_SQL_INTERNAL_PERF
}

// LoggingChannel implements the EventPayload interface.
func (m *AlterRole) LoggingChannel() logpb.Channel { return logpb.Channel_USER_ADMIN }

// LoggingChannel implements the EventPayload interface.
func (m *CreateRole) LoggingChannel() logpb.Channel { return logpb.Channel_USER_ADMIN }

// LoggingChannel implements the EventPayload interface.
func (m *DropRole) LoggingChannel() logpb.Channel { return logpb.Channel_USER_ADMIN }

// LoggingChannel implements the EventPayload interface.
func (m *RemoveZoneConfig) LoggingChannel() logpb.Channel { return logpb.Channel_OPS }

// LoggingChannel implements the EventPayload interface.
func (m *SetZoneConfig) LoggingChannel() logpb.Channel { return logpb.Channel_OPS }
