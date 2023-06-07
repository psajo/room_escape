package main

type RoomInfo struct {
	doorInfos [4]DoorInfo // 북동남서
	items     []ItemType
}

type DoorType string

type DoorInfo struct {
	doorType DoorType
	open     bool
	lock     bool
}

const (
	NoDoor DoorType = ""
	Wall   DoorType = "벽"
	Wood   DoorType = "나무문"
	Glass  DoorType = "유리문"
)

func newNoDoor() DoorInfo {
	return DoorInfo{doorType: NoDoor, open: true, lock: false}
}
func newWall() DoorInfo {
	return DoorInfo{doorType: Wall, open: false, lock: true}
}
func newWoodDoor() DoorInfo {
	return DoorInfo{doorType: Wood, open: false, lock: false}
}
func newLockedWoodDoor() DoorInfo {
	return DoorInfo{doorType: Wood, open: false, lock: true}
}
func newGlassDoor() DoorInfo {
	return DoorInfo{doorType: Glass, open: false, lock: true}
}

func newRooms1() [][]*RoomInfo {
	rooms := make([][]*RoomInfo, 6)

	for i := range rooms {
		rooms[i] = make([]*RoomInfo, 8)
	}

	rooms[4][1] = &RoomInfo{
		doorInfos: [4]DoorInfo{newNoDoor(), newWall(), newWall(), newWall()},
	}
	rooms[3][1] = &RoomInfo{doorInfos: [4]DoorInfo{
		newNoDoor(), newWall(), newNoDoor(), newWall()}}
	rooms[2][1] = &RoomInfo{doorInfos: [4]DoorInfo{newWall(), newGlassDoor(), newNoDoor(), newNoDoor()}}
	rooms[2][0] = &RoomInfo{
		doorInfos: [4]DoorInfo{newWall(), newNoDoor(), newWall(), newWall()},
		items:     []ItemType{HammerItem}}
	rooms[2][2] = &RoomInfo{
		doorInfos: [4]DoorInfo{newWall(), newNoDoor(), newWall(), newGlassDoor()},
	}
	rooms[2][3] = &RoomInfo{
		doorInfos: [4]DoorInfo{newNoDoor(), newWall(), newNoDoor(), newNoDoor()},
	}
	rooms[1][3] = &RoomInfo{
		doorInfos: [4]DoorInfo{newWall(), newNoDoor(), newNoDoor(), newWall()},
	}
	rooms[3][3] = &RoomInfo{
		doorInfos: [4]DoorInfo{newNoDoor(), newWall(), newNoDoor(), newWall()},
	}
	rooms[4][3] = &RoomInfo{
		doorInfos: [4]DoorInfo{newNoDoor(), newWoodDoor(), newWall(), newWall()},
	}
	rooms[4][4] = &RoomInfo{
		doorInfos: [4]DoorInfo{newWall(), newWall(), newNoDoor(), newWoodDoor()},
	}
	rooms[5][4] = &RoomInfo{
		doorInfos: [4]DoorInfo{newNoDoor(), newWall(), newWall(), newWall()},
		items:     []ItemType{KeyItem}}
	rooms[1][4] = &RoomInfo{doorInfos: [4]DoorInfo{newWoodDoor(), newWall(), newWall(), newNoDoor()}}
	rooms[0][4] = &RoomInfo{
		doorInfos: [4]DoorInfo{newWall(), newNoDoor(), newWoodDoor(), newWall()},
	}
	rooms[0][5] = &RoomInfo{
		doorInfos: [4]DoorInfo{newWall(), newNoDoor(), newWall(), newNoDoor()},
	}
	rooms[0][6] = &RoomInfo{
		doorInfos: [4]DoorInfo{newWall(), newLockedWoodDoor(), newWall(), newNoDoor()},
	}
	rooms[0][7] = &RoomInfo{
		doorInfos: [4]DoorInfo{newWall(), newWall(), newWall(), newLockedWoodDoor()},
	}

	return rooms
}

func (gameInfo *GameInfo) GetCurrentRoomInfo() *RoomInfo {
	y := gameInfo.player.position[0]
	x := gameInfo.player.position[1]
	currentRoom := gameInfo.rooms[y][x]
	return currentRoom
}

func (gameInfo *GameInfo) BreakGlassDoor() {
	var hammerIndex = -1
	for i, item := range gameInfo.player.items {
		if item == HammerItem {
			hammerIndex = i
			break
		}
	}
	if hammerIndex == -1 {
		return
	}

	doorInfos := &gameInfo.GetCurrentRoomInfo().doorInfos
	for i, doorInfo := range doorInfos {
		if doorInfo.doorType == Glass {
			doorInfos[i].lock = false
			doorInfos[i].open = true
			nextDoor := gameInfo.getNextRoomDoor(i)
			nextDoor.lock = false
			nextDoor.open = true
			gameInfo.player.items = append(gameInfo.player.items[:hammerIndex], gameInfo.player.items[hammerIndex+1:]...)
			PrintBreakGlassDoor()
		}
	}
}

func (gameInfo *GameInfo) getNextRoomDoor(doorDirection int) *DoorInfo {
	nextY := gameInfo.player.position[0]
	nextX := gameInfo.player.position[1]
	nextDoorDirection := (doorDirection + 2) % 4
	switch doorDirection {
	case 0: // 북
		nextY -= 1
	case 1: // 동
		nextX += 1
	case 2: // 서
		nextX -= 1
	case 3: //
		nextY += 1
	}
	return &gameInfo.rooms[nextY][nextX].doorInfos[nextDoorDirection]
}

func (gameInfo *GameInfo) UnLockDoor() {
	var keyItemIndex = -1
	for i, item := range gameInfo.player.items {
		if item == KeyItem {
			keyItemIndex = i
			break
		}
	}
	if keyItemIndex == -1 {
		return
	}

	doorInfos := &gameInfo.GetCurrentRoomInfo().doorInfos
	for i, doorInfo := range doorInfos {
		if doorInfo.doorType == Wood {
			doorInfos[i].lock = false
			doorInfos[i].open = true
			nextDoor := gameInfo.getNextRoomDoor(i)
			nextDoor.lock = false
			nextDoor.open = true
			gameInfo.player.items = append(gameInfo.player.items[:keyItemIndex], gameInfo.player.items[keyItemIndex+1:]...)
			PrintUnLockDoor()
		}
	}
}
