package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/matinbhdrn77/PhoneBook/cmd"
	"github.com/spf13/cobra"
)

func main() {
	const description = "PhoneBook Application"
	root := &cobra.Command{Short: description}

	trap := make(chan os.Signal, 1)
	signal.Notify(trap, syscall.SIGINT, syscall.SIGTERM)

	root.AddCommand(
		cmd.Server{}.Command(trap),
	)

	if err := root.Execute(); err != nil {
		log.Fatalf("failed to execute root command:\n%v", err)
	}
}
