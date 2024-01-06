# imports
import re
# classes
AreaDict = {
    "AreaName": "",
    "VersionNumber": 0,
    "AuthorNames": "",
    "RangesList": [], # [recommended min, recommended max, enforced min, enforced max]
    "ResetMessage": "",
    "FlagsList": [],
    "EconomyList": [], # [number of billions, extra ryo] this probably doesn't matter
    "ClimateList": [], # [Temp 0-4, precipitation 0-4, wind 0-4]
    "MobilesList": [],
    "ObjectsList": [],
    "RoomsList": [],
    "ShopsList": [],
    "RepairsList": [],
    "SpecialsList": []
}

def parseFile(file_data):
    # Area files are dumb and use the same delimiter for Area Sections and the start of each mob/obj/room
    # so I'm changing the base delimiter for Area Sections first
    file_data = re.sub("#([A-Z]+)", r"@\1", file_data)
    # now that the change to the area sections is done, we need to split on that delimiter
    simpleParsing = re.split("@", file_data)
    for x in simpleParsing:
        if re.match("^AREA", x):
            list = re.split(" ", x, 3)
            stripped = re.sub("\n+","",list[3])
            stripped = re.sub("~", "", stripped)
            AreaDict["AreaName"] = stripped
        elif re.match("^VERSION", x):
            list = re.split(" ", x, 1)
            stripped = re.sub("\n+", "", list[1])
            stripped = re.sub("~", "", stripped)
            AreaDict["VersionNumber"] = stripped
        elif re.match("^AUTHOR", x):
            list = re.split(" ", x, 1)
            stripped = re.sub("\n+", "", list[1])
            stripped = re.sub("~", "", stripped)
            AreaDict["AuthorNames"] = stripped
        elif re.match("^RANGES", x):
            list = re.split("\n", x, 3)
            ranges = list[1]
            rangeList = re.split(" ", ranges, 3)
            for rangeValue in rangeList:
                AreaDict["RangesList"].append(rangeValue)
        elif re.match("^RESETMSG", x):
            list = re.split(" ", x, 1)
            stripped = re.sub("\n+", "", list[1])
            stripped = re.sub("~", "", stripped)
            AreaDict["ResetMessage"] = stripped
        elif re.match("^FLAGS", x):
            list = re.split("\n", x, 3)
            flags = list[1]
            flagList = re.split(" ", flags, 3)
            for flagValue in flagList:
                AreaDict["FlagsList"].append(flagValue)
        elif re.match("^ECONOMY", x):
            list = re.split(" ", x, 2)
            billions = list[1]
            AreaDict["EconomyList"].append(billions)
            ryo = re.sub("\n+", "", list[2])
            AreaDict["EconomyList"].append(ryo)
        elif re.match("^CLIMATE", x):
            list = re.split(" ", x, 3)
            temp = list[1]
            AreaDict["ClimateList"].append(temp)
            wet = list[2]
            AreaDict["ClimateList"].append(wet)
            wind = re.sub("\n+", "", list[3])
            AreaDict["ClimateList"].append(wind)
        elif re.match("^MOBILES", x):
            list = re.split("#", x)
            for mob in list:
                if not re.match("^(MOBILES|0)", mob):
                    mobDict = {
                        "vnum": "",
                        "name": "",
                        "short": "",
                        "long": "",
                        "desc": "",
						"actFlags": "",
						"affectedFlags": "",
						"alignment": "",
						"complex": "",
						"level": "",
						"hitroll1": "",
						"armor": "",
						"mobHitdie": "",
						"mobDamdie": "",
						"gold": "",
						"exp": "",
						"pos": "",
						"defpos": "",
						"sex": "",
						"str": "",
						"int": "",
						"wis": "",
						"dex": "",
						"con": "",
						"cha": "",
						"save1": "",
						"save2": "",
						"save3": "",
						"save4": "",
						"save5": "",
						"race": "",
						"class": "",
						"height": "",
						"weight": "",
						"speaks": "",
						"speaking": "",
						"numAttacks": "",
						"hp": "",
						"cp": "",
						"sp": "",
						"hitroll2": "",
						"damroll": "",
						"xflas": "",
						"res": "",
						"imm": "",
						"sus": "",
						"attacks": "",
						"defences": "",
                        "programs": ""
                    }
                    detailsList = re.split("~", mob, 4)
                    mobDict["vnum"] = re.split("\n", detailsList[0])[0]
                    mobDict["name"] = re.split("\n", detailsList[0])[1]
                    mobDict["short"] = re.sub("\n", "", detailsList[1])
                    mobDict["long"] = re.sub("\n", "", detailsList[2])
                    mobDict["desc"] = re.sub("^\n", "", detailsList[3])
                    initialStatBlock = re.split(">", detailsList[4])[0]
                    # stat block
                    statLineList = re.split("\n", initialStatBlock, 8)
                    statLine1 = re.split(" ", statLineList[1], 3)
                    mobDict["actFlags"] = statLine1[0]
                    mobDict["affectedFlags"] = statLine1[1]
                    mobDict["alignment"] = statLine1[2]
                    mobDict["complex"] = statLine1[3]
                    statLine2 = re.split(" ", statLineList[2], 4)
                    mobDict["level"] = statLine2[0]
                    mobDict["hitroll1"] = statLine2[1]
                    mobDict["armor"] = statLine2[2]
                    mobDict["mobHitdie"] = statLine2[3]
                    mobDict["mobDamdie"] = statLine2[4]
                    statLine3 = re.split(" ", statLineList[3], 1)
                    mobDict["gold"] = statLine3[0]
                    mobDict["exp"] = statLine3[1]
                    statLine4 = re.split(" ", statLineList[4], 2)
                    mobDict["pos"] = statLine4[0]
                    mobDict["defpos"] = statLine4[1]
                    mobDict["sex"] = statLine4[2]
                    statLine5 = re.split(" ", statLineList[5], 6)
                    mobDict["str"] = statLine5[0]
                    mobDict["int"] = statLine5[1]
                    mobDict["wis"] = statLine5[2]
                    mobDict["dex"] = statLine5[3]
                    mobDict["con"] = statLine5[4]
                    mobDict["cha"] = statLine5[5]
                    mobDict["lck"] = statLine5[6]
                    statLine6 = re.split(" ", statLineList[6], 6)
                    mobDict["save1"] = statLine6[0]
                    mobDict["save2"] = statLine6[1]
                    mobDict["save3"] = statLine6[2]
                    mobDict["save4"] = statLine6[3]
                    mobDict["save5"] = statLine6[4]
                    statLine7 = re.split(" ", statLineList[7], 9)
                    mobDict["race"] = statLine7[0]
                    mobDict["class"] = statLine7[1]
                    mobDict["height"] = statLine7[2]
                    mobDict["weight"] = statLine7[3]
                    mobDict["speaks"] = statLine7[4]
                    mobDict["speaking"] = statLine7[5]
                    mobDict["numAttacks"] = statLine7[6]
                    mobDict["hp"] = statLine7[7]
                    mobDict["cp"] = statLine7[8]
                    mobDict["sp"] = statLine7[9]
                    statLine8 = re.split(" ", statLineList[8], 7)
                    mobDict["hitroll2"] = statLine8[0]
                    mobDict["damroll"] = statLine8[1]
                    mobDict["xflas"] = statLine8[2]
                    mobDict["res"] = statLine8[3]
                    mobDict["imm"] = statLine8[4]
                    mobDict["sus"] = statLine8[5]
                    mobDict["attacks"] = statLine8[6]
                    mobDict["defences"] = statLine8[7].strip()
                    if re.search(">", detailsList[4]):
                        mobDict["programs"] = (">" + re.split(">", detailsList[4], 1)[1]).rstrip("\n|")

                    AreaDict["MobilesList"].append(mobDict)
        elif re.match("^OBJECTS", x):
            list = re.split("#", x)
            for obj in list:
                if not re.match("^(OBJECTS|0)", obj):
                    objDict = {
                        "vnum": "",
                        "name": "",
                        "short": "",
                        "long": "",
                        "desc": "",
						"type": "",
						"flags": "",
						"wearFlags": "",
						"layers": "",
						"values": [],
						"weight": "",
						"cost": "",
						"rent": "",
						"E": [],
						"A": [],
                        "programs": ""
                    }
                    detailsList = re.split("~", obj, 4)
                    objDict["vnum"] = re.split("\n", detailsList[0])[0]
                    objDict["name"] = re.split("\n", detailsList[0])[1]
                    objDict["short"] = re.sub("\n", "", detailsList[1])
                    objDict["long"] = re.sub("\n", "", detailsList[2])
                    objDict["desc"] = re.sub("^\n", "", detailsList[3])
                    # stat block
                    initialStatBlock = re.split(">", detailsList[4])[0]
                    statLineList = re.split("\n", initialStatBlock, 4)
                    statLine1 = re.split(" ", statLineList[1], 3)
                    objDict["type"] = statLine1[0]
                    objDict["flags"] = statLine1[1]
                    objDict["wearFlags"] = statLine1[2]
                    if len(statLine1) >= 4:
                        objDict["layers"] = statLine1[3]
                    numSpaces = re.findall(" ", statLineList[2])
                    statLine2 = re.split(" ", statLineList[2], len(numSpaces))
                    for value in statLine2:
                        objDict["values"].append(value)
                    statLine3 = re.split(" ", statLineList[3], 3)
                    objDict["weight"] = statLine3[0]
                    objDict["cost"] = statLine3[1]
                    objDict["rent"] = statLine3[2]
                    # Extra Desc and Affects
                    EAList = re.split("\n?A\n", statLineList[4])
                    extraDescList = EAList[0]
                    if extraDescList != "":
                        extraDescGroups = re.split("E\n", extraDescList)
                        for extraDesc in extraDescGroups:
                            if (extraDesc != ""):
                                 extraDescDict = {"keywords": "","desc": ""}
                                 extraDescLines = re.split("~", extraDesc)
                                 extraDescDict["keywords"] = re.sub("\n", "", extraDescLines[0])
                                 extraDescDict["desc"] = re.sub("^\n", "", extraDescLines[1])
                                 objDict["E"].append(extraDescDict)
                    EAList.pop(0)
                    for affect in EAList:
                        affectDict = {"type": "", "value": ""}
                        affectList = re.split(" ", affect)
                        affectDict["type"] = affectList[0]
                        affectDict["value"] = re.sub("\n", "", affectList[1])
                        objDict["A"].append(affectDict)
                    if re.search(">", detailsList[4]):
                        objDict["programs"] = (">" + re.split(">", detailsList[4], 1)[1]).rstrip("\n|")

                    AreaDict["ObjectsList"].append(objDict)
        elif re.match("^ROOMS", x):
            list = re.split("\n#", x)
            for room in list:
                if not re.match("^(ROOMS|0)", room):
                    roomDict = {
                        "vnum": "",
                        "name": "",
                        "desc": "",
                        "area": "",
                        "flags": "",
                        "type": "",
                        "light": "",
                        "teleDelay": "",
                        "teleVnum": "",
                        "tunnel": "",
                        "D": [], # Exits (Doors)
                        "R": [], # Resets
                        "E": [], # Extra descs
                        "programs": ""
                    }
                    detailsList = re.split("~\n", re.sub("\nS$", "", room), 2)
                    roomInfoGroup = detailsList[0]
                    roomDict["vnum"] = re.split("\n", roomInfoGroup)[0]
                    roomDict["name"] = re.split("\n", roomInfoGroup)[1]
                    roomDict["desc"] = re.sub("\n$", "", detailsList[1])
                    # Stats and Exits
                    detProList = re.split("\n>", re.sub("^\n", "", detailsList[2]), 1) # split on programs
                    detailsGroup = detProList[0]
                    SDREList = re.split("\nE\n", detailsGroup) # split on extra descs
                    SDRGroup = SDREList[0]
                    SDRList = re.split("\nR", SDRGroup) # split on Resets
                    SDGroup = SDRList[0]
                    SDList = re.split("\nD", SDGroup) # split on Exits
                    statsList = re.split(" ", SDList[0])
                    roomDict["area"] = statsList[0]
                    roomDict["flags"] = statsList[1]
                    roomDict["type"] = statsList[2]
                    roomDict["light"] = statsList[3]
                    roomDict["teleDelay"] = statsList[4]
                    roomDict["teleVnum"] = statsList[5]
                    roomDict["tunnel"] = statsList[6]
                    SDList.pop(0)
                    # Exits list
                    if SDList != []:
                        for D in SDList:
                            exitDict = {"number": "", "desc": "", "keywords": "", "locks": "", "key": "", "toVnum": ""}
                            exitList = re.split("~", D)
                            exitStart = re.split("\n", exitList[0])
                            exitDict["number"] = re.sub("\n$", "", exitStart[0])
                            exitDict["desc"] = re.sub("\n$", "", exitStart[1], flags=re.MULTILINE)
                            exitDict["keywords"] = re.sub("\n", "", exitList[1])
                            exitFlagList = re.split(" ", exitList[2])
                            exitDict["locks"] = re.sub("\n", "", exitFlagList[0])
                            exitDict["key"] = exitFlagList[1]
                            exitDict["toVnum"] = exitFlagList[2]
                            roomDict["D"].append(exitDict)
                    SDRList.pop(0)
                    # Reset List
                    if SDRList != []:
                        for R in SDRList:
                            roomDict["R"].append("R" + R + "\n")
                    SDREList.pop(0)
                    # Extra desc list
                    if SDREList != []:
                        for E in SDREList:
                            extraDescDict = {"keywords": "", "desc": ""}
                            extraDescLines = re.split("~", E)
                            extraDescDict["keywords"] = re.sub("\n", "", extraDescLines[0])
                            extraDescDict["desc"] = re.sub("^\n", "", extraDescLines[1])
                            roomDict["E"].append(extraDescDict)
                    # Programs
                    if len(detProList) > 1:
                        roomDict["programs"] = (">" + detProList[1]).rstrip("\n|")

                    AreaDict["RoomsList"].append(roomDict)
        elif re.match("^SHOPS", x):
            list = re.split("\n", x)
            for shop in list:
                if not re.match("^(SHOPS|0)", shop):
                    if shop != "":
                        shopDict = {"keeperVnum": "", "trade": [], "profit": "", "profitSell": "",
                                    "openHour": "", "closeHour": "", "comment": ""}
                        shopMatch = re.match(
                            "(?P<keeperVnum>.{5})  (?P<trade0>.{3})(?P<trade1>.{3})(?P<trade2>.{3})(?P<trade3>.{3})"
                            "(?P<trade4>.{3})  (?P<profit>.{4})(?P<profitSell>.{4}) {7}(?P<open>.{3})(?P<close>.{3})"
                            " {4}; (?P<comment>.*)", shop)
                        shopDict["keeperVnum"] = shopMatch.group('keeperVnum').strip()
                        shopDict["trade"].append(shopMatch.group('trade0').strip())
                        shopDict["trade"].append(shopMatch.group('trade1').strip())
                        shopDict["trade"].append(shopMatch.group('trade2').strip())
                        shopDict["trade"].append(shopMatch.group('trade3').strip())
                        shopDict["trade"].append(shopMatch.group('trade4').strip())
                        shopDict["profit"] = shopMatch.group('profit').strip()
                        shopDict["profitSell"] = shopMatch.group('profitSell').strip()
                        shopDict["openHour"] = shopMatch.group('open').strip()
                        shopDict["closeHour"] = shopMatch.group('close').strip()
                        shopDict["comment"] = shopMatch.group('comment').strip()
                        AreaDict["ShopsList"].append(shopDict)
        elif re.match("^REPAIRS", x):
            list = re.split("\n", x)
            for repair in list:
                if not re.match("^(REPAIRS|0)", repair):
                    if repair != "":
                        repairDict = {"keeperVnum": "", "trade": [], "profit": "", "profitSell": "",
                                    "openHour": "", "closeHour": "", "comment": ""}
                        repairMatch = re.match(
                            "(?P<keeperVnum>.{5})  (?P<trade0>.{3})(?P<trade1>.{3})(?P<trade2>.{3})(?P<trade3>.{3})"
                            "(?P<trade4>.{3})  (?P<profit>.{4})(?P<profitSell>.{4}) {7}(?P<open>.{3})(?P<close>.{3})"
                            " {4}; (?P<comment>.*)", repair)
                        repairDict["keeperVnum"] = repairMatch.group('keeperVnum').strip()
                        repairDict["trade"].append(repairMatch.group('trade0').strip())
                        repairDict["trade"].append(repairMatch.group('trade1').strip())
                        repairDict["trade"].append(repairMatch.group('trade2').strip())
                        repairDict["trade"].append(repairMatch.group('trade3').strip())
                        repairDict["trade"].append(repairMatch.group('trade4').strip())
                        repairDict["profit"] = repairMatch.group('profit').strip()
                        repairDict["profitSell"] = repairMatch.group('profitSell').strip()
                        repairDict["openHour"] = repairMatch.group('open').strip()
                        repairDict["closeHour"] = repairMatch.group('close').strip()
                        repairDict["comment"] = repairMatch.group('comment').strip()
                        AreaDict["RepairsList"].append(repairDict)
        elif re.match("^SPECIALS", x):
            list = re.split("\n", x)
            for special in list:
                if not re.match("^(SPECIALS|0)", special):
                    if special != "" and special != "S" and special != "#$":
                        AreaDict["SpecialsList"].append(special)

    #print(AreaDict)
    return AreaDict

