package clientes

import (
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
	var (
		cliente *entity.ClienteSaldo
		err     error
	)

	if transacao.Tipo == string(constants.Debito) {
		cliente, err = usecase.repo.UpdateClienteSaldoByDebito(clienteID, transacao.Valor)
		if err != nil {
			return nil, constants.HTTPStatusUnprocessableEntity, err
		}
	} else {
		cliente, err = usecase.repo.UpdateClienteSaldoByCredito(clienteID, transacao.Valor)
		if err != nil {
			return nil, constants.HTTPStatusUnprocessableEntity, err
		}
	}

	if cliente.ID == 0 {
		return nil, constants.HTTPStatusNotFound, constantserrors.ErrClienteNotExist
	}

	err = usecase.repo.CreateTransaction(clienteID, transacao)
	if err != nil {
		return nil, constants.HTTPStatusUnprocessableEntity, err
	}

	return cliente, constants.HTTPStatusOK, nil
}

func (usecase *Usecase) GetBalance(clienteID int64) (*entity.Extrato, int, error) {
	var extrato entity.Extrato

	cliente, err := usecase.repo.GetSaldoByID(clienteID)
	if err != nil {
		return nil, constants.HTTPStatusInternalServerError, err
	}

	if cliente.ID == nil || *cliente.ID == 0 {
		return nil, constants.HTTPStatusNotFound, constantserrors.ErrClienteNotExist
	}

	ultimasTransacoes, err := usecase.repo.GetLastTransacoes(clienteID)
	if err != nil {
		return nil, constants.HTTPStatusInternalServerError, err
	}

	extrato.Saldo = *cliente
	extrato.UltimasTransacoes = *ultimasTransacoes

	extrato.Saldo.ID = nil

	return &extrato, constants.HTTPStatusOK, nil
}
