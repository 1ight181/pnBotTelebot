package providers

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

func resolveSecrets(cfg interface{}) error {
	val := reflect.ValueOf(cfg).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			original := field.String()
			if isSecretPath(original) {
				secretVal, err := readSecret(original)
				if err != nil {
					return fmt.Errorf("Не удалосьсь прочитать секрет: %q: %w", original, err)
				}
				field.SetString(secretVal)
			}

		case reflect.Struct:
			nested := field.Addr().Interface()
			if err := resolveSecrets(nested); err != nil {
				return err
			}

		case reflect.Ptr:
			if !field.IsNil() && field.Elem().Kind() == reflect.Struct {
				nested := field.Interface()
				if err := resolveSecrets(nested); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func isSecretPath(s string) bool {
	return strings.HasPrefix(s, "/run/secrets/")
}

func readSecret(path string) (string, error) {
	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}
