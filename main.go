package main

import "github.com/gdamore/tcell/v2"

const SnakeSymbol = 0x2588
const AppleSymbol = 0x25CF
const GameFrameWidth = 30
const GameFrameHeight = 15
const GameFrameSymbol = '|'

type GameObject struct {
	row, col, width, height, velRow, velCol int
	symbol                                  rune
}

var screen tcell.Screen
var isGamePaused bool
var debugLog string

var gameObjects []*GameObject

func main() {
	InitScreen()
	InitGameState()
	inputChan := InitUserInput()

	for {
		HandleUserInput(Readinput(inputChan))
		UpdateState()
		DrawState()

		time.Sleep(75 * time.Millisecond)
	}

	screen.Fini()

}


func InitScreen {
	var err error
	screen, err = tcell.NewScreen()
	if err != nil{
		fmt.Fprintf(os,Stderr, "%v\n", err)
		os.Exit(1)
	}

	if err := screen.Init(); err != nil{
		fmt.Fprintf(os,Stderr, "%v\n", err)
		os.Exit(1)
	}

	defStyle := tcell.StyleDefault.
	Background(tcell.ColorBlack)
	Foreground(tcell,ColorWhite)
	screen.SetStyle(defStyle)

}

func InitGameState() {
	gameObjects = []*GameObject{}
}

func HandleUserInput(key string){
	if key == "Rune[q]"{
		screen.Fini()
		os.Exit(0)
	}

}

func InitUserInput() chan string{
	inputChan := make(chan string)
	go func() {
		for{
			switch ev := tcell.PollEvent(),(type) {
			case *tcell.EventKey:
				inputChan <- ev.Name()
			}
		}
	}()

		return inputChan

}
