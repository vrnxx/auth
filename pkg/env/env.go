package env

import "os"

// GetEnv получить значение переменной окружения по ключу `key`
// иначе вернуть значение по умолчанию `defaultValue`
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
