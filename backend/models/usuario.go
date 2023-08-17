package models

import "time"

type Usuario struct {
	IDUsuario    int        `json:"id_usuario"`
	Nome         string     `json:"nome"`
	Email        string     `json:"email"`
	Senha        string     `json:"senha"`
	DTNascimento *time.Time `json:"dt_nascimento"`
	Bio          *string    `json:"bio"`
	Arroba       *string    `json:"arroba"`
	LinkPerfil   *string    `json:"link_perfil"`
	LinkCapa     *string    `json:"link_capa"`
	DataCriacao  time.Time  `json:"dt_criacao"`
	DataUpdate   time.Time  `json:"dt_update"`
}

type FeedUser struct {
	IDUsuario      int        `json:"id_usuario"`
	Nome           string     `json:"nome"`
	Email          string     `json:"email"`
	Senha          string     `json:"senha"`
	DTNascimento   *time.Time `json:"dt_nascimento"`
	Bio            *string    `json:"bio"`
	Arroba         *string    `json:"arroba"`
	LinkPerfil     string     `json:"link_perfil"`
	LinkCapa       string     `json:"link_capa"`
	DataCriacao    *time.Time `json:"dt_criacao"`
	DataUpdate     *time.Time `json:"dt_update"`
	IDPost         int        `json:"id_post"`
	PostTexto      string     `json:"post_texto"`
	PostImagem     *string    `json:"post_imagem"`
	QtdeComentario int        `json:"qtde_comentario"`
	QtdeCurtida    int        `json:"qtde_curtida"`
	Curtiu         *int       `json:"curtiu"`
}
