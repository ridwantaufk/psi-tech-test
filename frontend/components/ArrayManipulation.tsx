'use client'

import { useMemo } from 'react'
import { Card, Space, Tag } from 'antd'

const warna = ['merah', 'kuning', 'hijau', 'pink', 'ungu', 'maroon']
const baju = ['baju', 'celana', 'topi', 'jaket', 'sepatu']
const status = ['Diskon', 'Sale', 'Diskon', 'Sale', 'Sale']

function capital(s: string) {
    return s.charAt(0).toUpperCase() + s.slice(1).toLowerCase()
}

function manipulateArray(w: string[], b: string[], s: string[]): string[] {
    return w.map((warna, i) => {
        const baju = b[i % b.length]
        const status = s[i % s.length]
        return `${capital(baju)} ${capital(warna)} ${status}`
    })
}

export default function ArrayManipulation() {
    const items = useMemo(
        () => manipulateArray(warna, baju, status),
        []
    )

    return (
        <Card title="Hasil array manipulation" className="array-card" size="small">
            <Space wrap className="array-tag-list">
                {items.map((item, i) => (
                    <Tag
                        key={i}
                        color={item.toLowerCase().includes('diskon') ? 'blue' : 'green'}
                        className="array-tag"
                    >
                        {item}
                    </Tag>
                ))}
            </Space>
        </Card>
    )
}