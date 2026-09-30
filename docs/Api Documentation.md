# 📄 Ecommerce Golang — API Documentation (dành cho Front-end)

> **Base URL:** `http://localhost:8080/v1/api`  
> **Version:** 1.0  
> **Last Updated:** 2026-09-29

---

## Mục Lục

1. [Tổng Quan](#1-tổng-quan)
2. [Authentication APIs](#2-authentication-apis)
3. [User Management — Manager](#3-user-management--manager)
4. [User Profile — Client](#4-user-profile--client)
5. [Email Verification](#5-email-verification)
6. [Status Management](#6-status-management)
7. [Roles Management](#7-roles-management)
8. [Permission Management](#8-permission-management)
9. [Role-Permission Management](#9-role-permission-management)
10. [User-Role Management](#10-user-role-management)
11. [Product (Placeholder)](#11-product-placeholder)
12. [Health Check](#12-health-check)
13. [Mã Lỗi](#13-mã-lỗi)

---

## 1. Tổng Quan

### 1.1 Kiến trúc API

```mermaid
graph TB
    Client["🌐 Front-end Client"]
    
    subgraph API["API Gateway — /v1/api"]
        direction TB
        Public["🔓 Public Routes<br/>/common/..."]
        UserRoutes["🔐 User Routes<br/>/user/..."]
        ManagerRoutes["🔐 Manager Routes<br/>/manager/..."]
    end
    
    Client --> Public
    Client --> UserRoutes
    Client --> ManagerRoutes
    
    Public --> Auth["Authentication<br/>/common/authen"]
    
    UserRoutes --> UserProfile["User Profile<br/>/user/user"]
    UserRoutes --> UserEmail["Email Verify<br/>/user/email"]
    UserRoutes --> UserProduct["Product<br/>/user/product"]
    
    ManagerRoutes --> MgrUser["User CRUD<br/>/manager/user"]
    ManagerRoutes --> MgrStatus["Status CRUD<br/>/manager/status"]
    ManagerRoutes --> MgrRoles["Roles CRUD<br/>/manager/roles"]
    ManagerRoutes --> MgrPerm["Permission<br/>/manager/permission"]
    ManagerRoutes --> MgrRP["Role-Permission<br/>/manager/role_permission"]
    ManagerRoutes --> MgrUR["User-Role<br/>/manager/user_role"]
```

### 1.2 Xác thực (Authentication)

| Nhóm route | Middleware | Mô tả |
|------------|-----------|-------|
| `/v1/api/common/*` | Rate Limiting (Public) | Không cần token, public access |
| `/v1/api/user/*` | Rate Limiting (Private) + JWT Auth | Cần `Authorization: Bearer {token}` |
| `/v1/api/manager/*` | Rate Limiting (Private) + JWT Auth | Cần `Authorization: Bearer {token}` |

**Cách gửi token:**
```http
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

### 1.3 Format Response chuẩn

**✅ Success Response:**
```json
{
  "code": 200,
  "message": "Thông điệp thành công",
  "data": { ... }
}
```

**❌ Error Response:**
```json
{
  "code": 400,
  "status": "Bad Request",
  "message": "Mô tả chi tiết lỗi"
}
```

### 1.4 Format dữ liệu đặc biệt

| Trường | Format | Ví dụ |
|--------|--------|-------|
| `uuid` | UUID v4 | `"550e8400-e29b-41d4-a716-446655440000"` |
| `birth_date` | `dd-mm-yyyy` | `"15-03-1990"` |
| `email` | RFC 5322 | `"user@example.com"` |
| `phone_num` | VN phone | `"0912345678"` |

---

## 2. Authentication APIs

> **Base:** `/v1/api/common/authen`  
> **Access:** 🔓 Public (trừ Logout cần token)

---

### 2.1 Đăng ký — `POST /common/authen/register`

**Request Body:**
```json
{
  "id_status": 1,
  "user_name": "Nguyễn Văn A",
  "birth_date": "15-03-1990",
  "email": "nguyenvana@example.com",
  "phone_num": "0912345678",
  "address": "123 Đường ABC, Quận 1, TP.HCM",
  "password_hash": "MyStr0ngP@ssword",
  "avatar_url": "https://example.com/avatar.jpg"
}
```

| Trường | Kiểu | Bắt buộc | Validation |
|--------|------|----------|------------|
| `id_status` | `int` | ✅ | — |
| `user_name` | `string` | ✅ | min=2, max=100 |
| `birth_date` | `string` | ❌ | Format: `dd-mm-yyyy` |
| `email` | `string` | ✅ | Email hợp lệ |
| `phone_num` | `string` | ✅ | — |
| `address` | `string` | ✅ | — |
| `password_hash` | `string` | ✅ | — |
| `avatar_url` | `string` | ❌ | URL hợp lệ (nếu có) |

**Response:** `200 OK`
```json
{
  "code": 200,
  "message": "User registered successfully"
}
```

---

### 2.2 Đăng nhập — `POST /common/authen/login`

**Request Body:**
```json
{
  "account": "nguyenvana@example.com",
  "password": "MyStr0ngP@ssword"
}
```

| Trường | Kiểu | Bắt buộc | Mô tả |
|--------|------|----------|-------|
| `account` | `string` | ✅ | Email hoặc số điện thoại |
| `password` | `string` | ✅ | Mật khẩu |

**Response:** `200 OK`
```json
{
  "code": 200,
  "message": "Login successful",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
  }
}
```

> [!NOTE]
> Server cũng sẽ set 2 cookie HTTP-only: `access_token` và `refresh_token` (SameSite=Lax).

---

### 2.3 Đăng nhập Google — `POST /common/authen/login/google`

**Request:** Không cần body

**Response:** `307 Temporary Redirect`  
Redirect đến Google OAuth2 consent screen. Sau khi user đăng nhập Google, server sẽ callback và trả về token.

---

### 2.4 Refresh Token — `POST /common/authen/refresh`

**Request Body:**
```json
{
  "access_token": "eyJhbGciOi...",
  "refresh_token": "eyJhbGciOi..."
}
```

| Trường | Kiểu | Bắt buộc |
|--------|------|----------|
| `access_token` | `string` | ✅ |
| `refresh_token` | `string` | ✅ |

**Response:** `200 OK`
```json
{
  "code": 200,
  "message": "Refresh token successfully",
  "data": {
    "access_token": "eyJhbGciOiJNew...",
    "refresh_token": "eyJhbGciOiJNew..."
  }
}
```

---

### 2.5 Đăng xuất — `POST /common/authen/logout`

> 🔐 **Cần token:** `Authorization: Bearer {access_token}`

**Request Body:**
```json
{
  "access_token": "eyJhbGciOi...",
  "refresh_token": "eyJhbGciOi..."
}
```

**Response:** `200 OK`
```json
{
  "code": 200,
  "message": "Logout successful"
}
```

---

## 3. User Management — Manager

> **Base:** `/v1/api/manager/user`  
> **Access:** 🔐 JWT Required | Một số endpoint yêu cầu role **Admin**

---

### 3.1 Lấy tất cả users — `GET /manager/user`

> ⚠️ **Chỉ Admin** — Yêu cầu role Admin

**Request:** Không có params

**Response:** `200 OK`
```json
{
  "code": 200,
  "message": "Users retrieved successfully",
  "data": [
    {
      "uuid": "550e8400-e29b-41d4-a716-446655440000",
      "id_status": 1,
      "user_name": "Nguyễn Văn A",
      "email": "nguyenvana@example.com",
      "phone_num": "0912345678",
      "is_email_verified": true,
      "is_phonenum_verified": false,
      "address": "123 Đường ABC",
      "avatar_url": "https://example.com/avatar.jpg"
    }
  ]
}
```

---

### 3.2 Lấy user theo UUID — `GET /manager/user/{uuid}`

**Path Params:**

| Param | Kiểu | Validation |
|-------|------|------------|
| `uuid` | `string` | UUID format |

**Response:** `200 OK`
```json
{
  "code": 200,
  "message": "User retrieved successfully",
  "data": {
    "uuid": "550e8400-e29b-41d4-a716-446655440000",
    "id_status": 1,
    "user_name": "Nguyễn Văn A",
    "email": "nguyenvana@example.com",
    "phone_num": "0912345678",
    "is_email_verified": true,
    "is_phonenum_verified": false,
    "address": "123 Đường ABC",
    "avatar_url": ""
  }
}
```

---

### 3.3 Tìm user theo Email — `GET /manager/user/email/{email}`

**Path Params:**

| Param | Kiểu | Validation |
|-------|------|------------|
| `email` | `string` | Email format |

**Response:** `200 OK` — Tương tự response 3.2

---

### 3.4 Tìm user theo SĐT — `GET /manager/user/phone/{phone}`

**Path Params:**

| Param | Kiểu | Validation |
|-------|------|------------|
| `phone` | `string` | VN phone format |

**Response:** `200 OK` — Tương tự response 3.2

---

### 3.5 Tạo user mới — `POST /manager/user`

> ⚠️ **Chỉ Admin**

**Request Body:** Giống [Register Request Body](#21-đăng-ký--post-commonauthenregister)

**Response:** `201 Created`
```json
{
  "code": 201,
  "message": "User created successfully",
  "data": {
    "uuid": "550e8400-e29b-41d4-a716-446655440000"
  }
}
```

---

### 3.6 Cập nhật toàn bộ user — `PUT /manager/user/{uuid}`

**Path Params:** `uuid` (UUID format)

**Request Body:**
```json
{
  "id_status": 1,
  "user_name": "Nguyễn Văn A Updated",
  "birth_date": "15-03-1990",
  "email": "updated@example.com",
  "phone_num": "0987654321",
  "address": "456 Đường XYZ"
}
```

| Trường | Kiểu | Bắt buộc | Validation |
|--------|------|----------|------------|
| `id_status` | `int` | ✅ | — |
| `user_name` | `string` | ✅ | min=2, max=100 |
| `birth_date` | `string` | ❌ | `dd-mm-yyyy` |
| `email` | `string` | ✅ | Email hợp lệ |
| `phone_num` | `string` | ✅ | — |
| `address` | `string` | ✅ | — |

**Response:** `200 OK`
```json
{
  "code": 200,
  "message": "User updated successfully"
}
```

---

### 3.7 Cập nhật một phần user — `PATCH /manager/user/{uuid}`

**Path Params:** `uuid` (UUID format)

**Request Body:** Chỉ gửi trường cần cập nhật
```json
{
  "user_name": "Tên mới",
  "email": "newemail@example.com"
}
```

| Trường | Kiểu | Bắt buộc | Mô tả |
|--------|------|----------|-------|
| `id_status` | `int` | ❌ | Trạng thái mới |
| `user_name` | `string` | ❌ | min=2, max=100 |
| `birth_date` | `string` | ❌ | `dd-mm-yyyy` |
| `email` | `string` | ❌ | Email hợp lệ |
| `phone_num` | `string` | ❌ | — |
| `address` | `string` | ❌ | — |

**Response:** `200 OK`

---

### 3.8 Xóa user — `DELETE /manager/user/{uuid}`

**Path Params:** `uuid` (UUID format)

**Response:** `200 OK`
```json
{
  "code": 200,
  "message": "User deleted successfully"
}
```

---

## 4. User Profile — Client

> **Base:** `/v1/api/user/user`  
> **Access:** 🔐 JWT Required  
> **Mô tả:** API cho user tự quản lý thông tin cá nhân (giới hạn quyền)

---

### 4.1 Cập nhật toàn bộ profile — `PUT /user/user/{uuid}`

Giống [3.6 PUT /manager/user/{uuid}](#36-cập-nhật-toàn-bộ-user--put-manageruseruuid)

---

### 4.2 Cập nhật một phần profile — `PATCH /user/user/{uuid}`

Giống [3.7 PATCH /manager/user/{uuid}](#37-cập-nhật-một-phần-user--patch-manageruseruuid)

---

### 4.3 Xóa tài khoản — `DELETE /user/user/{uuid}`

Giống [3.8 DELETE /manager/user/{uuid}](#38-xóa-user--delete-manageruseruuid)

---

## 5. Email Verification

> **Base:** `/v1/api/user/email`  
> **Access:** 🔐 JWT Required

---

### 5.1 Gửi mã xác thực email — `GET /user/email/get_verification_code`

**Query Params:**

| Param | Kiểu | Bắt buộc | Validation |
|-------|------|----------|------------|
| `email` | `string` | ✅ | Email hợp lệ |

**Ví dụ:** `GET /user/email/get_verification_code?email=user@example.com`

**Response:** `200 OK`
```json
{
  "code": 200,
  "message": "Mã xác thực đã được gửi đến email của bạn"
}
```

---

### 5.2 Xác thực email — `POST /user/email/verify`

**Request Body:**
```json
{
  "email": "user@example.com",
  "otp": "123456"
}
```

| Trường | Kiểu | Bắt buộc | Validation |
|--------|------|----------|------------|
| `email` | `string` | ✅ | Email hợp lệ |
| `otp` | `string` | ✅ | Đúng 6 ký tự |

**Response:** `200 OK`
```json
{
  "code": 200,
  "message": "Email verified successfully"
}
```

---

### 5.3 Xác thực OTP qua email — `POST /user/email/verify-otp`

**Request Body:** Giống [5.2](#52-xác-thực-email--post-useremailverify)

**Response:** `200 OK`
```json
{
  "code": 200,
  "message": "OTP verified successfully"
}
```

---

## 6. Status Management

> **Base:** `/v1/api/manager/status`  
> **Access:** 🔐 JWT Required

---

### 6.1 Lấy tất cả status — `GET /manager/status`

**Response:** `200 OK`
```json
{
  "code": 200,
  "message": "All statuses retrieved successfully",
  "data": [
    {
      "id_status": 1,
      "name": "Active",
      "description": "Tài khoản đang hoạt động"
    },
    {
      "id_status": 2,
      "name": "Inactive",
      "description": "Tài khoản bị vô hiệu hóa"
    }
  ]
}
```

---

### 6.2 Lấy status theo ID — `GET /manager/status/{id}`

**Path Params:**

| Param | Kiểu | Validation |
|-------|------|------------|
| `id` | `int` | > 0 |

**Response:** `200 OK`
```json
{
  "code": 200,
  "message": "Status retrieved successfully",
  "data": {
    "id_status": 1,
    "name": "Active",
    "description": "Tài khoản đang hoạt động"
  }
}
```

---

### 6.3 Tạo status — `POST /manager/status`

**Request Body:**
```json
{
  "name": "Suspended",
  "description": "Tài khoản bị tạm ngưng"
}
```

| Trường | Kiểu | Bắt buộc | Validation |
|--------|------|----------|------------|
| `name` | `string` | ✅ | min=2, max=100 |
| `description` | `string` | ❌ | max=500 |

**Response:** `200 OK`
```json
{
  "code": 200,
  "message": "Status created successfully",
  "data": {
    "id": 3
  }
}
```

---

### 6.4 Cập nhật toàn bộ status — `PUT /manager/status/{id}`

**Path Params:** `id` (int, > 0)

**Request Body:**
```json
{
  "name": "Active Updated",
  "description": "Mô tả mới"
}
```

| Trường | Kiểu | Bắt buộc | Validation |
|--------|------|----------|------------|
| `name` | `string` | ✅ | min=2, max=100 |
| `description` | `string` | ❌ | max=500 |

---

### 6.5 Cập nhật một phần status — `PATCH /manager/status/{id}`

**Path Params:** `id` (int, > 0)  
**Request Body:** Chỉ gửi trường cần cập nhật

```json
{
  "name": "Updated Name"
}
```

---

### 6.6 Xóa status — `DELETE /manager/status/{id}`

**Path Params:** `id` (int, > 0)

**Response:** `200 OK`
```json
{
  "code": 200,
  "message": "Status deleted successfully"
}
```

---

## 7. Roles Management

> **Base:** `/v1/api/manager/roles`  
> **Access:** 🔐 JWT Required  
> **Pattern:** Giống [Status Management](#6-status-management)

---

### API Endpoints tổng hợp

| Method | Endpoint | Mô tả |
|--------|----------|-------|
| `GET` | `/manager/roles` | Lấy tất cả vai trò |
| `GET` | `/manager/roles/{id}` | Lấy vai trò theo ID |
| `POST` | `/manager/roles` | Tạo vai trò mới |
| `PUT` | `/manager/roles/{id}` | Cập nhật toàn bộ |
| `PATCH` | `/manager/roles/{id}` | Cập nhật một phần |
| `DELETE` | `/manager/roles/{id}` | Xóa vai trò |

### Request Body (POST / PUT):
```json
{
  "role_name": "Admin",
  "description": "Quản trị viên hệ thống"
}
```

| Trường | Kiểu | Bắt buộc | Validation |
|--------|------|----------|------------|
| `role_name` | `string` | ✅ | min=2, max=100 |
| `description` | `string` | ❌ | max=500 |

### Response Data (GET):
```json
{
  "role_id": 1,
  "role_name": "Admin",
  "description": "Quản trị viên hệ thống"
}
```

---

## 8. Permission Management

> **Base:** `/v1/api/manager/permission`  
> **Access:** 🔐 JWT Required  
> **Pattern:** Giống [Status Management](#6-status-management)

---

### API Endpoints tổng hợp

| Method | Endpoint | Mô tả |
|--------|----------|-------|
| `GET` | `/manager/permission` | Lấy tất cả quyền |
| `GET` | `/manager/permission/{id}` | Lấy quyền theo ID |
| `POST` | `/manager/permission` | Tạo quyền mới |
| `PUT` | `/manager/permission/{id}` | Cập nhật toàn bộ |
| `PATCH` | `/manager/permission/{id}` | Cập nhật một phần |
| `DELETE` | `/manager/permission/{id}` | Xóa quyền |

### Request Body (POST / PUT):
```json
{
  "action_name": "create_user",
  "description": "Quyền tạo người dùng"
}
```

| Trường | Kiểu | Bắt buộc | Validation |
|--------|------|----------|------------|
| `action_name` | `string` | ✅ | min=2, max=100 |
| `description` | `string` | ❌ | max=500 |

### Response Data (GET):
```json
{
  "action_id": 1,
  "action_name": "create_user",
  "description": "Quyền tạo người dùng"
}
```

---

## 9. Role-Permission Management

> **Base:** `/v1/api/manager/role_permission`  
> **Access:** 🔐 JWT Required

---

### 9.1 Lấy quyền theo vai trò — `GET /manager/role_permission/role/{id}`

**Path Params:** `id` (int, > 0) — ID vai trò

**Response:** `200 OK`
```json
{
  "code": 200,
  "message": "Permissions retrieved successfully",
  "data": [...]
}
```

---

### 9.2 Lấy vai trò theo quyền — `GET /manager/role_permission/permission/{id}`

**Path Params:** `id` (int, > 0) — ID quyền

**Response:** `200 OK`
```json
{
  "code": 200,
  "message": "Roles retrieved successfully",
  "data": [...]
}
```

---

### 9.3 Tạo liên kết vai trò-quyền — `POST /manager/role_permission`

**Request Body:**
```json
{
  "role_id": 1,
  "permission_id": 5
}
```

| Trường | Kiểu | Bắt buộc |
|--------|------|----------|
| `role_id` | `int` | ✅ |
| `permission_id` | `int` | ✅ |

**Response:** `201 Created`

---

### 9.4 Cập nhật liên kết — `PUT /manager/role_permission/{id}`

**Path Params:** `id` (int)

**Request Body:**
```json
{
  "role_id": 1,
  "permission_id": 3
}
```

---

### 9.5 Xóa liên kết — `DELETE /manager/role_permission/{id}`

**Path Params:** `id` (int)

**Response:** `200 OK`

---

## 10. User-Role Management

> **Base:** `/v1/api/manager/user_role`  
> **Access:** 🔐 JWT Required

---

### 10.1 Lấy vai trò của user — `GET /manager/user_role/user/{uuid}`

**Path Params:** `uuid` (UUID format)

**Response:** `200 OK`
```json
{
  "code": 200,
  "message": "Roles retrieved successfully",
  "data": [
    {
      "role_id": 1,
      "role_name": "Admin",
      "description": "Quản trị viên hệ thống"
    },
    {
      "role_id": 2,
      "role_name": "User",
      "description": "Người dùng thường"
    }
  ]
}
```

---

### 10.2 Lấy users theo vai trò — `GET /manager/user_role/role/{id}`

**Path Params:** `id` (int, > 0) — ID vai trò

**Response:** `200 OK`
```json
{
  "code": 200,
  "message": "Users retrieved successfully",
  "data": [
    {
      "uid": "550e8400-e29b-41d4-a716-446655440000",
      "user_name": "Nguyễn Văn A",
      "email": "nguyenvana@example.com",
      "phone_num": "0912345678",
      "address": "123 Đường ABC"
    }
  ]
}
```

---

### 10.3 Gán vai trò cho user — `POST /manager/user_role`

**Request Body:**
```json
{
  "id_role": 1,
  "uuid": "550e8400-e29b-41d4-a716-446655440000"
}
```

| Trường | Kiểu | Bắt buộc | Validation |
|--------|------|----------|------------|
| `id_role` | `int` | ✅ | — |
| `uuid` | `string` | ✅ | UUID format |

**Response:** `201 Created`

---

### 10.4 Cập nhật vai trò của user — `PUT /manager/user_role`

**Request Body:**
```json
{
  "id_role": 2,
  "uuid": "550e8400-e29b-41d4-a716-446655440000"
}
```

---

### 10.5 Xóa vai trò khỏi user — `DELETE /manager/user_role`

**Query Params:**

| Param | Kiểu | Bắt buộc | Validation |
|-------|------|----------|------------|
| `uuid` | `string` | ✅ | UUID format |
| `role_id` | `int` | ✅ | > 0 |

**Ví dụ:** `DELETE /manager/user_role?uuid=550e8400-...&role_id=1`

**Response:** `200 OK`

---

## 11. Product (Placeholder)

> **Base:** `/v1/api/user/product`  
> **Access:** 🔐 JWT Required  
> ⚠️ **Chưa có handler** — Các endpoint này đã được đăng ký route nhưng chưa có logic xử lý

| Method | Endpoint | Mô tả |
|--------|----------|-------|
| `GET` | `/user/product/search` | Tìm kiếm sản phẩm |
| `GET` | `/user/product/detail/{id}` | Chi tiết sản phẩm |
| `GET` | `/user/product/list` | Danh sách sản phẩm |

---

## 12. Health Check

> **Access:** 🔓 Public

---

### 12.1 Check Status — `GET /checkStatus`

**Response:** `200 OK`
```json
{
  "status": "ok"
}
```

---

### 12.2 Health Check — `GET /health` *(legacy router)*

**Response:** `200 OK`
```json
{
  "code": 200,
  "message": "Ứng dụng đang hoạt động"
}
```

---

### 12.3 Server Live — `GET /health/live` *(legacy router)*

**Response:** `200 OK`
```json
{
  "code": 200,
  "message": "Server đang hoạt động",
  "data": {
    "status": "alive",
    "uptime": "2h30m15s",
    "memory": {
      "used_mb": 45.2,
      "total_mb": 128.0
    },
    "goroutines": 15
  }
}
```

---

## 13. Mã Lỗi

### HTTP Status Codes

| Code | Status | Mô tả |
|------|--------|-------|
| `200` | OK | Thành công |
| `201` | Created | Tạo mới thành công |
| `307` | Temporary Redirect | Redirect (OAuth2) |
| `400` | Bad Request | Dữ liệu đầu vào không hợp lệ |
| `401` | Unauthorized | Chưa xác thực / Token hết hạn |
| `403` | Forbidden | Không có quyền truy cập |
| `404` | Not Found | Không tìm thấy resource |
| `409` | Conflict | Dữ liệu đã tồn tại (duplicate) |
| `429` | Too Many Requests | Vượt quá giới hạn request (Rate Limit) |
| `500` | Internal Server Error | Lỗi server |

### Validation Error Response

Khi dữ liệu đầu vào không hợp lệ:

```json
{
  "code": 400,
  "status": "Bad Request",
  "message": "Dữ liệu không hợp lệ: [chi tiết]"
}
```

Các loại lỗi validation phổ biến:

| Tag | Mô tả |
|-----|-------|
| `required` | Trường bắt buộc không được để trống |
| `email` | Phải là email hợp lệ |
| `uuid` | Phải đúng format UUID |
| `min=N` | Độ dài tối thiểu N ký tự |
| `max=N` | Độ dài tối đa N ký tự |
| `gt=0` | Phải lớn hơn 0 |
| `url` | Phải là URL hợp lệ |
| `len=6` | Phải đúng 6 ký tự |

---

## Bảng tổng hợp tất cả API Endpoints

| # | Method | Full Path | Auth | Mô tả |
|---|--------|-----------|------|-------|
| 1 | `POST` | `/v1/api/common/authen/register` | 🔓 | Đăng ký |
| 2 | `POST` | `/v1/api/common/authen/login` | 🔓 | Đăng nhập |
| 3 | `POST` | `/v1/api/common/authen/login/google` | 🔓 | Đăng nhập Google |
| 4 | `POST` | `/v1/api/common/authen/refresh` | 🔓 | Refresh token |
| 5 | `POST` | `/v1/api/common/authen/logout` | 🔐 | Đăng xuất |
| 6 | `GET` | `/v1/api/manager/user` | 🔐 Admin | Lấy tất cả users |
| 7 | `GET` | `/v1/api/manager/user/{uuid}` | 🔐 | User theo UUID |
| 8 | `GET` | `/v1/api/manager/user/email/{email}` | 🔐 | User theo email |
| 9 | `GET` | `/v1/api/manager/user/phone/{phone}` | 🔐 | User theo SĐT |
| 10 | `POST` | `/v1/api/manager/user` | 🔐 Admin | Tạo user |
| 11 | `PUT` | `/v1/api/manager/user/{uuid}` | 🔐 | Cập nhật toàn bộ user |
| 12 | `PATCH` | `/v1/api/manager/user/{uuid}` | 🔐 | Cập nhật một phần user |
| 13 | `DELETE` | `/v1/api/manager/user/{uuid}` | 🔐 | Xóa user |
| 14 | `PUT` | `/v1/api/user/user/{uuid}` | 🔐 | Cập nhật profile |
| 15 | `PATCH` | `/v1/api/user/user/{uuid}` | 🔐 | Cập nhật một phần profile |
| 16 | `DELETE` | `/v1/api/user/user/{uuid}` | 🔐 | Xóa tài khoản |
| 17 | `GET` | `/v1/api/user/email/get_verification_code` | 🔐 | Gửi mã xác thực |
| 18 | `POST` | `/v1/api/user/email/verify` | 🔐 | Xác thực email |
| 19 | `POST` | `/v1/api/user/email/verify-otp` | 🔐 | Xác thực OTP email |
| 20 | `GET` | `/v1/api/manager/status` | 🔐 | Tất cả status |
| 21 | `GET` | `/v1/api/manager/status/{id}` | 🔐 | Status theo ID |
| 22 | `POST` | `/v1/api/manager/status` | 🔐 | Tạo status |
| 23 | `PUT` | `/v1/api/manager/status/{id}` | 🔐 | Cập nhật status |
| 24 | `PATCH` | `/v1/api/manager/status/{id}` | 🔐 | Cập nhật 1 phần status |
| 25 | `DELETE` | `/v1/api/manager/status/{id}` | 🔐 | Xóa status |
| 26 | `GET` | `/v1/api/manager/roles` | 🔐 | Tất cả roles |
| 27 | `GET` | `/v1/api/manager/roles/{id}` | 🔐 | Role theo ID |
| 28 | `POST` | `/v1/api/manager/roles` | 🔐 | Tạo role |
| 29 | `PUT` | `/v1/api/manager/roles/{id}` | 🔐 | Cập nhật role |
| 30 | `PATCH` | `/v1/api/manager/roles/{id}` | 🔐 | Cập nhật 1 phần role |
| 31 | `DELETE` | `/v1/api/manager/roles/{id}` | 🔐 | Xóa role |
| 32 | `GET` | `/v1/api/manager/permission` | 🔐 | Tất cả permissions |
| 33 | `GET` | `/v1/api/manager/permission/{id}` | 🔐 | Permission theo ID |
| 34 | `POST` | `/v1/api/manager/permission` | 🔐 | Tạo permission |
| 35 | `PUT` | `/v1/api/manager/permission/{id}` | 🔐 | Cập nhật permission |
| 36 | `PATCH` | `/v1/api/manager/permission/{id}` | 🔐 | Cập nhật 1 phần perm |
| 37 | `DELETE` | `/v1/api/manager/permission/{id}` | 🔐 | Xóa permission |
| 38 | `GET` | `/v1/api/manager/role_permission/role/{id}` | 🔐 | Permissions theo role |
| 39 | `GET` | `/v1/api/manager/role_permission/permission/{id}` | 🔐 | Roles theo permission |
| 40 | `POST` | `/v1/api/manager/role_permission` | 🔐 | Tạo role-permission |
| 41 | `PUT` | `/v1/api/manager/role_permission/{id}` | 🔐 | Cập nhật role-perm |
| 42 | `DELETE` | `/v1/api/manager/role_permission/{id}` | 🔐 | Xóa role-permission |
| 43 | `GET` | `/v1/api/manager/user_role/user/{uuid}` | 🔐 | Roles của user |
| 44 | `GET` | `/v1/api/manager/user_role/role/{id}` | 🔐 | Users theo role |
| 45 | `POST` | `/v1/api/manager/user_role` | 🔐 | Gán role cho user |
| 46 | `PUT` | `/v1/api/manager/user_role` | 🔐 | Cập nhật user-role |
| 47 | `DELETE` | `/v1/api/manager/user_role` | 🔐 | Xóa user-role |
| 48 | `GET` | `/v1/api/user/product/search` | 🔐 | 🚧 Tìm sản phẩm |
| 49 | `GET` | `/v1/api/user/product/detail/{id}` | 🔐 | 🚧 Chi tiết SP |
| 50 | `GET` | `/v1/api/user/product/list` | 🔐 | 🚧 Danh sách SP |
| 51 | `GET` | `/v1/api/checkStatus` | 🔓 | Health check |
