import type { AxiosError } from 'axios';
import { ApiService } from '@/services/ApiService';

const createPost = async (token: string | undefined, form: {post_texto: string | null, file: File | null}) => {
    if(!token) return {status: false, error: 'Token não informado'}
    
    return await ApiService.post('/post', {post_texto: form.post_texto ?? '', file: form.file ?? undefined}, {headers: {Authorization: `${token}`}})
    .then(res => {
        return {status: true, data: res.data}
    }).catch((error: AxiosError) => {
        return {status: false, error: error.response?.data}
    })
}
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
const getComments = async (token: string | undefined, id_post: number, id_comentario: number, limit: number) => {
    if(!token) return {status: false, error: 'Token não informado'}
    
    return await ApiService.post('/list-comentario', {id_post, id_comentario, limit}, {headers: {Authorization: `${token}`}})
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
const getReplies = async (token: string | undefined, id_comentario: number, id_resp_comentario: number, limit: number) => {
    if(!token) return {status: false, error: 'Token não informado'}
    
    return await ApiService.post('/list-resp-comentario', {id_comentario, id_resp_comentario, limit}, {headers: {Authorization: `${token}`}})
    .then(res => {
        return {status: true, data: res.data}
    }).catch((error: AxiosError) => {
      return {status: false, error: error.response?.data}
    })
}
const createReply = async (token: string | undefined, id_comentario: number, resposta: string) => {
    if(!token) return {status: false, error: 'Token não informado'}
    
    return await ApiService.post('/resp-comentario', {id_comentario, resposta}, {headers: {Authorization: `${token}`}})
    .then(res => {
        return {status: true, data: res.data}
    }).catch((error: AxiosError) => {
      return {status: false, error: error.response?.data}
    })
}
const likeComment = async (token: string | undefined, id_comentario: number) => {
    if(!token) return {status: false, error: 'Token não informado'}
    
    return await ApiService.post('/curte-comentario', {id_comentario}, {headers: {Authorization: `${token}`}})
    .then(res => {
        return {status: true, data: res.data}
    }).catch((error: AxiosError) => {
      return {status: false, error: error.response?.data}
    })
}
const unlikeComment = async (token: string | undefined, id_comentario: number) => {
    if(!token) return {status: false, error: 'Token não informado'}
    
    return await ApiService.delete('/del-curtida-comentario', {data: {id_comentario}, headers: {Authorization: `${token}`}})
    .then(res => {
        return {status: true, data: res.data}
    }).catch((error: AxiosError) => {
      return {status: false, error: error.response?.data}
    })
}
const likeReply = async (token: string | undefined, id_resp_comentario: number) => {
    if(!token) return {status: false, error: 'Token não informado'}
    
    return await ApiService.post('/curte-resp-comentario', {id_resp_comentario}, {headers: {Authorization: `${token}`}})
    .then(res => {
        return {status: true, data: res.data}
    }).catch((error: AxiosError) => {
      return {status: false, error: error.response?.data}
    })
}
const unlikeReply = async (token: string | undefined, id_resp_comentario: number) => {
    if(!token) return {status: false, error: 'Token não informado'}
    
    return await ApiService.delete('/del-curtida-resp-comentario', {data: {id_resp_comentario}, headers: {Authorization: `${token}`}})
    .then(res => {
        return {status: true, data: res.data}
    }).catch((error: AxiosError) => {
      return {status: false, error: error.response?.data}
    })
}
const deletePost = async (token: string | undefined, id_post: number) => {
    if(!token) return {status: false, error: 'Token não informado'}
    
    return await ApiService.delete('/delete-post', {data: {id_post}, headers: {Authorization: `${token}`}})
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
    createComment,
    likeComment,
    unlikeComment,
    likeReply,
    unlikeReply,
    createReply,
    getReplies,
    createPost,
    deletePost
}