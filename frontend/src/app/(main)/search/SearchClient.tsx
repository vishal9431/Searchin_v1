'use client'

import SearchBar from '@/components/SearchBar'
import Link from 'next/link'
import { useSearchParams } from 'next/navigation'
import { useEffect, useState } from 'react'

interface SearchResult {
  ID: number
  Url: string
  Title: string
  Content: string
}

async function getSearchResults(query: string): Promise<SearchResult[]> {
  try {
    if (!query) return []

    let baseUrl = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:3001'
    baseUrl = baseUrl.replace(/\/+$/, '')
    if (!baseUrl.startsWith('http')) baseUrl = `http://${baseUrl}`

    const url = new URL(`${baseUrl}/api/search`)
    url.searchParams.append('q', query)

    const response = await fetch(url.toString(), {
      cache: 'no-store',
      headers: {
        'Content-Type': 'application/json',
        Accept: 'application/json',
      },
    })

    if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`)

    const data = await response.json()
    return Array.isArray(data) ? data : []
  } catch (error) {
    console.error('Search error:', error)
    return []
  }
}

export default function SearchClient() {
  const searchParams = useSearchParams()
  const query = searchParams.get('q') || ''
  const [results, setResults] = useState<SearchResult[]>([])
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    const fetchResults = async () => {
      setLoading(true)
      const data = await getSearchResults(query)
      setResults(data)
      setLoading(false)
    }

    if (query) {
      fetchResults()
    } else {
      setResults([])
    }
  }, [query])

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-4xl mx-auto px-4 py-8">
        <div className="mb-8 sticky top-0 bg-gray-50 pt-4 pb-2 z-10">
          <div className="flex items-center gap-4">
            <SearchBar initialValue={query} />
            {query && (
              <span className="text-sm text-gray-500">
                {results.length} {results.length === 1 ? 'result' : 'results'} found
              </span>
            )}
          </div>
        </div>

        {!query ? (
          <div className="min-h-[60vh] flex flex-col items-center justify-center p-6">
            <div className="w-full max-w-2xl">
              <SearchBar />
              <p className="mt-6 text-center text-gray-500 text-lg">
                Search across thousands of resources to find what you need
              </p>
            </div>
          </div>
        ) : loading ? (
          <div className="text-center text-gray-500">Loading...</div>
        ) : results.length > 0 ? (
          <div className="space-y-6">
            {results.map((result) => {
              const url = new URL(result.Url)
              return (
                <div
                  key={result.ID}
                  className="bg-white rounded-xl shadow-md hover:shadow-lg transition-shadow p-6 border border-gray-100"
                >
                  <div className="mb-2">
                    <Link href={result.Url} target="_blank" rel="noopener noreferrer">
                      <div className="flex items-center gap-2 text-sm text-gray-600">
                        <span className="text-green-700 font-medium break-all">
                          {url.hostname.replace('www.', '')}
                        </span>
                        <span className="text-gray-400">/</span>
                        <span className="text-gray-500 truncate">{url.pathname}</span>
                      </div>
                    </Link>
                  </div>

                  <Link href={result.Url} target="_blank" rel="noopener noreferrer">
                    <h2 className="text-lg sm:text-xl font-semibold text-indigo-700 hover:underline mb-2">
                      {result.Title || 'Untitled Document'}
                    </h2>
                  </Link>

                  {result.Content && (
                    <p className="text-gray-700 text-sm line-clamp-2">{result.Content}</p>
                  )}
                </div>
              )
            })}
          </div>
        ) : (
          <div className="text-center py-20">
            <h3 className="text-2xl font-semibold text-gray-700 mb-2">No results found</h3>
            <p className="text-gray-500 text-base">
              We couldn’t find any matches for <span className="font-medium text-gray-800">&quot;{query}&quot;</span>
            </p>
          </div>
        )}
      </div>
    </div>
  )
}
