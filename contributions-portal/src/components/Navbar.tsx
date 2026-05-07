'use client'

import { useAuthStore } from '@/store/auth'
import Link from 'next/link'
import { useRouter } from 'next/navigation'
import { useEffect, useState } from 'react'

export function Navbar() {
  const { contributor, clearContributor, isAuthenticated } = useAuthStore()
  const router = useRouter()
  const [mounted, setMounted] = useState(false)

  useEffect(() => {
    setMounted(true)
  }, [])

  function handleLogout() {
    clearContributor()
    router.push('/')
  }

  return (
    <div className="navbar bg-base-200 border-b border-base-300 px-4">
      <div className="navbar-start">
        <Link href="/" className="flex items-center gap-2">
          <span className="text-primary font-mono font-bold text-lg tracking-tight">
            {'<TecInfo />'}
          </span>
        </Link>
      </div>

      <div className="navbar-center hidden lg:flex">
        <ul className="menu menu-horizontal px-1 gap-1">
          <li>
            <Link href="/challenge/html-css" className="font-mono text-sm">
              Desafio HTML/CSS
            </Link>
          </li>
        </ul>
      </div>

      <div className="navbar-end gap-2">
        {/* Aguarda hidratação para evitar mismatch SSR/client com Zustand persist */}
        {!mounted ? (
          <div className="w-24 h-8 bg-base-300 rounded animate-pulse" />
        ) : isAuthenticated() ? (
          <>
            <div className="flex items-center gap-2">
              <div className="avatar placeholder">
                <div className="bg-primary text-primary-content rounded-full w-8">
                  <span className="text-xs font-bold">
                    {contributor?.full_name?.[0]?.toUpperCase() ?? '?'}
                  </span>
                </div>
              </div>
              <span className="text-sm hidden sm:inline text-base-content/70">
                {contributor?.full_name}
              </span>
            </div>
            <button onClick={handleLogout} className="btn btn-ghost btn-sm font-mono">
              Sair
            </button>
          </>
        ) : (
          <>
            <Link href="/login" className="btn btn-ghost btn-sm font-mono">
              Entrar
            </Link>
            <Link href="/register" className="btn btn-primary btn-sm font-mono">
              Cadastrar
            </Link>
          </>
        )}
      </div>
    </div>
  )
}
