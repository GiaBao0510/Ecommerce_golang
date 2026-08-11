package manager

type ManagerRouterGroup struct {
	AdminRouter
	UserRouter
	StatusRouter
	RolesRouter
	PermissionRouter
	RolePermissionRouter
	UserRoleRouter
}