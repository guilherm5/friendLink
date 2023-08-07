import type { AxiosError } from 'axios';
import { ApiService } from '@/services/ApiService';

const createStory = async (token: string | undefined, image: Blob) => {
    if(!token) return {status: false, error: 'Token não informado'}
    
    return await ApiService.post('/story', {story: image}, {headers: {Authorization: `${token}`}})
    .then(res => {
        return {status: true, data: res.data}
    }).catch((error: AxiosError) => {
        return {status: false, error: error.response?.data}
    })
}

export {
    createStory,
}