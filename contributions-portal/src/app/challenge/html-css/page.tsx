'use client'

import { AnalysisReport } from '@/components/AnalysisReport'
import { useSubmissions } from '@/hooks/useSubmissions'
import { ApiError, htmlCssApi } from '@/lib/api'
import { useAuthStore } from '@/store/auth'
import type { HtmlCssAnalysisReport } from '@/types'
import { html } from '@codemirror/lang-html'
import dynamic from 'next/dynamic'
import Link from 'next/link'
import { useCallback, useState } from 'react'

const CodeMirror = dynamic(() => import('@uiw/react-codemirror'), { ssr: false })

const STARTER_CODE = `<article>
  <!-- Substitua pelos seus dados reais -->
  <h3>Seu Nome Completo</h3>
  <p>Uma frase sobre você ❤️</p>

  <h4>Programming languages I use</h4>
  <section>
    <div>JavaScript</div>
    <div>Python</div>
  </section>

  <h4 title="social-links-tecinfo">Social Links</h4>
  <section>
    <a href="https://linkedin.com/in/seu-usuario" target="_blank">
      <img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/linkedin/linkedin-original.svg" />
      LinkedIn
    </a>
    <a href="https://github.com/seu-usuario" target="_blank">
      <img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/github/github-original.svg" />
      GitHub
    </a>
  </section>
</article>

<style>
  article {
    /* Adicione seus estilos aqui */
  }
</style>`

