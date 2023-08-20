package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

func main() {
	const description = "PhoneBook Application"
	root := &cobra.Command{Short: description}

	trap := make(chan os.Signal, 1)
	signal.Notify(trap, syscall.SIGINT, syscall.SIGTERM)

}
