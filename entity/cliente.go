package entity

type ClienteSaldo struct {
	ID     int64 `json:"id,omitempty"`
	Saldo  int64 `json:"saldo"`
	Limite int64 `json:"limite"`
}
