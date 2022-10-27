package main

import (
	"fmt"
	"strings"
)

type Cases struct {
	MAX_W int
	MAX_H int
}

type Case func(h int, w int) bool

func (c *Cases) Case1(h int, w int) bool {
	return h < w
}

func (c *Cases) Case2(h int, w int) bool {
	return h == w
}

func (c *Cases) Case3(h int, w int) bool {
	return c.MAX_H-h-1 == w
}

func (c *Cases) Case4(h int, w int) bool {
	return h+w < 30
}

func (c *Cases) Case5(h int, w int) bool {
	return h == w/2
}

func (c *Cases) Case6(h int, w int) bool {
	return w < 10 || h < 10
}

func (c *Cases) Case7(h int, w int) bool {
	return w > 15 && h > 15
}

func (c *Cases) Case8(h int, w int) bool {
	return w == 0 || h == 0
}

func (c *Cases) Case9(h int, w int) bool {
	return h-w >= 11 || w-h >= 11
}

func (c *Cases) Case11(h int, w int) bool {
	return w == 1 || w == c.MAX_W-2 || h == 1 || h == c.MAX_H-2
}

func (c *Cases) Case18(h int, w int) bool {
	return !(h > 1 && w > 1 || (h == 0 && w == 0))
}

func (c *Cases) Case19(h int, w int) bool {
	return w == 0 || w == c.MAX_W-1 || h == 0 || h == c.MAX_H-1
}

func (c *Cases) Case23(h int, w int) bool {
	return w%2 == 0 && h%3 == 0
}

func (c *Cases) Case24(h int, w int) bool {
	return w == h || h+w == c.MAX_H-1
}

func (c *Cases) Case25(h int, w int) bool {
	return w%6 == 0 || h%6 == 0
}

func (c *Cases) Render(f Case) {
	for h := 0; h < c.MAX_H; h++ {
		str := strings.Builder{}
		for w := 0; w < c.MAX_W; w++ {
			if f(h, w) {
				str.WriteString("# ")
			} else {
				str.WriteString(". ")
			}
		}
		fmt.Println(str.String())
	}
	fmt.Println()
}

func main() {
	cases := Cases{25, 25}

	cases.Render(cases.Case1)
	cases.Render(cases.Case2)
	cases.Render(cases.Case3)
	cases.Render(cases.Case4)
	cases.Render(cases.Case5)
	cases.Render(cases.Case6)
	cases.Render(cases.Case7)
	cases.Render(cases.Case8)
	cases.Render(cases.Case9)
	cases.Render(cases.Case11)
	cases.Render(cases.Case18)
	cases.Render(cases.Case19)
	cases.Render(cases.Case23)
	cases.Render(cases.Case24)
	cases.Render(cases.Case25)
}
