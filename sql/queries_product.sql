--##################  Category

-- name: GetAllCategories :many
SELECT * FROM categories;

-- name: GetCategoryByID :one
SELECT * FROM categories WHERE id_cate = $1;

-- name: GetCategoryByName :one
SELECT * FROM categories WHERE "name" = $1;

-- name: CreateCategory :exec
INSERT INTO categories(name, description) VALUES($1,$2);

-- name: DeleteCategory :execresult
DELETE FROM categories WHERE id_cate = $1;

-- name: UpdateCategoryByID_PUT :execresult
UPDATE categories 
SET 	"name" = $1, 
		description = $2
WHERE id_cate = $3;

--##################  sub_categories

-- name: GetAllSubCategories :many
SELECT * FROM sub_categories;

-- name: GetSubCategoryByID :one
SELECT * FROM sub_categories WHERE id_sub_cate = $1;

-- name: GetSubCategoryByName :one
SELECT * FROM sub_categories WHERE "name" = $1;

-- name: CreateSubCategory :exec
INSERT INTO sub_categories(id_cate, "name", description) VALUES($1, $2, $3);

-- name: DeleteSubCategory :execresult
DELETE FROM sub_categories WHERE id_sub_cate = $1;

-- name: UpdateSubCategoryByID_PUT :execresult
UPDATE sub_categories 
SET 	"name" = $1, 
		description = $2
WHERE id_cate = $3;

-- ##################  product_detail

-- name: GetAllProductDetails :many
SELECT * FROM product_detail;

-- name: GetProductDetailByID :one
SELECT * FROM product_detail WHERE id_product = $1;

-- name: CreateProductDetail :exec
INSERT INTO product_detail(id_product, price, like_count, quantity, production_date, expiration_date) VALUES($1, $2, $3, $4, $5, $6);

-- name: DeleteProductDetail :execresult
DELETE FROM product_detail WHERE id_product = $1;

-- name: UpdateProductDetailByID_PUT :execresult
UPDATE product_detail 
SET 	price = $1, 
		like_count = $2, 
		quantity = $3, 
		production_date = $4, 
		expiration_date = $5
WHERE id_product = $6;

-- name: UpdateProductPriceByID_PATCH :execresult
UPDATE product_detail 
SET 	price = $1
WHERE id_product = $2;

-- name: UpdateProductLikeCountByID_PATCH :execresult
UPDATE product_detail 
SET 	like_count = $1
WHERE id_product = $2;

-- name: UpdateProductQuantityByID_PATCH :execresult
UPDATE product_detail 
SET 	quantity = $1
WHERE id_product = $2;

-- ################### Brand
-- name: GetAllBrands :many
SELECT * FROM brand;

-- name: GetBrandByID :one
SELECT * FROM brand WHERE id_brand = $1;

-- name: GetBrandByName :one
SELECT * FROM brand WHERE "name" = $1;

-- name: CreateBrand :exec
INSERT INTO brand("name") VALUES($1);

-- name: DeleteBrand :execresult
DELETE FROM brand WHERE id_brand = $1;

-- name: UpdateBrandByID_PUT :execresult
UPDATE brand 
SET 	"name" = $1
WHERE id_brand = $2;

-- ################### shop_owner

-- name: GetAllShopOwners :many
SELECT * FROM shop_owner;

-- name: GetShopOwnerByID :one
SELECT * FROM shop_owner WHERE id_shop_owner = $1;

-- name: GetShopOwnerByName :one
SELECT * FROM shop_owner WHERE "name" = $1;

-- name: CreateShopOwner :exec
INSERT INTO shop_owner(uuid, "name", description, avatar_url) VALUES($1, $2, $3, $4);

-- name: DeleteShopOwner :execresult
DELETE FROM shop_owner WHERE id_shop_owner = $1;

-- name: UpdateShopOwnerByID_PUT :execresult
UPDATE shop_owner 
SET 	uuid = $1, 
		"name" = $2, 
		description = $3, 
		avatar_url = $4
WHERE id_shop_owner = $5;

-- name: UpdateShopOwnerNameByID_PATCH :execresult
UPDATE shop_owner 
SET 	"name" = $1
WHERE id_shop_owner = $2;

-- name: UpdateShopOwnerDescriptionByID_PATCH :execresult
UPDATE shop_owner 
SET 	description = $1
WHERE id_shop_owner = $2;

-- name: UpdateShopOwnerAvatarByID_PATCH :execresult
UPDATE shop_owner 
SET 	avatar_url = $1
WHERE id_shop_owner = $2;

-- ################### SHOP_OWNER_NOTIFICATION