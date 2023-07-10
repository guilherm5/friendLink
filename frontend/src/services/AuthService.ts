import type { AxiosError } from 'axios';
import { ApiService } from '@/services/ApiService';

const handleSignup = async (form: { nome: string, email: string, senha: string}) => {
    return await ApiService.post('/new-user', form)
    .then(res => {
        return {status: true, data: res.data}
    }).catch((error: AxiosError) => {
      return {status: false, error: error}
    })
}

const handleSignin = async (form: { email: string, senha: string}) => {
    return await ApiService.post('/login', form)
    .then(res => {
        return {status: true, data: res.data}
    }).catch((error: AxiosError) => {
      return {status: false, error: error}
    })
}

export {
    handleSignup,
    handleSignin
}