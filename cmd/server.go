package cmd

import (
	"os"

	"github.com/matinbhdrn77/PhoneBook/internal/config"
	"github.com/matinbhdrn77/PhoneBook/pkg/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type Server struct{}

func (cmd Server) Command(trap chan os.Signal) *cobra.Command {
	run := func(_ *cobra.Command, _ []string) {
		cmd.run(config.Load(true), trap)
	}

	return &cobra.Command{
		Use:   "server",
		Short: "run PhoneBook server",
		Run:   run,
	}
}

func (cmd *Server) run(cfg *config.Config, trap chan os.Signal) {
	logger := logger.NewZap(cfg.Logger)

	logger.Error("Implement me")

	field := zap.String("signal trap", (<-trap).String())
	logger.Info("exiting by recieving unix signal:", field)
}
