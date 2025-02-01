package messages

const (
	ErrorConnect   = "Ошибка подключения к базе данных: %v"
	ErrorCreate    = "Ошибка создания пользователя: %v"
	ErrorGet       = "Ошибка получения пользователя: %v"
	ErrorUpdate    = "Ошибка обновления пользователя: %v"
	ErrorDelete    = "Ошибка удаления пользователя: %v"
	ErrorMigration = "Ошибка применения миграций: %v"
	ErrorGetSQLDB  = "Ошибка получения *sql.DB: %v"
)
