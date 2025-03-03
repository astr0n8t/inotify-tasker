package internal

import (
	"log"
	"os"
	"os/signal"

	"github.com/fsnotify/fsnotify"

	"github.com/astr0n8t/inotify-tasker/config"
)

// Kills the current sessioon and tries to start a new one
func Reload(e fsnotify.Event) {
	log.Printf("Config file changed: %v Reloading config", e.Name)
	log.Printf("Reloaded config: %v", e.Name)
}

// Runs inotify-tasker
func Run() {

	// Make sure we can load config
	config := config.Config()
	log.Printf("Loaded config file %v", config.ConfigFileUsed())

	// Add config hot reloading
	config.OnConfigChange(Reload)
	config.WatchConfig()
	log.Printf("Watching config for changes")


	// Don't exit until we receive stop from the OS
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+c to exit")
	<-stop
}
