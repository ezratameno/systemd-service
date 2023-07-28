package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	// load command line arguments
	name := flag.String("name", "world", "name to print")
	flag.Parse()

	log.Printf("Starting sleep service for %s", *name)

	// setup signal catching
	sigs := make(chan os.Signal, 1)

	// catch all signals since not explicitly listing
	signal.Notify(sigs)
	//signal.Notify(sigs,syscall.SIGQUIT)

	// method invoked upon seeing signal
	go func() {
		for sig := range sigs {
			log.Printf("RECEIVED SIGNAL: %s", sig)

			switch sig {
			case syscall.SIGURG:
				log.Printf("ignoring sigurg")
			default:
				AppCleanup()
				os.Exit(1)
			} // switch
		} // for

	}()

	// infinite print loop
	for {
		log.Printf("hello %s", *name)

		// wait random number of milliseconds
		Nsecs := rand.Intn(3000)
		log.Printf("About to sleep %dms before looping again", Nsecs)
		time.Sleep(time.Millisecond * time.Duration(Nsecs))
	}
	return nil
}

func AppCleanup() {
	log.Println("CLEANUP APP BEFORE EXIT!!!")
}
