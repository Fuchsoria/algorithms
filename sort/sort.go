package sort

type Sort struct {
	arr []int
}

func (s *Sort) swap(a int, b int) {
	s.arr[a], s.arr[b] = s.arr[b], s.arr[a]
}

func (s *Sort) binarySearch(key int, low int, high int) int {
	if high <= low {
		if key > s.arr[low] {
			return low + 1
		} else {
			return low
		}
	}

	mid := (low + high) / 2

	if key > s.arr[mid] {
		return s.binarySearch(key, mid+1, high)
	}

	return s.binarySearch(key, low, mid-1)
}

func (s *Sort) Bubble(input []int) []int {
	s.arr = append([]int{}, input...)
	N := len(input)

	for j := N - 1; j >= 0; j-- {
		for i := 0; i < j; i++ {
			if s.arr[i] > s.arr[i+1] {
				s.swap(i, i+1)
			}
		}
	}

	return s.arr
}

func (s *Sort) Insertion(input []int) []int {
	s.arr = append([]int{}, input...)
	N := len(input)

	for i := 1; i < N; i++ {
		for j := i - 1; j >= 0 && s.arr[j] > s.arr[j+1]; j-- {
			s.swap(j, j+1)
		}
	}

	return s.arr
}

func (s *Sort) InsertionShift(input []int) []int {
	s.arr = append([]int{}, input...)
	N := len(input)

	var j int
	for i := 1; i < N; i++ {
		K := s.arr[i]

		for j = i - 1; j >= 0 && s.arr[j] > K; j-- {
			s.arr[j+1] = s.arr[j]
		}

		s.arr[j+1] = K
	}

	return s.arr
}

func (s *Sort) InsertionBinarySearch(input []int) []int {
	s.arr = append([]int{}, input...)
	N := len(input)

	var j int
	for i := 1; i < N; i++ {
		K := s.arr[i]
		P := s.binarySearch(K, 0, i-1)

		for j = i - 1; j >= P; j-- {
			s.arr[j+1] = s.arr[j]
		}

		s.arr[j+1] = K
	}

	return s.arr
}

func (s *Sort) Shell(input []int) []int {
	s.arr = append([]int{}, input...)
	N := len(input)

	for gap := N / 2; gap > 0; gap /= 2 {
		for i := gap; i < N; i++ {
			for j := i; j >= gap && s.arr[j-gap] > s.arr[j]; j -= gap {
				s.swap(j-gap, j)
			}
		}
	}

	return s.arr
}
