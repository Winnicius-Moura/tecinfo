import type { CheckResult, HtmlCssAnalysisReport } from '@/types'

interface Props {
  report: HtmlCssAnalysisReport
}

export function AnalysisReport({ report }: Props) {
  const pct = Math.round(report.percentage)
  const ringColor =
    report.approved
      ? 'text-success'
      : pct >= 50
        ? 'text-warning'
        : 'text-error'

  return (
    <div className="space-y-6">
      {/* Score card */}
      <div className="card bg-base-200 border border-base-300">
        <div className="card-body items-center text-center gap-4">
          <div
            className={`radial-progress ${ringColor} font-mono text-2xl font-bold`}
            style={
              {
                '--value': pct,
                '--size': '7rem',
                '--thickness': '6px',
              } as React.CSSProperties
            }
            role="progressbar"
            aria-valuenow={pct}
            aria-valuemin={0}
            aria-valuemax={100}
          >
            {pct}%
          </div>

          <div>
            <p className="text-base-content/60 font-mono text-sm">
              {report.score} / {report.max_score} pontos
            </p>
            <div className={`badge mt-2 font-mono ${report.approved ? 'badge-success' : 'badge-error'}`}>
              {report.approved ? '✓ Aprovado' : '✗ Reprovado'}
            </div>
          </div>

          {!report.approved && (
            <p className="text-sm text-base-content/50">
              Mínimo para aprovação: <span className="text-warning font-bold">70%</span>
            </p>
          )}
        </div>
      </div>

      {/* Passed checks */}
      {report.passed_checks?.length > 0 && (
        <CheckSection title="Verificações aprovadas" items={report.passed_checks} variant="success" />
      )}

      {/* Failed checks */}
      {report.failed_checks?.length > 0 && (
        <CheckSection title="Verificações reprovadas" items={report.failed_checks} variant="error" />
      )}
    </div>
  )
}

function CheckSection({
  title,
  items,
  variant,
}: {
  title: string
  items: CheckResult[]
  variant: 'success' | 'error'
}) {
  const icon = variant === 'success' ? '✓' : '✗'
  const badgeClass = variant === 'success' ? 'badge-success' : 'badge-error'
  const textClass = variant === 'success' ? 'text-success' : 'text-error'

  return (
    <div>
      <h3 className={`font-mono font-semibold mb-2 ${textClass}`}>
        {icon} {title} ({items.length})
      </h3>
      <div className="space-y-2">
        {items.map((item) => (
          <div key={item.rule} className="collapse collapse-arrow bg-base-200 border border-base-300">
            <input type="checkbox" />
            <div className="collapse-title text-sm font-mono flex items-center gap-2">
              <span className={`badge badge-sm ${badgeClass}`}>
                {item.points}/{item.max_points}
              </span>
              {item.rule}
            </div>
            {(item.expected || item.actual || item.diff) && (
              <div className="collapse-content font-mono text-xs space-y-1">
                {item.expected && (
                  <p>
                    <span className="text-base-content/40">esperado:</span>{' '}
                    <span className="text-info">{item.expected}</span>
                  </p>
                )}
                {item.actual && (
                  <p>
                    <span className="text-base-content/40">recebido:</span>{' '}
                    <span className="text-warning">{item.actual}</span>
                  </p>
                )}
                {item.diff && (
                  <pre className="bg-base-300 rounded p-2 whitespace-pre-wrap">{item.diff}</pre>
                )}
              </div>
            )}
          </div>
        ))}
      </div>
    </div>
  )
}
