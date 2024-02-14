package clientes

import (
	constantserrors "api.default.marincor.com/app/errors"
	"api.default.marincor.com/app/repository/clientes"
	"api.default.marincor.com/config/constants"
	"api.default.marincor.com/entity"
)

type Validator struct {
	clientesRepo *clientes.Repository
	tiposValidos *constants.TipoTransacao
}

func NewValidator() *Validator {
	return &Validator{
		clientesRepo: clientes.New(),
	}
}

func (v *Validator) ValidateCreateTransaction(clienteID int64, transacao *entity.TransacaoRequest) (int, error) {
	if clienteID == 0 {
		return constants.HTTPStatusBadRequest, constantserrors.ErrClienteIDMissing
	}

	if transacao.Valor == 0 {
		return constants.HTTPStatusBadRequest, constantserrors.ErrValorMissing
	}

	if transacao.Valor < 0 {
		return constants.HTTPStatusUnprocessableEntity, constantserrors.ErrInvalidTransacaoValor
	}

	if transacao.Tipo == "" {
		return constants.HTTPStatusBadRequest, constantserrors.ErrTipoMissing
	}

	if transacao.Descrição == "" {
		return constants.HTTPStatusUnprocessableEntity, constantserrors.ErrDescricaoMissing
	}

	if len(transacao.Descrição) > constants.MaxLengthDescricao {
		return constants.HTTPStatusUnprocessableEntity, constantserrors.ErrDescricaoTooLarge
	}

	if !v.tiposValidos.IsValid(constants.TipoTransacao(transacao.Tipo)) {
		return constants.HTTPStatusUnprocessableEntity, constantserrors.ErrInvalidTransacaoTipo
	}

	return constants.HTTPStatusOK, nil
}

func (v *Validator) ValidateGetBalance(clienteID int64) (int, error) {
	if clienteID == 0 {
		return constants.HTTPStatusNotFound, constantserrors.ErrClienteIDMissing
	}

	return constants.HTTPStatusOK, nil
}
