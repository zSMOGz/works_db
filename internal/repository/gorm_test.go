package repository

import (
	"testing"
	"works_db/internal/database"
	"works_db/internal/messages"
)

const (
	testUserNameGORM = "Тестовый пользователь"
	updatedNameGORM  = "Обновленный пользователь"
)

func TestGormRepository(t *testing.T) {
	db, err := database.ConnectGORM()
	if err != nil {
		t.Fatalf(messages.ErrorConnect, err)
	}
	repo := NewGormRepository(db)

	// Тест создания
	user, err := repo.Create(testUserNameGORM)
	if err != nil {
		t.Errorf(messages.ErrorCreate, err)
	}

	// Тест получения
	retrievedUser, err := repo.Get(user.ID)
	if err != nil || retrievedUser.Name != testUserNameGORM {
		t.Errorf(messages.ErrorGet, err)
	}

	// Тест обновления
	err = repo.Update(user.ID, updatedNameGORM)
	if err != nil {
		t.Errorf(messages.ErrorUpdate, err)
	}

	// Тест удаления
	err = repo.Delete(user.ID)
	if err != nil {
		t.Errorf(messages.ErrorDelete, err)
	}
}
