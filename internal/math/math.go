package math

import (
	"errors"
	"sync"
)

var ErrDivisionByZero = errors.New("деление на ноль")

// Sum возвращает сумму двух чисел
func Sum(a, b int) int {
	return a + b
}

// Divide делит число a на число b и возвращает результат или ошибку, если b равно нулю
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

// SumSlice возвращает сумму всех элементов слайса
func SumSlice(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// CountUniqueValues возвращает количество уникальных значений в мапе
func CountUniqueValues(m map[string]int) int {
	unique := make(map[int]bool)
	for _, v := range m {
		unique[v] = true
	}
	return len(unique)
}

// SumConcurrent суммирует числа из слайса, используя горутины
func SumConcurrent(numbers []int) int {
	var wg sync.WaitGroup
	sumCh := make(chan int, len(numbers))

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			sumCh <- n
		}(num)
	}

	// Закрываем канал после завершения всех горутин
	go func() {
		wg.Wait()
		close(sumCh)
	}()

	sum := 0
	for n := range sumCh {
		sum += n
	}

	return sum
}
