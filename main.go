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

		// draw the house
		drawHouse(canvas)

		// move to the right so we can draw the next house
		canvas.Move(drawing.Right, 9)
		canvas.Turn(drawing.Left)
	}

	fmt.Println("Name the file:")
	fileName := easyinput.TakeInput()
	if fileName == "" {
		fmt.Println("Must provide a filename!")
		return
	}

	canvas.SaveImage(fileName + ".jpeg")
}

func drawHouse(canvas *drawing.Canvas) {
	canvas.Move(drawing.Straight, 2)

	// draw green door
	canvas.DrawRectangle(drawing.Green, 1, 2)
	canvas.Move(drawing.Backwards, 2)

	canvas.Move(drawing.Right, 3)
	canvas.Move(drawing.Right, 5)

	// draw red frame
	canvas.DrawRectangle(drawing.Red, 5, 5)

	// draw blue roof
	canvas.DrawTriangle(drawing.Blue, 5)

	// move to lower-left corner of house
	canvas.Move(drawing.Backwards, 5)
	canvas.Turn(drawing.Backwards)
}
