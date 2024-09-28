package engine

import "fmt"

type Card struct {
	Suit     string
	Rank     string
	Value    int
	IsHidden bool
}

func (c *Card) String() string {
	return fmt.Sprintf("%v of %v", c.Rank, c.Suit)
}
