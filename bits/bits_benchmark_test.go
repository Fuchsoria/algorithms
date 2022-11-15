package bits

import "testing"

func BenchmarkKingMoves(b *testing.B) {
	chess := Chess{}

	for _, bm := range casesKing {
		b.Run(bm.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = chess.GetKingBitboardMoves(bm.Input)
			}
		})
	}
}

func BenchmarkHorseMoves(b *testing.B) {
	chess := Chess{}

	for _, bm := range casesKing {
		b.Run(bm.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = chess.GetHorseBitboardMoves(bm.Input)
			}
		})
	}
}

func BenchmarkCount(b *testing.B) {
	chess := Chess{}

	for _, bm := range casesKing {
		b.Run(bm.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = chess.Popcnt(chess.GetKingBitboardMoves(bm.Input))
			}
		})
	}

	for _, bm := range casesHorse {
		b.Run(bm.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = chess.Popcnt(chess.GetHorseBitboardMoves(bm.Input))
			}
		})
	}
}

func BenchmarkOptCount(b *testing.B) {
	chess := Chess{}

	for _, bm := range casesKing {
		b.Run(bm.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = chess.PopcntOpt(chess.GetKingBitboardMoves(bm.Input))
			}
		})
	}

	for _, bm := range casesHorse {
		b.Run(bm.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = chess.PopcntOpt(chess.GetHorseBitboardMoves(bm.Input))
			}
		})
	}
}
