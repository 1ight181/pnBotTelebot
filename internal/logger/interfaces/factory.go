package interfaces

type LoggerFactory[ContextOpts any, Logger any] interface {
	NewBaseLogger() *Logger
	NewLoggerWithContext(opts ContextOpts) *Logger
}
