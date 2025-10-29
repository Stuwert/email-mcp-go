package main

import (
	"fmt"

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

func main() {
	done := make(chan struct{})

	// 1. Initiate server
	// 2. Initiate connection to database
	// 3. Register Tools
	// 4. Register Resource
	// 5. Register Prompts

	server := mcp.NewServer(stdio.NewStdioServerTransport())

	// Tool Example
	err := server.RegisterTool("hello", "Say hello to a person", func(arguments MyFunctionArguments) (*mcp.ToolResponse, error) {
		return mcp.NewToolResponse(mcp.NewTextContent(fmt.Sprintf("Hello, %server!", arguments.Submitter))), nil
	})

	if err != nil {
		panic(err)
	}

	// Resource Example
	err = server.RegisterResource("test://resource", "resource_test", "This is a test resource", "application/json", func() (*mcp.ResourceResponse, error) {
		return mcp.NewResourceResponse(mcp.NewTextEmbeddedResource("test://resource", "This is a test resource", "application/json")), nil
	})

	if err != nil {
		panic(err)
	}

	// Prompt Example
	err = server.RegisterPrompt("prompt_test", "this is a test prompt", func(arguments Content) (*mcp.PromptResponse, error) {
		return mcp.NewPromptResponse("description", mcp.NewPromptMessage(mcp.NewTextContent(fmt.Sprintf("Hello, %server", arguments.Title)), mcp.RoleUser)), nil
	})

	if err != nil {
		panic(err)
	}

	err = server.Serve()

	if err != nil {
		panic(err)
	}

	<-done
}
