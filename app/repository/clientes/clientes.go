package clientes

import (
	"api.default.marincor.com/adapters/database"
	"api.default.marincor.com/entity"
)

type Repository struct{}

func New() *Repository {
	return &Repository{}
}

func (repo *Repository) CreateTransaction(clienteID int64, transacao entity.TransacaoRequest) error {
	return database.Exec("INSERT INTO transacoes (id_cliente, valor, tipo, descricao) VALUES ($1, $2, $3, $4)",
		clienteID, transacao.Valor, transacao.Tipo, transacao.Descrição)
}

func (repo *Repository) GetClienteByID(clienteID int64) (*entity.ClienteSaldo, error) {
	return database.Query("SELECT id, saldo, limite FROM clientes WHERE id = $1", new(entity.ClienteSaldo), clienteID)
}

func (repo *Repository) UpdateClienteSaldo(clienteID int64, saldo int64) (*entity.ClienteSaldo, error) {
	return database.Query("UPDATE clientes SET saldo = $1 WHERE id = $2 RETURNING saldo, limite", new(entity.ClienteSaldo), saldo, clienteID)
}
