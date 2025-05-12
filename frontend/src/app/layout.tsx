import type { Metadata } from 'next'
import { Epilogue, Noto_Sans } from 'next/font/google'
import './globals.css'

const epilogue = Epilogue({ 
  subsets: ['latin'],
  variable: '--font-epilogue',
  weight: ['400', '500', '700', '900']
})

const notoSans = Noto_Sans({ 
  subsets: ['latin'],
  variable: '--font-noto-sans',
  weight: ['400', '500', '700', '900']
})

export const metadata: Metadata = {
  title: 'Searchin',
  description: 'Search the web with Searchin',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body className={`${epilogue.variable} ${notoSans.variable} font-sans bg-white min-h-screen flex flex-col`}>
        {children}
      </body>
    </html>
  )
}