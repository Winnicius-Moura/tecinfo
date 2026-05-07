import type {
  ApiResponse,
  Contributor,
  ContributorLoginPayload,
  ContributorRegistrationPayload,
  HtmlCssAnalysisReport,
  HtmlCssSubmission,
  HtmlCssSubmitPayload,
} from '@/types'

const BASE = process.env.NEXT_PUBLIC_API_URL ?? 'http://localhost:8080/contributions-analysis/api/v1'

export class ApiError extends Error {
  constructor(
    public readonly statusCode: number,
    message: string,
  ) {
    super(message)
    this.name = 'ApiError'
  }
}

async function request<T>(path: string, init?: RequestInit): Promise<T> {
  const res = await fetch(`${BASE}${path}`, {
    headers: { 'Content-Type': 'application/json', ...init?.headers },
    ...init,
  })

  const body: ApiResponse<T> = await res.json()

  if (!body.success) {
    throw new ApiError(res.status, body.message)
  }

  return body.data
}

// ─── Contributor ──────────────────────────────────────────────────────────────

export const contributorApi = {
  register: (payload: ContributorRegistrationPayload) =>
    request<Contributor>('/contributor/register', {
      method: 'POST',
      body: JSON.stringify(payload),
    }),

  login: (payload: ContributorLoginPayload) =>
    request<Contributor>('/contributor/login', {
      method: 'PUT',
      body: JSON.stringify(payload),
    }),
}

// ─── HTML/CSS Challenge ───────────────────────────────────────────────────────

export const htmlCssApi = {
  submit: (payload: HtmlCssSubmitPayload) =>
    request<HtmlCssAnalysisReport>('/test/html-css/submit', {
      method: 'POST',
      body: JSON.stringify(payload),
    }),

  submissions: (contributorId: string) =>
    request<HtmlCssSubmission[]>(
      `/test/html-css/submissions?contributor_id=${contributorId}`,
    ),
}
