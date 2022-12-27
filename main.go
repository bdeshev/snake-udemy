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

func UpdateState() {
	if isGamePaused{
		return
	}
	for i := range gameObjects {
		gameObjects[i].row += gameObjects[i].velRow
		gameObjects[i].col += gameObjects[i].velCol
	}
}

func DrawState() {
	if isGamePaused {
		return
	}

	screen.Clear()
	PrintString(0, 0, debugLog)
	PrintGameFrame()
	
	for _,obj := range gameObjects {
		PrintFilledRect(obj.row, obj.col, obj.width, obj.height, obj.symbol)
}
	screen.Show()
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

func Readinput(inputChan chan string) string{
	var key string
	select {
	case key = <-inputChan:
	default:
		key = ""
	}

	return key

}

func PrintGameFrame(){
	screenWidth, screenHeight := screen.Size()
	row,col := screenHeight/2 - GameFrameHeight/2 - 1, screenWidth/2 - GameFrameWidth/2 - 1 
	width, jeight := GameFrameWidth+2, GameFrameHeight+2

	PrintUnfilledRect(row, col, width, height, GameFrameSymbol)
}

func PrintStringCentered(row, col int, str string) {
	col = col - len(str)/2
	PrintString(row, col, str)
}

func PrintString(row, col int, str string) {
	for _,c := range str {
		PrintFilledRect(row, col, 1, 1, c)
		col += 1
	}
}

func PrintFilledRect(row, col, width, height int, ch rune){
	for r := 0; r < height; r++{
		for c := 0; c < width; c++ {
			screen.SetContent(col+c, row+r, ch, nil, tcell.StyleDefault)
		}
	}
}

func PrintUnfiledRect(row,col,width,height int, ch rune){
	for c := 0; c < width; c++{
		screen.SetContent(col+c, row, ch, nil, tcell.StyleDefault)
	}
	
	for c := 0; c < width; c++{
		screen.SetContent(col+c, row+height-1, ch, nil, tcell.StyleDefault)
	}

	for c := 0; c < width; c++{
		screen.SetContent(col, row, ch, nil, tcell.StyleDefault)
		screen.SetContent(col+width-1, row, ch, nil, tcell.StyleDefault)
	}
}