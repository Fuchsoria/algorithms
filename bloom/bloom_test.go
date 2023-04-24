package bloom

import (
	"testing"
)

var words = []string{
	"abound", "abounds", "abundance",
	"abundant", "accessible", "bloom",
	"blossom", "bolster", "bonny",
	"bonus", "bonuses", "coherent",
	"cohesive", "colorful", "comely",
	"comfort", "gems", "generosity",
	"generous", "generously", "genial",
	"bluff", "cheater", "hate",
	"humanity",
	"hurt", "nuke", "gloomy",
	"facebook", "twitter",
}

func TestBloomFilter(t *testing.T) {
	t.Run("Multitest", func(t *testing.T) {
		bloom := NewBloomFilter(100)

		for _, word := range words {
			bloom.Insert(word)
		}
	})
}
