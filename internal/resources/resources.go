package resources

import (
	mcp "github.com/metoro-io/mcp-golang"
)

type Resources struct{}

func NewResourceHandlers() *Resources {
	return &Resources{}
}

func (r *Resources) HelloResource() (*mcp.ResourceResponse, error) {
	return mcp.NewResourceResponse(mcp.NewTextEmbeddedResource("test://resource", "This is a test resource", "application/json")), nil
}
