package entity

import "time"

type Saldo struct {
	Total       int64     `json:"total"`
	DataExtrato time.Time `json:"data_extrato"`
	Limite      int64     `json:"limite"`
}

type UltimasTransacoes struct {
	Valor       int64     `json:"valor"`
	Tipo        string    `json:"tipo"`
	Descricao   string    `json:"descricao"`
	RealizadaEm time.Time `json:"realizada_em"`
}

type Extrato struct {
	Saldo             Saldo               `json:"saldo"`
	UltimasTransacoes []UltimasTransacoes `json:"ultimas_transacoes"`
}
