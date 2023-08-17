package models

import "time"

type Post struct {
	IDPost      int       `json:"id_post"`
	PostTexto   *string   `json:"post_texto"`
	IDUsuarioPT int       `json:"id_usuario_Pt"`
	UUIDPost    string    `json:"uuid_post"`
	PostImagem  string    `json:"post_imagem"`
	DataUpdate  time.Time `json:"dt_update"`
	DataCriacao time.Time `json:"dt_criacao"`
}
