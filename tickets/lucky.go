package main

type Lucky struct {
	count int
}

func (l *Lucky) countRecursive(n int, sumA int, sumB int) {
	if n == 0 {
		if sumA == sumB {
			l.count++
		}
		return
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			l.countRecursive(n-1, sumA+i, sumB+j)
		}
	}
}

func (l *Lucky) GetLuckyCount(n int) int {
	l.count = 0

	l.countRecursive(n, 0, 0)

	return l.count
}
