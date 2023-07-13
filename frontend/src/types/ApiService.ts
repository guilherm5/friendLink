import type { AxiosError } from "axios"

type DefaultResponse = {
    status: boolean, 
    error?: AxiosError, 
    data?: any,
}

export type {
    DefaultResponse
}