package callback

import (
	"errors"
	deps "pnBot/internal/bot/processors/dependencies"
	"strings"

	"gopkg.in/telebot.v3"
)

type CallbackProcessor struct {
	dependencies *deps.ProcessorDependencies
}

func NewCallbackProcessor(dependencies *deps.ProcessorDependencies) *CallbackProcessor {
	return &CallbackProcessor{
		dependencies: dependencies,
	}
}

func (cp *CallbackProcessor) ProcessCallback(c telebot.Context) (err error) {
	defer func() {
		respErr := c.Respond(&telebot.CallbackResponse{})
		if respErr != nil && err == nil {
			err = respErr
		}
	}()

	processingText := cp.dependencies.TextProvider.GetText("processing")
	message, err := c.Bot().Send(c.Chat(), processingText, &telebot.ReplyMarkup{RemoveKeyboard: true})
	if err != nil {
		return err
	}
	defer func() {
		c.Bot().Delete(message)
	}()

	rawData := c.Callback().Data
	data := strings.TrimPrefix(rawData, "\f")

	switch data {
	case "subscribe":
		return cp.ProcessSubscribe(c)
	case "last":
		return cp.ProccesLast(c)
	case "next":
		return cp.ProccesNext(c)
	case "menu":
		return cp.ProccesMenu(c)
	case "unsubscribe":
		return cp.ProcessUnsubscribe(c)
	case "filter_settings":
		return cp.ProcessFilterSettings(c)
	case "frequency_settings":
		return cp.ProcessFrequencySettings(c)
	default:
		if strings.HasPrefix(data, "filter|") {
			return cp.ProcessFilterToggle(c, data)
		}
		if strings.HasPrefix(data, "apply_filter|") {
			return cp.ProcessApplyFilter(c, data)
		}
		return errors.New("Получен неизвестный callback: " + data)
	}

}
