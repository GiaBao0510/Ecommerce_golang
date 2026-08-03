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

/*_______________ Bảng User 4 ___________________*/
-- name: CreateUser :exec
INSERT INTO "user"(uuid, id_status, user_name, birth_date, email, phone_num, address, password_hash, avatar_url)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);

-- name: UserEmailExists :one
SELECT EXISTS(SELECT 1 FROM "user" WHERE email = $1);

-- name: UserPhoneExists :one
SELECT EXISTS(SELECT 1 FROM "user" WHERE phone_num = $1);

-- name: GetUID_PasswordHashByEmail :one
SELECT uuid, password_hash FROM "user" WHERE email = $1;

-- name: GetUID_PasswordHashByPhone :one
SELECT uuid, password_hash FROM "user" WHERE phone_num = $1;

-- name: GetAllUsers :many
SELECT * FROM "user";

-- name: GetUserByID :one
SELECT * FROM "user" WHERE uuid = $1;

-- name: GetUserByEmail :one
SELECT * FROM "user" WHERE email = $1;

-- name: GetUserByPhone :one
SELECT * FROM "user" WHERE phone_num = $1;

-- name: UpdateUser_PUT :execresult
UPDATE "user" 
	SET id_status = $1, 
		user_name = $2, 
		birth_date = $3, 
		email = $4, phone_num = $5, address = $6 
	WHERE uuid = $7;

-- name: UpdateUser_PATCH :execresult
UPDATE "user" 
	SET id_status = COALESCE($1, id_status), 
		user_name = COALESCE($2, user_name), 
		birth_date = COALESCE($3, birth_date), 
		email = COALESCE($4, email), 
		phone_num = COALESCE($5, phone_num), 
		address = COALESCE($6, address),
		updated_at = NOW() 
	WHERE uuid = $7;

-- name: UpdateUserPassword_PATCH :execresult
UPDATE "user"
	SET password_hash = $1,
		updated_at = NOW()
	WHERE uuid = $2;

-- name: UpdateUserAvatar_PATCH :execresult
UPDATE "user"
	SET avatar_url = $1,
		updated_at = NOW()
	WHERE uuid = $2;

-- name: DeleteUser :execresult
DELETE FROM "user" WHERE uuid = $1;

-- name: VerifyEmail :execresult
UPDATE "user"
	SET is_email_verified = TRUE,
		updated_at = NOW()
	WHERE uuid = $1;

-- name: VerifyPhone :execresult
UPDATE "user"
	SET is_phonenum_verified = TRUE,
		updated_at = NOW()
	WHERE uuid = $1;

/*_______________ Bảng User-Roles 5 ___________________*/
-- name: CreateUserRole :exec
INSERT INTO user_role(uuid, role_id) VALUES($1, $2);

-- name: GetUserByRoleID :many
SELECT u.uuid, u.user_name, u.email, u.phone_num, u.address
FROM "user" u 
	JOIN user_role ur ON u.uuid = ur.uuid 
	JOIN role r ON r.role_id = ur.role_id
	WHERE r.role_id = $1;

-- name: GetRolesByUserID :many
SELECT r.role_id, r.role_name, r.description
FROM "user" u 
	JOIN user_role ur ON u.uuid = ur.uuid 
	JOIN role r ON r.role_id = ur.role_id
	WHERE u.uuid = $1;

-- name: UpdateUserRoleByUserID_PUT :exec
UPDATE user_role SET role_id = $1 WHERE uuid = $2;

-- name: DeleteUserRole :exec
DELETE FROM user_role WHERE uuid = $1 AND role_id = $2;