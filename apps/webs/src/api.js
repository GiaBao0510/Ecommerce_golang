const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/v1/api'

async function request(path, options = {}) {
  const response = await fetch(`${API_BASE_URL}${path}`, {
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json',
      ...options.headers,
    },
    ...options,
  })

  const contentType = response.headers.get('content-type') || ''
  const body = contentType.includes('application/json')
    ? await response.json()
    : await response.text()

  if (!response.ok) {
    const message = typeof body === 'object'
      ? body.message || body.error || 'Yêu cầu không thành công.'
      : body || 'Yêu cầu không thành công.'
    throw new Error(message)
  }

  return body
}

export const authApi = {
  register: (payload) => request('/common/authen/register', {
    method: 'POST',
    body: JSON.stringify(payload),
  }),
  login: (payload) => request('/common/authen/login', {
    method: 'POST',
    body: JSON.stringify(payload),
  }),
  googleLoginUrl: `${API_BASE_URL}/common/authen/login/google`,
}
