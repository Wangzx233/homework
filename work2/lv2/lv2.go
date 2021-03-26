package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

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

func (p Pokers)Print()  {
	for _, i2 := range p {
		fmt.Print(i2.PokerSelf()," ")
	}
	fmt.Println()
}

////这个是花色优先
//func (p Pokers) Less(i, j int) bool {
//	if p[i].Flower == p[j].Flower{
//		if p[i].Num < p[j].Num {
//			return true
//		}
//		return false
//	}
//	if p[i].Flower > p[j].Flower {
//		return false
//	}
//	return true
//}

//这个是数值优先
func (p Pokers) Less(i, j int) bool {
	if p[i].Num == p[j].Num{
		if p[i].Flower < p[j].Flower {
			return true
		}
		return false
	}
	if p[i].Num > p[j].Num {
		return false
	}
	return true
}

func (p Pokers) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Pokers) Len()  int{
	return len(p)
}

func init()  {
	rand.Seed(time.Now().Unix())
}

func Shuffle(p []Poker)  {
	//Fisher-Yates随机置乱
	for i := len(p) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		p[i], p[num] = p[num], p[i]
	}
}

func main() {
	pokers := CreatePokers()
	fmt.Println("洗牌前")
	pokers.Print()
	fmt.Println("洗牌后")
	Shuffle(pokers)
	pokers.Print()

	//
	fmt.Println("排序后")
	sort.Sort(pokers)
	pokers.Print()

}
