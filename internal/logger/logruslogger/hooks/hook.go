package hooks

import (
	extractoriface "pnBot/internal/logger/interfaces"

	"github.com/sirupsen/logrus"
)

// Хук, который вытаскивает поля из context.Context с помощью ContextExtractor
type ContextHook struct {
	Extractor extractoriface.ContextExtractor
}

func (h *ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels // Применяется ко всем уровням логов
}

func (h *ContextHook) Fire(entry *logrus.Entry) error {
	ctx := entry.Context
	if ctx == nil || h.Extractor == nil {
		return nil
	}

	fields := h.Extractor.Extract(ctx)
	for k, v := range fields {
		entry.Data[k] = v
	}
	return nil
}
