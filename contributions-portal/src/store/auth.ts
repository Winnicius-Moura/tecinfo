'use client'

import type { Contributor } from '@/types'
import { create } from 'zustand'
import { persist } from 'zustand/middleware'

interface AuthState {
  contributor: Contributor | null
  setContributor: (contributor: Contributor) => void
  clearContributor: () => void
  isAuthenticated: () => boolean
}

export const useAuthStore = create<AuthState>()(
  persist(
    (set, get) => ({
      contributor: null,

      setContributor: (contributor) => {
        if (contributor.token && typeof window !== 'undefined') {
          localStorage.setItem('token', contributor.token)
        }
        set({ contributor })
      },

      clearContributor: () => {
        if (typeof window !== 'undefined') {
          localStorage.removeItem('token')
        }
        set({ contributor: null })
      },

      isAuthenticated: () => get().contributor !== null,
    }),
    {
      name: 'tecinfo-auth',
    },
  ),
)
