package main

import (
	"fmt"
	"os"
	"os/exec"
)

type RoomInfo string

const (
	maxRoomRowIndex = 5
	maxRoomColIndex = 6
)

var roomInfo [6][7]RoomInfo
var gameMap [13][15]string
var player = [2]int{4,1}
var inven [3]string

func initMap() {
	for y := 0; y < len(roomInfo); y++ {
		for x := 0; x < len(roomInfo[y]); x++ {
			roomInfo[y][x] = "00000"
		}
	}

	for y := 0; y < len(gameMap); y++ {
		for x := 0; x < len(gameMap[y]); x++ {
			gameMap[y][x] = "⬛"
		}
	}
	
}

func setupRoomInfo() {
	// 방유무, 문방향, 문상태, 아이템
	roomInfo[4][1] = "1000"
	roomInfo[3][1] = "1000"
	roomInfo[2][1] = "1240"
	roomInfo[2][0] = "1001"
	roomInfo[2][2] = "1440"
	roomInfo[2][3] = "1000"
	roomInfo[3][3] = "1000"
	roomInfo[4][3] = "1220"
	roomInfo[4][4] = "1420"
	roomInfo[5][4] = "1002"
	roomInfo[1][3] = "1000"
	roomInfo[1][4] = "1120"
	roomInfo[0][4] = "1320"
	roomInfo[0][5] = "1230"
	roomInfo[0][6] = "1433"
}

func convertIndex(roomY int, roomX int) (mapY int, mapX int) {
	return roomY*2 + 1, roomX*2 + 1
}

func setupDoor() {
	for y := 0; y < len(roomInfo); y++ {
		for x := 0; x < len(roomInfo[y]); x++ {
			hasRoom := string(roomInfo[y][x][0])
			doorDirection := string(roomInfo[y][x][1])
			doorStatus := string(roomInfo[y][x][2])

			if hasRoom == "0" || doorDirection == "0" {
				continue
			}

			mapY, mapX := convertIndex(y, x)
			var doorY int
			var doorX int

			if doorDirection == "1" {
				doorY = mapY - 1
				doorX = mapX
			} else if doorDirection == "2" {
				doorY = mapY
				doorX = mapX + 1

			} else if doorDirection == "3" {
				doorY = mapY + 1
				doorX = mapX

			} else {
				doorY = mapY
				doorX = mapX - 1
			}

			if doorStatus == "1" {
				gameMap[doorY][doorX] = "🟫"
			} else if doorStatus == "2" {
				gameMap[doorY][doorX] = "🚪"
			} else if doorStatus == "3" {
				gameMap[doorY][doorX] = "🔒"
			} else if doorStatus == "4" {
				gameMap[doorY][doorX] = "🧊"
			}else {
				gameMap[doorY][doorX] = "🟫"
			}
		}
	}
}

func setupRoom() {
	for y := 0; y < len(roomInfo); y++ {
		for x := 0; x < len(roomInfo[y]); x++ {
			hasRoom := string(roomInfo[y][x][0])
			item := string(roomInfo[y][x][3])

			if hasRoom == "0" {
				continue
			}

			mapY, mapX := convertIndex(y, x)

			// 위왼
			gameMap[mapY-1][mapX-1] = "🟫"
			// 위중 문
			gameMap[mapY-1][mapX] = "🟫"
			// 위오
			gameMap[mapY-1][mapX+1] = "🟫"

			// 중왼
			gameMap[mapY][mapX-1] = "🟫"
			// 중중 아이템
			if item == "1" {
				gameMap[mapY][mapX] = "🔨"
			} else if item == "2" {
				gameMap[mapY][mapX] = "🔑"
			} else if item == "3" {
				gameMap[mapY][mapX] = "👍"
			} else {
				gameMap[mapY][mapX] = "⬜️"
			}
			// 중오
			gameMap[mapY][mapX+1] = "🟫"

			// 아왼
			gameMap[mapY+1][mapX-1] = "🟫"
			// 아중
			gameMap[mapY+1][mapX] = "🟫"
			// 아오
			gameMap[mapY+1][mapX+1] = "🟫"
		}
	}
}

func setupPlayer() {
	mapY, mapX :=convertIndex(player[0],player[1])
	gameMap[mapY][mapX] = "🤔"
}

func setupGameMap() {
	// 바탕
	for y := 0; y < len(gameMap); y++ {
		for x := 0; x < len(gameMap[y]); x++ {
			gameMap[y][x] = "⬛"
		}
	}

	// 벽 및 아이템
	setupRoom()

	// 문
	setupDoor()

	// 플레이어
	setupPlayer()
}

func printMap() {
	clearConsole()
	for y := 0; y < len(gameMap); y++ {
		for x := 0; x < len(gameMap[y]); x++ {
			//  r :=[]rune(string(gameMap[y][x]))[0]
			fmt.Print(gameMap[y][x])
		}
		fmt.Println()
	}
}

func printInven() {
	fmt.Print("items: ")
	for _,item := range inven{
		if item == "0" {
			continue
		}
		if item =="1" {
			fmt.Print("🔨")
		}else if item == "2" {
			fmt.Print("🔑")
		}
	}
	fmt.Println()
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printClearGame() {
	fmt.Println("🎉🎉🎉 Game Clear 🎉🎉🎉")
}
