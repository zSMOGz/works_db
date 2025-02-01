package main

import (
	"testing"
	"works_db/internal/database"
	"works_db/internal/repository"
)

const (
	testUserID   = 1
	testUserName = "Тестовый пользователь"
	updatedName  = "Обновленный пользователь"
)

func TestMainOperations(t *testing.T) {
	// Подключение к тестовой базе данных SQL
	dbSQL, err := database.ConnectSQL()
	if err != nil {
		t.Fatalf(ErrorConnect, err)
	}
	defer dbSQL.Close()

	sqlRepo := repository.NewSQLRepository(dbSQL)

	// Тест создания
	err = sqlRepo.Create(testUserID, testUserName)
	if err != nil {
		t.Errorf(ErrorCreate, err)
	}

	// Тест получения
	_, name, err := sqlRepo.Get(testUserID)
	if err != nil || name != testUserName {
		t.Errorf(ErrorGet, err)
	}

	// Тест обновления
	err = sqlRepo.Update(testUserID, updatedName)
	if err != nil {
		t.Errorf(ErrorUpdate, err)
	}

	// Тест удаления
	err = sqlRepo.Delete(testUserID)
	if err != nil {
		t.Errorf(ErrorDelete, err)
	}

	// Подключение к тестовой базе данных GORM
	dbGORM, err := database.ConnectGORM()
	if err != nil {
		t.Fatalf(ErrorConnect, err)
	}

	gormRepo := repository.NewGormRepository(dbGORM)

	// Тест создания
	user, err := gormRepo.Create(testUserName)
	if err != nil {
		t.Errorf(ErrorCreate, err)
	}

	// Тест получения
	retrievedUser, err := gormRepo.Get(user.ID)
	if err != nil || retrievedUser.Name != testUserName {
		t.Errorf(ErrorGet, err)
	}

	// Тест обновления
	err = gormRepo.Update(user.ID, updatedName)
	if err != nil {
		t.Errorf(ErrorUpdate, err)
	}

	// Тест удаления
	err = gormRepo.Delete(user.ID)
	if err != nil {
		t.Errorf(ErrorDelete, err)
	}
}
