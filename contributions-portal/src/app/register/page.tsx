'use client'

import { ApiError, contributorApi } from '@/lib/api'
import { useAuthStore } from '@/store/auth'
import Link from 'next/link'
import { useRouter } from 'next/navigation'
import { useState } from 'react'

export default function RegisterPage() {
  const router = useRouter()
  const setContributor = useAuthStore((s) => s.setContributor)

  const [form, setForm] = useState({ full_name: '', email: '', password: '' })
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    setLoading(true)
    setError(null)

    try {
      const contributor = await contributorApi.register(form)
      setContributor(contributor)
      router.push('/challenge/html-css')
    } catch (err) {
      setError(err instanceof ApiError ? err.message : 'Erro inesperado. Tente novamente.')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="min-h-[calc(100vh-4rem)] flex items-center justify-center px-4 py-12">
      <div className="card bg-base-200 border border-base-300 w-full max-w-md">
        <div className="card-body gap-6">
          <div className="text-center">
            <h1 className="text-2xl font-bold font-mono">Criar conta</h1>
            <p className="text-base-content/50 text-sm mt-1">
              Cadastre-se para enviar suas contribuições
            </p>
          </div>

          {error && (
            <div role="alert" className="alert alert-error text-sm font-mono">
              <span>✗ {error}</span>
            </div>
          )}

          <form onSubmit={handleSubmit} className="space-y-4">
            <fieldset className="fieldset">
              <legend className="fieldset-legend font-mono text-xs">Nome completo</legend>
              <input
                type="text"
                required
                placeholder="Seu Nome Completo"
                className="input input-bordered w-full font-mono"
                value={form.full_name}
                onChange={(e) => setForm({ ...form, full_name: e.target.value })}
              />
            </fieldset>

            <fieldset className="fieldset">
              <legend className="fieldset-legend font-mono text-xs">E-mail</legend>
              <input
                type="email"
                required
                placeholder="seu@email.com"
                className="input input-bordered w-full font-mono"
                value={form.email}
                onChange={(e) => setForm({ ...form, email: e.target.value })}
              />
            </fieldset>

            <fieldset className="fieldset">
              <legend className="fieldset-legend font-mono text-xs">Senha</legend>
              <input
                type="password"
                required
                minLength={6}
                placeholder="••••••••"
                className="input input-bordered w-full font-mono"
                value={form.password}
                onChange={(e) => setForm({ ...form, password: e.target.value })}
              />
            </fieldset>

            <button
              type="submit"
              disabled={loading}
              className="btn btn-primary w-full font-mono"
            >
              {loading ? <span className="loading loading-spinner loading-sm" /> : 'Cadastrar'}
            </button>
          </form>

          <p className="text-center text-sm text-base-content/50">
            Já tem conta?{' '}
            <Link href="/login" className="link link-primary font-mono">
              Entrar
            </Link>
          </p>
        </div>
      </div>
    </div>
  )
}
