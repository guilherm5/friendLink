package models

type Pagination struct {
	IDPost         int    `json:"id_post"`
	PostImagem     string `json:"post_imagem"`
	PostTexto      string `json:"post_texto"`
	Nome           string `json:"nome"`
	Arroba         string `json:"arroba"`
	LinkPerfil     string `json:"link_perfil"`
	QtdeComentario int    `json:"qtde_comentario"`
	QtdeCurtida    int    `json:"qtde_curtida"`
}
