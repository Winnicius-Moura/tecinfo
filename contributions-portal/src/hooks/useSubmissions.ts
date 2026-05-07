import { htmlCssApi } from '@/lib/api'
import useSWR from 'swr'

export function useSubmissions(contributorId: string | undefined) {
  return useSWR(
    contributorId ? ['submissions', contributorId] : null,
    ([, id]) => htmlCssApi.submissions(id),
    { revalidateOnFocus: false },
  )
}
