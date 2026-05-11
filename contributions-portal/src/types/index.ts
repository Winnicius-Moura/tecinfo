// ─── Contributor ─────────────────────────────────────────────────────────────

export interface Contributor {
  id: string
  counter: number
  full_name: string
  email: string
  token?: string
  created_at: string
  updated_at: string
}

export interface ContributorRegistrationPayload {
  full_name: string
  email: string
}

export interface ContributorLoginPayload {
  email: string
  password: string
}

// ─── HTML/CSS Challenge ───────────────────────────────────────────────────────

export interface CheckResult {
  rule: string
  passed: boolean
  points: number
  max_points: number
  expected?: string
  actual?: string
  diff?: string
}

export interface HtmlCssAnalysisReport {
  score: number
  max_score: number
  percentage: number
  approved: boolean
  passed_checks: CheckResult[]
  failed_checks: CheckResult[]
}

export interface HtmlCssSubmission {
  id: string
  contributor_id: string
  analysis_result_id: string
  html_content: string
  created_at: string
  updated_at: string
}

export interface HtmlCssSubmitPayload {
  contributor_id: string
  html_content: string
}

// ─── Local Submission (Prisma DB) ────────────────────────────────────────────

export interface LocalSubmission {
  id: string
  contributorId: string
  analysisResultId: string | null
  htmlContent: string
  percentage: number | null
  approved: boolean | null
  reportJson: string | null
  createdAt: string
}



export interface ApiResponse<T> {
  status: number
  success: boolean
  message: string
  data: T
}
