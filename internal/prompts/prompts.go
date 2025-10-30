package prompts

import (
	"fmt"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/stuwert/email-mcp/internal/types"
)

type Prompts struct{}

func NewPromptHandlers() *Prompts {
	return &Prompts{}
}

func (p *Prompts) HelloPrompt(arguments types.Content) (*mcp.PromptResponse, error) {
	return mcp.NewPromptResponse("description", mcp.NewPromptMessage(mcp.NewTextContent(fmt.Sprintf("Hello, %server", arguments.Title)), mcp.RoleUser)), nil
}
