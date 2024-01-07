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
	"slices"
	"encoding/json"
)

var AllItems []Item
var ReleaseItems []Item
var ProtoItems []Item
var Weapons []Item
var Armor []Item
var OtherItems []Item
var RetiredLimitedImmItems = []int{284,2550,1815}

type Item struct{
	Vnum int `json:"vnum"`
	UpgradeVnum int `json:"upgrade_vnum"`
	Name string `json:"name"`
	Type string `json:"type"`
	Layer int `json:"layer"`
	Cost int `json:"cost"`
	FixedCost int `json:"fixed_cost"`
	Weight int `json:"weight"`
	LevelReq int `json:"level_req"`
	LegacyLevelReq int `json:"legacy_level_req"`
	Flags []string `json:"flags"`
	ShortDesc string `json:"short_desc"`
	LongDesc string `json:"long_desc"`
	WearLocs []string `json:"wear_locs"`
	Affects []ItemAffect `json:"affects"`
	Values []int `json:"values"`
	AreaOrigin string `json:"area_origin"`
	ActionDesc string `json:"action_desc"`
}

type ItemAffect struct{
	AffectName string `json:"ItemAffectName"`
	AffectVal int	`json:"ItemAffectValue"`
}



func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func IS_SET(flag int, bit int) bool{
	//I hate that this is prevalent enough in smaug code to need to recreate it here just to translate
	return ((flag) & (bit) == bit)
}

func SaveItem(theItem Item){
	AllItems = append(AllItems, theItem)
	if (slices.Contains(theItem.Flags, "prototype")){
		ProtoItems = append(ProtoItems, theItem)
	}else{
		ReleaseItems = append(ReleaseItems, theItem)
		switch theItem.Type{
		case "weapon":
			Weapons = append(Weapons, theItem)
		case "armor":
			Armor = append(Armor, theItem)
		default:
			OtherItems = append(OtherItems, theItem)
		}
	}
}

func GetWearLocs(thatNum int)[]string{
	var wearLocs []string
	for x:=0;x<32;x++{
		if (IS_SET(thatNum, 1<<x)){
			wearLocs = append(wearLocs, W_flags[x])
		}
	}
	return wearLocs
}

func GetItemFlags(thatNum int)[]string{
	var itemFlags []string
	for x:=0;x<32;x++{
		if (IS_SET(thatNum, 1<<x)){
			itemFlags = append(itemFlags, O_flags[x])
		}
	}
	return itemFlags
}

func nmsplitter(line string, callSize int) []int {
	//fmt.Println("Call size: "+strconv.Itoa(callSize))
	//fmt.Println("Line size: "+strconv.Itoa(len(line)))
	//fmt.Println("Line: " + line)
	spraySize := max(callSize, len(line))
	parsedArray := make([]int, spraySize)
	var err error
	tempSpray := strings.Split(line," ")
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
								tempLine := nmsplitter(itemReader.Text(),4)
								currentItem.Type = O_types[tempLine[0]]
								if (slices.Contains(RetiredLimitedImmItems, tempLine[0])){
									currentItem.Type = "Restricted-"+currentItem.Type
								}
								currentItem.Flags = GetItemFlags(tempLine[1])
								currentItem.WearLocs = GetWearLocs(tempLine[2])
								currentItem.Layer = tempLine[3]
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
							case 8: //Affects
								if(itemReader.Text() == "A"){
									itemReader.Scan()
									tempLine := nmsplitter(itemReader.Text(), 2)
									var tempAffect ItemAffect
									tempAffect.AffectName = A_types[tempLine[0]]
									tempAffect.AffectVal = tempLine[1] 
									currentItem.Affects = append(currentItem.Affects, tempAffect)
								}
								//itemLine++
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

func printBasicStats(){
	p := message.NewPrinter(language.English)

	reportStr := "All items: %v\nReleased Items: %v\nProto Items: %v\nWeapons: %v\nArmor: %v\nAll other items: %v"
	report := p.Sprintf(reportStr,
			len(AllItems),
			len(ReleaseItems),
			len(ProtoItems),
			len(Weapons),
			len(Armor),
			len(OtherItems))
	fmt.Println(report)
}

func printItems(numToPrint int){
	p := message.NewPrinter(language.English)
	//sortSkillsByUse()

	for i:=0; i<numToPrint; i++{
		displayItem := AllItems[i]
		reportStr := "vnum: %v\t->Upgrade_Vnum: %v\nName: %v\nShort Description: %v\nLong Description:%v"
		reportStr += "\nType:%v\tCost/FixedCost: %v/%v\nWear Locs:%v\nFlags:%v\nArea File:%v"
		report := p.Sprintf(reportStr,
				displayItem.Vnum,
				displayItem.UpgradeVnum,
				displayItem.Name,
				displayItem.ShortDesc,
				displayItem.LongDesc,
				displayItem.Type,
				displayItem.Cost,
				displayItem.FixedCost,
				displayItem.WearLocs,
				displayItem.Flags,
				displayItem.AreaOrigin)
		fmt.Println(report)
	}
}

func printAllItems(){
	printItems(len(AllItems))
}

func writeItemsToFilesFullSep(){

}

func writeAllItemsFile(){
	jsonAllItems, _ := json.MarshalIndent(AllItems, ""," ")
	filename := "AllItems.json"
    destination, err := os.Create(filename)
    if err != nil {
        fmt.Println("os.Create:", err)
        return
    }
    defer destination.Close()

    fmt.Fprintf(destination, "%s\n", jsonAllItems)
    //fmt.Fprintf(destination, "Using fmt.Fprintf in %s\n", filename)
}

func writeWeaponsFile(){

}

func writeArmorFile(){

}

func writeOtherItemsFile(){

}
