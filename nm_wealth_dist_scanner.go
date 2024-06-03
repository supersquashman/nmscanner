package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
	"path/filepath"
	"io/fs"
	"strings"
	"sort"
	"strconv"
	"golang.org/x/text/language"
	"golang.org/x/text/number"
    "golang.org/x/text/message"
	"time"
)


var Players []Player
var currentPlayers []string

var VillageLookup = []string{"Leaf","Sand","Stone","Mist","Sound","Cloud"}
var skillTracker map[string]Jutsu
var currentPlayerSkillTracker map[string]Jutsu

var skillKeys []string
var currentSkillKeys []string

var HowFarBackWeGo = 6 //in months

type Player struct{
	Name string
	AllMoney int
	Bank int
	Wallet int
	Legacy int
	CurrentLevel int
	OverallLevel int `default:"0"`
	Village string
	LogoutRoom int
	PlayedTime int
}

type Jutsu struct{
	Leaf int `default:"0"`
	Sand int `default:"0"`
	Mist int `default:"0"`
	Stone int `default:"0"`
	Cloud int `default:"0"`
	Sound int `default:"0"`
	TotalCount int `default:"0"`
}

func DetermineActivePlayers(currentPath string, timeframe int){
	currentTime := time.Now()
	cutoff := currentTime.AddDate(0,(timeframe*-1),0).Unix()

	filepath.WalkDir(currentPath, func (Fpath string, di fs.DirEntry, err error) error {
		if !di.IsDir(){
			info, erri := di.Info()

			pFile, err := os.Open(Fpath)
			if err != nil {
				log.Fatalf("Error opening file: %v", err)
			}
			defer pFile.Close()

			// Read the file content
			//content, err := ioutil.ReadAll(pFile)
			playerReader := bufio.NewScanner(pFile)
			playerReader.Split(bufio.ScanLines)
  
			for playerReader.Scan() {
				if strings.Contains(playerReader.Text(), "LastOn"){
					lastonFields := strings.Fields(playerReader.Text())
					lastTime, erri := strconv.ParseInt(lastonFields[1], 10, 64)
					if (lastTime >= cutoff){
						currentPlayers = append(currentPlayers,info.Name())
					}
					if erri != nil {err = erri}
				}
			}

			pFile.Close()

			if erri != nil {err = erri}
		}
		return err
	})
}

func isCurrentPlayer(playerName string) bool{
	for _,pname := range currentPlayers{
		if pname==playerName{
			//fmt.Println(playerName+" is a current player")
			return true
		}
	}
	return false
}

