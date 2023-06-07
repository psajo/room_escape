package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var moveCommands = []string{
	"북쪽으로 가", "북", "북 가",
	"동쪽으로 가", "동", "동 가",
	"남쪽으로 가", "남", "남 가",
	"서쪽으로 가", "서", "서 가"}
var openDoorCommands = []string{
	"나무문 열기", "나무문 열어", "나무문 열",
	"유리문 열기", "유리문 열어", "유리문 열",
	"잠긴문 열기", "잠긴문 열어", "잠긴문 열"}
var useItemCommands = []string{"망치 사용 유리문", "열쇠 사용 잠긴문"}

var getItemCommands = []string{"망치 줍", "망치 줍다", "열쇠 줍", "열쇠 줍다"}

type CommandType int

const (
	UnknownCommand CommandType = iota
	MoveCommand
	GetItemCommand
	UseItemCommand
	OpenDoorCommand
)

func InputCommand() string {
	var userCommand string
	fmt.Print("-> ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		userCommand = scanner.Text()
	}
	return userCommand
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (gameinfo *GameInfo) isGoal() bool {
	playerY := gameinfo.player.position[0]
	playerX := gameinfo.player.position[1]
	if playerY == gameinfo.goal[0] && playerX == gameinfo.goal[1] {
		return true
	}
	return false
}

func (gameInfo *GameInfo) HandleCommand(userCommand string) {
	switch {
	case contains(moveCommands, userCommand):
		gameInfo.handleMoveCommand(userCommand)
		if gameInfo.isGoal() {
			gameInfo.gameStatus = GameClear
		}
	case contains(openDoorCommands, userCommand):
		gameInfo.handleOpenDoorCommand(userCommand)
	case contains(useItemCommands, userCommand):
		gameInfo.handleUseItemCommand(userCommand)
	case contains(getItemCommands, userCommand):
		gameInfo.getItem(userCommand)
	}
}

func (gameInfo *GameInfo) handleMoveCommand(userCommand string) {
	var direction = 0
	var position = gameInfo.player.position
	var nextY = position[0]
	var nextX = position[1]
	switch {
	case strings.Contains(userCommand, "북"):
		direction = 0
		nextY -= 1
	case strings.Contains(userCommand, "동"):
		direction = 1
		nextX += 1
	case strings.Contains(userCommand, "남"):
		direction = 2
		nextY += 1
	case strings.Contains(userCommand, "서"):
		direction = 3
		nextX -= 1
	}
	doorInfo := gameInfo.GetCurrentRoomInfo().doorInfos[direction]

	switch doorInfo.doorType {
	case NoDoor:
		gameInfo.moveNextRoom(nextY, nextX)
	case Wall:
		PrintCantMoveForWall()
	case Wood:
		if doorInfo.lock {
			PrintCantMoveForLockWoodDoor()
		} else if doorInfo.open {
			gameInfo.moveNextRoom(nextY, nextX)
		} else {
			PrintCantMoveForCloseWoodDoor()
		}
	case Glass:
		if doorInfo.open {
			PrintMoveNextRoom()
			gameInfo.moveNextRoom(nextY, nextX)
		} else {
			PrintCantMoveForGlassDoor()
		}
	}

}
func (gameInfo *GameInfo) handleOpenDoorCommand(userCommand string) {
	doorInfos := &gameInfo.GetCurrentRoomInfo().doorInfos
	for i, doorInfo := range doorInfos {
		if doorInfo.doorType == Wood {
			doorInfos[i].open = true
			gameInfo.getNextRoomDoor(i).open = true
			PrintOpenWoodDoor()
			break
		}
	}
}
func (gameInfo *GameInfo) handleUseItemCommand(userCommand string) {
	switch {
	case strings.Contains(userCommand, "망치"):
		gameInfo.BreakGlassDoor()
	case strings.Contains(userCommand, "열쇠"):
		gameInfo.UnLockDoor()
	}
}

func (gameInfo *GameInfo) moveNextRoom(nextY int, nextX int) {
	position := &gameInfo.player.position
	position[0] = nextY
	position[1] = nextX
	PrintMoveNextRoom()
}
