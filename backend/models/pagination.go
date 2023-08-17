package models

import "time"

type Pagination struct {
	IDPost         int        `json:"id_post"`
	IDUser         int        `json:"id_user"`
	PostImagem     *string    `json:"post_imagem"`
	PostTexto      *string    `json:"post_texto"`
	Nome           string     `json:"nome"`        //para puxar nome do usuario
	Arroba         string     `json:"arroba"`      //para puxar nome do arroba
	LinkPerfil     string     `json:"link_perfil"` //para puxar perfil do usuario
	Curtiu         *int       `json:"curtiu"`
	Seguindo       *int       `json:"seguindo"`
	QtdeComentario int        `json:"qtde_comentario"`
	QtdeCurtida    int        `json:"qtde_curtida"`
	DTCriacao      *time.Time `json:"dt_criacao"`
}
