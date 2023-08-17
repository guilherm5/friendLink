package models

import "time"

type ObjectUserStory struct {
	IDStory        int        `json:"story"`
	LinkStory      string     `json:"link_story"`
	IDUsuarioStory int        `json:"id_usuario_story"`
	DTCriacao      *time.Time `json:"dt_criacao"`
	DTExclusao     *time.Time `json:"dt_exclusao"`
	IDUsuario      int        `json:"id_usuario"`
	Nome           string     `json:"nome"`
	LinkPerfil     *string    `json:"link_perfil"`
	Arroba         *string    `json:"arroba"`
	Story          *string    `json:"stories"`
}
