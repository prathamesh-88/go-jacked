package main

import (
	"fmt"
	"strings"

	"github.com/prathamesh-88/go-jacked/cmd/engine"
	"github.com/prathamesh-88/go-jacked/utils"
)

var logger = utils.Logger{}

func main() {
	player1 := &engine.Player{
		Name:     "Prathamesh",
		IsDealer: false,
	}
	player2 := &engine.Player{
		Name:     "Kajol",
		IsDealer: false,
	}

	players := []*engine.Player{player1, player2}

	game := engine.Game{
		NDecks:             8,
		IsDAS:              true,
		IsDealerStopSoft17: true,
	}
	game.StartNewGame(players)
	game.GetGameStatus()
	for _, player := range players {
		if player.IsDealer {
			continue
		}
		logger.Debug(fmt.Sprintf("It's %s's turn", player.Name))
		logger.Debug(player.GetHandsString())
		playerHand := player.Hands[0]
		for {
			if playerHand.IsBlackJack {
				logger.Info("Congratulations! You hit blackjack!")
				break
			}
			var playerChoice string
			logger.Info("What do you wish to do? Hit[H] / Stay[S]")
			fmt.Scanln(&playerChoice)
			playerChoice = strings.ToLower(playerChoice)
			if playerChoice == "h" {
				playerHand.GetNewCard(&game.Deck, &game.UsedCards, false)
				logger.Info(fmt.Sprintf("Updated hand: %s", player.GetHandsString()))
				if playerHand.IsBust {
					logger.Info("You went bust! Better luck next hand!")
					break
				}

				if playerHand.IsBlackJack {
					logger.Info("You hit 21 ")
					break
				}
			} else if playerChoice == "s" {
				break
			} else {
				logger.Debug("Please choose a valid option")
			}
		}
	}

	dealer := game.GetDealer()
	dealer.GetHandsString()
	isDealerBust, dealerScore := game.RunDealersHand()
	logger.Info(fmt.Sprintf("Dealer's final score: %d", dealerScore))
	for _, player := range game.Players {
		if player.IsDealer {
			continue
		}
		for idx, hand := range player.Hands {
			if hand.IsBust {
				logger.Info(fmt.Sprintf("%s's Hand %d was bust", player.Name, idx+1))
			}

			if isDealerBust {
				logger.Info(fmt.Sprintf("%s wins hand %d", player.Name, idx+1))
				continue
			}

			if hand.Score > dealerScore {
				logger.Info(fmt.Sprintf("%s wins hand %d", player.Name, idx+1))
			} else if hand.Score < dealerScore {
				logger.Info(fmt.Sprintf("%s lost hand %d", player.Name, idx+1))
			} else {
				logger.Info(fmt.Sprintf("%s's hand %d was a push!", player.Name, idx+1))
			}
		}
	}

}
