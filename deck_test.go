package deck

import (
	"fmt"
	"testing"
)

// compile-time check of whether deck implements Deck.
var _ Deck = (*deck)(nil)

func TestCount(t *testing.T) {
	t.Parallel()

	tc := []struct {
		desc   string
		jokers int
		want   int
	}{
		{
			desc:   "no jokers should be 52 cards",
			jokers: 0,
			want:   52,
		},
		{
			desc:   "one joker should be 53 cards",
			jokers: 1,
			want:   53,
		},
		{
			desc:   "four jokers should be 56 cards",
			jokers: 4,
			want:   56,
		},
	}

	for _, test := range tc {
		t.Run(test.desc, func(t *testing.T) {
			d := NewDeck(test.jokers)

			if d.Count() != test.want {
				t.Errorf("got %v, want %v", d.Count(), test.want)
			}
		})
	}
}

func TestRemainingCount_Draw(t *testing.T) {
	t.Parallel()

	tc := []struct {
		desc string
		draw int
		want int
	}{
		{
			desc: "after drawing zero 52 should remain",
			draw: 0,
			want: 52,
		},
		{
			desc: "after drawing once 51 should remain",
			draw: 1,
			want: 51,
		},
		{
			desc: "after drawing ten 42 should remain",
			draw: 10,
			want: 42,
		},
		{
			desc: "after drawing 52 0 should remain",
			draw: 52,
			want: 0,
		},
	}

	for _, test := range tc {
		t.Run(fmt.Sprintf("Draw %s", test.desc), func(t *testing.T) {
			d := NewDeck(0)

			for i := 0; i < test.draw; i++ {
				_, _ = d.Draw()
			}

			if d.RemainingCount() != test.want {
				t.Errorf("got %v, want %v", d.RemainingCount(), test.want)
			}
		})

		t.Run(fmt.Sprintf("DrawX %s", test.desc), func(t *testing.T) {
			d := NewDeck(0)
			_, _ = d.DrawX(test.draw)

			if d.RemainingCount() != test.want {
				t.Errorf("got %v, want %v", d.RemainingCount(), test.want)
			}
		})
	}
}

func TestRemainingCount_Discard(t *testing.T) {
	t.Parallel()

	tc := []struct {
		desc string
		draw int
		want int
	}{
		{
			desc: "after discarding zero 52 should remain",
			draw: 0,
			want: 52,
		},
		{
			desc: "after discarding once 51 should remain",
			draw: 1,
			want: 51,
		},
		{
			desc: "after discarding ten 42 should remain",
			draw: 10,
			want: 42,
		},
		{
			desc: "after discarding 52 0 should remain",
			draw: 52,
			want: 0,
		},
	}

	for _, test := range tc {
		t.Run(fmt.Sprintf("Discard %s", test.desc), func(t *testing.T) {
			d := NewDeck(0)

			for i := 0; i < test.draw; i++ {
				_ = d.Discard()
			}

			if d.RemainingCount() != test.want {
				t.Errorf("got %v, want %v", d.RemainingCount(), test.want)
			}
		})

		t.Run(fmt.Sprintf("DiscardX %s", test.desc), func(t *testing.T) {
			d := NewDeck(0)
			_ = d.DiscardX(test.draw)

			if d.RemainingCount() != test.want {
				t.Errorf("got %v, want %v", d.RemainingCount(), test.want)
			}
		})
	}
}

func TestDraw(t *testing.T) {
	t.Parallel()

	t.Run("should be able to draw from fresh deck", func(t *testing.T) {
		d := NewDeck(0)

		if got, err := d.Draw(); err != nil {
			t.Errorf("Draw() = (%v, %v), want no error", got, err)
		} else if got == nil {
			t.Error("Draw() = (nil, nil), want non-nil card")
		}
	})
}

func TestDrawError(t *testing.T) {
	t.Parallel()

	t.Run("should fail to draw with no remaining cards", func(t *testing.T) {
		d := NewDeck(0)

		for d.RemainingCount() != 0 {
			if got, err := d.Draw(); err != nil {
				t.Errorf("Draw() = (%v, %v), want no error", got, err)
			} else if got == nil {
				t.Error("Draw() = (nil, nil), want non-nil card")
			}
		}

		if got, err := d.Draw(); err == nil {
			t.Errorf("Draw() = (%v, nil), want error", got)
		} else if got != nil {
			t.Errorf("Draw() = (%v, %v), want nil card", got, err)
		}
	})
}

