'use client'

import { Table, Input, Button, Image, Tooltip, Row, Col, Typography } from 'antd'
import { SearchOutlined, ReloadOutlined } from '@ant-design/icons'
import type { ColumnsType } from 'antd/es/table'
import { ManipulatedUser } from '../types'

const { Text } = Typography

interface Props {
    data: ManipulatedUser[]
    loading: boolean
    search: string
    page: number
    pageSize: number
    onSearch: (val: string) => void
    onPageChange: (page: number, pageSize: number) => void
    onRefresh: () => void
}

const columns: ColumnsType<ManipulatedUser> = [
    {
        title: 'Nama',
        dataIndex: 'name',
        key: 'name',
        sorter: (a, b) => a.name.localeCompare(b.name),
    },
    {
        title: 'Umur',
        dataIndex: 'age',
        key: 'age',
        width: 75,
        sorter: (a, b) => a.age - b.age,
    },
    {
        title: 'Alamat',
        dataIndex: 'location',
        key: 'location',
        render: (loc: string) => (
            <Tooltip title={loc}>
                <Text type="secondary" className="address-detail-text">
                    (Alamat Detail)
                </Text>
            </Tooltip>
        ),
    },
    {
        title: 'Email',
        dataIndex: 'email',
        key: 'email',
    },
    {
        title: 'No. Telepon 1',
        dataIndex: 'phone',
        key: 'phone',
    },
    {
        title: 'No. Telepon 2',
        dataIndex: 'cell',
        key: 'cell',
    },
    {
        title: 'Gambar',
        dataIndex: 'picture',
        key: 'picture',
        render: (pics: string[]) => (
            <Image
                src={pics[1]}
                width={55}
                height={55}
                style={{ borderRadius: 4, objectFit: 'cover' }}
                preview={{ src: pics[0] }}
                alt="foto"
            />
        ),
    },
]

export default function UserTable({
    data,
    loading,
    search,
    page,
    pageSize,
    onSearch,
    onPageChange,
    onRefresh,
}: Props) {
    return (
        <>
            <Row gutter={12} className="table-toolbar" align="middle">
                <Col flex="auto">
                    <Input
                        className="search-input"
                        placeholder="Cari nama, email, atau lokasi"
                        prefix={<SearchOutlined className="search-icon" />}
                        value={search}
                        onChange={e => onSearch(e.target.value)}
                        allowClear
                    />
                </Col>
                <Col>
                    <Button
                        className="refresh-button"
                        type="primary"
                        icon={<ReloadOutlined />}
                        onClick={onRefresh}
                        loading={loading}
                    >
                        + New Data
                    </Button>
                </Col>
            </Row>
            <Table
                className="user-table"
                columns={columns}
                dataSource={data}
                loading={loading}
                rowKey="email"
                size="middle"
                bordered
                pagination={{
                    current: page,
                    pageSize,
                    total: data.length,
                    showSizeChanger: true,
                    pageSizeOptions: ['5', '10', '20', '50', '100'],
                    onChange: onPageChange,
                }}
            />
        </>
    )
}