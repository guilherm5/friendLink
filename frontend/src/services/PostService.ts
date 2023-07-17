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
const likePost = async (token: string | undefined, id_post: number) => {
    if(!token) return {status: false, error: 'Token não informado'}
    
    return await ApiService.post('/curtida', {id_post}, {headers: {Authorization: `${token}`}})
    .then(res => {
        return {status: true, data: res.data}
    }).catch((error: AxiosError) => {
      return {status: false, error: error.response?.data}
    })
}
const unlikePost = async (token: string | undefined, id_post: number) => {
    if(!token) return {status: false, error: 'Token não informado'}
    
    return await ApiService.delete('/curtida', {data: {id_post}, headers: {Authorization: `${token}`}})
    .then(res => {
        return {status: true, data: res.data}
    }).catch((error: AxiosError) => {
      return {status: false, error: error.response?.data}
    })
}
const getComments = async (token: string | undefined, id_post: number) => {
    if(!token) return {status: false, error: 'Token não informado'}
    
    return await ApiService.post('/list-comentario', {id_post}, {headers: {Authorization: `${token}`}})
    .then(res => {
        return {status: true, data: res.data}
    }).catch((error: AxiosError) => {
      return {status: false, error: error.response?.data}
    })
}
const createComment = async (token: string | undefined, id_post: number, comentario: string) => {
    if(!token) return {status: false, error: 'Token não informado'}
    
    return await ApiService.post('/comentario', {id_post, comentario}, {headers: {Authorization: `${token}`}})
    .then(res => {
        return {status: true, data: res.data}
    }).catch((error: AxiosError) => {
      return {status: false, error: error.response?.data}
    })
}

export {
    getPosts,
    likePost,
    unlikePost,
    getComments,
    createComment
}