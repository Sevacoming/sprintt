package main

import "testing"

func TestGenerateRandomElementsZeroSize(t *testing.T) {
	got := generateRandomElements(0)
	if len(got) != 0 {
		t.Fatalf("ожидали пустой слайс при size=0, получили длину %d", len(got))
	}
}

func TestGenerateRandomElementsNegativeSize(t *testing.T) {
	got := generateRandomElements(-10)
	if len(got) != 0 {
		t.Fatalf("ожидали пустой слайс при size<0, получили длину %d", len(got))
	}
}

func TestGenerateRandomElementsPositiveSize(t *testing.T) {
	const size = 100
	got := generateRandomElements(size)

	if len(got) != size {
		t.Fatalf("ожидали длину %d, получили %d", size, len(got))
	}

	for i, v := range got {
		if v <= 0 {
			t.Fatalf("ожидали только положительные числа, элемент %d = %d", i, v)
		}
	}
}

func TestMaximumEmptySlice(t *testing.T) {
	if got := maximum(nil); got != 0 {
		t.Fatalf("для nil-слайса ожидаем 0, получили %d", got)
	}

	if got := maximum([]int{}); got != 0 {
		t.Fatalf("для пустого слайса ожидаем 0, получили %d", got)
	}
}

func TestMaximumSingleElement(t *testing.T) {
	nums := []int{42}
	if got := maximum(nums); got != 42 {
		t.Fatalf("для одного элемента ожидаем 42, получили %d", got)
	}
}

func TestMaximumSeveralElements(t *testing.T) {
	nums := []int{1, 10, 3, 7, 9}
	if got := maximum(nums); got != 10 {
		t.Fatalf("ожидали максимум 10, получили %d", got)
	}
}

func TestMaximumAllEqual(t *testing.T) {
	nums := []int{5, 5, 5, 5}
	if got := maximum(nums); got != 5 {
		t.Fatalf("ожидали максимум 5, получили %d", got)
	}
}

// Небольшой sanity-check для maxChunks: результат должен совпадать с maximum.
func TestMaxChunksMatchesMaximum(t *testing.T) {
	nums := generateRandomElements(10_000)

	max1 := maximum(nums)
	max2 := maxChunks(nums)

	if max1 != max2 {
		t.Fatalf("maxChunks != maximum: %d vs %d", max1, max2)
	}
}
