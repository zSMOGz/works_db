package repository

import (
	"testing"

	"works_db/internal/database"
	"works_db/internal/messages"
)

const (
	testUserIDSQL   = 1
	testUserNameSQL = "Тестовый пользователь"
	updatedNameSQL  = "Обновленный пользователь"
)

func TestSQLRepository(t *testing.T) {
	db, err := database.ConnectSQL()
	if err != nil {
		t.Fatalf(messages.ErrorConnect, err)
	}
	defer db.Close()
	repo := NewSQLRepository(db)

	// Тест создания
	err = repo.Create(testUserIDSQL, testUserNameSQL)
	if err != nil {
		t.Errorf(messages.ErrorCreate, err)
	}

	// Тест получения
	_, name, err := repo.Get(testUserIDSQL)
	if err != nil || name != testUserNameSQL {
		t.Errorf(messages.ErrorGet, err)
	}

	// Тест обновления
	err = repo.Update(testUserIDSQL, updatedNameSQL)
	if err != nil {
		t.Errorf(messages.ErrorUpdate, err)
	}

	// Тест удаления
	err = repo.Delete(testUserIDSQL)
	if err != nil {
		t.Errorf(messages.ErrorDelete, err)
	}
}
