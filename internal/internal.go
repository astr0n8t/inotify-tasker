package internal

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/fsnotify/fsnotify"

	"github.com/astr0n8t/inotify-tasker/config"
)


// Runs dishook
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
