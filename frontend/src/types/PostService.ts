type CommentResponse = {
    id_comentario: number,
    id_post: number,
    id_usuario_cmt: number,
    comentario: string,
    dt_comentario: string,
    nome: string,
    link_perfil: string,
    arroba: string,
}
type PostResponse = {
    id_post: number, 
    uuid_post: string,
    post_texto?: string, 
    post_imagem?: string,
    nome: string,
    arroba: string,
    link_perfil: string,
    qtde_comentario: number,
    qtde_curtida: number,
    curtiu: boolean,
}
type Post = PostResponse & {
    comentarios?: CommentResponse[],
    loadingComentarios?: boolean,
    carregadoUmaVez?: boolean,
}
export type {
    PostResponse,
    CommentResponse,
    Post
}