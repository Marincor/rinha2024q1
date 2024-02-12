package clientes

import (
	"math"

	constantserrors "api.default.marincor.com/app/errors"
	"api.default.marincor.com/app/repository/clientes"
	"api.default.marincor.com/config/constants"
	"api.default.marincor.com/entity"
)

type Usecase struct {
	repo *clientes.Repository
}

func New() *Usecase {
	return &Usecase{
		repo: clientes.New(),
	}
}

func (usecase *Usecase) CreateTransaction(clienteID int64, transacao entity.TransacaoRequest) (*entity.ClienteSaldo, int, error) {
	var saldoAfterTransaction int64

	cliente, err := usecase.repo.GetClienteByID(clienteID)
	if err != nil {
		return nil, constants.HTTPStatusInternalServerError, err
	}

	if cliente.ID == 0 {
		return nil, constants.HTTPStatusNotFound, constantserrors.ErrClienteNotExist
	}

	if transacao.Tipo == string(constants.Debito) {
		saldoToLimit := int64(math.Abs(float64(cliente.Saldo - transacao.Valor)))
		if saldoToLimit > cliente.Limite {
			return nil, constants.HTTPStatusUnprocessableEntity, constantserrors.ErrInvalidTransacaoValor
		}

		saldoAfterTransaction = cliente.Saldo - transacao.Valor
	} else {
		saldoAfterTransaction = cliente.Saldo + transacao.Valor
	}

	err = usecase.repo.CreateTransaction(clienteID, transacao)
	if err != nil {
		return nil, constants.HTTPStatusInternalServerError, err
	}

	clienteSaldo, err := usecase.repo.UpdateClienteSaldo(clienteID, saldoAfterTransaction)
	if err != nil {
		return nil, constants.HTTPStatusInternalServerError, err
	}

	return clienteSaldo, constants.HTTPStatusOK, nil
}
