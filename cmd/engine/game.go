package engine

import (
	"fmt"

	"github.com/prathamesh-88/go-jacked/utils"
)

var logger = utils.Logger{}

type Game struct {
	NDecks             int
	Deck               []Card
	UsedCards          []int
	Players            []*Player
	IsDAS              bool // Is Double after split allowed
	IsDealerStopSoft17 bool
}

func (g *Game) StartNewGame(players []*Player) {
	g.GetNewDeck(g.NDecks)
	g.Players = players

	for _, player := range players {
		newHand := GetNewHand(player)
		newHand.GetNewCard(&g.Deck, &g.UsedCards, false)
		newHand.GetNewCard(&g.Deck, &g.UsedCards, false)
		player.Hands = append(player.Hands, newHand)
	}

	dealer := Player{
		Name:     "Dealer",
		IsDealer: true,
	}
	dealerHand := GetNewHand(&dealer)
	dealerHand.GetNewCard(&g.Deck, &g.UsedCards, false)
	dealerHand.GetNewCard(&g.Deck, &g.UsedCards, true)
	dealer.Hands = append(dealer.Hands, dealerHand)
	g.Players = append(g.Players, &dealer)
}

func (g *Game) GetNewDeck(numDecks int) {
	deck := []Card{}
	for i := 0; i < numDecks; i++ {
		for _, suit := range Suits {
			for rankIndex, rank := range Ranks {
				card := Card{Suit: suit, Rank: rank, Value: Values[rankIndex]}
				deck = append(deck, card)
			}
		}
	}
	g.Deck = deck
}

func (g *Game) GetDealer() *Player {
	return g.Players[len(g.Players)-1]
}

func (g *Game) RunDealersHand() (bool, int) {
	dealer := g.GetDealer()
	dealerHand := dealer.Hands[0]
	dealerHand.Cards[1].IsHidden = false
	for dealerHand.Score <= 17 {
		dealerHand.GetNewCard(&g.Deck, &g.UsedCards, false)
	}
	return dealerHand.IsBust, dealerHand.Score
}

func (g *Game) GetGameStatus() {
	for index, player := range g.Players {
		logger.Debug(fmt.Sprintf("(Player %d) %s has the following cards:", index+1, player.Name))
		for idx, hand := range player.Hands {
			logger.Debug(fmt.Sprintf("Hand %d: Total current score: %d", idx+1, hand.Score))
			cardsString := ""
			for _, card := range hand.Cards {
				cardsString += fmt.Sprintf("%s of %s | ", card.Rank, card.Suit)
			}
			logger.Debug(cardsString + "\n")
		}
	}
}
