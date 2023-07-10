import type { AxiosError } from 'axios';
import { ApiService } from '@/services/ApiService';
import type { FirstStepInfoUpdate } from '@/types/UserService';

const handleInfoUpdate = async (token: string | undefined, form: FirstStepInfoUpdate) => {
    if(!token) return {status: false, error: 'Token não informado'}
    return await ApiService.put('/info-user', form, {headers: {'Authorization': token}})
    .then(res => {
        return {status: true, data: res.data}
    }).catch((error: AxiosError) => {
      return {status: false, error: error.response?.data}
    })
}

export {
    handleInfoUpdate
}