import React, { useState, useEffect } from 'react'
import { Users, Shield, UserCheck, FileText } from 'lucide-react'
import { usersApi, rolesApi, groupsApi, auditApi } from '../services/api'

interface Stats {
  users: number
  roles: number
  groups: number
  auditLogs: number
}

export default function DashboardPage() {
  const [stats, setStats] = useState<Stats>({ users: 0, roles: 0, groups: 0, auditLogs: 0 })
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchStats = async () => {
      try {
        const [users, roles, groups, audit] = await Promise.all([
          usersApi.getUsers(),
          rolesApi.getRoles(),
          groupsApi.getGroups(),
          auditApi.getAuditLogs({ limit: 1 })
        ])

        setStats({
          users: users.length || 0,
          roles: roles.length || 0,
          groups: groups.length || 0,
          auditLogs: audit.total || 0
        })
      } catch (error) {
        console.error('Failed to fetch stats:', error)
      } finally {
        setLoading(false)
      }
    }

    fetchStats()
  }, [])

  const statCards = [
    {
      name: 'Total Users',
      value: stats.users,
      icon: Users,
      color: 'bg-blue-500',
      bgColor: 'bg-blue-50'
    },
    {
      name: 'Roles',
      value: stats.roles,
      icon: Shield,
      color: 'bg-green-500',
      bgColor: 'bg-green-50'
    },
    {
      name: 'Groups',
      value: stats.groups,
      icon: UserCheck,
      color: 'bg-purple-500',
      bgColor: 'bg-purple-50'
    },
    {
      name: 'Audit Events',
      value: stats.auditLogs,
      icon: FileText,
      color: 'bg-orange-500',
      bgColor: 'bg-orange-50'
    }
  ]

  if (loading) {
    return (
      <div className="flex items-center justify-center h-64">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
      </div>
    )
  }

  return (
    <div>
      <div className="mb-8">
        <h1 className="text-2xl font-bold text-gray-900">Dashboard</h1>
        <p className="mt-1 text-sm text-gray-600">
          Overview of your identity and access management system
        </p>
      </div>

      <div className="grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-4">
        {statCards.map((item) => (
          <div key={item.name} className={`card ${item.bgColor}`}>
            <div className="flex items-center">
              <div className="flex-shrink-0">
                <div className={`p-3 rounded-lg ${item.color}`}>
                  <item.icon className="h-6 w-6 text-white" />
                </div>
              </div>
              <div className="ml-5 w-0 flex-1">
                <dl>
                  <dt className="text-sm font-medium text-gray-500 truncate">
                    {item.name}
                  </dt>
                  <dd className="text-lg font-medium text-gray-900">
                    {item.value.toLocaleString()}
                  </dd>
                </dl>
              </div>
            </div>
          </div>
        ))}
      </div>

      <div className="mt-8 grid grid-cols-1 gap-5 lg:grid-cols-2">
        <div className="card">
          <h3 className="text-lg font-medium text-gray-900 mb-4">Recent Activity</h3>
          <div className="space-y-3">
            <div className="flex items-center justify-between py-2 border-b border-gray-100">
              <span className="text-sm text-gray-600">System initialized</span>
              <span className="text-xs text-gray-400">Just now</span>
            </div>
            <div className="flex items-center justify-between py-2 border-b border-gray-100">
              <span className="text-sm text-gray-600">Admin user created</span>
              <span className="text-xs text-gray-400">2 minutes ago</span>
            </div>
            <div className="flex items-center justify-between py-2">
              <span className="text-sm text-gray-600">Database schema initialized</span>
              <span className="text-xs text-gray-400">5 minutes ago</span>
            </div>
          </div>
        </div>

        <div className="card">
          <h3 className="text-lg font-medium text-gray-900 mb-4">System Status</h3>
          <div className="space-y-3">
            <div className="flex items-center justify-between">
              <span className="text-sm text-gray-600">Authentication Service</span>
              <span className="badge badge-success">Online</span>
            </div>
            <div className="flex items-center justify-between">
              <span className="text-sm text-gray-600">Database Connection</span>
              <span className="badge badge-success">Connected</span>
            </div>
            <div className="flex items-center justify-between">
              <span className="text-sm text-gray-600">Cache Service</span>
              <span className="badge badge-success">Active</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}