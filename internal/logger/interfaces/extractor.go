package interfaces

import "context"

// ContextExtractor извлекает поля из контекста для логирования.
type ContextExtractor interface {
	Extract(ctx context.Context) map[string]interface{}
}
