package engine

import "fmt"

type Player struct {
	Name     string
	Hands    []*Hand
	IsDealer bool
}

func (p *Player) GetHandsString() string {
	handsString := ""

	for _, hand := range p.Hands {
		handsString += fmt.Sprintf("%s\n", hand.String())
	}

	return handsString
}
