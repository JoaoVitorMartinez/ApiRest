package models

type Jogo struct {
	Id        int     `json:"Id"`
	Nome      string  `json:"Nome"`
	Descricao string  `json:"Descricao"`
	Preco     float64 `json:"Preco"`
	Nota      float32 `json:"Nota"`
}

var Jogos []Jogo
