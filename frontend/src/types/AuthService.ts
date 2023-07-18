import type { DefaultResponse } from "./ApiService"

type LoginResponse = DefaultResponse & {
    data?: {
        logged?: string,
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