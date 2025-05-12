import { NextResponse } from 'next/server'

export async function GET(request: Request) {
  try {
    const { searchParams } = new URL(request.url)
    const query = searchParams.get('q')
    
    if (!query) {
      return NextResponse.json(
        { error: 'Missing search query' },
        { status: 400 }
      )
    }

    // Mock data - replace with actual search API call
    const mockResults = Array(5).fill(null).map((_, i) => ({
      title: `Result ${i+1} for "${query}"`,
      url: `https://example.com/search?q=${encodeURIComponent(query)}&result=${i+1}`,
      description: `This is a sample result #${i+1} for your search "${query}". In a real app, this would be actual search results from your database or API.`
    }))

    // Simulate delay
    await new Promise(resolve => setTimeout(resolve, 300))

    return NextResponse.json({ results: mockResults })
  } catch (error) {
    console.error('Search API error:', error)
    return NextResponse.json(
      { error: 'Internal server error' },
      { status: 500 }
    )
  }
}