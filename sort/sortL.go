package sort

type SortL struct {
}

func (s *SortL) getMax(arr []int) int {
	n := len(arr)
	max := arr[0]

	for i := 1; i < n; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}

func (s *SortL) radixCount(input []int, n int, exp int) {
	output := make([]int, n)
	count := make([]int, 10)

	// store count of occurrences in count[]
	for i := 0; i < n; i++ {
		count[(input[i]/exp)%10]++
	}

	// change count[i] so that count[i] now contains actual position of this digit in output[]
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// Build the output array
	for i := n - 1; i >= 0; i-- {
		output[count[(input[i]/exp)%10]-1] = input[i]
		count[(input[i]/exp)%10]--
	}

	// Copy the output array to arr[], so that arr[] now contains sorted numbers according to current digit
	for i := 0; i < n; i++ {
		input[i] = output[i]
	}
}

func (s *SortL) Radix(input []int) []int {
	array := append([]int{}, input...)
	n := len(array)

	// Find the maximum number to know number of digits
	m := s.getMax(array)

	for exp := 1; m/exp > 0; exp *= 10 {
		s.radixCount(array, n, exp)
	}

	return array
}

func (s *SortL) Bucket(input []int, chunk int) (result []int) {
	array := append([]int{}, input...)
	result = make([]int, 0)

	var max, min int
	for _, n := range array {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	nBuckets := int(max-min)/chunk + 1
	buckets := make([][]int, nBuckets)
	for i := 0; i < nBuckets; i++ {
		buckets[i] = make([]int, 0)
	}

	for _, n := range array {
		idx := int(n-min) / chunk
		buckets[idx] = append(buckets[idx], n)
	}

	for _, bucket := range buckets {
		if len(bucket) > 0 {
			sort := Sort{}
			bucket = sort.Quick(bucket)
			result = append(result, bucket...)
		}
	}

	return result
}

func (s *SortL) Counting(input []int, max int) []int {
	array := append([]int{}, input...)
	n := len(array)

	output := make([]int, n)
	count := make([]int, max)

	for i := 0; i < n; i++ {
		count[array[i]]++
	}

	for i := 1; i < max; i++ {
		count[i] += count[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		output[count[array[i]]-1] = array[i]
		count[array[i]]--
	}

	for i := 0; i < n; i++ {
		array[i] = output[i]
	}

	return array
}
