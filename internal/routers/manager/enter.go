package manager

type ManagerRouterGroup struct {
	AdminRouter
	UserRouter
	StatusRouter
	CategoryRouter
	SubCategoryRouter
	ProductDetailRouter
	BrandRouter
	ShopOwnerRouter
	RolesRouter
	PermissionRouter
	RolePermissionRouter
	UserRoleRouter
}
