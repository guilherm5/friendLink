import type { AxiosError } from "axios"

type DefaultResponse = {
    status: boolean, 
    error?: AxiosError, 
    data?: any,
}
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
    DefaultResponse,
    LoginResponse,
    Auth
}