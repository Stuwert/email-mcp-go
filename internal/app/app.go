package app

import (
	"github.com/stuwert/email-mcp/internal/prompts"
	"github.com/stuwert/email-mcp/internal/resources"
	"github.com/stuwert/email-mcp/internal/tools"
)

// type Tool struct {
// 	Hello HelloTool
// }

// App Owns associating the handlers. Resource handler, prompt handler, tools handler
// Main instruments the routes during the boot step.
// Routes.SetupHandlers(app)

type Application struct {
	Prompts   *prompts.Prompts
	Tools     *tools.Tools
	Resources *resources.Resources
}

func NewApplication() (*Application, error) {

	app := &Application{
		Prompts:   prompts.NewPromptHandlers(),
		Tools:     tools.NewToolHandlers(),
		Resources: resources.NewResourceHandlers(),
	}

	return app, nil
}
