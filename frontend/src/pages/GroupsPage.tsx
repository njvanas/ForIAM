import React, { useState, useEffect } from 'react'
import { Plus, Edit, Trash2, Search } from 'lucide-react'
import { groupsApi } from '../services/api'

interface Group {
  id: string
  name: string
  description?: string
  created_at: string
}

export default function GroupsPage() {
  const [groups, setGroups] = useState<Group[]>([])
  const [loading, setLoading] = useState(true)
  const [searchTerm, setSearchTerm] = useState('')

  useEffect(() => {
    fetchGroups()
  }, [])

  const fetchGroups = async () => {
    try {
      const data = await groupsApi.getGroups()
      setGroups(data || [])
    } catch (error) {
      console.error('Failed to fetch groups:', error)
      setGroups([])
    } finally {
      setLoading(false)
    }
  }

  const filteredGroups = groups.filter(group =>
    group.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
    group.description?.toLowerCase().includes(searchTerm.toLowerCase())
  )

  if (loading) {
    return (
      <div className="flex items-center justify-center h-64">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
      </div>
    )
  }

  return (
    <div>
      <div className="sm:flex sm:items-center sm:justify-between mb-6">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Groups</h1>
          <p className="mt-1 text-sm text-gray-600">
            Manage user groups and memberships
          </p>
        </div>
        <div className="mt-4 sm:mt-0">
          <button className="btn btn-primary">
            <Plus className="h-4 w-4 mr-2" />
            Add Group
          </button>
        </div>
      </div>

      <div className="card">
        <div className="mb-4">
          <div className="relative">
            <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 h-4 w-4 text-gray-400" />
            <input
              type="text"
              placeholder="Search groups..."
              className="form-input pl-10"
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
            />
          </div>
        </div>

        <div className="overflow-x-auto">
          <table className="table">
            <thead>
              <tr>
                <th>Name</th>
                <th>Description</th>
                <th>Created</th>
                <th>Actions</th>
              </tr>
            </thead>
            <tbody>
              {filteredGroups.length === 0 ? (
                <tr>
                  <td colSpan={4} className="text-center py-8 text-gray-500">
                    {searchTerm ? 'No groups found matching your search.' : 'No groups found. Create your first group to get started.'}
                  </td>
                </tr>
              ) : (
                filteredGroups.map((group) => (
                  <tr key={group.id}>
                    <td className="font-medium">{group.name}</td>
                    <td>{group.description || '-'}</td>
                    <td>{new Date(group.created_at).toLocaleDateString()}</td>
                    <td>
                      <div className="flex space-x-2">
                        <button className="btn btn-outline btn-sm">
                          <Edit className="h-3 w-3" />
                        </button>
                        <button className="btn btn-danger btn-sm">
                          <Trash2 className="h-3 w-3" />
                        </button>
                      </div>
                    </td>
                  </tr>
                ))
              )}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  )
}