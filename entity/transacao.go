package entity

type TransacaoRequest struct {
	Valor     int64  `json:"valor"`
	Tipo      string `json:"tipo"`
	Descrição string `json:"descricao"`
}
