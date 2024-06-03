package main

var AreaMaps []AreaMap
var CurrentArea AreaMap

/*====Direction Reference:
[FLAGS KEY ROOMNUM]
D0 = north
D1 = east
D2 = south
D3 = west
D4 = up
D5 = down
D6 = northeast
D7 = northwest
D8 = southeast
D9 = southwest
D10 = special/keyword
=====End Direction Reference*/
type MapRow struct{
	Rooms []RoomData
}

type AreaMap struct{
	Name string
	MapRows []MapRow
}

type RoomData struct{
	Name string
	RoomNumber int
	North int
	East int
	South int
	West int
	Up int
	Down int
	Northeast int
	Southeast int
	Northwest int
	Southwest int
	Special int
	Display string
}


func AddRoomToAreaMap(currentRoom RoomData){
	var north = 0
	var east = 0
	var south = 0
	var west = 0
	var x_coord = -1
	var y_coord = -1

	for row_ind, row := range CurrentArea.MapRows{
		for room_ind, room := range row.Rooms{
			if(room.RoomNumber == currentRoom.North){
				if(row_ind < len(CurrentArea.MapRows)-1){
					nextRow := make([]RoomData, len(row))
					rowAfter := make([]RoomData, len(row))
					nextRow[room_ind].Display = "|"
					rowAfter[room_ind] = currentRoom
					CurrentArea.MapRows = append(CurrentArea.MapRows, nextRow, rowAfter)
				} else {
					CurrentArea.MapRow[row_ind+1][room_ind].Display = "|"
					CurrentArea.MapRow[row_ind+2][room_ind] = currentRoom
				}
			}
		}
	}
}