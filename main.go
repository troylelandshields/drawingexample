package main

import (
	"fmt"

	"github.com/troylelandshields/drawing"
	"github.com/troylelandshields/easyinput"
)

func main() {
	canvas := drawing.NewCanvas(2000, 1000)

	canvas.Move(drawing.Left, 10)
	canvas.Turn(drawing.Right)

	numberOfHouses := 5
	for i := 1; i <= numberOfHouses; i++ {
		fmt.Println("Drawing house", i)
		drawHouse(canvas)
	}

	fmt.Println("Name the file:")
	fileName := easyinput.TakeInput()
	canvas.SaveImage(fileName)
}

func drawHouse(canvas *drawing.Canvas) {
	canvas.Move(drawing.Straight, 2)
	canvas.DrawRectangle(drawing.Green, 1, 2)
	canvas.Move(drawing.Backwards, 2)

	canvas.Move(drawing.Right, 3)
	canvas.Move(drawing.Right, 5)

	canvas.DrawRectangle(drawing.Red, 5, 5)
	canvas.DrawTriangle(drawing.Blue, 5)

	canvas.Move(drawing.Backwards, 5)
	canvas.Move(drawing.Left, 9)
	canvas.Turn(drawing.Left)
}
