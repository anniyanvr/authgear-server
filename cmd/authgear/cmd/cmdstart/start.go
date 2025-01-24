//go:build !authgearlite
// +build !authgearlite

package cmdstart

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/authgear/authgear-server/cmd/authgear/cmd"
	"github.com/authgear/authgear-server/cmd/authgear/server"
)

var cmdStart = &cobra.Command{
	Use:   "start [main|resolver|admin]...",
	Short: "Start specified servers",
	Run: func(cmd *cobra.Command, args []string) {
		ctrl := &server.Controller{}

		serverTypes := args
		if len(serverTypes) == 0 {
			// Default to start all servers
			serverTypes = []string{"main", "resolver", "admin"}
		}
		for _, typ := range serverTypes {
			switch typ {
			case "main":
				ctrl.ServeMain = true
			case "resolver":
				ctrl.ServeResolver = true
			case "admin":
				ctrl.ServeAdmin = true
			default:
				log.Fatalf("unknown server type: %s", typ)
			}
		}

		ctrl.Start(cmd.Context())
	},
}

func init() {
	cmd.Root.AddCommand(cmdStart)
}
