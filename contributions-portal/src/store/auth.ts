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

      setContributor: (contributor) => set({ contributor }),

      clearContributor: () => set({ contributor: null }),

      isAuthenticated: () => get().contributor !== null,
    }),
    {
      name: 'tecinfo-auth',
    },
  ),
)
