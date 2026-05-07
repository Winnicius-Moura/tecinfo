import { Navbar } from '@/components/Navbar'
import type { Metadata } from 'next'
import { Geist, Geist_Mono } from 'next/font/google'
import './globals.css'

const geistSans = Geist({
  variable: '--font-geist-sans',
  subsets: ['latin'],
})

const geistMono = Geist_Mono({
  variable: '--font-geist-mono',
  subsets: ['latin'],
})

export const metadata: Metadata = {
  title: 'TecInfo — Contributions Portal',
  description: 'Plataforma de análise de exercícios do curso TecInfo 2025–2026',
}

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <html
      lang="pt-BR"
      data-theme="dark"
      className={`${geistSans.variable} ${geistMono.variable} h-full antialiased`}
    >
      <body className="min-h-full flex flex-col bg-base-100 text-base-content">
        <Navbar />
        <main className="flex-1">{children}</main>
        <footer className="footer footer-center py-6 text-base-content/40 text-sm">
          <p>TecInfo 2025–2026 — Feito para ensinar open source na prática</p>
        </footer>
      </body>
    </html>
  );
}
