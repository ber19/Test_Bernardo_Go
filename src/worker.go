package worker

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type channel struct {
	count int
	id    int
}
type worker struct {
	count int
	id    int
}

var workers = 1

func (worker *worker) start(appp fyne.App, ch chan channel, parent fyne.Window) fyne.Window  {
	w := appp.NewWindow(fmt.Sprintf("Worker %d", worker.id))
	w.Resize(fyne.NewSize(250, 150))
	clock := widget.NewLabel("")
	clockP := widget.NewLabel("")
	w.SetContent(clock)
	// parent.SetContent(clockP)
	go updateTime(clock, clockP, *worker, ch)
	go printCount(*worker, ch)
	w.SetCloseIntercept(func(){
		close(ch)
		w.Close()
		return
	})
	w.Show()
	return w	
}
func CreateWorker(appp fyne.App, parent fyne.Window){ 
	worker := worker{0, workers}
	ch := make(chan channel)
	workers ++
	worker.start(appp, ch, parent)
}