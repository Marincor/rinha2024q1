package constants

import "errors"

var (
	ErrInvalidUserAgent       = errors.New("invalid user agent")
	ErrMissingUserAgent       = errors.New("user agent is missing")
	ErrDatabaseNotConnected   = errors.New("database is not connected")
	ErrAssertDBResponse       = errors.New("error while asserting database response")
	ErrErrorTooLargeEOFBuffer = errors.New("error when reading request headers: EOF. Buffer size")
)

var (
	ErrClienteIDMissing = errors.New("cliente id is missing")
	ErrValorMissing     = errors.New("valor is missing")
	ErrTipoMissing      = errors.New("tipo is missing")
	ErrDescricaoMissing = errors.New("descricao is missing")
)

var (
	ErrInvalidTransacaoTipo  = errors.New("transacao tipo invalid")
	ErrDescricaoTooLarge     = errors.New("descricao too large")
	ErrInvalidTransacaoValor = errors.New("transacao valor invalid")
)

var (
	ErrClienteNotExist = errors.New("cliente not exist")
)
