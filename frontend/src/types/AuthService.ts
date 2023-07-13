import type { DefaultResponse } from "./ApiService"

type LoginResponse = DefaultResponse & {
    data?: {
        logged?: string,
    }
}
type Auth = {
    token: string | undefined,
    exp: number,
    Nome: string,
}
export type {
    LoginResponse,
    Auth
}