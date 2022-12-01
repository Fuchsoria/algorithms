package sort

type Sort struct {
	arr []int
}

func (s *Sort) swap(a int, b int) {
	s.arr[a], s.arr[b] = s.arr[b], s.arr[a]
}

// search for index of highest item in array till j
func (s *Sort) findMax(j int) int {
	maxIndex := 0
	for i := 1; i <= j; i++ {
		if s.arr[i] > s.arr[maxIndex] {
			maxIndex = i
		}
	}

	return maxIndex
}

func (s *Sort) binarySearch(key int, startIndex int, endIndex int) int {
	// Recursion end condition
	if endIndex <= startIndex {
		if key > s.arr[startIndex] {
			return startIndex + 1
		}

		return startIndex
	}

	midIndex := (startIndex + endIndex) / 2

	if key > s.arr[midIndex] {
		// Return right part
		return s.binarySearch(key, midIndex+1, endIndex)
	}

	// Return left part
	return s.binarySearch(key, startIndex, midIndex-1)
}

func (s *Sort) heapify(root, size int) {
	X := root

	// calc Left item
	L := 2*X + 1

	// calc Right item
	R := L + 1

	// checking for highest item
	if L < size && s.arr[L] > s.arr[X] {
		X = L
	}

	// checking for highest item
	if R < size && s.arr[R] > s.arr[X] {
		X = R
	}

	// end recursion if nothing anymore to swap
	if X == root {
		return
	}

	// swap root with highest item
	s.swap(root, X)
	// start heap recheck again
	s.heapify(X, size)
}

func (s *Sort) split(L, R int) int {
	P := s.arr[R]
	M := L - 1

	for j := L; j <= R; j++ {
		if s.arr[j] <= P {
			M++
			s.swap(M, j)
		}
	}

	return M
}

func (s *Sort) merge(L, M, R int) {
	T := make([]int, R-L+1)

	a := L
	b := M + 1
	m := 0

	for a <= M && b <= R {
		if s.arr[a] < s.arr[b] {
			T[m] = s.arr[a]
			m++
			a++
		} else {
			T[m] = s.arr[b]
			m++
			b++
		}
	}

	for a <= M {
		T[m] = s.arr[a]
		m++
		a++
	}

	for b <= R {
		T[m] = s.arr[b]
		m++
		b++
	}

	for i := L; i <= R; i++ {
		s.arr[i] = T[i-L]
	}
}

func (s *Sort) Bubble(input []int) []int {
	s.arr = append([]int{}, input...)
	N := len(input)

	// going through array from end to start
	for j := N - 1; j >= 0; j-- {
		// moving highest item to end
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

	// select item and move them to sorted part
	for i := 1; i < N; i++ {
		// find right order for moved item in sorted part via continuously swap
		for j := i - 1; j >= 0 && s.arr[j] > s.arr[j+1]; j-- {
			s.swap(j, j+1)
		}
	}

	return s.arr
}

func (s *Sort) InsertionShift(input []int) []int {
	s.arr = append([]int{}, input...)
	N := len(input)

	// init var for sub loop to use it in second case
	var j int
	// loop through unsorted part
	for i := 1; i < N; i++ {
		// select item
		K := s.arr[i]

		// shift items
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
		// binary search for right place
		P := s.binarySearch(K, 0, i-1)

		// shift untill finded place
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

	// set gap which will be divided by 2 every iteration
	for gap := N / 2; gap > 0; gap /= 2 {
		//
		for i := gap; i < N; i++ {
			for j := i; j >= gap && s.arr[j-gap] > s.arr[j]; j -= gap {
				s.swap(j-gap, j)
			}
		}
	}

	return s.arr
}

func (s *Sort) Selection(input []int) []int {
	s.arr = append([]int{}, input...)
	N := len(input)

	// run sort from end to start
	for j := N - 1; j > 0; j-- {
		// swap current item with max element
		s.swap(j, s.findMax(j))
	}

	return s.arr
}

func (s *Sort) Heap(input []int) []int {
	s.arr = append([]int{}, input...)
	N := len(input)

	// h = parent item of last item and go through all items recursively
	for h := N/2 - 1; h >= 0; h-- {
		// recheck heap recursively
		s.heapify(h, len(input))
	}

	// run sort from end to start
	for j := N - 1; j > 0; j-- {
		// swap first index which is HIGHEST because of heap structure with the last one
		s.swap(0, j)
		// recheck heap structure
		s.heapify(0, j)
	}

	return s.arr
}

func (s *Sort) quickSort(L, R int) {
	if L >= R {
		return
	}

	M := s.split(L, R)
	s.quickSort(L, M-1)
	s.quickSort(M+1, R)
}

func (s *Sort) Quick(input []int) []int {
	s.arr = append([]int{}, input...)
	N := len(input)

	s.quickSort(0, N-1)

	return s.arr
}

func (s *Sort) mergeSort(L, R int) {
	if L >= R {
		return
	}

	M := (L + R) / 2

	s.mergeSort(L, M)
	s.mergeSort(M+1, R)
	s.merge(L, M, R)
}

func (s *Sort) Merge(input []int) []int {
	s.arr = append([]int{}, input...)

	s.mergeSort(0, len(s.arr)-1)

	return s.arr
}
