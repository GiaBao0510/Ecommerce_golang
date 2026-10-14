package manager

type ManagerRouterGroup struct {
	AdminRouter
	UserRouter
	StatusRouter
	CategoryRouter
	SubCategoryRouter
	ProductDetailRouter
	RolesRouter
	PermissionRouter
	RolePermissionRouter
	UserRoleRouter
}
