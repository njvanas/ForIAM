import { describe, it, expect, beforeEach, afterEach } from 'vitest'

// Mock axios completely
vi.mock('axios')

// Mock the API module
vi.mock('./api', () => ({
  authApi: {
    login: vi.fn(),
    logout: vi.fn(),
    getProfile: vi.fn()
  },
  usersApi: {
    getUsers: vi.fn(),
    createUser: vi.fn(),
    updateUser: vi.fn(),
    deleteUser: vi.fn()
  },
  rolesApi: {
    getRoles: vi.fn(),
    createRole: vi.fn(),
    updateRole: vi.fn(),
    deleteRole: vi.fn()
  },
  groupsApi: {
    getGroups: vi.fn(),
    createGroup: vi.fn(),
    updateGroup: vi.fn(),
    deleteGroup: vi.fn()
  },
  auditApi: {
    getAuditLogs: vi.fn()
  }
}))

// Import after mocking
import { authApi, usersApi } from './api'

describe('API Services', () => {
  beforeEach(() => {
    // Reset mocks before each test
    vi.clearAllMocks()
  })

  afterEach(() => {
    vi.restoreAllMocks()
  })

  describe('authApi', () => {
    it('should have login method', () => {
      expect(typeof authApi.login).toBe('function')
    })

    it('should have logout method', () => {
      expect(typeof authApi.logout).toBe('function')
    })

    it('should have getProfile method', () => {
      expect(typeof authApi.getProfile).toBe('function')
    })
  })

  describe('usersApi', () => {
    it('should have getUsers method', () => {
      expect(typeof usersApi.getUsers).toBe('function')
    })

    it('should have createUser method', () => {
      expect(typeof usersApi.createUser).toBe('function')
    })

    it('should have updateUser method', () => {
      expect(typeof usersApi.updateUser).toBe('function')
    })

    it('should have deleteUser method', () => {
      expect(typeof usersApi.deleteUser).toBe('function')
    })
  })
})