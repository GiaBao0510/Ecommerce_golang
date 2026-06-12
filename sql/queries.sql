--// File: sql/queries.sql: tệp tin này chủ yếu chứa các câu lệnh thực thi

/*_______________ Bảng Status 0 ___________________*/
-- name: CreateStatus :exec
INSERT INTO status(name,description,updated_at,deleted_at) VALUES($1,$2,$3,$4) RETURNING id_status;

-- name: GetStatusByID :one
SELECT * FROM status WHERE id_status = $1;

-- name: GetAllStatus :many
SELECT * FROM status;

-- name: UpdateStatus :exec
UPDATE status SET name = $1, description = $2, updated_at = $3 WHERE id_status = $4;

-- name: DeleteStatus :exec
DELETE FROM status WHERE id_status = $1;

/*_______________ Bảng Role 1 ___________________*/
-- name: CreateRole :exec
INSERT INTO role(role_name,description) VALUES($1,$2) RETURNING role_id;

-- name: GetRoleByID :one
SELECT * FROM role WHERE role_id = $1;

-- name: GetAllRoles :many
SELECT * FROM role;

-- name: UpdateRole :exec
UPDATE role SET role_name = $1, description = $2 WHERE role_id = $3;

-- name: DeleteRole :exec
DELETE FROM role WHERE role_id = $1;

/*_______________ Bảng permissions 2 ___________________*/
-- name: CreatePermission :exec
INSERT INTO permissions(action_name, description, updated_at, deleted_at) VALUES($1,$2,$3,$4) RETURNING action_id;

-- name: GetPermissionByID :one
SELECT * FROM permissions WHERE action_id = $1;

-- name: GetAllPermissions :many
SELECT * FROM permissions;

-- name: UpdatePermission :exec
UPDATE permissions SET action_name = $1, description = $2 WHERE action_id = $3;

-- name: DeletePermission :exec
DELETE FROM permissions WHERE action_id = $1;

