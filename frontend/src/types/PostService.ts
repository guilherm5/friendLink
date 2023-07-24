type CommentResponse = {
    id_comentario: number,
    id_post: number,
    id_usuario_cmt: number,
    comentario: string,
    dt_comentario: string,
    nome: string,
    link_perfil: string,
    arroba: string,
    qtde_resposta_comentario: number,
    qtde_curtida_comentario: number,
    curtiu_comentario: boolean,
}
type ReplyResponse = {
    id_resp_comentario: number,
    id_comentario_rp: number,
    id_usuario_rp: number,
    resposta: string,
    dt_resp_comentario: string,
    nome: string,
    link_perfil: string,
    arroba: string,
    qtde_curtida_resp_comentario: number,
    curtiu_resp_comentario: boolean,
}
type Comment = CommentResponse & {
    lastReplyId?: number,
    respostas?: ReplyResponse[],
    loadingRespostas?: boolean,
    carregadoUmaVez?: boolean,
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
    comentarios?: Comment[],
    loadingComentarios?: boolean,
    carregadoUmaVez?: boolean,
}
export type {
    PostResponse,
    CommentResponse,
    Post,
    Comment,
    ReplyResponse
}