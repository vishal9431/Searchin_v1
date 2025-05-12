import { Suspense } from 'react'
import SearchClient from './SearchClient'

export default function SearchPage() {
  return (
    <Suspense fallback={<div className="text-center text-gray-500 py-10">Loading search...</div>}>
      <SearchClient />
    </Suspense>
  )
}
