export interface ManipulatedUser {
  name: string
  location: string
  email: string
  age: number
  phone: string
  cell: string
  picture: string[]
}

export interface ApiResponse {
  data: ManipulatedUser[]
  results: number
  page: string
}