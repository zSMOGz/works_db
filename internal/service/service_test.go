package service

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockServer создает тестовый сервер, который возвращает фиксированный ответ
func MockServer(response string, statusCode int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(response))
	}))
}

func TestGetUserFromAPI(t *testing.T) {
	// Создаем мок-сервер, который возвращает фиксированный JSON
	mockResponse := `{"id": 1, "name": "Тестовый пользователь"}`
	mockServer := MockServer(mockResponse, http.StatusOK)
	defer mockServer.Close()

	// Создаем HTTP-клиент и вызываем тестируемую функцию
	client := &http.Client{}
	user, err := GetUserFromAPI(client, mockServer.URL, 1)
	if err != nil {
		t.Fatalf("Ошибка при вызове GetUserFromAPI: %v", err)
	}

	// Проверяем результат
	if user.ID != 1 || user.Name != "Тестовый пользователь" {
		t.Errorf("Ожидался пользователь с ID=1 и именем 'Тестовый пользователь', получен %+v", user)
	}
}
