package main

import (
	"fmt"
	"os"
	"os/exec"
)

func PrintGameClear() {
	fmt.Println("탈출 성공!")
}

func PrintPlayerDead() {
	fmt.Println("플레이어 사망!")
}

func (gameInfo *GameInfo) PrintGameInfo() {
	gameInfo.GetCurrentRoomInfo()
	y := gameInfo.player.position[0]
	x := gameInfo.player.position[1]
	currentRoom := gameInfo.rooms[y][x]
	fmt.Println("==================== 방 정보 ====================")
	printAllDirections(currentRoom.doorInfos)
	printRoomItems(currentRoom.items)
	fmt.Println("=================================================")
	fmt.Println("==================== 내 정보 ====================")
	printPlayerItems(gameInfo.player.items)
	fmt.Println("=================================================")
}

func printPlayerItems(items []ItemType) {
	fmt.Print("내 아이템: ")
	if len(items) == 0 {
		fmt.Print("없음")
	} else {
		fmt.Printf("%v", items)
	}
	fmt.Println()
}

func sprintDoor(doorInfo DoorInfo) string {
	var text = ""
	switch doorInfo.doorType {
	case Wall:
		text = "벽"
	case NoDoor:
		text = "다음방"
	case Wood:
		if doorInfo.lock {
			text = "잠긴문"
		} else if doorInfo.open {
			text = "열린나무문"
		} else {
			text = "닫힌나무문"
		}
	case Glass:
		if doorInfo.open {
			text = "깨진유리문"
		} else {
			text = "유리문"
		}
	}
	return text
}

func printAllDirections(doorInfos [4]DoorInfo) {
	fmt.Println("북: " + sprintDoor(doorInfos[0]))
	fmt.Println("동: " + sprintDoor(doorInfos[1]))
	fmt.Println("남: " + sprintDoor(doorInfos[2]))
	fmt.Println("서: " + sprintDoor(doorInfos[3]))
}

func printRoomItems(items []ItemType) {
	fmt.Print("방 아이템: ")
	if items == nil {
		fmt.Print("없음")
	} else if len(items) == 0 {
		fmt.Print("없음")
	} else {
		fmt.Printf("%v", items)
	}
	fmt.Println()
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func PrintMoveNextRoom() {
	fmt.Println("다음 방으로 이동했다.")
}

func PrintCantMoveForWall() {
	fmt.Println("막혀 있다.")
}

func PrintCantMoveForLockWoodDoor() {
	fmt.Println("잠긴문이 막고 있다.")
}

func PrintCantMoveForCloseWoodDoor() {
	fmt.Println("나무문이 닫혀 있다.")
}

func PrintCantMoveForGlassDoor() {
	fmt.Println("유리문이 막고 있다.")
}

func PrintBreakGlassDoor() {
	fmt.Println("유리문을 파괴 했다.")
}

func PrintGetItem(item string) {
	fmt.Println(item + "획득.")
}

func PrintUnLockDoor() {
	fmt.Println("잠긴문을 열었다.")
}

func PrintOpenWoodDoor() {
	fmt.Println("나무문을 열었다.")
}
