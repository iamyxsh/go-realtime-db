import { create } from 'zustand'
import { apiService } from '../services'

interface UserState {
  name: string
  email: string
  token: string
  signup: (name: string, email: string, password: string) => Promise<void>
  signin: (email: string, password: string) => Promise<void>
}

export const useUserStore = create<UserState>((set) => ({
  name: '',
  email: '',
  token: '',
  signup: async (name: string, email: string, password: string) => {
    try {
      const token = await apiService.signup(name, email, password)

      set({ name, email, token })
    } catch (err) {
      throw Error(err as string)
    }
  },
  signin: async (email: string, password: string) => {
    try {
      const token = await apiService.signin(email, password)
      set({ email, token })
    } catch (err) {
      throw Error(err as string)
    }
  },
}))
