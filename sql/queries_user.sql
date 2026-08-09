/*_______________ Bảng User 4 ___________________*/
-- name: CreateUser :exec
INSERT INTO "user"(uuid, id_status, user_name, birth_date, email, phone_num, address, password_hash, avatar_url)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);

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

-- name: UpdateUserRoleByUserID_PUT :execresult
UPDATE user_role SET role_id = $1 WHERE uuid = $2;

-- name: DeleteUserRole :execresult
DELETE FROM user_role WHERE uuid = $1 AND role_id = $2;