package models

type Config struct {
	Bot           Bot           `mapstructure:"bot"`
	DataBase      DataBase      `mapstructure:"db"`
	AdminPanel    AdminPanel    `mapstructure:"admin_panel"`
	ImageUploader ImageUploader `mapstructure:"image_uploader"`
	Notifier      Notifier      `mapstructure:"notifier"`
	Smtp          Smtp          `mapstructure:"smtp"`
	Cache         Cache         `mapstructure:"cache"`
	SpamManager   SpamManager   `mapstructure:"spam_manager"`
}

// Validate проверяет корректность структуры Config, делегируя валидацию полю Bot.
// Возвращает ошибку, если конфигурация Bot некорректна, или nil, если все проверки пройдены.
func (c *Config) Validate() error {
	if err := c.Bot.Validate(); err != nil {
		return err
	}
	return nil
}