export default function HtmlCssChallengePage() {
  const { contributor, isAuthenticated } = useAuthStore()
  const { data: submissions, mutate } = useSubmissions(contributor?.id)

  const [code, setCode] = useState(STARTER_CODE)
  const [report, setReport] = useState<HtmlCssAnalysisReport | null>(null)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const [activeTab, setActiveTab] = useState<'ticket' | 'editor' | 'history'>('ticket')

  const handleChange = useCallback((value: string) => setCode(value), [])

  async function handleSubmit() {
    if (!contributor) return
    setLoading(true)
    setError(null)
    setReport(null)

    try {
      const result = await htmlCssApi.submit({
        contributor_id: contributor.id,
        html_content: code,
      })
      setReport(result)
      mutate()
      setActiveTab('editor')
    } catch (err) {
      setError(err instanceof ApiError ? err.message : 'Erro ao enviar. Tente novamente.')
    } finally {
      setLoading(false)
    }
  }

  if (!isAuthenticated()) {
    return (
      <div className="min-h-[calc(100vh-4rem)] flex items-center justify-center px-4">
        <div className="card bg-base-200 border border-base-300 max-w-md w-full">
          <div className="card-body items-center text-center gap-4">
            <span className="text-4xl">🔒</span>
            <h2 className="text-xl font-bold font-mono">Acesso restrito</h2>
            <p className="text-base-content/60 text-sm">
              Você precisa estar logado para acessar os desafios.
            </p>
            <div className="flex gap-3">
              <Link href="/login" className="btn btn-primary btn-sm font-mono">
                Entrar
              </Link>
              <Link href="/register" className="btn btn-outline btn-sm font-mono">
                Cadastrar
              </Link>
            </div>
          </div>
        </div>
      </div>
    )
  }

  return (
    <div className="max-w-7xl mx-auto px-4 py-8 space-y-6">
      {/* Header */}
      <div className="flex items-start justify-between flex-wrap gap-4">
        <div>
          <div className="flex items-center gap-2 mb-1">
            <span className="badge badge-primary badge-outline font-mono text-xs">
              Desafio #01
            </span>
            <span className="badge badge-ghost font-mono text-xs">HTML/CSS</span>
          </div>
          <h1 className="text-2xl font-bold font-mono">Card de Apresentação</h1>
          <p className="text-base-content/50 text-sm mt-1">
            Crie seu card pessoal seguindo o template especificado
          </p>
        </div>

        <div className="flex items-center gap-2 text-sm text-base-content/50 font-mono">
          <div className="w-2 h-2 rounded-full bg-success animate-pulse" />
          Logado como{' '}
          <span className="text-base-content font-bold">{contributor?.full_name}</span>
        </div>
      </div>

      {/* Tabs */}
      <div role="tablist" className="tabs tabs-border">
        <button
          role="tab"
          className={`tab font-mono ${activeTab === 'ticket' ? 'tab-active' : ''}`}
          onClick={() => setActiveTab('ticket')}
        >
          📋 Ticket
        </button>
        <button
          role="tab"
          className={`tab font-mono ${activeTab === 'editor' ? 'tab-active' : ''}`}
          onClick={() => setActiveTab('editor')}
        >
          💻 Editor
        </button>
        <button
          role="tab"
          className={`tab font-mono ${activeTab === 'history' ? 'tab-active' : ''}`}
          onClick={() => setActiveTab('history')}
        >
          📜 Histórico
          {submissions && submissions.length > 0 && (
            <span className="badge badge-sm badge-primary ml-1">{submissions.length}</span>
          )}
        </button>
      </div>

      {/* Tab: Ticket */}
      {activeTab === 'ticket' && (
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
          <div className="lg:col-span-2 space-y-4">
            <div className="card bg-base-200 border border-base-300">
              <div className="card-body gap-4">
                <div className="flex items-center justify-between flex-wrap gap-2">
                  <h2 className="font-bold font-mono text-lg flex items-center gap-2">
                    🎯 Objetivo
                  </h2>
                  <a
                    href="https://github.com/Winnicius-Moura/tecinfo/blob/main/tickets/galeria-ticket.md"
                    target="_blank"
                    rel="noreferrer"
                    className="btn btn-outline btn-xs font-mono gap-1"
                  >
                    📄 Ler ticket completo
                  </a>
                </div>
                <p className="text-base-content/80 text-sm leading-relaxed">
                  Crie um <strong>card de apresentação pessoal</strong> em HTML/CSS que será exibido
                  na galeria do repositório TecInfo após aprovação via pull request.
                </p>

                <div className="divider my-0" />

                <h2 className="font-bold font-mono text-lg">📐 Estrutura obrigatória</h2>
                <div className="mockup-code text-xs">
                  <pre data-prefix="1"><code>{`<article>`}</code></pre>
                  <pre data-prefix="2"><code>{`  <h3>Seu Nome Completo</h3>`}</code></pre>
                  <pre data-prefix="3"><code>{`  <p>Uma frase sobre você</p>`}</code></pre>
                  <pre data-prefix="4"><code>{`  <h4>Programming languages I use</h4>`}</code></pre>
                  <pre data-prefix="5"><code>{`  <section>`}</code></pre>
                  <pre data-prefix="6"><code>{`    <div>Linguagem 1</div>`}</code></pre>
                  <pre data-prefix="7"><code>{`    <div>Linguagem 2</div>`}</code></pre>
                  <pre data-prefix="8"><code>{`  </section>`}</code></pre>
                  <pre data-prefix="9"><code>{`  <h4 title="social-links-tecinfo">Social Links</h4>`}</code></pre>
                  <pre data-prefix="10"><code>{`  <section>`}</code></pre>
                  <pre data-prefix="11"><code>{`    <a href="..." target="_blank">`}</code></pre>
                  <pre data-prefix="12"><code>{`      <img src="devicon-url" />`}</code></pre>
                  <pre data-prefix="13"><code>{`      LinkedIn`}</code></pre>
                  <pre data-prefix="14"><code>{`    </a>`}</code></pre>
                  <pre data-prefix="15"><code>{`  </section>`}</code></pre>
                  <pre data-prefix="16"><code>{`</article>`}</code></pre>
                </div>

                <div className="divider my-0" />

                <h2 className="font-bold font-mono text-lg">📊 Critérios de pontuação</h2>
                <div className="overflow-x-auto">
                  <table className="table table-sm font-mono text-xs">
                    <thead>
                      <tr>
                        <th>Categoria</th>
                        <th className="text-right">Pontos</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td>Estrutura & Semântica</td>
                        <td className="text-right text-info">38 pts</td>
                      </tr>
                      <tr>
                        <td>Badges de linguagens</td>
                        <td className="text-right text-info">22 pts</td>
                      </tr>
                      <tr>
                        <td>Links sociais</td>
                        <td className="text-right text-info">22 pts</td>
                      </tr>
                      <tr>
                        <td>Fidelidade ao CSS</td>
                        <td className="text-right text-info">18 pts</td>
                      </tr>
                      <tr className="font-bold border-t border-base-300">
                        <td>Total</td>
                        <td className="text-right text-primary">100 pts</td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
            </div>
          </div>

          {/* Sidebar */}
          <div className="space-y-4">
            <div className="card bg-base-200 border border-base-300">
              <div className="card-body gap-3">
                <h3 className="font-bold font-mono">📌 Regras</h3>
                <ul className="text-sm text-base-content/70 space-y-2">
                  <li className="flex gap-2">
                    <span className="text-success mt-0.5">✓</span>
                    Mínimo de <strong>70%</strong> para aprovação
                  </li>
                  <li className="flex gap-2">
                    <span className="text-success mt-0.5">✓</span>
                    Use ícones do{' '}
                    <a
                      href="https://devicon.dev"
                      target="_blank"
                      rel="noreferrer"
                      className="link link-primary"
                    >
                      Devicon
                    </a>
                  </li>
                  <li className="flex gap-2">
                    <span className="text-success mt-0.5">✓</span>
                    Links com <code className="badge badge-ghost badge-sm">target="_blank"</code>
                  </li>
                  <li className="flex gap-2">
                    <span className="text-success mt-0.5">✓</span>
                    Inclua um bloco <code className="badge badge-ghost badge-sm">&lt;style&gt;</code>
                  </li>
                  <li className="flex gap-2">
                    <span className="text-warning mt-0.5">!</span>
                    Use <strong>seus dados reais</strong>
                  </li>
                </ul>
              </div>
            </div>

            <div className="card bg-warning/10 border border-warning/30">
              <div className="card-body gap-2 p-4">
                <h3 className="font-bold font-mono text-warning text-sm">⚠️ Avaliação</h3>
                <p className="text-xs text-base-content/60">
                  Esta submissão conta para sua nota. Envie apenas quando estiver satisfeito com o resultado.
                </p>
              </div>
            </div>

            <button
              onClick={() => setActiveTab('editor')}
              className="btn btn-primary w-full font-mono"
            >
              Ir para o editor →
            </button>
          </div>
        </div>
      )}

      {/* Tab: Editor */}
      {activeTab === 'editor' && (
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
          {/* Editor */}
          <div className="space-y-3">
            <div className="flex items-center justify-between">
              <h2 className="font-bold font-mono text-sm text-base-content/60 uppercase tracking-wide">
                Seu código
              </h2>
              <span className="font-mono text-xs text-base-content/40">
                {code.length} chars
              </span>
            </div>

            <div className="rounded-lg overflow-hidden border border-base-300">
              <CodeMirror
                value={code}
                height="520px"
                extensions={[html()]}
                onChange={handleChange}
                theme="dark"
                basicSetup={{
                  lineNumbers: true,
                  foldGutter: true,
                  highlightActiveLine: true,
                  autocompletion: true,
                }}
              />
            </div>

            {error && (
              <div role="alert" className="alert alert-error text-sm font-mono">
                <span>✗ {error}</span>
              </div>
            )}

            <button
              onClick={handleSubmit}
              disabled={loading}
              className="btn btn-primary w-full font-mono"
            >
              {loading ? (
                <>
                  <span className="loading loading-spinner loading-sm" />
                  Analisando...
                </>
              ) : (
                '🚀 Enviar para análise'
              )}
            </button>
          </div>

          {/* Result */}
          <div className="space-y-3">
            <h2 className="font-bold font-mono text-sm text-base-content/60 uppercase tracking-wide">
              Resultado da análise
            </h2>

            {!report && !loading && (
              <div className="flex flex-col items-center justify-center h-64 border border-dashed border-base-300 rounded-lg text-base-content/30 gap-3">
                <span className="text-5xl">🔍</span>
                <p className="font-mono text-sm">Envie seu código para ver o resultado</p>
              </div>
            )}

            {loading && (
              <div className="flex flex-col items-center justify-center h-64 border border-base-300 rounded-lg gap-3">
                <span className="loading loading-dots loading-lg text-primary" />
                <p className="font-mono text-sm text-base-content/50">Analisando seu código...</p>
              </div>
            )}

            {report && !loading && <AnalysisReport report={report} />}
          </div>
        </div>
      )}

      {/* Tab: History */}
      {activeTab === 'history' && (
        <div className="space-y-4">
          <h2 className="font-bold font-mono">Suas submissões anteriores</h2>

          {!submissions && (
            <div className="flex justify-center py-12">
              <span className="loading loading-dots loading-lg text-primary" />
            </div>
          )}

          {submissions?.length === 0 && (
            <div className="flex flex-col items-center justify-center py-16 text-base-content/30 gap-3">
              <span className="text-5xl">📭</span>
              <p className="font-mono text-sm">Nenhuma submissão ainda</p>
            </div>
          )}

          {submissions && submissions.length > 0 && (
            <div className="space-y-3">
              {submissions.map((s, i) => (
                <div key={s.id} className="card bg-base-200 border border-base-300">
                  <div className="card-body p-4 flex-row items-center justify-between flex-wrap gap-2">
                    <div className="flex items-center gap-3">
                      <span className="font-mono text-base-content/40 text-sm">#{submissions.length - i}</span>
                      <div>
                        <p className="font-mono text-sm">{s.id}</p>
                        <p className="text-xs text-base-content/40 font-mono">
                          {new Date(s.created_at).toLocaleString('pt-BR')}
                        </p>
                      </div>
                    </div>
                    <button
                      onClick={() => {
                        setCode(s.html_content)
                        setActiveTab('editor')
                      }}
                      className="btn btn-ghost btn-xs font-mono"
                    >
                      Carregar código
                    </button>
                  </div>
                </div>
              ))}
            </div>
          )}
        </div>
      )}
    </div>
  )
}
