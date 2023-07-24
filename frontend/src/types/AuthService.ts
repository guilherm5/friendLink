import type { DefaultResponse } from "./ApiService"

type LoginResponse = Omit<DefaultResponse, 'data'> & {
    data?: {
        logged?: string,
        refresh?: string,
    }
}
type Auth = {
    token: string | undefined,
    refreshToken: string | undefined,
    exp: number,
    Nome: string,
    ID: number,
}
export type {
    LoginResponse,
    Auth
}