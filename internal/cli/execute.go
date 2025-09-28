package cli

import "context"

// Execute starts the program.
//
// Execute runs the root command that will start the application. The root
// command will process the command line to determine what action to perform
// such as running a sub-command.
//
// If a command encounters an error, that error is returned to the caller.
func Execute(ctx context.Context) error {
	return rootCommand.ExecuteContext(ctx)
}
