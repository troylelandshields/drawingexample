package main

import "github.com/troylelandshields/drawing"

func main() {
	canvas := drawing.NewCanvas(1000, 1000)

	canvas.Move(drawing.Straight, 2)
	canvas.DrawRectangle(drawing.Green, 1, 2)
	canvas.Move(drawing.Backwards, 2)

	canvas.Move(drawing.Right, 3)
	canvas.Move(drawing.Right, 5)

	canvas.DrawRectangle(drawing.Red, 5, 5)
	canvas.DrawTriangle(drawing.Blue, 5)

	canvas.SaveImage("img.png")
}
