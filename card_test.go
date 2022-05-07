package deck

import "testing"

// compile-time check of whether card implements Card.
var _ Card = (*card)(nil)

func TestIsJoker(t *testing.T) {
	t.Parallel()

	tc := []struct {
		desc string
		card Card
		want bool
	}{
		{
			desc: "two of hearts is not a joker",
			card: NewCard(Two, Hearts),
			want: false,
		},
		{
			desc: "joker of hearts is a joker",
			card: NewCard(JokerValue, Hearts),
			want: true,
		},
		{
			desc: "joker of joker is a joker",
			card: NewCard(JokerValue, JokerSuit),
			want: true,
		},
		{
			desc: "two of joker is a joker",
			card: NewCard(Two, JokerSuit),
			want: true,
		},
	}

	for _, test := range tc {
		t.Run(test.desc, func(t *testing.T) {
			got := test.card.IsJoker()
			if got != test.want {
				t.Errorf("%v.IsJoker() = %v, want %v", test.card, got, test.want)
			}
		})
	}
}
