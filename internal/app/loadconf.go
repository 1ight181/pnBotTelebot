package app

import (
	"fmt"
	"log"
	"os"
	models "pnBot/internal/config/models"
	confprov "pnBot/internal/config/providers/interfaces"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type AppConfigOptions struct {
	Provider    confprov.ConfigProvider
	FileName    string
	FileType    string
	EnvVar      string
	DefaultPath string
}

func loadAppConfig(opts AppConfigOptions) (*models.Config, error) {
	path := os.Getenv(opts.EnvVar)
	if path == "" {
		path = opts.DefaultPath
		log.Printf("Переменная окружения %s не задана, используется путь: %s", opts.EnvVar, path)
	}

	err := godotenv.Load(path + "/.env")
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла")
	}

	config, err := opts.Provider.Load(
		path,
		opts.FileName,
		opts.FileType,
	)

	if err != nil {
		switch e := err.(type) {
		case *os.PathError:
			return nil, fmt.Errorf("ошибка пути: %w", e)
		case *viper.ConfigFileNotFoundError:
			return nil, fmt.Errorf("конфигурационный файл не найден: %w", e)
		case *viper.ConfigParseError:
			return nil, fmt.Errorf("ошибка при разборе конфигурации: %w", e)
		default:
			return nil, fmt.Errorf("неизвестная ошибка: %w", e)
		}
	}

	return config, nil
}
