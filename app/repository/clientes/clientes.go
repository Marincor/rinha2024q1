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

func (repo *Repository) UpdateClienteSaldoByDebito(clienteID int64, valor int64) (*entity.ClienteSaldo, error) {
	return database.Query("UPDATE clientes SET saldo = saldo - $1 WHERE id = $2 RETURNING id, saldo, limite", new(entity.ClienteSaldo), valor, clienteID)
}

func (repo *Repository) UpdateClienteSaldoByCredito(clienteID int64, valor int64) (*entity.ClienteSaldo, error) {
	return database.Query("UPDATE clientes SET saldo = saldo + $1 WHERE id = $2 RETURNING id, saldo, limite", new(entity.ClienteSaldo), valor, clienteID)
}

func (repo *Repository) GetLastTransacoes(clienteID int64) (*[]entity.UltimasTransacoes, error) {
	output := []entity.UltimasTransacoes{}

	response, err := database.Query("SELECT valor, tipo, descricao, realizada_em FROM transacoes WHERE id_cliente = $1 ORDER BY realizada_em DESC LIMIT 10", &output, clienteID)

	return response, err
}

func (repo *Repository) GetSaldoByID(clienteID int64) (*entity.Saldo, error) {
	return database.Query("SELECT id, saldo as total, limite, CURRENT_TIMESTAMP as data_extrato FROM clientes WHERE id = $1", new(entity.Saldo), clienteID)
}
