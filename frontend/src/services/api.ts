import axios from 'axios'
import { API_ENDPOINT } from '../constants/env'

export class APIService {
  private endpoint: string

  constructor(endpoint: string) {
    this.endpoint = endpoint
  }

  async signup(name: string, email: string, password: string): Promise<string> {
    try {
      const res = await axios.post(`${this.endpoint}/api/signup`, {
        name,
        email,
        password,
      })

      return res.data.payload
    } catch (err) {
      throw Error(err as string)
    }
  }

  async signin(email: string, password: string): Promise<string> {
    try {
      const res = await axios.post(`${this.endpoint}/api/signin`, {
        email,
        password,
      })

      return res.data.payload
    } catch (err) {
      throw Error(err as string)
    }
  }
}

export const apiService = new APIService(API_ENDPOINT)
