package models

import "time"

type User struct {
	IDUsuario    int        `json:"id_usuario"`
	Nome         string     `json:"nome"`
	Email        string     `json:"email"`
	Senha        string     `json:"senha"`
	DTNascimento *time.Time `json:"dt_nascimento"`
	Bio          *string    `json:"bio"`
}

type UserLogin struct {
	IDUsuario int    `json:"id_usuario"`
	Email     string `json:"email"`
	Senha     string `json:"senha"`
}
