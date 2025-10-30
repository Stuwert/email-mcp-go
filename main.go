package main

import (
	mcp "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
	"github.com/stuwert/email-mcp/internal"
	"github.com/stuwert/email-mcp/internal/app"
)

// app.RegisterTools()
// // Loops over generic tools in array/slice, registers them with the appropriate resource.
// // Generic Tools Struct that calls register tool under the hood when it's attached.

// app.RegisterResources()
// app.RegisterPrompts()
// app.ConnectToDatabase()

func main() {
	// The point of main:
	// 1. Handle deferring instantiation of external resources (database)
	// 2. Parse Environment Variables as necessary
	// 3. Execute the

	done := make(chan struct{})

	// 1. Initiate server
	// 2. Initiate connection to database
	// 3. Register Tools
	// 4. Register Resource
	// 5. Register Prompts

	server := mcp.NewServer(stdio.NewStdioServerTransport())

	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	internal.RegisterPrompts(server, app)
	internal.RegisterResources(server, app)
	internal.RegisterTools(server, app)

	<-done
}
