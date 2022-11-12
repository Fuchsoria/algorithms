package arrays

type array[T any] struct {
	size int
	arr  []T
}

func (a *array[T]) isEmpty() bool {
	return a.size == 0
}

func (a *array[T]) Size() int {
	return a.size
}

func (a *array[T]) Len() int {
	return len(a.arr)
}

func (a *array[T]) Get(index int) T {
	return a.arr[index]
}

func (a *array[T]) GetArray() []T {
	return a.arr[:a.size]
}

func (a *array[T]) add(item T) {
	a.arr[a.Size()-1] = item
}

func (a *array[T]) addByIndex(item T, index int) {
	copyArr := make([]T, len(a.arr))
	copy(copyArr, a.arr)

	newArr := make([]T, 0)
	newArr = append(newArr, copyArr[:index]...)
	newArr = append(newArr, item)
	newArr = append(newArr, copyArr[index:]...)

	a.arr = newArr
	a.size++
}

func (a *array[T]) Remove(index int) T {
	value := a.arr[index]

	a.arr = append(a.arr[:index], a.arr[index+1:]...)

	a.size--

	return value
}

type singleArray[T any] struct {
	array[T]
}

func (s *singleArray[T]) Resize() {
	newArr := make([]T, s.Size()+1)

	copy(newArr, s.arr)

	s.arr = newArr
	s.size++
}

func (s *singleArray[T]) Add(item T) {
	s.Resize()
	s.add(item)
}

func (s *singleArray[T]) AddByIndex(item T, index int) {
	s.addByIndex(item, index)
}

func NewSingleArr[T any]() *singleArray[T] {
	return &singleArray[T]{array[T]{arr: make([]T, 0), size: 0}}
}

type vectorArray[T any] struct {
	vector int
	array[T]
}

func (s *vectorArray[T]) resize() {
	newArr := make([]T, s.size+s.vector)

	copy(newArr, s.arr)

	s.arr = newArr
}

func (s *vectorArray[T]) Add(item T) {
	if s.Size() == s.Len() {
		s.resize()
	}

	s.size++
	s.add(item)
}

func (s *vectorArray[T]) AddByIndex(item T, index int) {
	s.addByIndex(item, index)
}

func NewVectorArr[T int](vector int) *vectorArray[T] {
	return &vectorArray[T]{vector, array[T]{size: 0, arr: make([]T, 0)}}
}

type factorArray[T int] struct {
	factor int
	array[T]
}

func (s *factorArray[T]) resize() {
	if s.isEmpty() {
		s.arr = make([]T, 1)

		return
	}

	newArr := make([]T, s.size*s.factor)

	copy(newArr, s.arr)

	s.arr = newArr
}

func (s *factorArray[T]) Add(item T) {
	if s.size == s.Len() {
		s.resize()
	}

	s.size++
	s.add(item)
}

func (s *factorArray[T]) AddByIndex(item T, index int) {
	s.addByIndex(item, index)
}

func NewFactorArr[T int]() *factorArray[T] {
	return &factorArray[T]{2, array[T]{size: 0, arr: make([]T, 0)}}
}

type matrixArray[T int] struct {
	vector int
	size   int
	arr    *singleArray[*vectorArray[T]]
}

func (s *matrixArray[T]) Get(index int) T {
	return s.arr.Get(index / s.vector).Get(index % s.vector)
}

func (s *matrixArray[T]) Add(item T) {
	if s.size == s.arr.Size()*s.vector {
		s.arr.Add(NewVectorArr[T](s.vector))
	}

	s.arr.Get(s.size / s.vector).Add(item)
	s.size++
}

func NewMatrixArr[T int]() *matrixArray[T] {
	return &matrixArray[T]{100, 0, NewSingleArr[*vectorArray[T]]()}
}

type ArrayList[T any] struct {
	slice []T
}

func (s *ArrayList[T]) Add(item T) {
	s.slice = append(s.slice, item)
}

func (s *ArrayList[T]) AddByIndex(item T, index int) {
	copySlice := make([]T, len(s.slice))
	copy(copySlice, s.slice)

	newSlice := make([]T, 0)
	newSlice = append(newSlice, copySlice[:index]...)
	newSlice = append(newSlice, item)
	newSlice = append(newSlice, copySlice[index:]...)

	s.slice = newSlice
}

func (s *ArrayList[T]) Get(index int) T {
	return s.slice[index]
}

func (s *ArrayList[T]) Remove(index int) T {
	value := s.slice[index]

	s.slice = append(s.slice[:index], s.slice[index+1:]...)

	return value
}

func (s *ArrayList[T]) GetArray() []T {
	return s.slice
}

func NewArrayList[T any]() *ArrayList[T] {
	return &ArrayList[T]{[]T{}}
}
