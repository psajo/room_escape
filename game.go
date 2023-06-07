package main

type GameStatus int

const (
	GamePlaying GameStatus = iota
	GameClear
	PlayerDead
)

func StartGame() {
	clearConsole()
	gameInfo := initStage1()
	for gameInfo.gameStatus == GamePlaying {
		gameInfo.PrintGameInfo()
		userCommand := InputCommand()
		clearConsole()
		gameInfo.HandleCommand(userCommand)
	}
	switch gameInfo.gameStatus {
	case GameClear:
		PrintGameClear()
	}
}

type GameInfo struct {
	rooms      [][]*RoomInfo
	goal       [2]int
	gameStatus GameStatus
	player     PlayerInfo
}

func initStage1() GameInfo {
	gameInfo := GameInfo{}
	gameInfo.rooms = newRooms1()
	gameInfo.goal = [2]int{0, 7}
	gameInfo.player = PlayerInfo{position: [2]int{4, 1}}

	return gameInfo
}
