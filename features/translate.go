package features

func Translate(source string, lang string) string {
	if lang == "" {
		lang = "ru"
	}
	v := map[string](map[string]string){
		"Войти":                      {"eng": "Login", "ru": "Войти", "chn": "登录"},
		"Выйти":                      {"eng": "Logout", "ru": "Выйти", "chn": "出口"},
		"Главная":                    {"eng": "Main", "ru": "Главная", "chn": "主"},
		"Язык":                       {"eng": "Language", "ru": "Язык", "chn": "语言"},
		"Тема":                       {"eng": "Topic", "ru": "Тема", "chn": "话题"},
		"Пост":                       {"eng": "Post", "ru": "Пост", "chn": "文章"},
		"Содержимое":                 {"eng": "Content", "ru": "Содержимое", "chn": "内容"},
		"Отправить":                  {"eng": "Send", "ru": "Отправить", "chn": "发送"},
		"Подписчики":                 {"eng": "Followers", "ru": "Подписчики", "chn": "关注者"},
		"Подписки":                   {"eng": "Following", "ru": "Подписки", "chn": "下列"},
		"Подписаться":                {"eng": "Follow", "ru": "Подписаться", "chn": "跟随"},
		"Отписаться":                 {"eng": "Unfollow", "ru": "Отписаться", "chn": "取消关注"},
		"Регистрация":                {"eng": "Registration", "ru": "Регистрация", "chn": "注册"},
		"Нет подписчиков":            {"eng": "No followers ;(", "ru": "Нет подписчиков ;(", "chn": "没有追随者 ;("},
		"Нет подписок":               {"eng": "No following", "ru": "Нет подписок", "chn": "没有下列"},
		"Новый пост":                 {"eng": "Create Post", "ru": "Новый пост", "chn": "写个帖子"},
		"Пользователя не существует": {"eng": "User doesn't exist", "ru": "Пользователя не существует", "chn": "用户不存在"},
		"Логин":         {"eng": "Login", "ru": "Логин", "chn": "昵称"},
		"Пароль":        {"eng": "Password", "ru": "Пароль", "chn": "密码"},
		"Удалить пост":  {"eng": "Delete Post", "ru": "Удалить пост", "chn": "删除文章"},
		"Комментарий":   {"eng": "Comment", "ru": "Комментарий", "chn": "评论"},
		"Удалить":       {"eng": "Delete", "ru": "Удалить", "chn": "删除"},
		"Редактировать": {"eng": "Edit", "ru": "Редактировать", "chn": "编辑"}}
	return v[source][lang]
}
