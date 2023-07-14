package models

type Pagination struct {
	IDPost      int    `json:"id_post"`
	PostImagem  string `json:"post_imagem"`
	PostTexto   string `json:"post_texto"`
	IDUsuarioPT int    `json:"id_usuario_Pt"`
	UUIDPost    string `json:"uuid_post"`
	Nome        string `json:"nome"`
	Arroba      string `json:"arroba"`
	LinkPerfil  string `json:"link_perfil"`
}
