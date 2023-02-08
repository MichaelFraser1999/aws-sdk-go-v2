// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ApplicationStatus string

// Enum values for ApplicationStatus
const (
	ApplicationStatusActivated   ApplicationStatus = "ACTIVATED"
	ApplicationStatusStarting    ApplicationStatus = "STARTING"
	ApplicationStatusStopped     ApplicationStatus = "STOPPED"
	ApplicationStatusStopping    ApplicationStatus = "STOPPING"
	ApplicationStatusFailed      ApplicationStatus = "FAILED"
	ApplicationStatusRegistering ApplicationStatus = "REGISTERING"
	ApplicationStatusDeleting    ApplicationStatus = "DELETING"
	ApplicationStatusUnknown     ApplicationStatus = "UNKNOWN"
)

// Values returns all known values for ApplicationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ApplicationStatus) Values() []ApplicationStatus {
	return []ApplicationStatus{
		"ACTIVATED",
		"STARTING",
		"STOPPED",
		"STOPPING",
		"FAILED",
		"REGISTERING",
		"DELETING",
		"UNKNOWN",
	}
}

type ApplicationType string

// Enum values for ApplicationType
const (
	ApplicationTypeHana ApplicationType = "HANA"
)

// Values returns all known values for ApplicationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ApplicationType) Values() []ApplicationType {
	return []ApplicationType{
		"HANA",
	}
}

type ComponentStatus string

// Enum values for ComponentStatus
const (
	ComponentStatusActivated ComponentStatus = "ACTIVATED"
)

// Values returns all known values for ComponentStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ComponentStatus) Values() []ComponentStatus {
	return []ComponentStatus{
		"ACTIVATED",
	}
}

type ComponentType string

// Enum values for ComponentType
const (
	ComponentTypeHana ComponentType = "HANA"
)

// Values returns all known values for ComponentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ComponentType) Values() []ComponentType {
	return []ComponentType{
		"HANA",
	}
}

type CredentialType string

// Enum values for CredentialType
const (
	CredentialTypeAdmin CredentialType = "ADMIN"
)

// Values returns all known values for CredentialType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CredentialType) Values() []CredentialType {
	return []CredentialType{
		"ADMIN",
	}
}

type DatabaseStatus string

// Enum values for DatabaseStatus
const (
	DatabaseStatusRunning  DatabaseStatus = "RUNNING"
	DatabaseStatusStarting DatabaseStatus = "STARTING"
	DatabaseStatusStopped  DatabaseStatus = "STOPPED"
	DatabaseStatusWarning  DatabaseStatus = "WARNING"
	DatabaseStatusUnknown  DatabaseStatus = "UNKNOWN"
)

// Values returns all known values for DatabaseStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DatabaseStatus) Values() []DatabaseStatus {
	return []DatabaseStatus{
		"RUNNING",
		"STARTING",
		"STOPPED",
		"WARNING",
		"UNKNOWN",
	}
}

type DatabaseType string

// Enum values for DatabaseType
const (
	DatabaseTypeSystem DatabaseType = "SYSTEM"
	DatabaseTypeTenant DatabaseType = "TENANT"
)

// Values returns all known values for DatabaseType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (DatabaseType) Values() []DatabaseType {
	return []DatabaseType{
		"SYSTEM",
		"TENANT",
	}
}

type FilterOperator string

// Enum values for FilterOperator
const (
	FilterOperatorEquals              FilterOperator = "Equals"
	FilterOperatorGreaterThanOrEquals FilterOperator = "GreaterThanOrEquals"
	FilterOperatorLessThanOrEquals    FilterOperator = "LessThanOrEquals"
)

// Values returns all known values for FilterOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FilterOperator) Values() []FilterOperator {
	return []FilterOperator{
		"Equals",
		"GreaterThanOrEquals",
		"LessThanOrEquals",
	}
}

type HostRole string

// Enum values for HostRole
const (
	HostRoleLeader  HostRole = "LEADER"
	HostRoleWorker  HostRole = "WORKER"
	HostRoleStandby HostRole = "STANDBY"
	HostRoleUnknown HostRole = "UNKNOWN"
)

// Values returns all known values for HostRole. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (HostRole) Values() []HostRole {
	return []HostRole{
		"LEADER",
		"WORKER",
		"STANDBY",
		"UNKNOWN",
	}
}

type OperationStatus string

// Enum values for OperationStatus
const (
	OperationStatusInprogress OperationStatus = "INPROGRESS"
	OperationStatusSuccess    OperationStatus = "SUCCESS"
	OperationStatusError      OperationStatus = "ERROR"
)

// Values returns all known values for OperationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OperationStatus) Values() []OperationStatus {
	return []OperationStatus{
		"INPROGRESS",
		"SUCCESS",
		"ERROR",
	}
}

type PermissionActionType string

// Enum values for PermissionActionType
const (
	PermissionActionTypeRestore PermissionActionType = "RESTORE"
)

// Values returns all known values for PermissionActionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PermissionActionType) Values() []PermissionActionType {
	return []PermissionActionType{
		"RESTORE",
	}
}