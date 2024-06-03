package main

import (
    "fmt"
    "os"
	"golang.org/x/text/language"
    "golang.org/x/text/message"
)

func WritePlayerLogoutDataToFile(){
	p := message.NewPrinter(language.English)
	//[name, level, hours, legacy count,logout room, logout area name]
	//var player_line_items [6]string
		filename := "PlayerLogoutInfo.csv"
		destination, err := os.Create(filename)
		if err != nil {
			fmt.Println("os.Create:", err)
			return
		}
	for _,player := range Players{

		logout_zone := GetAreaNameByVnum(player.LogoutRoom)

		lineStr := "%v,%v,%v,%v,%v,%v"
		line := p.Sprintf(lineStr,
			player.Name,
			player.CurrentLevel,
			player.PlayedTime,
			player.Legacy,
			player.LogoutRoom,
			logout_zone)
		fmt.Println(line)
		
		fmt.Fprintf(destination, "%s\n", line)
		
	}


	defer destination.Close()

	
}

func GetAreaNameByVnum(roomNum int)string{
	areaName := "error"
	for _,area := range AreasList{
		if roomNum > area.Low_Vnum && roomNum < area.High_Vnum{
			areaName = area.Filename
			return areaName
		}
	}
	return areaName
}

