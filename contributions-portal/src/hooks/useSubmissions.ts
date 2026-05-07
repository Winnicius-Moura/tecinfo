import type { LocalSubmission } from '@/types'
import useSWR from 'swr'

const fetcher = (url: string) => fetch(url).then((r) => r.json()) as Promise<LocalSubmission[]>

export function useSubmissions(contributorId: string | undefined) {
  return useSWR(
    contributorId ? `/api/submissions?contributor_id=${contributorId}` : null,
    fetcher,
    { revalidateOnFocus: false },
  )
}
