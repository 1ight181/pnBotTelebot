package models

type Config struct {
	Bot Bot `mapstructure:"bot"`
}

// Validate проверяет корректность структуры Config, делегируя валидацию полю Bot.
// Возвращает ошибку, если конфигурация Bot некорректна, или nil, если все проверки пройдены.
func (c *Config) Validate() error {
	if err := c.Bot.Validate(); err != nil {
		return err
	}
	return nil
}
