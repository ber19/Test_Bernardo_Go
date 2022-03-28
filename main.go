package main

import (
	worker "bernardo/test/src"
	"fmt"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()
	w := app.NewWindow("Master")
	w.SetMaster()
	w.Resize(fyne.NewSize(250, 100))
	w.CenterOnScreen()

	w.SetContent(widget.NewButton("Worker", func(){worker.CreateWorker(app, w)}))

	w.Show()
	app.Run()
	fmt.Println("Proceso finalizado")
	time.Sleep(5 * time.Second)
}