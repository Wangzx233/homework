package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)
func init() {
	rand.Seed(time.Now().Unix())
}
const (
	//黑桃
	Spade = 0
	//红桃
	Hearts = 1
	//梅花
	Club = 2
	//方块
	Diamond = 3
)

type Poker struct {
	Num int
	Flower int
}

func (p Poker)PokerSelf()string  {
	var buffer string

	switch p.Flower {
	case Spade:
		buffer += "♤"
	case Hearts:
		buffer += "♡"
	case Club:
		buffer += "♧"
	case Diamond:
		buffer += "♢"
	}
	switch p.Num {
	case 13:
		buffer += "2"
	case 12:
		buffer += "A"
	case 11:
		buffer += "K"
	case 10:
		buffer += "Q"
	case 9:
		buffer += "J"
	default:
		buffer += strconv.Itoa(p.Num+2)
	}

	return buffer
}

func CreatePokers()(pokers Pokers)  {
	for i := 1; i < 14; i++ {
		for j := 0; j < 4; j++ {
				pokers = append(pokers,Poker{
					Num:    i,
					Flower: j,
				})
		}
	}
	return
}

type Pokers []Poker

func (p Pokers)Print()  {
	for _, i2 := range p {
		fmt.Print(i2.PokerSelf()," ")
	}
	fmt.Println()
}
func main() {
	pokers:=CreatePokers()
	Random(pokers,3)
	pokers.Print()
}

//Fisher-Yates随机置乱算法:随机抽一个数与最后一个数交换，然后再随机抽一个数与倒数第二个数交换，如此重复直到第一个数。随机抽的数由时间随机
func Random(strings []Poker, length int) (Pokers, error) {

	for i := len(strings) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		strings[i], strings[num] = strings[num], strings[i]
	}

	var pokers Pokers
	for i := 0; i < length; i++ {
		pokers = append(pokers,Poker{
			Num:    i,
		})
	}
	return pokers, nil
}