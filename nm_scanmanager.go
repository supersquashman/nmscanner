package main

import (
	"fmt"
)

func run_wealth_track(){
	skillTracker = make(map[string]Jutsu)
	currentPlayerSkillTracker = make(map[string]Jutsu)
	pathBase := "player"
	DetermineActivePlayers(pathBase, HowFarBackWeGo)
	LoadPlayers(pathBase)
	//CheckFinances()
	//JutsuUsageData()
	//CurrentUserJutsuUsageData()
}

func run_item_track(){
	pathBase := "area"
	fmt.Println(O_flags[4])
	LoadItemInfo(pathBase)
	PrintItems(20)
	PrintBasicStats()
	fmt.Println(len(AllItems))
	WriteAllItemsFile()
}

func run_player_logout_track(){
	playerPathBase := "player"
	areaPathBase := "area"
	LoadItemInfo(areaPathBase)

	skillTracker = make(map[string]Jutsu)
	currentPlayerSkillTracker = make(map[string]Jutsu)
	DetermineActivePlayers(playerPathBase, HowFarBackWeGo)
	LoadPlayers(playerPathBase)


	WritePlayerLogoutDataToFile()
}

func main(){
	//run_item_track()
	//run_wealth_track()
	run_player_logout_track()
}