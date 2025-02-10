package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// User представляет структуру пользователя
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetUserFromAPI получает данные пользователя по ID из внешнего API
func GetUserFromAPI(client *http.Client, apiURL string, userID int) (*User, error) {
	resp, err := client.Get(fmt.Sprintf("%s/users/%d", apiURL, userID))
	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе к API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API вернул код ошибки: %d", resp.StatusCode)
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("ошибка при декодировании ответа: %v", err)
	}

	return &user, nil
}
