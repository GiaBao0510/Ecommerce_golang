import { useState } from 'react'
import { Navigate, Route, Routes, useLocation, useNavigate } from 'react-router-dom'
import { authApi } from './api'

const initialRegister = {
  user_name: '',
  email: '',
  phone_num: '',
  address: '',
  password_hash: '',
}

function AuthPage() {
  const location = useLocation()
  const navigate = useNavigate()
  const [mode, setMode] = useState(location.pathname === '/register' ? 'register' : 'login')
  const [loginForm, setLoginForm] = useState({ account: '', password: '' })
  const [registerForm, setRegisterForm] = useState(initialRegister)
  const [status, setStatus] = useState({ type: '', message: '' })
  const [submitting, setSubmitting] = useState(false)

  const switchMode = (nextMode) => {
    setMode(nextMode)
    setStatus({ type: '', message: '' })
    navigate(nextMode === 'register' ? '/register' : '/login', { replace: true })
  }

  const updateForm = (setter, field) => (event) => {
    setter((current) => ({ ...current, [field]: event.target.value }))
  }

  const handleLogin = async (event) => {
    event.preventDefault()
    setSubmitting(true)
    setStatus({ type: '', message: '' })

    try {
      const result = await authApi.login(loginForm)
      sessionStorage.setItem('auth_result', JSON.stringify(result))
      navigate('/dashboard', { state: { message: 'Đăng nhập thành công. Chào mừng bạn trở lại.' } })
    } catch (error) {
      setStatus({ type: 'error', message: error.message || 'Đăng nhập thất bại.' })
    } finally {
      setSubmitting(false)
    }
  }

  const handleRegister = async (event) => {
    event.preventDefault()
    setSubmitting(true)
    setStatus({ type: '', message: '' })

    try {
      await authApi.register({
        ...registerForm,
        id_status: 1,
      })
      setStatus({ type: 'success', message: 'Đăng ký thành công. Bạn có thể đăng nhập ngay.' })
      setLoginForm({ account: registerForm.email, password: '' })
      setMode('login')
      navigate('/login', { replace: true })
    } catch (error) {
      setStatus({ type: 'error', message: error.message || 'Đăng ký thất bại.' })
    } finally {
      setSubmitting(false)
    }
  }

  const handleGoogleLogin = () => {
    const form = document.createElement('form')
    form.method = 'POST'
    form.action = authApi.googleLoginUrl
    document.body.appendChild(form)
    form.submit()
  }

  return (
    <main className="auth-shell">
      <section className="brand-panel">
        <div className="brand-mark" aria-hidden="true">A/</div>
        <p className="eyebrow">Atelier commerce</p>
        <h1>Một tài khoản cho mọi lần mua sắm.</h1>
        <p className="brand-copy">
          Không gian tài khoản gọn gàng để bạn theo dõi hành trình của mình,
          từ lần ghé thăm đầu tiên đến đơn hàng tiếp theo.
        </p>
        <div className="brand-footnote">
          <span className="signal-dot" />
          <span>Hệ thống đang sẵn sàng</span>
        </div>
      </section>

      <section className="form-panel">
        <div className="form-header">
          <span className="section-kicker">Welcome back</span>
          <h2>{mode === 'login' ? 'Đăng nhập' : 'Tạo tài khoản'}</h2>
          <p>
            {mode === 'login'
              ? 'Tiếp tục đến không gian cá nhân của bạn.'
              : 'Bắt đầu một trải nghiệm mua sắm riêng dành cho bạn.'}
          </p>
        </div>

        <div className="mode-switch" role="tablist" aria-label="Chế độ tài khoản">
          <button className={mode === 'login' ? 'active' : ''} onClick={() => switchMode('login')} role="tab" type="button">
            Đăng nhập
          </button>
          <button className={mode === 'register' ? 'active' : ''} onClick={() => switchMode('register')} role="tab" type="button">
            Đăng ký
          </button>
        </div>

        {status.message && <div className={`status ${status.type}`} role="alert">{status.message}</div>}

        {mode === 'login' ? (
          <form onSubmit={handleLogin} className="auth-form">
            <label>
              Email hoặc số điện thoại
              <input value={loginForm.account} onChange={updateForm(setLoginForm, 'account')} placeholder="you@example.com" autoComplete="username" required />
            </label>
            <label>
              Mật khẩu
              <input type="password" value={loginForm.password} onChange={updateForm(setLoginForm, 'password')} placeholder="Nhập mật khẩu" autoComplete="current-password" required />
            </label>
            <button className="primary-button" disabled={submitting} type="submit">
              {submitting ? 'Đang xử lý...' : 'Đăng nhập'}
              <span aria-hidden="true">→</span>
            </button>
          </form>
        ) : (
          <form onSubmit={handleRegister} className="auth-form">
            <label>
              Họ và tên
              <input value={registerForm.user_name} onChange={updateForm(setRegisterForm, 'user_name')} placeholder="Nguyễn Văn A" minLength="2" required />
            </label>
            <div className="field-grid">
              <label>
                Email
                <input type="email" value={registerForm.email} onChange={updateForm(setRegisterForm, 'email')} placeholder="you@example.com" autoComplete="email" required />
              </label>
              <label>
                Số điện thoại
                <input value={registerForm.phone_num} onChange={updateForm(setRegisterForm, 'phone_num')} placeholder="0912..." autoComplete="tel" required />
              </label>
            </div>
            <label>
              Địa chỉ
              <input value={registerForm.address} onChange={updateForm(setRegisterForm, 'address')} placeholder="Địa chỉ nhận hàng" required />
            </label>
            <label>
              Mật khẩu
              <input type="password" value={registerForm.password_hash} onChange={updateForm(setRegisterForm, 'password_hash')} placeholder="Tối thiểu 8 ký tự" autoComplete="new-password" minLength="8" required />
            </label>
            <button className="primary-button" disabled={submitting} type="submit">
              {submitting ? 'Đang tạo tài khoản...' : 'Tạo tài khoản'}
              <span aria-hidden="true">→</span>
            </button>
          </form>
        )}

        <div className="divider"><span>hoặc tiếp tục với</span></div>
        <button className="google-button" onClick={handleGoogleLogin} type="button">
          <span className="google-icon" aria-hidden="true">G</span>
          Đăng nhập bằng Google
        </button>
        <p className="legal-copy">Bằng việc tiếp tục, bạn đồng ý với các điều khoản sử dụng của Atelier.</p>
      </section>
    </main>
  )
}

function Dashboard() {
  const location = useLocation()
  const message = location.state?.message || 'Bạn đã đăng nhập thành công.'

  return (
    <main className="dashboard-shell">
      <div className="dashboard-topbar">
        <div className="brand-mark small">A/</div>
        <span className="status-pill"><span className="signal-dot" /> Đã xác thực</span>
      </div>
      <section className="welcome-card">
        <span className="section-kicker">Account home</span>
        <h1>Chào mừng bạn đến Atelier.</h1>
        <p>{message}</p>
        <div className="dashboard-actions">
          <a href="http://localhost:8080/v1/api/checkStatus" target="_blank" rel="noreferrer">Kiểm tra API <span>↗</span></a>
          <a href="/login">Đăng xuất khỏi màn hình <span>→</span></a>
        </div>
      </section>
    </main>
  )
}

export default function App() {
  return (
    <Routes>
      <Route path="/" element={<Navigate to="/login" replace />} />
      <Route path="/login" element={<AuthPage />} />
      <Route path="/register" element={<AuthPage />} />
      <Route path="/dashboard" element={<Dashboard />} />
      <Route path="*" element={<Navigate to="/login" replace />} />
    </Routes>
  )
}
