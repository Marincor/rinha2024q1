package constants

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	Port                      = "9090"
	LogDataKey                = "payload"
	LogSeverityKey            = "log_severity"
	MainLoggerName            = "health"
	MainServiceName           = MainLoggerName + "_api"
	MaxResquestLimit          = 2
	AccessTokenExpirationTime = 15
	SignedURLExp              = 60
	Audience                  = "https://iam.services.marincor.com"
	DefaultLimit              = 10
	DefaultOffset             = 0
	DefaultOTPGenerator       = "iam"
	DefaultContextTimeout     = 100 * time.Millisecond
)

const (
	Local      = "local"
	Staging    = "staging"
	Production = "production"
	Test       = "test"
)

var (
	Debug, _     = strconv.ParseBool(os.Getenv("DEBUG"))
	GcpProjectID = os.Getenv("PROJECT")
	SecretPrefix = os.Getenv("SEC_PREFIX")
	Environment  = os.Getenv("ENVIRONMENT")
	Prefork      = strings.ToLower(os.Getenv("PREFORK")) != "false"
)

var (
	AllowedContentTypes   = []string{fiber.MIMEApplicationJSON}
	AllowedHeaders        = "X-Session-Id, Authorization, Content-Type, Accept, Origin"
	AllowedMethods        = "GET,POST,OPTIONS"
	AllowedOrigins        = "https://tbd, https://tbd"
	AllowedStageOrigins   = "https://localhost:3000, http://localhost:3000"
	AllowedUnthrottledIPs = []string{"127.0.0.1"}
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
