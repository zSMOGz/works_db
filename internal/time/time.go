package time

import (
	"time"
)

// TimeProvider предоставляет текущее время
type TimeProvider interface {
	Now() time.Time
}

// RealTimeProvider реализует TimeProvider с использованием реального времени
type RealTimeProvider struct{}

func (RealTimeProvider) Now() time.Time {
	return time.Now()
}

// TimeOfDay возвращает текущее время суток
func TimeOfDay(provider TimeProvider) string {
	now := provider.Now()
	hour := now.Hour()
	switch {
	case hour < 6:
		return "Ночь"
	case hour < 12:
		return "Утро"
	case hour < 18:
		return "День"
	default:
		return "Вечер"
	}
}
