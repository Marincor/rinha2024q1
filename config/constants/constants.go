package constants

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	Port                  = "9090"
	LogDataKey            = "payload"
	LogSeverityKey        = "log_severity"
	MaxResquestLimit      = 2
	DefaultContextTimeout = 100 * time.Millisecond
)

const (
	Local      = "local"
	Staging    = "staging"
	Production = "production"
	Test       = "test"
)

var (
	Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	Prefork  = strings.ToLower(os.Getenv("PREFORK")) != "false"
)

var (
	AllowedContentTypes = []string{fiber.MIMEApplicationJSON}
	AllowedHeaders      = "X-Session-Id, Authorization, Content-Type, Accept, Origin"
	AllowedMethods      = "GET,POST,OPTIONS"
)

const (
	TemplatesFolder = "templates"
)

type LoggingSeverity string

const (
	SeverityDebug     LoggingSeverity = "debug"
	SeverityInfo      LoggingSeverity = "info"
	SeverityWarning   LoggingSeverity = "warning"
	SeverityError     LoggingSeverity = "error"
	SeverityCritical  LoggingSeverity = "critical"
	SeverityEmergency LoggingSeverity = "emergency"
)

const (
	MaxLengthDescricao = 10
)

type TipoTransacao string

const (
	Credito TipoTransacao = "c"
	Debito  TipoTransacao = "d"
)

func (*TipoTransacao) IsValid(tipo TipoTransacao) bool {
	return tipo == Credito || tipo == Debito
}
