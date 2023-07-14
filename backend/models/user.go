package models

import "time"

type User struct {
	IDUsuario    int        `json:"id_usuario"`
	Nome         string     `json:"nome"`
	Email        string     `json:"email"`
	Senha        string     `json:"senha"`
	DTNascimento *time.Time `json:"dt_nascimento"`
	Bio          *string    `json:"bio"`
	LinkPerfil   string     `json:"link_perfil"`
	LinkCapa     string     `json:"link_capa"`
	Arroba       *string    `json:"arroba"`
}
