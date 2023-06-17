package i18n

import "context"

const defaultLang = "ru"

// Translate ...
func Translate(ctx context.Context, key string, defaultString string) string {
	lang := GetLangFromContext(ctx)

	v, ok := dictionary[lang][key]
	if ok {
		return v
	}
	return defaultString
}

var dictionary = map[string]map[string]string{
	"ru": {
		"ICS_file_for_game_already_exists": "ICS файл для игры уже существует",
		"ICS_file_name_already_exists":     "ICS файл с таким названием уже существует",
		"ICS_file_not_found":               "ICS файл не найден",
	},
}
