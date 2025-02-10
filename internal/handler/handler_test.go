package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	// Создаем запрос с методом GET и пустым телом
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatalf("Не удалось создать запрос: %v", err)
	}

	// Создаем ResponseRecorder для записи ответа
	rr := httptest.NewRecorder()

	// Вызываем обработчик напрямую
	HelloHandler(rr, req)

	// Проверяем код ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Ожидался код ответа %d, получен %d", http.StatusOK, status)
	}

	// Проверяем содержимое ответа
	expected := map[string]string{"message": "Привет, мир!"}
	var actual map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&actual); err != nil {
		t.Fatalf("Не удалось декодировать ответ: %v", err)
	}

	if actual["message"] != expected["message"] {
		t.Errorf("Ожидалось сообщение %q, получено %q", expected["message"], actual["message"])
	}
}

func TestGoodbyeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/goodbye", nil)
	if err != nil {
		t.Fatalf("Не удалось создать запрос: %v", err)
	}

	rr := httptest.NewRecorder()
	GoodbyeHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Ожидался код ответа %d, получен %d", http.StatusOK, status)
	}

	expected := map[string]string{"message": "Пока, мир!"}
	var actual map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&actual); err != nil {
		t.Fatalf("Не удалось декодировать ответ: %v", err)
	}

	if actual["message"] != expected["message"] {
		t.Errorf("Ожидалось сообщение %q, получено %q", expected["message"], actual["message"])
	}
}
