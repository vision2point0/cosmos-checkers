package module

import (
	"time"

	modulev1 "github.com/alice/checkers/api/module/v1"
	"github.com/cosmos/cosmos-sdk/client" // Import transaction types
	"github.com/spf13/cobra"
)

func CmdCreateGame() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-game [player1] [player2]",
		Short: "Create a new checkers game",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := client.GetClientContextFromCmd(cmd) // Get the client context

			req := &modulev1.ReqCheckersTorram{
				Player1:   args[0],
				Player2:   args[1],
				StartTime: time.Now().Unix(),                // current timestamp
				EndTime:   time.Now().Add(time.Hour).Unix(), // example: 1 hour later
			}

			// Example: Sending the request via an RPC call (you may need to adjust this based on your architecture)
			// Here we're assuming you have a function like SendCreateGameRequest to handle this
			_, err := SendCreateGameRequest(ctx, req) // This function would need to be implemented
			if err != nil {
				return err // Return error if the request fails
			}

			cmd.Println("Game created successfully!")
			return nil
		},
	}
	return cmd
}

// Dummy function to illustrate sending a request
func SendCreateGameRequest(ctx client.Context, req *modulev1.ReqCheckersTorram) (*modulev1.ResCheckersTorram, error) {
	// Implement your RPC call logic here, for example, using the context to send the transaction
	// This is just a placeholder for demonstration purposes
	return &modulev1.ResCheckersTorram{Success: true, Message: "Game created!"}, nil
}
