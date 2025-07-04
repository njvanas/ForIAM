import { describe, it, expect, beforeEach, afterEach } from 'vitest'
import axios from 'axios'
import { authApi, usersApi } from './api'

// Mock axios
vi.mock('axios')
const mockedAxios = vi.mocked(axios, true)

describe('API Services', () => {
  beforeEach(() => {
    // Reset mocks before each test
    vi.clearAllMocks()
    
    // Mock axios.create to return a mock instance
    mockedAxios.create.mockReturnValue({
      interceptors: {
        request: { use: vi.fn() },
        response: { use: vi.fn() }
      },
      post: vi.fn(),
      get: vi.fn(),
      put: vi.fn(),
      delete: vi.fn()
    } as any)
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