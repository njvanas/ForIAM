import { describe, it, expect } from 'vitest'
import { render, screen } from '@testing-library/react'
import { BrowserRouter } from 'react-router-dom'
import App from './App'

// Mock the auth context
const MockAuthProvider = ({ children }: { children: React.ReactNode }) => {
  return <div data-testid="mock-auth-provider">{children}</div>
}

// Mock the components to avoid complex dependencies
vi.mock('./contexts/AuthContext', () => ({
  AuthProvider: MockAuthProvider,
  useAuth: () => ({
    isAuthenticated: false,
    loading: false,
    user: null,
    login: vi.fn(),
    logout: vi.fn()
  })
}))

vi.mock('./pages/LoginPage', () => ({
  default: () => <div data-testid="login-page">Login Page</div>
}))

describe('App', () => {
  it('renders without crashing', () => {
    render(
      <BrowserRouter>
        <App />
      </BrowserRouter>
    )
    
    expect(screen.getByTestId('mock-auth-provider')).toBeInTheDocument()
  })

  it('shows login page when not authenticated', () => {
    render(
      <BrowserRouter>
        <App />
      </BrowserRouter>
    )
    
    expect(screen.getByTestId('login-page')).toBeInTheDocument()
  })
})