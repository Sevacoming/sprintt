package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE           = 10_000_000 // размер слайса для бенчмарка
	CHUNKS         = 8          // количество частей
	maxRandomValue = 1_000_000  // верхняя граница случайных чисел
)

// generateRandomElements генерирует слайс из size положительных целых чисел.
// При size <= 0 возвращает пустой слайс.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	result := make([]int, size)
	for i := range result {
		// Генерируем числа в диапазоне [1, maxRandomValue]
		result[i] = rand.Intn(maxRandomValue) + 1
	}

	return result
}

// maximum находит максимальный элемент в слайсе.
// Для пустого или nil-слайса возвращает 0.
func maximum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxVal := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}

	return maxVal
}

// maxChunks делит слайс на CHUNKS частей, в каждой части находит максимум
// в отдельной горутине, а затем возвращает максимум из максимумов.
// Для пустого или nil-слайса возвращает 0.
// Если длина меньше CHUNKS — просто использует maximum() целиком.
func maxChunks(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// Маленький слайс нет смысла резать на части
	if n <= CHUNKS {
		return maximum(nums)
	}

	chunkSize := n / CHUNKS
	maxValues := make([]int, CHUNKS)

	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		// фиксируем значения индексов для замыкания
		chunkIndex := i
		start := chunkIndex * chunkSize
		end := start + chunkSize

		// Последний кусок забирает «хвост», если n не делится нацело
		if chunkIndex == CHUNKS-1 {
			end = n
		}

		go func(idx, from, to int) {
			defer wg.Done()
			maxValues[idx] = maximum(nums[from:to])
		}(chunkIndex, start, end)
	}

	wg.Wait()

	// Ищем максимум среди максимумов
	return maximum(maxValues)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	data := generateRandomElements(SIZE)
	if len(data) == 0 {
		fmt.Println("сгенерирован пустой слайс, нечего обрабатывать")
		return
	}

	start := time.Now()
	maxSeq := maximum(data)
	elapsedSeq := time.Since(start).Microseconds()

	start = time.Now()
	maxParallel := maxChunks(data)
	elapsedParallel := time.Since(start).Microseconds()

	fmt.Printf("Последовательный максимум: %d, время: %d мкс\n", maxSeq, elapsedSeq)
	fmt.Printf("Параллельный максимум (%d кусков): %d, время: %d мкс\n", CHUNKS, maxParallel, elapsedParallel)
}
