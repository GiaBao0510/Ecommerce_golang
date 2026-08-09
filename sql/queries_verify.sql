/*_______ User Verification Queries _________*/

-- name: UserEmailExists :one
SELECT EXISTS(SELECT 1 FROM "user" WHERE email = $1);

-- name: UserPhoneExists :one
SELECT EXISTS(SELECT 1 FROM "user" WHERE phone_num = $1);

-- name: UserEmailExists_HasNotBeenVerified :one
SELECT EXISTS(SELECT 1 FROM "user" WHERE email = $1 AND is_email_verified = FALSE);

-- name: UserPhoneExists_HasNotBeenVerified :one
SELECT EXISTS(SELECT 1 FROM "user" WHERE phone_num = $1 AND is_phonenum_verified = FALSE);

-- name: VerifyEmail :execresult
UPDATE "user"
	SET is_email_verified = TRUE,
		updated_at = NOW()
	WHERE email = $1;

-- name: VerifyPhone :execresult
UPDATE "user"
	SET is_phonenum_verified = TRUE,
		updated_at = NOW()
	WHERE phone_num = $1;