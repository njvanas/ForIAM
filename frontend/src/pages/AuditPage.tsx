import React, { useState, useEffect } from 'react'
import { Search, Filter } from 'lucide-react'
import { auditApi } from '../services/api'

interface AuditLog {
  id: string
  action: string
  resource?: string
  user_id?: string
  ip_address?: string
  status: string
  created_at: string
}

export default function AuditPage() {
  const [logs, setLogs] = useState<AuditLog[]>([])
  const [loading, setLoading] = useState(true)
  const [searchTerm, setSearchTerm] = useState('')

  useEffect(() => {
    fetchAuditLogs()
  }, [])

  const fetchAuditLogs = async () => {
    try {
      const data = await auditApi.getAuditLogs()
      setLogs(data?.logs || [])
    } catch (error) {
      console.error('Failed to fetch audit logs:', error)
      setLogs([])
    } finally {
      setLoading(false)
    }
  }

  const filteredLogs = logs.filter(log =>
    log.action.toLowerCase().includes(searchTerm.toLowerCase()) ||
    log.resource?.toLowerCase().includes(searchTerm.toLowerCase()) ||
    log.status.toLowerCase().includes(searchTerm.toLowerCase())
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
      <div className="mb-6">
        <h1 className="text-2xl font-bold text-gray-900">Audit Logs</h1>
        <p className="mt-1 text-sm text-gray-600">
          Track all system activities and security events
        </p>
      </div>

      <div className="card">
        <div className="mb-4 flex space-x-4">
          <div className="flex-1 relative">
            <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 h-4 w-4 text-gray-400" />
            <input
              type="text"
              placeholder="Search audit logs..."
              className="form-input pl-10"
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
            />
          </div>
          <button className="btn btn-outline">
            <Filter className="h-4 w-4 mr-2" />
            Filter
          </button>
        </div>

        <div className="overflow-x-auto">
          <table className="table">
            <thead>
              <tr>
                <th>Action</th>
                <th>Resource</th>
                <th>Status</th>
                <th>IP Address</th>
                <th>Timestamp</th>
              </tr>
            </thead>
            <tbody>
              {filteredLogs.length === 0 ? (
                <tr>
                  <td colSpan={5} className="text-center py-8 text-gray-500">
                    {searchTerm ? 'No audit logs found matching your search.' : 'No audit logs available.'}
                  </td>
                </tr>
              ) : (
                filteredLogs.map((log) => (
                  <tr key={log.id}>
                    <td className="font-medium">{log.action}</td>
                    <td>{log.resource || '-'}</td>
                    <td>
                      <span className={`badge ${
                        log.status === 'success' ? 'badge-success' : 
                        log.status === 'error' ? 'badge-error' : 
                        'badge-warning'
                      }`}>
                        {log.status}
                      </span>
                    </td>
                    <td>{log.ip_address || '-'}</td>
                    <td>{new Date(log.created_at).toLocaleString()}</td>
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