'use client'

import Link from 'next/link'
import { useGalleryWebSocket } from '@/hooks/useGallery'

export default function GalleryPage() {
  const { cards, isConnected } = useGalleryWebSocket('ws://localhost:8081/ws')

  return (
    <div className="min-h-screen bg-base-100">
      <div className="max-w-7xl mx-auto px-4 py-8 space-y-8">
        <div className="flex flex-col md:flex-row justify-between items-center gap-4">
          <div>
            <h1 className="text-3xl font-bold font-mono text-primary flex items-center gap-3">
              <span>🖼️</span> Galeria TecInfo 2026 Fundamentos de Criação de Site
            </h1>
            <p className="text-base-content/60 mt-2 font-mono text-sm">
              Acompanhe em tempo real os testes de HTML/CSS aprovados na nossa plataforma.
            </p>
          </div>

          <div className="flex items-center gap-4">
            <div className="flex items-center gap-2 font-mono text-xs">
              <span className="relative flex h-3 w-3">
                {isConnected && (
                  <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-success opacity-75"></span>
                )}
                <span
                  className={`relative inline-flex rounded-full h-3 w-3 ${
                    isConnected ? 'bg-success' : 'bg-error'
                  }`}
                ></span>
              </span>
              <span className={isConnected ? 'text-success' : 'text-error'}>
                {isConnected ? 'Ao Vivo' : 'Desconectado'}
              </span>
            </div>
            
            <Link href="/" className="btn btn-outline btn-sm font-mono">
              Voltar ao Início
            </Link>
          </div>
        </div>

        <div className="divider"></div>

        {cards.length === 0 && (
          <div className="flex flex-col items-center justify-center py-20 text-base-content/30 gap-4">
            {isConnected ? (
              <>
                <span className="loading loading-ring loading-lg text-primary"></span>
                <p className="font-mono text-sm">Aguardando os primeiros envios aprovados...</p>
              </>
            ) : (
              <>
                <span className="text-5xl">🔌</span>
                <p className="font-mono text-sm">Tentando conectar ao servidor...</p>
              </>
            )}
          </div>
        )}

        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
          {cards.map((card, i) => {
            // Inject standard body style to override user's body
            const injectedCss = `
              <style>
                html {
                  height: 100% !important;
                  width: 100% !important;
                  background: transparent !important;
                }
                body {
                  margin: 0 !important;
                  display: flex !important;
                  justify-content: center !important;
                  align-items: center !important;
                  height: 100% !important;
                  width: 100% !important;
                  box-sizing: border-box !important;
                  overflow: hidden !important; /* disable scroll */
                  border-radius: 1rem !important; /* Bordas arredondadas no card do aluno */
                }
                
                /* Força o nowrap para melhorar a aparência fiel ao card-preview */
                h1, h2, h3, h4, h5, h6, p, span, div.badge, a {
                  white-space: nowrap !important;
                }

                /* Optional scrollbar hiding for clean look */
                ::-webkit-scrollbar { display: none; }
              </style>
            `
            const finalHtml = `${card.html_content}${injectedCss}`

            const dateObj = new Date(card.approved_at)
            const formattedDate = dateObj.toLocaleDateString('pt-BR', {
              day: '2-digit',
              month: 'long',
              year: 'numeric'
            })

            return (
              <div
                key={`${card.contributor_id}-${i}`}
                className="card bg-base-100 shadow-md hover:shadow-lg transition-all duration-300 w-full flex flex-col p-1 gap-2"
              >
                {/* Iframe wrapper for isolation with padding 1 as requested */}
                <div className="w-full relative flex items-center justify-center aspect-square ">
                  <iframe
                    srcDoc={finalHtml}
                    className="w-full rounded h-full border-0 bg-transparent"
                    sandbox="allow-popups allow-popups-to-escape-sandbox allow-scripts allow-top-navigation-by-user-activation"
                    title={`Submissão de ${card.contributor_id}`}
                  />
                </div>
                {/* Card Footer with Date */}
                <div className="px-2 pb-1 text-xs font-mono text-base-content/60 text-center">
                  {formattedDate}
                </div>
              </div>
            )
          })}
        </div>
      </div>
    </div>
  )
}
