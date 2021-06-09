package api

type Usuario struct {
	Id    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"e-mail"`
}
