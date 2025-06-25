package providers

import (
	"pnBot/internal/config/models"
	"strings"

	"github.com/spf13/viper"
)

type ViperConfigProvider struct{}

func (v *ViperConfigProvider) Load(
	path,
	filename,
	configType string,
) (*models.Config, error) {
	viper.SetConfigName(filename)
	viper.AddConfigPath(path)
	viper.SetConfigType(configType)

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &models.Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	if err := resolveSecrets(config); err != nil {
		return nil, err
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return config, nil
}
