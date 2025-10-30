package tools

import (
	"fmt"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/stuwert/email-mcp/internal/types"
)

type Tools struct{}

func NewToolHandlers() *Tools {
	return &Tools{}
}

func (t *Tools) HelloTool(arguments types.MyFunctionArguments) (*mcp.ToolResponse, error) {
	return mcp.NewToolResponse(mcp.NewTextContent(fmt.Sprintf("Hello, %server!", arguments.Submitter))), nil
}
