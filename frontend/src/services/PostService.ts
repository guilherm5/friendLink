import type { AxiosError } from 'axios';
import { ApiService } from '@/services/ApiService';

const getPosts = async (token: string | undefined, id_post: number, limit_post: number) => {
    if(!token) return {status: false, error: 'Token não informado'}
    
    return await ApiService.post('/list-post', {id_post, limit_post}, {headers: {Authorization: `${token}`}})
    .then(res => {
        return {status: true, data: res.data}
    }).catch((error: AxiosError) => {
      return {status: false, error: error.response?.data}
    })
}

export {
    getPosts
}