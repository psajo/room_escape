package main

import "strings"

type PlayerInfo struct {
	position [2]int
	items    []ItemType
}

func (gameInfo *GameInfo) getItem(userCommand string) {
	switch {
	case strings.Contains(userCommand, "망치"):
		if gameInfo.removeRoomItem("망치") != -1 {
			gameInfo.player.items = append(gameInfo.player.items, HammerItem)
			PrintGetItem("망치")
		}
	case strings.Contains(userCommand, "열쇠"):
		if gameInfo.removeRoomItem("열쇠") != -1 {
			gameInfo.player.items = append(gameInfo.player.items, KeyItem)
			PrintGetItem("열쇠")
		}
	}
}

func (gameInfo *GameInfo) removeRoomItem(item ItemType) int {
	currentRoom := gameInfo.GetCurrentRoomInfo()
	for i, roomItem := range currentRoom.items {
		if roomItem == item {
			currentRoom.items = append(currentRoom.items[:i], currentRoom.items[i+1:]...)
			return i
		}
	}
	return -1
}
