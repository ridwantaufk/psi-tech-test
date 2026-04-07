'use client'

import { Typography } from 'antd'
import UserTable from '../components/UserTable'
import ArrayManipulation from '../components/ArrayManipulation'
import { useUsers } from '../hooks/useUsers'

const { Title } = Typography

export default function HomePage() {
  const {
    users,
    loading,
    search,
    setSearch,
    page,
    pageSize,
    setPage,
    setPageSize,
    fetchData,
  } = useUsers()

  return (
    <div style={{ padding: '32px 40px' }}>
      <Title level={4} style={{ marginBottom: 16 }}>
        List
      </Title>

      <UserTable
        data={users}
        loading={loading}
        search={search}
        page={page}
        pageSize={pageSize}
        onSearch={setSearch}
        onPageChange={(p, ps) => {
          setPage(p)
          setPageSize(ps)
        }}
        onRefresh={fetchData}
      />

      <ArrayManipulation />
    </div>
  )
}