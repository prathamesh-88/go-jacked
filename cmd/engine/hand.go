package engine

import (
	"fmt"
	"math/rand"
	"slices"
)

type Hand struct {
	Cards         []Card
	Player        *Player
	Score         int
	HasAce        bool
	IsDealer      bool
	IsBlackJack   bool
	IsBust        bool
	IsSplittable  bool
	IsCompleted   bool
	TotalBet      int
	Insurance     int
	BustProb      float32
	BlackJackProb float32
}

func (h *Hand) String() string {
	cards := ""

	for _, card := range h.Cards {
		if card.IsHidden {
			cards += "Hidden Card | "
		} else {
			cards += fmt.Sprintf("%v | ", card.String())
		}
	}

	handValue := fmt.Sprintf("Total score: %v", h.Score)
	return cards + handValue
}

func (h *Hand) GetHandValue() int {
	handValue := 0
	for _, card := range h.Cards {
		if !card.IsHidden {
			handValue += card.Value
		}
	}

	if h.HasAce && handValue > 21 {
		handValue -= 10
	}

	h.Score = handValue

	return handValue
}

func (h *Hand) GetNewCard(deck *[]Card, usedCards *[]int, isHidden bool) {
	currentCardIndex := rand.Intn(len(*deck))
	for slices.Contains(*usedCards, currentCardIndex) {
		fmt.Println("Random function caused collision!")
		currentCardIndex = rand.Intn(52*8 + 1)
	}
	currentCard := (*deck)[currentCardIndex]

	if currentCard.Rank == "A" {
		h.HasAce = true
	}
	currentCard.IsHidden = isHidden

	*usedCards = append(*usedCards, currentCardIndex)
	h.Cards = append(h.Cards, currentCard)

	h.GetHandValue()

	if h.Score == 21 {
		h.IsBlackJack = true
	}

	if h.Score > 21 {
		h.IsBust = true
	}
}

func (h *Hand) Split(deck *[]Card, usedCards *[]int) {
	if len(h.Cards) != 2 {
		logger.Debug("Non splittable hand")
		return
	}

	if h.Player.IsDealer {
		logger.Debug("Dealer can't split")
		return
	}

	newHand := GetNewHand(h.Player)
	newHand.Cards = append(newHand.Cards, h.Cards[1])
	newHand.GetNewCard(deck, usedCards, false)

	h.Cards = []Card{h.Cards[0]}
	h.GetNewCard(deck, usedCards, false)

	h.Player.Hands = append(h.Player.Hands, newHand)
}

func GetNewHand(player *Player) *Hand {
	return &Hand{
		Player:        player,
		Cards:         []Card{},
		Score:         0,
		HasAce:        false,
		BustProb:      0,
		BlackJackProb: 0,
	}
}
