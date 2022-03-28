package worker

import (
	"fmt"
	"time"

	"fyne.io/fyne/widget"
)

func updateTime(clock *widget.Label, clockP *widget.Label, wait time.Duration, worker worker, ch chan channel) {
	for range time.Tick(time.Second) {
		select {
		case <-ch:
			return
		default:
			worker.count++
			clock.SetText(fmt.Sprint(worker.count))
			clockP.SetText(fmt.Sprint(worker.count))
			ch <- channel{worker.count, worker.id}
		}
	}
}

func printCount(worker worker, ch chan channel) {
	for range time.Tick(time.Second) {
		select {
		case <-ch:
			return
		default:
			for elem := range ch {
				fmt.Printf("Worker: %d Cuenta: %d\n", elem.id, elem.count)
			}
		}
	}
}