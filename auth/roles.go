package auth

type BuiltInRole string

const (
	// RoleProjectOwner project owner role
	RoleProjectOwner BuiltInRole = "owner"
	// RoleProjectMember project member
	RoleProjectMember BuiltInRole = "member"
	// RoleProjectViewer project viewer - read only maximum level
	RoleProjectViewer BuiltInRole = "viewer"
	// RoleProjectSupport support access for the project
	RoleProjectSupport BuiltInRole = "support"
)
