import type { Metadata } from 'next'
import Header from '@/components/Header'
// import './globals.css'

export const metadata: Metadata = {
  title: 'Searchin - Search the Web',
  description: 'A modern search engine for the web',
}

export default function MainLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <div className="flex flex-col min-h-screen">
      <Header />
      <main className="flex-1">
        {children}
      </main>
      <Footer />
    </div>
  )
}

function Footer() {
  return (
    <footer className="border-t border-[#f0f2f4] py-6 mt-8">
      <div className="container mx-auto px-10">
        <div className="flex flex-col md:flex-row justify-between items-center">
          <div className="flex items-center gap-4 mb-4 md:mb-0">
            <div className="size-4">
              <svg viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
                <g clipPath="url(#clip0_6_535)">
                  <path
                    fillRule="evenodd"
                    clipRule="evenodd"
                    d="M47.2426 24L24 47.2426L0.757355 24L24 0.757355L47.2426 24ZM12.2426 21H35.7574L24 9.24264L12.2426 21Z"
                    fill="currentColor"
                  ></path>
                </g>
              </svg>
            </div>
            <span className="text-[#111417] font-medium">Searchin</span>
          </div>
          <div className="flex gap-6">
            <a href="#" className="text-[#647587] text-sm hover:text-[#111417]">Privacy</a>
            <a href="#" className="text-[#647587] text-sm hover:text-[#111417]">Terms</a>
            <a href="#" className="text-[#647587] text-sm hover:text-[#111417]">Settings</a>
          </div>
        </div>
      </div>
    </footer>
  )
}