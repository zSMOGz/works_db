package time

import (
	"testing"
	"time"
)

// MockTimeProvider реализует TimeProvider с фиксированным временем
type MockTimeProvider struct {
	time time.Time
}

func (m MockTimeProvider) Now() time.Time {
	return m.time
}

func TestTimeOfDay(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected string
	}{
		{"Ночь", time.Date(2023, 1, 1, 3, 0, 0, 0, time.UTC), "Ночь"},
		{"Утро", time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC), "Утро"},
		{"День", time.Date(2023, 1, 1, 15, 0, 0, 0, time.UTC), "День"},
		{"Вечер", time.Date(2023, 1, 1, 21, 0, 0, 0, time.UTC), "Вечер"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			provider := MockTimeProvider{time: tt.time}
			result := TimeOfDay(provider)
			if result != tt.expected {
				t.Errorf("TimeOfDay(%v) = %s; ожидалось %s", tt.time, result, tt.expected)
			}
		})
	}
}
