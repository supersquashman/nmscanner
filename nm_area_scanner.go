package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
	"path/filepath"
	"io/fs"
	"strings"
	//"sort"
	"strconv"
	"golang.org/x/text/language"
	//"golang.org/x/text/number"
    "golang.org/x/text/message"
	//"time"
	"unicode/utf8"
)

var AllItems []Item
var ReleaseItems []Item
var ProtoItems []Item
var Weapons []Item
var Armor []Item
var OtherItems []Item

type Item struct{
	Vnum int
	UpgradeVnum int
	Name string
	Type string
	Layer int
	Cost int
	FixedCost int
	Weight int
	LevelReq int
	LegacyLevelReq int
	Flags []string
	ShortDesc string
	LongDesc string
	WearLocs []string
	Affects []string
	Values []int
	AreaOrigin string
	ActionDesc string
}



func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func SaveItem(theItem Item){
	AllItems = append(AllItems, theItem)
}

func nmsplitter(line string, callSize int) []int {
	//fmt.Println("Call size: "+strconv.Itoa(callSize))
	//fmt.Println("Line size: "+strconv.Itoa(len(line)))
	//fmt.Println("Line: " + line)
	spraySize := max(callSize, len(line))
	parsedArray := make([]int, spraySize)
	var err error
	tempSpray := strings.Split(line,"'")
	for indx, val := range tempSpray{
		parsedArray[indx],err = strconv.Atoi(val)
	}
	if(err!=nil){}
	return parsedArray
}

func loadItemInfo(currentPath string){
	filepath.WalkDir(currentPath, func (Fpath string, di fs.DirEntry, err error) error {
		if !di.IsDir() && !strings.Contains(di.Name(),".bak") && strings.Contains(di.Name(), ".are") {
			
			//var tempPlayer Player
			newVnum := 0
			info, err := di.Info()
			var erri error
			fmt.Println("Opening file: "+info.Name())
			areaName := info.Name()
			//isCurrent := isCurrentPlayer(tempPlayer.Name)
			readingObjects := false
			var currentItem Item

			areaFile, err := os.Open(Fpath)
			if err != nil {
				log.Fatalf("Error opening file: %v", err)
			}
			defer areaFile.Close()

			// Read the file content
			//content, err := ioutil.ReadAll(areaFile)
			itemReader := bufio.NewScanner(areaFile)
			itemReader.Split(bufio.ScanLines)
  
			for itemReader.Scan() {
				if (strings.Contains(itemReader.Text(),"#OBJECTS")){
					readingObjects = true
					itemReader.Scan()
					//newVnum, err = strconv.Atoi(trimFirstRune(itemReader.Text()))
				}else if (strings.Contains(itemReader.Text(),"#ROOMS")){
					readingObjects = false
				}

				//sameObject := true
				itemLine := 0
				 

				for (readingObjects){
					itemReader.Scan()
					//currentVnum = newVnum
					//for (sameObject){
						if (strings.Contains(itemReader.Text(),"#") && !strings.Contains(itemReader.Text(),"#ROOMS")){
							//fmt.Println(itemReader.Text())
							newVnum, err = strconv.Atoi(trimFirstRune(itemReader.Text()))
							//sameObject = false
							SaveItem(currentItem)
							var tempItem Item
							tempItem.Vnum = newVnum
							tempItem.AreaOrigin = areaName
							currentItem = tempItem
							itemLine=0
						}else if(strings.Contains(itemReader.Text(),"#ROOMS")){
							readingObjects=false
						}else{
							switch itemLine{
							case 0:
								currentItem.Name=itemReader.Text()
								itemLine++
							case 1:
								//fmt.Println(itemReader.Text())
								currentItem.ShortDesc=itemReader.Text()
								itemLine++
							case 2:
								currentItem.LongDesc=itemReader.Text()
								itemLine++
							case 3:
								currentItem.ActionDesc=itemReader.Text()
								itemLine++
							case 4: //[type,flags,wearflags, optional:layers]
								tempLine := strings.Split(itemReader.Text(),"")
								//currentItem.Type,erri = strconv.Atoi(tempLine[0])
								//currentItem.Flags,erri = strconv.Atoi(tempLine[1])
								//currentItem.WearLocs,erri = strconv.Atoi(tempLine[2])
								if(len(tempLine)>3){
									currentItem.Layer,erri = strconv.Atoi(tempLine[3])
								}else{
									currentItem.Layer = 0
								}
								itemLine++
							case 5: //values
								//currentItem.Name=itemReader.Text()
								value0,value1,value2,value3,value4 := 0,0,0,0,0
								tempLine := nmsplitter(itemReader.Text(), 5)
								value0 = tempLine[0]
								value1 = tempLine[1]
								value2 = tempLine[2]
								value3 = tempLine[3]
								value4 = tempLine[4]
								currentItem.Values = append(currentItem.Values,value0,value1,value2,value3,value4)
								itemLine++
							case 6: //[cost, fixedcost, level, legacylevel]
								tempLine := nmsplitter(itemReader.Text(), 4)
								currentItem.Cost = tempLine[0]
								currentItem.FixedCost = tempLine[1]
								currentItem.LevelReq = tempLine[2]
								currentItem.LegacyLevelReq = tempLine[3]
								itemLine++
							case 7: //[upgradeVnum, weight]
								tempLine := nmsplitter(itemReader.Text(), 2)
								currentItem.UpgradeVnum = tempLine[0]
								currentItem.Weight = tempLine[1]
								itemLine++
							case 8:
								//currentItem.Name=itemReader.Text()
								itemLine++
							case 9:
								//currentItem.Name=itemReader.Text()
								itemLine++
							case 10:
								//currentItem.Name=itemReader.Text()
								itemLine++
							default:
								//fmt.Println(itemReader.Text())
								itemLine++
							}
						}
						//itemReader.Scan()
					//}
				}
			}
		//fmt.Println(itemReader.Text())

		areaFile.Close()
		if erri != nil {err = erri}
		}/*else if di.IsDir(){
			info, erri := di.Info()
				loadPlayers(currentPath+"/"+info.Name())
			if erri != nil {err = erri}
		}*/
		return err
	})
}

func printItems(){
	
	p := message.NewPrinter(language.English)
	//sortSkillsByUse()

	for i:=0; i<len(AllItems); i++{
		displayItem := AllItems[i]
		//reportStr := "%v: %v vnum\tName\nMist\tStone\tSand\tSound\tCloud\n%v\t%v\t%v\t%v\t%v\t%v"
		reportStr := "vnum: %v\t->Upgrade_Vnum: %v\nName: %v\nShort Description: %v\nLong Description:%v"
		reportStr += "\nCost/FixedCost: %v/%v"
		report := p.Sprintf(reportStr,
				displayItem.Vnum,
				displayItem.UpgradeVnum,
				displayItem.Name,
				displayItem.ShortDesc,
				displayItem.LongDesc,
				displayItem.Cost,
				displayItem.FixedCost)
		fmt.Println(report)
	}
}