func TestDrawX(t *testing.T) {
	t.Parallel()

	tc := []struct {
		desc string
		draw int
	}{
		{
			desc: "should be able to draw 0 from fresh deck",
			draw: 0,
		},
		{
			desc: "should be able to draw 10 from fresh deck",
			draw: 10,
		},
		{
			desc: "should be able to draw 52 from fresh deck",
			draw: 52,
		},
	}

	for _, test := range tc {
		t.Run(test.desc, func(t *testing.T) {
			d := NewDeck(0)

			if got, err := d.DrawX(test.draw); err != nil {
				t.Errorf("DrawX(%v) = (%v, %v), want no error", test.draw, got, err)
			} else if got == nil {
				t.Errorf("DrawX(%v) = (nil, nil), want non-nil cards", test.draw)
			} else if len(got) != test.draw {
				t.Errorf("DrawX(%v) = len %v, want len %v", test.draw, len(got), test.draw)
			}
		})
	}
}

func TestDrawXError(t *testing.T) {
	t.Parallel()

	t.Run("should fail to draw 53 with fresh deck", func(t *testing.T) {
		d := NewDeck(0)

		if got, err := d.DrawX(53); err == nil {
			t.Errorf("Draw() = (%v, nil), want error", got)
		} else if got != nil {
			t.Errorf("Draw() = (%v, %v), want nil cards", got, err)
		}
	})
}

func TestDiscard(t *testing.T) {
	t.Parallel()

	t.Run("should be able to discard from fresh deck", func(t *testing.T) {
		d := NewDeck(0)

		if err := d.Discard(); err != nil {
			t.Errorf("Discard() = %v, want no error", err)
		}
	})
}

func TestDiscardError(t *testing.T) {
	t.Parallel()

	t.Run("should fail to discard with no remaining cards", func(t *testing.T) {
		d := NewDeck(0)

		for d.RemainingCount() != 0 {
			if err := d.Discard(); err != nil {
				t.Errorf("Discard() = %v, want no error", err)
			}
		}

		if err := d.Discard(); err == nil {
			t.Error("Discard() = nil, want error")
		}
	})
}

func TestDiscardX(t *testing.T) {
	t.Parallel()

	tc := []struct {
		desc string
		draw int
	}{
		{
			desc: "should be able to discard 0 from fresh deck",
			draw: 0,
		},
		{
			desc: "should be able to discard 10 from fresh deck",
			draw: 10,
		},
		{
			desc: "should be able to discard 52 from fresh deck",
			draw: 52,
		},
	}

	for _, test := range tc {
		t.Run(test.desc, func(t *testing.T) {
			d := NewDeck(0)

			if err := d.DiscardX(test.draw); err != nil {
				t.Errorf("DiscardX(%v) = %v, want no error", test.draw, err)
			}
		})
	}
}

func TestDiscardXError(t *testing.T) {
	t.Parallel()

	t.Run("should fail to discard 53 with fresh deck", func(t *testing.T) {
		d := NewDeck(0)

		if err := d.DiscardX(53); err == nil {
			t.Error("Discard(53) = nil, want error")
		}
	})
}

func TestShuffle(t *testing.T) {
	t.Parallel()

	tc := []struct {
		desc string
		draw int
		want int
	}{
		{
			desc: "remaining count should be 52 after drawing 0 and shuffling",
			draw: 0,
			want: 52,
		},
		{
			desc: "remaining count should be 42 after drawing 10 and shuffling",
			draw: 10,
			want: 42,
		},
		{
			desc: "remaining count should be 0 after drawing 52 and shuffling",
			draw: 52,
			want: 0,
		},
	}

	for _, test := range tc {
		t.Run(test.desc, func(t *testing.T) {
			d := NewDeck(0)

			_, err := d.DrawX(test.draw)
			if err != nil {
				t.Fatalf("d.DrawX(%v) = %v, want no error", test.draw, err)
			}

			d.Shuffle(1)

			if d.RemainingCount() != test.want {
				t.Errorf("got %v, want %v", d.RemainingCount(), test.want)
			}
		})
	}
}

func TestFullShuffle(t *testing.T) {
	t.Parallel()

	tc := []struct {
		desc string
		draw int
	}{
		{
			desc: "remaining count should be 52 after drawing 0 and full shuffling",
			draw: 0,
		},
		{
			desc: "remaining count should be 52 after drawing 10 and full shuffling",
			draw: 10,
		},
		{
			desc: "remaining count should be 52 after drawing 52 and full shuffling",
			draw: 52,
		},
	}

	for _, test := range tc {
		t.Run(test.desc, func(t *testing.T) {
			d := NewDeck(0)

			_, err := d.DrawX(test.draw)
			if err != nil {
				t.Fatalf("d.DrawX(%v) = %v, want no error", test.draw, err)
			}

			d.FullShuffle()

			if d.RemainingCount() != 52 {
				t.Errorf("got %v, want 52", d.RemainingCount())
			}
		})
	}
}
