package models

import "time"

type Comentario struct {
	IDComentario              int        `json:"id_comentario"`
	Comentario                string     `json:"comentario"`
	IDPostCMT                 int        `json:"id_post_cmt"`
	IDUsuarioCmt              int        `json:"id_usuario_cmt"`
	DTComentario              *time.Time `json:"dt_comentario"`
	Nome                      string     `json:"nome"`        //para puxar nome do usuario
	Arroba                    string     `json:"arroba"`      //para puxar arroba do usuario
	LinkPerfil                *string    `json:"link_perfil"` //para puxar link do usuario
	IDPost                    int        `json:"id_post"`     //para puxar id do post
	QtdeRespostaComentario    int        `json:"qtde_resposta_comentario"`
	QtdeCurtidaComentario     int        `json:"qtde_curtida_comentario"`
	CurtiuComentario          *int       `json:"curtiu_comentario"`
	QtdeCurtidaRespComentario *int       `json:"qtde_curtida_resp_comentario"`
	CurtiuRespComentario      *int       `json:"curtiu_resp_comentario"`
}

type RespComment struct {
	IDComentario                int     `json:"id_comentario"`
	IDResposta                  int     `json:"id_resp_comentario"`
	Resposta                    string  `json:"resposta"`
	DTRespCometario             string  `json:"dt_resp_comentario"`
	IDComentarioRP              int     `json:"id_comentario_rp"`
	IDUsuarioRP                 int     `json:"id_usuario_rp"`
	IDPostRP                    int     `json:"id_post_rp"`
	Nome                        string  `json:"nome"`        //para puxar nome do usuario
	Arroba                      string  `json:"arroba"`      //para puxar arroba do usuario
	LinkPerfil                  *string `json:"link_perfil"` //para puxar link do usuario
	IDPost                      int     `json:"id_post"`     //para puxar id do post
	QtdeCurtidadeRespComentario *int    `json:"qtde_curtida_resp_comentario"`
	CurtiuRespComentario        *int    `json:"curtiu_resp_comentario"`
}
