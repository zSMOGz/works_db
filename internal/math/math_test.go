package math

import (
	"errors"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"положительные числа", 2, 3, 5},
		{"отрицательные числа", -1, -1, -2},
		{"положительное и отрицательное число", 5, -3, 2},
		{"ноль", 0, 0, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Sum(test.a, test.b)
			if result != test.expected {
				t.Errorf("Add(%d, %d) = %d; ожидалось %d", test.a, test.b, result, test.expected)
			}
		})
	}
}

func TestAddWithQuick(t *testing.T) {
	// Определяем функцию, которая будет проверять корректность Add
	checkAdd := func(a, b int) bool {
		result := Sum(a, b)
		return result == a+b
	}

	// Используем quick.Check для автоматической генерации тестов
	if err := quick.Check(checkAdd, nil); err != nil {
		t.Errorf("Ошибка в функции Add: %v", err)
	}
}

func TestAddWithQuickConfig(t *testing.T) {
	checkAdd := func(a, b int) bool {
		result := Sum(a, b)
		return result == a+b
	}

	// Настройка параметров генерации данных
	config := &quick.Config{
		MaxCount: 1000, // Количество итераций
		Values: func(values []reflect.Value, rand *rand.Rand) {
			// Генерация случайных чисел в диапазоне от -100 до 100
			values[0] = reflect.ValueOf(rand.Intn(201) - 100)
			values[1] = reflect.ValueOf(rand.Intn(201) - 100)
		},
	}

	if err := quick.Check(checkAdd, config); err != nil {
		t.Errorf("Ошибка в функции Add: %v", err)
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a, b        int
		expected    int
		expectedErr error
	}{
		{"успешное деление", 10, 2, 5, nil},
		{"деление на ноль", 10, 0, 0, ErrDivisionByZero},
		{"деление с отрицательным числом", -10, 2, -5, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)

			// Проверяем, что ошибка соответствует ожидаемой
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("Divide(%d, %d) вернула ошибку: %v, ожидалось: %v", tt.a, tt.b, err, tt.expectedErr)
			}

			// Если ошибки нет, проверяем результат
			if err == nil && result != tt.expected {
				t.Errorf("Divide(%d, %d) = %d; ожидалось %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestSumSlice(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{"положительные числа", []int{1, 2, 3, 4}, 10},
		{"отрицательные числа", []int{-1, -2, -3, -4}, -10},
		{"смешанные числа", []int{-1, 2, -3, 4}, 2},
		{"пустой слайс", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumSlice(tt.numbers)
			if result != tt.expected {
				t.Errorf("SumSlice(%v) = %d; ожидалось %d", tt.numbers, result, tt.expected)
			}
		})
	}
}

func TestCountUniqueValues(t *testing.T) {
	tests := []struct {
		name     string
		m        map[string]int
		expected int
	}{
		{"уникальные значения", map[string]int{"a": 1, "b": 2, "c": 3}, 3},
		{"дубликаты значений", map[string]int{"a": 1, "b": 1, "c": 2}, 2},
		{"пустая мапа", map[string]int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountUniqueValues(tt.m)
			if result != tt.expected {
				t.Errorf("CountUniqueValues(%v) = %d; ожидалось %d", tt.m, result, tt.expected)
			}
		})
	}
}

func TestSumConcurrent(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{"положительные числа", []int{1, 2, 3, 4}, 10},
		{"отрицательные числа", []int{-1, -2, -3, -4}, -10},
		{"смешанные числа", []int{-1, 2, -3, 4}, 2},
		{"пустой слайс", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumConcurrent(tt.numbers)
			if result != tt.expected {
				t.Errorf("SumConcurrent(%v) = %d; ожидалось %d", tt.numbers, result, tt.expected)
			}
		})
	}
}