func LoadPlayers(currentPath string){
	filepath.WalkDir(currentPath, func (Fpath string, di fs.DirEntry, err error) error {
		if !di.IsDir(){
			var tempPlayer Player
			info, erri := di.Info()
			tempPlayer.Name = info.Name()
			isCurrent := isCurrentPlayer(tempPlayer.Name)

			pFile, err := os.Open(Fpath)
			if err != nil {
				log.Fatalf("Error opening file: %v", err)
			}
			defer pFile.Close()

			// Read the file content
			//content, err := ioutil.ReadAll(pFile)
			playerReader := bufio.NewScanner(pFile)
			playerReader.Split(bufio.ScanLines)
  
			for playerReader.Scan() {
				//fmt.Println(playerReader.Text())
				
				if strings.Contains(playerReader.Text(), "Bank"){
					tempPlayer.Bank, err = strconv.Atoi(strings.Fields(playerReader.Text())[1])
				}
				if strings.Contains(playerReader.Text(), "Gold"){
					tempPlayer.Wallet,err = strconv.Atoi(strings.Fields(playerReader.Text())[1])
				}
				if strings.Contains(playerReader.Text(), "Level"){
					levelFields := strings.Fields(playerReader.Text())
					if levelFields[0]=="Level"{
						tempPlayer.CurrentLevel, err = strconv.Atoi(levelFields[1])
						tempPlayer.OverallLevel += tempPlayer.CurrentLevel
					}
				}
				if strings.Contains(playerReader.Text(), "Class"){
					classFields := strings.Fields(playerReader.Text())
					if classFields[0]=="Class"{
						villVal, erri := strconv.Atoi(classFields[1])
						tempPlayer.Village=VillageLookup[villVal] 
						if erri != nil {err = erri}
					}
				}
				if strings.Contains(playerReader.Text(), "Legacy"){
					legacyFields := strings.Fields(playerReader.Text())
					if legacyFields[0]=="Legacy"{
						tempPlayer.Legacy, err = strconv.Atoi(legacyFields[1])
						tempPlayer.OverallLevel += tempPlayer.Legacy*300
					}
				}
				if strings.Contains(playerReader.Text(),"Room"){
					roomFields := strings.Fields(playerReader.Text())
					if roomFields[0]=="Room"{
						tempPlayer.LogoutRoom, err = strconv.Atoi(roomFields[1])
					}
				}
				if strings.Contains(playerReader.Text(),"Played"){
					timePlayedFields := strings.Fields(playerReader.Text())
					if timePlayedFields[0]=="Played"{
						tempPlayer.PlayedTime, err = strconv.Atoi(timePlayedFields[1])
					}
				}
				
				if strings.Contains(playerReader.Text(), "Skill"){
					skillFields := strings.Fields(playerReader.Text())
					if skillFields[0]=="Skill"{
						skillName:=strings.Split(playerReader.Text(),"'")[1]
						//fmt.Println(skillName)
						skillItem := skillTracker[skillName]
						switch villVal := tempPlayer.Village; villVal{
						case "Leaf":
							skillItem.Leaf++
						case "Sand":
							skillItem.Sand++
						case "Sound":
							skillItem.Sound++
						case "Mist":
							skillItem.Mist++
						case "Stone":
							skillItem.Stone++
						case "Cloud":
							skillItem.Cloud++
						}
						skillItem.TotalCount++
						skillTracker[skillName] = skillItem

						if isCurrent{
							skillItem := currentPlayerSkillTracker[skillName]
							switch villVal := tempPlayer.Village; villVal{
							case "Leaf":
								skillItem.Leaf++
							case "Sand":
								skillItem.Sand++
							case "Sound":
								skillItem.Sound++
							case "Mist":
								skillItem.Mist++
							case "Stone":
								skillItem.Stone++
							case "Cloud":
								skillItem.Cloud++
							}
							skillItem.TotalCount++
							currentPlayerSkillTracker[skillName] = skillItem
						}
					}
				}
			}

			pFile.Close()

			tempPlayer.AllMoney = tempPlayer.Bank + tempPlayer.Wallet

			Players = append(Players, tempPlayer)

			if erri != nil {err = erri}
		}/*else if di.IsDir(){
			info, erri := di.Info()
				loadPlayers(currentPath+"/"+info.Name())
			if erri != nil {err = erri}
		}*/
		return err
	})
}

func sortPlayersByWealth(){
	sort.Slice(Players, func(i, j int) bool {
		return Players[i].AllMoney > Players[j].AllMoney
	  })
}

func sortSkillsByUse(){
	for key := range skillTracker {
        skillKeys = append(skillKeys, key)
    }
    sort.SliceStable(skillKeys, func(i, j int) bool{
        return skillTracker[skillKeys[i]].TotalCount > skillTracker[skillKeys[j]].TotalCount
    })
}

func sortCurrentSkillsByUse(){
	for key := range currentPlayerSkillTracker {
        currentSkillKeys = append(currentSkillKeys, key)
    }
    sort.SliceStable(currentSkillKeys, func(i, j int) bool{
        return currentPlayerSkillTracker[currentSkillKeys[i]].TotalCount > currentPlayerSkillTracker[currentSkillKeys[j]].TotalCount
    })
}

func dirTest(currentPath string){
	filepath.WalkDir(currentPath, func (Fpath string, di fs.DirEntry, err error) error {
		if !di.IsDir(){
			//var tempPlayer Player
			info, erri := di.Info()
			//tempPlayer.Name = info.Name()

			fmt.Println(info.Name())
			if erri != nil {err = erri}
		}
		return err
	})
}

