package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var userInput string

func StartGame() {
	initMap()
	setupRoomInfo()
	setupGameMap()
	printMap()
	printInven()
	for {
		gameStatus := playGame()
		printMap()
		printInven()
		if gameStatus == ClearGame {
			break
		}
	}
	printClearGame()
}

func inputCommand() {
	fmt.Print("-> ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		userInput = scanner.Text()
		fmt.Println("user input: " + userInput)
	}
}
func movePlayer() bool {
	playerY := player[0]
	playerX := player[1]
	info := roomInfo[playerY][playerX]
	doorDirection := string(info[1])
	doorStatus := string(info[2])

	switch {
	case strings.Contains(userInput,"북"):
		if playerY - 1 < 0 {
			return false
		}
		nextRoomInfo := string(roomInfo[playerY-1][playerX])
		nextRoomStatus := string(nextRoomInfo[0])

		if nextRoomStatus =="0" {
			return false
		}

		if doorDirection != "1" || (doorStatus == "1" || doorStatus == "0") {
			player[0] = playerY-1
			return true
		}

	case strings.Contains(userInput,"동"):
		if playerX + 1 > maxRoomColIndex {
			return false
		}
		nextRoomInfo := string(roomInfo[playerY][playerX+1])
		nextRoomStatus := string(nextRoomInfo[0])

		if nextRoomStatus =="0" {
			return false
		}
		
		if doorDirection != "2" || (doorStatus == "1" || doorStatus == "0") {
			player[1] = playerX+1
			return true
		}

	case strings.Contains(userInput,"남"):
		if playerY + 1 > maxRoomRowIndex {
			return false
		}
		nextRoomInfo := string(roomInfo[playerY+1][playerX])
		nextRoomStatus := string(nextRoomInfo[0])

		if nextRoomStatus =="0" {
			return false
		}

		if doorDirection != "3" || (doorStatus == "1" || doorStatus == "0") {
			player[0] = playerY+1
			return true
		}

	case strings.Contains(userInput,"서"):
		if playerX -1 < 0 {
			return false
		}
		nextRoomInfo := string(roomInfo[playerY][playerX-1])
		nextRoomStatus := string(nextRoomInfo[0])

		if nextRoomStatus =="0" {
			return false
		}

		if doorDirection != "4" || (doorStatus == "1" || doorStatus == "0") {
			player[1] = playerX-1
			return true
		}
	}

	return false
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func getItem() {
	info:=roomInfo[player[0]][player[1]]
	item := string(info[3])

	if item == "0" {
		return 
	}

	for i,v := range inven {
		if v =="" {
			inven[i] = item
			runes := []rune(info)
			runes[3] = '0'
			roomInfo[player[0]][player[1]] = RoomInfo(string(runes))
			fmt.Println(string(runes))
			fmt.Println(roomInfo[player[0]][player[1]])
			break
		}
	}
}

// func setDoorStatus(y int, x int, doorStatus string) {

// 	runes := []rune(info)
// 			runes[2] = '0'
// 			roomInfo[y][y] = RoomInfo(string(runes))
// }

func useItem() {
	info:=roomInfo[player[0]][player[1]]
	doorStatus := string(info[2])
	doorDirection := string(info[1])
	var usedItem string
	switch {
	case strings.Contains(userInput, "망치"):
		if doorStatus != "4" || !contains(inven[:], "1") {
			return
		}
		usedItem= "1"
		// 현재 문
		runes := []rune(info)
		runes[2] = '0'
		roomInfo[player[0]][player[1]] = RoomInfo(string(runes))

		// 다음 문
		var nextY = player[0]
		var nextX = player[1]
		if doorDirection =="1" {
			nextY = player[0] - 1
		}else if doorDirection == "2" {
			nextX = player[1] +1
		}else if doorDirection == "3" {
			nextY = player[0] +1
		}else {
			nextX = player[1] -1
		}
		nextInfo := roomInfo[nextY][nextX]
		nextRunes := []rune(nextInfo)
		nextRunes[2] = '0'
		roomInfo[nextY][nextX] = RoomInfo(string(nextRunes))
		
	case strings.Contains(userInput, "열쇠"):
		if doorStatus != "3" || !contains(inven[:], "2") {
			return 
		}
		usedItem="2"
		// 현재 문
		runes := []rune(info)
		runes[2] = '1'
		roomInfo[player[0]][player[1]] = RoomInfo(string(runes))

		// 다음 문
		var nextY = player[0]
		var nextX = player[1]
		if doorDirection =="1" {
			nextY = player[0] - 1
		}else if doorDirection == "2" {
			nextX = player[1] +1
		}else if doorDirection == "3" {
			nextY = player[0] +1
		}else {
			nextX = player[1] -1
		}
		nextInfo := roomInfo[nextY][nextX]
		nextRunes := []rune(nextInfo)
		nextRunes[2] = '1'
		roomInfo[nextY][nextX] = RoomInfo(string(nextRunes))
	}

	if usedItem == "" {
		return
	}

	for i,v := range inven {
		if v == usedItem {
			inven[i] = ""
		}
	}
}

func openDoor() {
	info:=roomInfo[player[0]][player[1]]
	doorStatus := string(info[2])
	doorDirection := string(info[1])

	if doorStatus != "2" {
		return
	}

	// 현재 문
	runes := []rune(info)
	runes[2] = '1'
	roomInfo[player[0]][player[1]] = RoomInfo(string(runes))

	// 다음 문
	var nextY = player[0]
	var nextX = player[1]
	if doorDirection =="1" {
		nextY = player[0] - 1
	}else if doorDirection == "2" {
		nextX = player[1] +1
	}else if doorDirection == "3" {
		nextY = player[0] +1
	}else {
		nextX = player[1] -1
	}
	nextInfo := roomInfo[nextY][nextX]
	nextRunes := []rune(nextInfo)
	nextRunes[2] = '1'
	roomInfo[nextY][nextX] = RoomInfo(string(nextRunes))
}

func handleCommand() {
	var moveCommands = []string{
		"북쪽으로 가","북","북 가",
		"동쪽으로 가","동","동 가",
		"남쪽으로 가","남","남 가",
		"서쪽으로 가","서","서 가"}
	var doorCommands = []string {
		"나무문 열기", "나무문 열어", "나무문 열",
		"유리문 열기", "유리문 열어", "유리문 열",
		"잠긴문 열기", "잠긴문 열어", "잠긴문 열"}
	var useItemCommands = []string {"망치 사용 유리문", "열쇠 사용 잠긴문"}

	switch {
	case contains(moveCommands, userInput):
		movePlayer()
		getItem()
	case contains(useItemCommands, userInput) :
		useItem()
	case contains(doorCommands, userInput) :
		openDoor()
	}
}

type GameStatus int

const (
	ClearGame GameStatus = iota
	Playing
)

func playGame() GameStatus {
	inputCommand()
	handleCommand()
	setupGameMap()
	return isGameEnd()
}

func isGameEnd() GameStatus {
	for _,v := range inven {
		if v == "3" {
			return ClearGame
		}
	}
	return Playing
}
