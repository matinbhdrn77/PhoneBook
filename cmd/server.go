package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

type Server struct{}

func (cmd Server) Command(trap chan os.Signal) *cobra.Command {
	return nil
}

func (cmd *Server) run(cfg *config.Config, trap chan os.Signal) {

}
