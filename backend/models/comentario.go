package models

import "time"

type Comentario struct {
	IDComentario int       `json:"id_comentario"`
	Comentario   string    `json:"comentario"`
	IDPostCMT    int       `json:"id_post_cmt"`
	IDUsuarioCmt int       `json:"id_usuario_cmt"`
	DTComentario time.Time `json:"dt_comentario"`
	Nome         string    `json:"nome"`
	LinkPerfil   *string   `json:"link_perfil"`
	IDPost       int       `json:"id_post"`
}
