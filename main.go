package main

import (
	"fmt"
	wa "github.com/radovskyb/watcher"
	"github.com/synw/streload/ws"
	"os"
	"time"
)

var w = wa.New()

func reloader(folderpath string) {
	err := w.AddRecursive(folderpath)
	if err != nil {
		panic("Can not add path " + folderpath)
	}
	w.FilterOps(wa.Write, wa.Create, wa.Move, wa.Remove, wa.Rename)
	// lauch listener
	go func() {
		for {
			select {
			case e := <-w.Event:
				msg := "Change detected in " + e.Path
				ws.SendMsg(msg)
				fmt.Println(msg)
			case err := <-w.Error:
				msg := "Watcher error " + err.Error()
				fmt.Println(msg)
			case <-w.Closed:
				msg := "Watcher closed"
				fmt.Println(msg)
				return
			}
		}
	}()
	fmt.Println("Watching " + folderpath + " ...")
	// start listening
	err = w.Start(time.Millisecond * 200)
	if err != nil {
		panic("Error starting the watcher")
	}

}

func main() {
	filepath := os.Args[1]
	go ws.RunWs()
	reloader(filepath)
}
