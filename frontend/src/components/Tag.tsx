import Link from 'next/link'

interface TagProps {
  text: string
}

export default function Tag({ text }: TagProps) {
  return (
    <Link
      href={`/search?q=${encodeURIComponent(text)}`}
      className="flex h-8 shrink-0 items-center justify-center gap-x-2 rounded-full bg-[#f0f2f4] px-4 text-[#111417] text-sm font-medium leading-normal hover:bg-[#e4e6e8] transition-colors duration-200"
    >
      {text}
    </Link>
  )
}
