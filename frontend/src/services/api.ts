import axios from 'axios'

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080'

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
})

// Add auth token to requests
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

// Handle auth errors
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token')
      window.location.href = '/login'
    }
    return Promise.reject(error)
  }
)

export const authApi = {
  login: async (email: string, password: string) => {
    const response = await api.post('/auth/login', { email, password })
    return response.data
  },
  
  logout: async () => {
    await api.post('/auth/logout')
  },
  
  getProfile: async () => {
    const response = await api.get('/auth/profile')
    return response.data
  }
}

export const usersApi = {
  getUsers: async () => {
    const response = await api.get('/users')
    return response.data
  },
  
  createUser: async (userData: any) => {
    const response = await api.post('/users', userData)
    return response.data
  },
  
  updateUser: async (id: string, userData: any) => {
    const response = await api.put(`/users/${id}`, userData)
    return response.data
  },
  
  deleteUser: async (id: string) => {
    await api.delete(`/users/${id}`)
  }
}

export const rolesApi = {
  getRoles: async () => {
    const response = await api.get('/roles')
    return response.data
  },
  
  createRole: async (roleData: any) => {
    const response = await api.post('/roles', roleData)
    return response.data
  },
  
  updateRole: async (id: string, roleData: any) => {
    const response = await api.put(`/roles/${id}`, roleData)
    return response.data
  },
  
  deleteRole: async (id: string) => {
    await api.delete(`/roles/${id}`)
  }
}

export const groupsApi = {
  getGroups: async () => {
    const response = await api.get('/groups')
    return response.data
  },
  
  createGroup: async (groupData: any) => {
    const response = await api.post('/groups', groupData)
    return response.data
  },
  
  updateGroup: async (id: string, groupData: any) => {
    const response = await api.put(`/groups/${id}`, groupData)
    return response.data
  },
  
  deleteGroup: async (id: string) => {
    await api.delete(`/groups/${id}`)
  }
}

export const auditApi = {
  getAuditLogs: async (params?: any) => {
    const response = await api.get('/audit', { params })
    return response.data
  }
}

export default api