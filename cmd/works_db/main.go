package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"

	"works_db/internal/config"
	"works_db/internal/database"
	"works_db/internal/repository"
)

const (
	ErrorConnect             = "Ошибка подключения к базе данных: %v"
	ErrorCreate              = "Ошибка создания пользователя: %v"
	ErrorGet                 = "Ошибка получения пользователя: %v"
	ErrorUpdate              = "Ошибка обновления пользователя: %v"
	ErrorDelete              = "Ошибка удаления пользователя: %v"
	ErrorMigration           = "Ошибка применения миграций: %v"
	ErrorGetSQLDB            = "Ошибка получения *sql.DB: %v"
	MessageMigrationSuccess  = "Миграции успешно применены!"
	MessageInsertSuccessSQL  = "Запись успешно вставлена (SQL)"
	MessageReadSuccessSQL    = "Прочитанная запись (SQL): ID=%d, Name=%s"
	MessageUpdateSuccessSQL  = "Запись успешно обновлена (SQL)"
	MessageDeleteSuccessSQL  = "Запись успешно удалена (SQL)"
	MessageInsertSuccessGORM = "Запись успешно вставлена (GORM): %d"
	MessageReadSuccessGORM   = "Прочитанная запись (GORM): ID=%d, Name=%s"
	MessageUpdateSuccessGORM = "Запись успешно обновлена (GORM)"
	MessageDeleteSuccessGORM = "Запись успешно удалена (GORM)"
)

func main() {
	dbSQL, err := database.ConnectSQL()
	if err != nil {
		log.Printf(ErrorConnect, err)
		return
	}
	defer dbSQL.Close()

	if err := goose.Up(dbSQL, config.MigrationsPath); err != nil {
		log.Fatalf(ErrorMigration, err)
	}

	sqlRepo := repository.NewSQLRepository(dbSQL)

	err = sqlRepo.Create(1, "Валерий Михайлович")
	if err != nil {
		log.Printf(ErrorCreate, err)
		return
	}
	log.Println(MessageInsertSuccessSQL)

	id, name, err := sqlRepo.Get(1)
	if err != nil {
		log.Printf(ErrorGet, err)
		return
	}
	log.Printf(MessageReadSuccessSQL, id, name)

	// Пример использования sqlRepo для обновления записи
	err = sqlRepo.Update(id, "Сергей Петрович")
	if err != nil {
		log.Printf(ErrorUpdate, err)
		return
	}
	log.Println(MessageUpdateSuccessSQL)

	// Пример использования sqlRepo для удаления записи
	err = sqlRepo.Delete(id)
	if err != nil {
		log.Printf(ErrorDelete, err)
		return
	}
	log.Println(MessageDeleteSuccessSQL)

	dbGORM, err := database.ConnectGORM()
	if err != nil {
		log.Printf(ErrorConnect, err)
		return
	}

	gormRepo := repository.NewGormRepository(dbGORM)

	// Пример использования gormRepo для создания записи
	user, err := gormRepo.Create("Евгений Фёдорович")
	if err != nil {
		log.Printf(ErrorCreate, err)
		return
	}
	log.Printf(MessageInsertSuccessGORM, user.ID)

	// Пример использования gormRepo для получения записи
	user, err = gormRepo.Get(user.ID)
	if err != nil {
		log.Printf(ErrorGet, err)
		return
	}
	log.Printf(MessageReadSuccessGORM, user.ID, user.Name)

	// Пример использования gormRepo для обновления записи
	err = gormRepo.Update(user.ID, "Фёдор Евгеньевич")
	if err != nil {
		log.Printf(ErrorUpdate, err)
		return
	}
	log.Println(MessageUpdateSuccessGORM)

	// Пример использования gormRepo для удаления записи
	err = gormRepo.Delete(user.ID)
	if err != nil {
		log.Printf(ErrorDelete, err)
		return
	}
	log.Println(MessageDeleteSuccessGORM)
}
