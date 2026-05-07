import Link from 'next/link'

export default function HomePage() {
  return (
    <div className="min-h-[calc(100vh-4rem)] flex flex-col">
      {/* Hero */}
      <section className="flex-1 flex flex-col items-center justify-center text-center px-4 py-20 gap-6">
        <div className="badge badge-outline badge-primary font-mono">TecInfo 2025–2026</div>

        <h1 className="text-4xl sm:text-5xl font-bold font-mono leading-tight max-w-2xl">
          Aprenda open source{' '}
          <span className="text-primary">na prática</span>
        </h1>

        <p className="text-base-content/60 max-w-lg text-lg">
          Faça o fork, contribua com código real, envie pull requests e receba
          feedback!
        </p>

        <div className="flex flex-wrap gap-3 justify-center">
          <Link href="/register" className="btn btn-primary font-mono">
            Começar agora
          </Link>
          <Link href="/challenge/html-css" className="btn btn-outline font-mono">
            Ver desafio
          </Link>
        </div>
      </section>

      {/* Features */}
      <section className="border-t border-base-300 bg-base-200">
        <div className="max-w-5xl mx-auto px-4 py-16 grid grid-cols-1 sm:grid-cols-3 gap-6">
          <FeatureCard
            icon="⚡"
            title="Análise automática"
            description="Seu código é analisado em tempo real por um motor similar ao HackerRank."
          />
          <FeatureCard
            icon="🔀"
            title="Fluxo real de open source"
            description="Fork → branch → commit → pull request. Exatamente como nos grandes projetos."
          />
          <FeatureCard
            icon="🏆"
            title="Nota por contribuição"
            description="Suas submissões são avaliadas e pontuadas. Cada PR aprovado conta."
          />
        </div>
      </section>

      {/* CTA */}
      <section className="border-t border-base-300 py-16 text-center px-4">
        <h2 className="text-2xl font-bold font-mono mb-4">
          Pronto para o primeiro desafio?
        </h2>
        <p className="text-base-content/60 mb-6">
          Cadastre-se, leia o ticket e envie seu HTML/CSS para análise.
        </p>
        <Link href="/register" className="btn btn-primary font-mono">
          Criar conta
        </Link>
      </section>
    </div>
  )
}

function FeatureCard({
  icon,
  title,
  description,
}: {
  icon: string
  title: string
  description: string
}) {
  return (
    <div className="card bg-base-100 border border-base-300 p-6 gap-3">
      <span className="text-3xl">{icon}</span>
      <h3 className="font-bold font-mono">{title}</h3>
      <p className="text-base-content/60 text-sm">{description}</p>
    </div>
  )
}

