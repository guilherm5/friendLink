type FirstStepInfoUpdate = {
    foto_perfil: null | File, 
    foto_capa: null | File, 
    bio: string, 
    arroba: string
}

type UserResponse = {
    id_usuario: number,
    nome: string,
    email: string,
    arroba: string,
    link_perfil: string,
    link_capa: string,
    dt_criacao: string,
}

export type {
    FirstStepInfoUpdate,
    UserResponse
}