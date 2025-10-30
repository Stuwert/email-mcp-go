package internal

import (
	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/stuwert/email-mcp/internal/app"
)

func RegisterTools(server *mcp_golang.Server, app *app.Application) error {

	err := server.RegisterTool("hello", "Say hello to a person", app.Tools.HelloTool)

	if err != nil {
		return err
	}

	return nil
}

func RegisterPrompts(server *mcp_golang.Server, app *app.Application) error {
	err := server.RegisterPrompt("prompt_test", "this is a test prompt", app.Prompts.HelloPrompt)

	if err != nil {
		return err
	}

	return nil
}

func RegisterResources(server *mcp_golang.Server, app *app.Application) error {
	err := server.RegisterResource("test://resource", "resource_test", "This is a test resource", "application/json", app.Resources.HelloResource)

	if err != nil {
		return err
	}

	return nil
}
