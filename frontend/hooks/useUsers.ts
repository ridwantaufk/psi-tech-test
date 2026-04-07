'use client'

import { useState, useEffect, useMemo } from 'react'
import { ManipulatedUser } from '../types'
import { getExternalUsers } from '../services/api'

export function useUsers() {
  const [users, setUsers] = useState<ManipulatedUser[]>([])
  const [loading, setLoading] = useState(false)
  const [page, setPage] = useState(1)
  const [pageSize, setPageSize] = useState(10)
  const [search, setSearch] = useState('')

  const fetchData = async () => {
    setLoading(true)
    try {
      const res = await getExternalUsers(pageSize, page)
      setUsers(res.data || [])
    } catch (err) {
      console.error('gagal fetch:', err)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchData()
  }, [page, pageSize])

  const filtered = useMemo(() => {
    if (!search.trim()) return users
    const q = search.toLowerCase()
    return users.filter(u =>
      u.name.toLowerCase().includes(q) ||
      u.email.toLowerCase().includes(q) ||
      u.location.toLowerCase().includes(q)
    )
  }, [users, search])
  return {
    users: filtered,
    loading,
    search,
    setSearch,
    page,
    pageSize,
    setPage,
    setPageSize,
    fetchData,
  }
}