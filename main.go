package main

import (
	"fmt"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

func main() {
	done := make(chan struct{})

	server := mcp.NewServer(stdio.NewStdioServerTransport())

	err := server.RegisterTool("hello", "Say hello to a person", func(arguments myFunctionArguments) (*mcp.ToolResponse, error) {
		return mcp.NewToolResponse(mcp.NewTextContent(fmt.Sprintf("Hello, %server!", arguments.Submitter))), nil
	})

	if err != nil {
		panic(err)
	}

	<-done
}
