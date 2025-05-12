import SearchBar from '@/components/SearchBar'
import Tag from '@/components/Tag'

export default function HomePage() {
  const trendingSearches = [
    'Brexit',
    'FIFA World Cup',
    'Bitcoin',
    'Game of Thrones',
    'iPhone 13',
    'New York Fashion Week',
    'Plastic pollution',
    'SpaceX',
    'Vegan recipes',
    'Yoga'
  ]

  const exploreMore = [
    'Images',
    'Video',
    'News',
    'Shopping',
    'Books',
    'Finance'
  ]

  return (
    <div className="px-10 flex-1 flex justify-center py-5">
      <div className="flex flex-col max-w-[960px] flex-1">
        <h1 className="text-[#111417] text-[28px] font-bold leading-tight px-4 text-center pb-3 pt-5">
          Searchin
        </h1>
        
        <SearchBar />
        
        <section className="px-4">
          <h2 className="text-[#111417] text-lg font-bold leading-tight tracking-[-0.015em] pb-2 pt-4">
            Trending searches
          </h2>
          <div className="flex gap-3 p-3 flex-wrap">
            {trendingSearches.map((search, index) => (
              <Tag key={index} text={search} />
            ))}
          </div>
        </section>
        
        <section className="px-4">
          <h2 className="text-[#111417] text-lg font-bold leading-tight tracking-[-0.015em] pb-2 pt-4">
            Explore more
          </h2>
          <div className="flex gap-3 p-3 flex-wrap">
            {exploreMore.map((item, index) => (
              <Tag key={index} text={item} />
            ))}
          </div>
        </section>
      </div>
    </div>
  )
}