func CheckFinances(){
	sortPlayersByWealth()
	var applicablePlayers = 0
	var totalPlayers = len(Players)
	var totalRyo = 0
	var applicableTotalRyo = 0

	var legacyPlayers = 0
	var legacyPlayerWealth = 0
	var displayKey = "+_+ = Can Buy a House\n ** = Has used Legacy at least once"

	for _, plyr := range Players{
		if plyr.AllMoney >= 50{
			applicablePlayers++
			applicableTotalRyo += plyr.AllMoney
		}
		if plyr.Legacy > 0{
			legacyPlayers++
			legacyPlayerWealth+=plyr.AllMoney
		}
		totalRyo += plyr.AllMoney
	}

	 p := message.NewPrinter(language.English)

	for i:=0; i<applicablePlayers/10; i++{
		displayName := Players[i].Name
		if Players[i].Legacy > 0{
			displayName = "**"+displayName+"**"
		}
		if Players[i].AllMoney > 100000{
			displayName = "+_+"+displayName+"+_+"
		}
		theLine := p.Sprintf("%v.) %v: %v ryo total\n(Wallet: %v ryo;  Bank: %v ryo)", 
			i+1, 
			displayName, 
			number.Decimal(Players[i].AllMoney), 
			number.Decimal(Players[i].Wallet), 
			number.Decimal(Players[i].Bank))
		fmt.Println(theLine)
	}
	reportStr := "Total players: %v\nPlayers counted (>= 50 ryo): %v\nTotal ryo counted: %v\n(Total Ryo not counted: %v)\nAveraged Wealth:%v"
	reportStr += "\n========\nLegacy Players: %v\nLegacy Player Wealth: %v\n(Avg Legacy Wealth: %v)"
	report := p.Sprintf(reportStr,
			number.Decimal(totalPlayers),
			number.Decimal(applicablePlayers),
			number.Decimal(applicableTotalRyo),
			number.Decimal(totalRyo-applicableTotalRyo),
			number.Decimal(applicableTotalRyo/applicablePlayers),
			number.Decimal(legacyPlayers),
			number.Decimal(legacyPlayerWealth),
			number.Decimal(legacyPlayerWealth/legacyPlayers))
	fmt.Println(report)
	fmt.Println(displayKey)
	//fmt.print(sortedPlayers())
}

func JutsuUsageData(){
	
	p := message.NewPrinter(language.English)
	sortSkillsByUse()

	for i:=0; i<100; i++{
		skillName := skillKeys[i]
		skillEntry := skillTracker[skillName]
		reportStr := "%v: %v total\nLeaf\tMist\tStone\tSand\tSound\tCloud\n%v\t%v\t%v\t%v\t%v\t%v"
		report := p.Sprintf(reportStr,
				skillName,
				number.Decimal(skillEntry.TotalCount),
				number.Decimal(skillEntry.Leaf),
				number.Decimal(skillEntry.Mist),
				number.Decimal(skillEntry.Stone),
				number.Decimal(skillEntry.Sand),
				number.Decimal(skillEntry.Sound),
				number.Decimal(skillEntry.Cloud))
		fmt.Println(report)
	}
}

func CurrentUserJutsuUsageData(){
	p := message.NewPrinter(language.English)
	sortCurrentSkillsByUse()


	report := p.Sprintf("Total Active Players: %v", number.Decimal(len(currentPlayers)))
	fmt.Println(report)

	for i:=0; i<100; i++{
		skillName := currentSkillKeys[i]
		skillEntry := currentPlayerSkillTracker[skillName]
		reportStr := "%v: %v total\nLeaf\tMist\tStone\tSand\tSound\tCloud\n%v\t%v\t%v\t%v\t%v\t%v"
		report := p.Sprintf(reportStr,
				skillName,
				number.Decimal(skillEntry.TotalCount),
				number.Decimal(skillEntry.Leaf),
				number.Decimal(skillEntry.Mist),
				number.Decimal(skillEntry.Stone),
				number.Decimal(skillEntry.Sand),
				number.Decimal(skillEntry.Sound),
				number.Decimal(skillEntry.Cloud))
		fmt.Println(report)
	}
}


