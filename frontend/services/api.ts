import axios from 'axios'
import { ApiResponse } from '../types'

const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080',
  withCredentials: true,
})

export const getExternalUsers = async (results = 10, page = 1): Promise<ApiResponse> => {
  const res = await api.get(`/api/users/external?results=${results}&page=${page}`)
  return res.data
}

export const login = async (username: string, password: string) => {
  const res = await api.post('/auth/login', { username, password })
  return res.data
}

export const checkout = async (data: { harga_barang: number; voucher_code?: string }) => {
  const res = await api.post('/api/checkout', data)
  return res.data
}