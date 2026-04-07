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
    <main className="home-page">
      <header className="home-header">
        <Title level={3} className="home-title">
          Daftar isi data
        </Title>
      </header>

      <section className="home-section">
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
      </section>

      <ArrayManipulation />
    </main>
  )
}