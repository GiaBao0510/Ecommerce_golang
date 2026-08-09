--// File: sql/queries.sql: tệp tin này chủ yếu chứa các câu lệnh thực thi

/*_______________ Bảng Status 0 ___________________*/
-- name: CreateStatus :exec
INSERT INTO status(name,description,updated_at,deleted_at) VALUES($1,$2,$3,$4) RETURNING id_status;

-- name: GetStatusByID :one
SELECT * FROM status WHERE id_status = $1;

-- name: GetAllStatus :many
SELECT * FROM status;

-- name: UpdateStatus_PUT :execresult
UPDATE status SET name = $1, description = $2, updated_at = NOW() WHERE id_status = $3;

-- name: UpdateStatus_PATCH :execresult
UPDATE status SET name = COALESCE($1, name), description =  COALESCE($2, description), updated_at = NOW() WHERE  id_status = $3;

-- name: DeleteStatus :execresult
DELETE FROM status WHERE id_status = $1;

/*_______________ Bảng Role 1 ___________________*/
-- name: CreateRole :exec
INSERT INTO role(role_name,description) VALUES($1,$2) RETURNING role_id;

-- name: GetRoleByID :one
SELECT * FROM role WHERE role_id = $1;

-- name: GetAllRoles :many
SELECT * FROM role;

-- name: UpdateRole_PUT :execresult
UPDATE role SET role_name = $1, description = $2 WHERE role_id = $3;

-- name: UpdateRole_PATCH :execresult
UPDATE role SET role_name = COALESCE($1, role_name), description = COALESCE($2, description) WHERE role_id = $3;

-- name: DeleteRole :execresult
DELETE FROM role WHERE role_id = $1;

/*_______________ Bảng permissions 2 ___________________*/
-- name: CreatePermission :exec
INSERT INTO permissions(action_name, description) VALUES($1,$2) RETURNING action_id;

-- name: GetPermissionByID :one
SELECT * FROM permissions WHERE action_id = $1;

-- name: GetAllPermissions :many
SELECT * FROM permissions;

-- name: UpdatePermission_PUT :execresult
UPDATE permissions SET action_name = $1, description = $2, updated_at = NOW() WHERE action_id = $3;

-- name: UpdatePermission_PATCH :execresult
UPDATE permissions SET 
	action_name = COALESCE($1, action_name), 
	description = COALESCE($2, description), 
	updated_at = NOW()
WHERE action_id = $3;

-- name: DeletePermission :execresult
DELETE FROM permissions WHERE action_id = $1;

/*_______________ Bảng Role_Permissions 3 ___________________*/
-- name: CreateRolePermission :exec
INSERT INTO role_permission(action_id, role_id) VALUES($1,$2);

-- name: GetPermissionsByRoleID :many
SELECT p.action_id, p.action_name, p.description
FROM role r 
	INNER JOIN role_permission rp ON r.role_id = rp.role_id
	INNER JOIN permissions p ON p.action_id = rp.action_id
WHERE r.role_id = $1;

-- name: GetRolesByPermissionID :many
SELECT r.role_id, r.role_name, r.description
FROM role r 
	INNER JOIN role_permission rp ON r.role_id = rp.role_id
	INNER JOIN permissions p ON p.action_id = rp.action_id
WHERE p.action_id = $1;

-- name: DeleteRolePermission :execresult
DELETE FROM role_permission WHERE action_id = $1 AND role_id = $2;

-- name: UpdateRolePermissionByRoleID_PUT :execresult
UPDATE role_permission 
	SET action_id = $1, 
	role_id = $2,
	updated_at = NOW() 
WHERE action_id = $1 AND role_id = $2;