def encodeFile(modAreaDict):
    # Area, Version, Author, Ranges, RegetMsg, Flags, Economy, and Climate
    rangesFormat = " ".join(modAreaDict["RangesList"])
    flagsFormat = " ".join(modAreaDict["FlagsList"])
    economyFormat = " ".join(modAreaDict["EconomyList"])
    climateFormat = " ".join(modAreaDict["ClimateList"])
    firstSection = "#AREA   {}~\n\n\n\n#VERSION {}\n#AUTHOR {}~\n\n#RANGES\n{}\n$\n\n#RESETMSG {}~\n\n" \
                   "#FLAGS\n{}\n\n#ECONOMY {}\n\n#CLIMATE {}\n\n".format(
        modAreaDict["AreaName"], modAreaDict["VersionNumber"], modAreaDict["AuthorNames"], rangesFormat,
        modAreaDict["ResetMessage"], flagsFormat, economyFormat, climateFormat)
    # Mobiles
    mobList = []
    for mob in modAreaDict["MobilesList"]:
        statBlockFormat = "{} {} {} {}\n{} {} {} {} {}\n{} {}\n{} {} {}\n{} {} {} {} {} {} {}\n{} {} {} {} {}\n" \
                          "{} {} {} {} {} {} {} {} {} {}\n{} {} {} {} {} {} {} {}".format(
            mob["actFlags"], mob["affectedFlags"],mob["alignment"], mob["complex"],
            mob["level"], mob["hitroll1"],mob["armor"], mob["mobHitdie"],
            mob["mobDamdie"], mob["gold"], mob["exp"], mob["pos"], mob["defpos"],
            mob["sex"], mob["str"], mob["int"], mob["wis"], mob["dex"], mob["con"],
            mob["cha"], mob["lck"], mob["save1"], mob["save2"], mob["save3"],
            mob["save4"], mob["save5"], mob["race"], mob["class"], mob["height"],
            mob["weight"], mob["speaks"], mob["speaking"], mob["numAttacks"], mob["hp"],
            mob["cp"], mob["sp"], mob["hitroll2"], mob["damroll"], mob["xflas"],
            mob["res"], mob["imm"], mob["sus"], mob["attacks"], mob["defences"])
        # Programs
        programs = mob["programs"]
        if programs != "":
            programs += "\n|\n"
        # Combine it all
        individualMob = "#{}\n{}~\n{}~\n{}\n~\n{}~\n{}\n{}".format(
            mob["vnum"], mob["name"], mob["short"], mob["long"], mob["desc"], statBlockFormat, programs)
        mobList.append(individualMob)
    allMobs = "".join(mobList)
    mobSection = "#MOBILES\n{}#0\n\n\n".format(allMobs)
    # Objects
    objList = []
    for obj in modAreaDict["ObjectsList"]:
        statBlockFormat = "{} {} {} {}\n{} {} {} {}\n{} {} {}".format(
            obj["type"], obj["flags"], obj["wearFlags"], obj["layers"],
            obj["values"][0], obj["values"][1], obj["values"][2], obj["values"][3],
            obj["weight"], obj["cost"], obj["rent"])
        EAFormat = ""
        # E
		#below line might be needed later, but I don't think so
        #if "E" in obj["statBlock"]:
        if "E" in obj:
            extraDescList = obj["E"]
            E = ""
            if extraDescList != "":
                eList = []
                for extraDesc in extraDescList:
                    extraDescFormat = "E\n{}~\n{}~\n".format(extraDesc["keywords"], extraDesc["desc"])
                    eList.append(extraDescFormat)
                E = "".join(eList)
            EAFormat += E
        # A
        if "A" in obj:
            affectsList = obj["A"]
            A = ""
            if affectsList != "":
                aList = []
                for affect in affectsList:
                    affectFormat = "A\n{} {}\n".format(affect["type"], affect["value"])
                    aList.append(affectFormat)
                A = "".join(aList)
            EAFormat += A
        # Programs
        programs = obj["programs"]
        if programs != "":
            programs += "\n|\n"
        # Combine it all
        individualObj = "#{}\n{}~\n{}~\n{}~\n{}~\n{}\n{}{}".format(
            obj["vnum"], obj["name"], obj["short"], obj["long"], obj["desc"], statBlockFormat, EAFormat, programs)
        objList.append(individualObj)
    allObjs = "".join(objList)
    objSection = "#OBJECTS\n{}#0\n\n\n".format(allObjs)
    # Rooms
    roomList = []
    for room in modAreaDict["RoomsList"]:
        statBlockFormat = "{} {} {} {} {} {} {}".format(
            room["area"], room["flags"], room["type"], room["light"], room["teleDelay"], room["teleVnum"],
            room["tunnel"])
        DREFormat = ""
        # D
        if "D" in room:
            exitList = room["D"]
            D = ""
            if exitList != "":
                dList = []
                for exit in exitList:
                    exitFormat = "D{}\n{}~\n{}~\n{} {} {}\n".format(
                        exit["number"], exit["desc"], exit["keywords"], exit["locks"], exit["key"], exit["toVnum"])
                    dList.append(exitFormat)
                D = "".join(dList)
            DREFormat += D
        # R
        if "R" in room:
            resetList = room["R"]
            R = ""
            if extraDescList != "":
                R = "".join(resetList)
            DREFormat += R
        # E
        if "E" in room:
            extraDescList = room["E"]
            E = ""
            if extraDescList != "":
                eList = []
                for extraDesc in extraDescList:
                    extraDescFormat = "E\n{}~\n{}~\n".format(extraDesc["keywords"], extraDesc["desc"])
                    eList.append(extraDescFormat)
                E = "".join(eList)
            DREFormat += E
        # Programs
        programs = room["programs"]
        if programs != "":
            programs += "\n|\n"
        # Combine it all
        individualRoom = "#{}\n{}~\n{}\n~\n{}\n{}{}S\n".format(
            room["vnum"], room["name"], room["desc"], statBlockFormat, DREFormat, programs)
        roomList.append(individualRoom)
    allRooms = "".join(roomList)
    roomSection = "#ROOMS\n{}#0\n\n\n".format(allRooms)
    # Shops
    shopsList = []
    for shop in modAreaDict["ShopsList"]:
        shopsFormat = "{:>5}  {:>3}{:>3}{:>3}{:>3}{:>3}  {:>4}{:>4}       {:>3}{:>3}    ; {}\n".format(
            shop["keeperVnum"], shop["trade"][0], shop["trade"][1], shop["trade"][2], shop["trade"][3],
            shop["trade"][4], shop["profit"], shop["profitSell"], shop["openHour"], shop["closeHour"],
            shop["comment"])
        shopsList.append(shopsFormat)
    allShops = "".join(shopsList)
    # Repairs
    repairsList = []
    for repair in modAreaDict["RepairsList"]:
        repairsFormat = "{:>5}  {:>3}{:>3}{:>3}{:>3}{:>3}  {:>4}{:>4}       {:>3}{:>3}    ; {}\n".format(
            repair["keeperVnum"], repair["trade"][0], repair["trade"][1], repair["trade"][2], repair["trade"][3],
            repair["trade"][4], repair["profit"], repair["profitSell"], repair["openHour"], repair["closeHour"],
            repair["comment"])
        repairsList.append(repairsFormat)
    allRepairs = "".join(repairsList)
    # Specials
    allSpecials = "".join(modAreaDict["SpecialsList"])
    lastSection = "#SHOPS\n{}0\n\n\n#REPAIRS\n{}0\n\n\n#SPECIALS\n{}S\n\n\n#$".format(allShops, allRepairs, allSpecials)
    completeFile = firstSection+mobSection+objSection+roomSection+lastSection
    return completeFile

# other