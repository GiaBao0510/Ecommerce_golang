package util

// LogoURL là link logo của dự án e-commerce.
// Bạn có thể chuyển thành biến môi trường (os.Getenv) nếu muốn cấu hình linh hoạt hơn.
const LogoURL = "https://res.cloudinary.com/hk2vtuba/image/upload/v1785711493/Gemini_Generated_Image_82ea7s82ea7s82ea_hgmqlg.png"

// OTP_CodeSendingTemplate trả về một chuỗi HTML đại diện cho mẫu email gửi mã OTP.
func OTP_CodeSendingTemplate(otp string) string {
	return `<!DOCTYPE html>
<html lang="vi" xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<title>Mã xác thực OTP</title>
<style>
  body, table, td, a { -webkit-text-size-adjust: 100%; -ms-text-size-adjust: 100%; }
  table, td { mso-table-lspace: 0pt; mso-table-rspace: 0pt; }
  img { -ms-interpolation-mode: bicubic; border: 0; outline: none; text-decoration: none; }
  body { margin: 0; padding: 0; width: 100% !important; height: 100% !important; background-color: #F3F4F6; font-family: 'Segoe UI', Roboto, Helvetica, Arial, sans-serif; }

  @media only screen and (max-width: 600px) {
    .email-container { width: 100% !important; }
    .fluid-padding { padding-left: 20px !important; padding-right: 20px !important; }
    .otp-code { font-size: 30px !important; letter-spacing: 6px !important; }
  }
</style>
</head>
<body style="margin:0; padding:0; background-color:#F3F4F6;">
  <!-- Preheader (ẩn, hiển thị ở phần preview trong hộp thư) -->
  <div style="display:none; max-height:0; overflow:hidden; mso-hide:all;">
    Mã xác thực OTP của bạn sẽ hết hạn sau 5 phút. Vui lòng không chia sẻ mã này với bất kỳ ai.
  </div>

  <table role="presentation" width="100%" cellpadding="0" cellspacing="0" style="background-color:#F3F4F6;">
    <tr>
      <td align="center" style="padding: 32px 16px;">

        <table role="presentation" class="email-container" width="600" cellpadding="0" cellspacing="0" style="width:600px; max-width:600px; background-color:#FFFFFF; border-radius:16px; overflow:hidden; box-shadow: 0 4px 24px rgba(17, 24, 39, 0.08);">

          <!-- Header: Logo -->
          <tr>
            <td align="center" style="background-color:#111827; padding: 28px 24px;">
              <img src="` + LogoURL + `" alt="Logo" width="150" style="display:block; max-width:150px; height:auto;">
            </td>
          </tr>

          <!-- Accent bar -->
          <tr>
            <td style="height:4px; line-height:4px; font-size:0; background: linear-gradient(90deg, #F97316 0%, #FB923C 50%, #FDBA74 100%);">&nbsp;</td>
          </tr>

          <!-- Body content -->
          <tr>
            <td class="fluid-padding" style="padding: 40px 48px 24px 48px;">
              <h1 style="margin:0 0 8px 0; font-size:22px; line-height:28px; color:#111827; font-weight:700;">
                Xác thực tài khoản của bạn
              </h1>
              <p style="margin:0 0 20px 0; font-size:15px; line-height:24px; color:#4B5563;">
                Xin chào,<br>
                Chúng tôi nhận được yêu cầu xác thực cho tài khoản của bạn. Vui lòng sử dụng mã OTP bên dưới để hoàn tất quá trình xác thực và tiếp tục mua sắm.
              </p>
            </td>
          </tr>

          <!-- OTP Box -->
          <tr>
            <td align="center" style="padding: 0 48px 24px 48px;">
              <table role="presentation" width="100%" cellpadding="0" cellspacing="0">
                <tr>
                  <td align="center" style="background-color:#FFF7ED; border:1px dashed #FDBA74; border-radius:12px; padding: 24px;">
                    <p style="margin:0 0 8px 0; font-size:12px; letter-spacing:1px; color:#9A3412; text-transform:uppercase; font-weight:600;">
                      Mã xác thực của bạn
                    </p>
                    <span class="otp-code" style="display:inline-block; font-size:36px; line-height:1; font-weight:800; letter-spacing:10px; color:#EA580C; font-family: 'Courier New', Consolas, monospace;">
                      ` + otp + `
                    </span>
                  </td>
                </tr>
              </table>
            </td>
          </tr>

          <!-- Expiry note -->
          <tr>
            <td class="fluid-padding" style="padding: 0 48px 24px 48px;">
              <p style="margin:0; font-size:14px; line-height:22px; color:#6B7280; text-align:center;">
                ⏱ Mã có hiệu lực trong <strong style="color:#111827;">5 phút</strong> kể từ khi email này được gửi.
              </p>
            </td>
          </tr>

          <!-- Security warning -->
          <tr>
            <td class="fluid-padding" style="padding: 0 48px 32px 48px;">
              <table role="presentation" width="100%" cellpadding="0" cellspacing="0" style="background-color:#F9FAFB; border-left:4px solid #F97316; border-radius:6px;">
                <tr>
                  <td style="padding: 14px 18px;">
                    <p style="margin:0; font-size:13px; line-height:20px; color:#4B5563;">
                      🔒 <strong>Lưu ý bảo mật:</strong> Không chia sẻ mã này với bất kỳ ai, kể cả nhân viên hỗ trợ. Nếu bạn không thực hiện yêu cầu này, vui lòng bỏ qua email hoặc liên hệ đội ngũ hỗ trợ ngay lập tức.
                    </p>
                  </td>
                </tr>
              </table>
            </td>
          </tr>

          <!-- Divider -->
          <tr>
            <td style="border-top:1px solid #E5E7EB;"></td>
          </tr>

          <!-- Footer -->
          <tr>
            <td align="center" style="padding: 24px 48px 32px 48px;">
              <p style="margin:0 0 6px 0; font-size:12px; line-height:18px; color:#9CA3AF;">
                Đây là email tự động, vui lòng không trả lời trực tiếp email này.
              </p>
              <p style="margin:0; font-size:12px; line-height:18px; color:#9CA3AF;">
                © 2026 E-Commerce Store. Mọi quyền được bảo lưu.
              </p>
            </td>
          </tr>

        </table>

      </td>
    </tr>
  </table>
</body>
</html>`
}