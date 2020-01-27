package miner

import "fmt"

//Finder function to find ores and send to channel.
func Finder(oreChannel chan string) func(mine []string) {
	return func(mine []string) {
		for _, item := range mine {
			if item == "ore" {
				oreChannel <- item //send item on oreChannel
			}
		}
	}
}

//Breaker of ore from the given channel
func Breaker(oreChannel chan string, minedOreChan chan string) func(n int) {
	return func(n int) {
		for i := 0; i < n; i++ {
			foundOre := <-oreChannel //read from oreChannel
			fmt.Println("From Finder: ", foundOre)
			minedOreChan <- "minedOre" //send to minedOreChan
		}
	}
}

//Smelter smelts ore from the minedOrdeChannel
func Smelter(minedOreChan chan string) func(n int) {
	return func(n int) {
		for i := 0; i < n; i++ {
			minedOre := <-minedOreChan //read from minedOreChan
			fmt.Println("From Miner: ", minedOre)
			fmt.Println("From Smelter: Ore is smelted")
		}
	}
}
