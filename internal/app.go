package internal

import (
	mcp "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

type MyFunctionArguments struct {
	Submitter string  `json:"submitter" jsonschema:"required,description=The name of the thing calling this tool (openai, google, claude, etc)"`
	Content   Content `json:"content" jsonschema:"required,description=The content of the message"`
}

type Content struct {
	Title       string  `json:"title" jsonschema:"required,description=The title to submit"`
	Description *string `json:"description" jsonschema:"description=The description to submit"`
}

// type Tool struct {
// 	Hello HelloTool
// }

// App Owns associating the handlers. Resource handler, prompt handler, tools handler
// Main instruments the routes during the boot step.
// Routes.SetupHandlers(app)

type Application struct{}

func NewApplication() (err error) {

	server := mcp.NewServer(stdio.NewStdioServerTransport())

	// Two Questions.
	// 1. Why does it think it's redefining the error down below?
	// 2. Why does starting the app struggle with references? I'm missing something about the abstraction
	err = app.RegisterTools(server)

	if err != nil {
		return err
	}
}

func (a *App) RegisterTools(server *mcp.Server) (err error) {
	// Tool Example
	err := server.RegisterTool("hello", "Say hello to a person", helloTool)

	if err != nil {
		panic(err)
	}
}

func (App) registerResources() {

}

func (App) registerPrompts() {

}
