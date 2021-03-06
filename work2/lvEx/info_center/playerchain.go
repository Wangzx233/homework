package info_center

import (
	"fmt"
	"work2/lvEx/util"
)

type PlayerChain struct {
	Players []Player
}

func InitPlayer(num int)*PlayerChain {
	var re PlayerChain
	var p []Player
	for i := 0; i < num; i++ {
		p = append(p, NewPlayer(i))
	}
	re.Players = p
	return &re
}

func NewPlayer(i int) Player {
	return Player{
		Name:     util.RandString(),
		Hp:       3,
		ID:       i,
		Hero:     ChooseHero("甄姬"),
		HandCard: nil,
	}
}

func (p PlayerChain)PrintPlayer()  {
	for i, i2 := range p.Players {
		fmt.Println("player:",i,i2)
	}
}

func (p *PlayerChain)ChooseNextPlayer(now *Player) *Player {
	for i, i2 := range p.Players {
		if i2.ID == now.ID{
			if p.Len()==i+1{
				fmt.Println("下一轮,下一个是-->",p.Players[0].Name)
				return &p.Players[0]
			}
			fmt.Println("下一个是-->",p.Players[i+1].Name)
			return &p.Players[i+1]
		}
	}
	return &Player{}
}

func (p PlayerChain)Len()int {
	return len(p.Players)
}

func (p *PlayerChain)Killed(id int) {
	fmt.Println("----------人被杀就会死----------")
	for i, i2 := range p.Players {
		if i2.ID == id {
			i2.PrintPlayer()
			fmt.Println("被杀死了....")
			p.Players = append(p.Players[:i], p.Players[i+1:]...)
		}
	}
}

