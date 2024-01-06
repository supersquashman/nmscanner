package main
/*
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

func loadPlayers(currentPath string){
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

			players = append(players, tempPlayer)

			if erri != nil {err = erri}
		}
		return err
	})
}
*/ 