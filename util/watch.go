package util

import (
    "log"
    "fmt"
    "github.com/fsnotify/fsnotify"
)

func StartWatcher(config Config) {

    fmt.Println(config)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				//log.Println("event:", event)
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("File created:", event.Name)
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("Write:", event.Name)
				}
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					log.Println("File removed:", event.Name)
				}
				if event.Op&fsnotify.Rename == fsnotify.Rename {
					log.Println("File renamed:", event.Name)
				}
				//if event.Op&fsnotify.Chmod == fsnotify.Chmod {
					//log.Println("Chmod:", event.Name)
				//}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(config.Paths.INpath)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
