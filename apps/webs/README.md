# Frontend xác thực người dùng

## Yêu cầu

- Node.js 18+
- Backend Go chạy tại `http://localhost:8080`
- CORS backend cho phép `http://localhost:5173`

## Chạy local

```bash
cd apps/webs
npm install
copy .env.example .env
npm run dev
```

Mở địa chỉ Vite in trong terminal, thường là `http://localhost:5173`.

Có thể đổi API backend trong `.env`:

```dotenv
VITE_API_BASE_URL=http://localhost:8080/v1/api
```

## Các luồng đã có

- Đăng ký: `POST /common/authen/register`
- Đăng nhập: `POST /common/authen/login`
- Đăng nhập Google: mở `POST /common/authen/login/google` bằng form điều hướng trình duyệt
- Sau đăng nhập thường, ứng dụng chuyển đến `/dashboard` và hiển thị thông báo thành công.
- Lỗi từ backend được hiển thị trực tiếp trên form.

## Lưu ý Google OAuth

Backend hiện phải đăng ký callback Google tại:

```text
GET /v1/api/common/authen/google/callback
```

và sau callback nên redirect về frontend, ví dụ:

```text
http://localhost:5173/dashboard?oauth=success
```

Frontend đã có nút bắt đầu OAuth. Nếu callback backend chưa được đăng ký hoặc Google OAuth chưa cấu hình đúng redirect URI, luồng Google sẽ không hoàn tất dù nút trên frontend hoạt động.
