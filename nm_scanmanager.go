package main

import(
	"fmt"
)

func run_wealth_track(){
	skillTracker = make(map[string]Jutsu)
	currentPlayerSkillTracker = make(map[string]Jutsu)
	pathBase := "player"
	determineActivePlayers(pathBase, HowFarBackWeGo)
	loadPlayers(pathBase)
	//checkFinances()
	//JutsuUsageData()
	CurrentUserJutsuUsageData()
}

func run_item_track(){
	pathBase := "area"
	fmt.Println(O_flags[4])
	loadItemInfo(pathBase)
	printItems(20)
	printBasicStats()
	fmt.Println(len(AllItems))
	writeAllItemsFile()
}

func main(){
	run_item_track()
}