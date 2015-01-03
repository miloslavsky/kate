package features

func Strings(sess, lang string) (string, string) {
	var Str, URL string
	if sess == "" {
		Str, URL = Translate("Войти", lang), "/login"
	} else {
		Str, URL = Translate("Выйти", lang), "/logout"
	}
	return Str, URL
}
