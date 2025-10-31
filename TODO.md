# Email MCP Server - Implementation TODOs

A sequential checklist for building and testing a Go MCP server that expands email sorter functionality.

## Phase 1: Foundation & Setup

- [x] **TODO 1: Create main.go entry point**
  - Create a `main.go` file in the root directory
  - Set up basic package structure with main function
  - Import necessary MCP packages from `github.com/metoro-io/mcp-golang`
  - Add placeholder log statement to verify basic compilation works
  - Test: `go run main.go` should compile and exit successfully

- [x] **TODO 2: Initialize basic MCP server**
    - Create a level of abstraction so that the app.go file handles registering tools, resources and prompts
    - Main file starts the app, records any errors and handles deferring

- [x] **TODO 3: Hook this into a local chat instance and validate whether or not it's running
    - Figure out how to hook up server into local running instance (can I throw it into cursor?)
    - Check that resources and hello functionality works. 

// Cursor said it works. Need to try adding this to Claude locally.

- [ ] **TODO 4: Add server configuration**
  - Create configuration structure (using environment variables or config file)
  - Support for reading configuration on startup
  - Add basic validation for required config values
  - Test: Server should fail gracefully with clear error messages if config is missing

## Phase 2: First Tool Implementation

- [ ] **TODO 4: Create tool package structure**
  - Create `internal/tools/email.go` file
  - Set up package structure for tool definitions
  - Add basic tool interface/structure following mcp-golang patterns
  - Test: Package should compile without errors

- [ ] **TODO 5: Implement simplest tool - "parse_email_address"**
  - Create a tool that takes a raw email address string as input
  - Parse and extract: local part, domain, display name (if present)
  - Return structured JSON with parsed components
  - Register tool with MCP server
  - Test: Use MCP client or test harness to call tool and verify output

- [ ] **TODO 6: Add input validation and error handling**
  - Add validation for tool inputs (required fields, format checks)
  - Return appropriate error messages for invalid inputs
  - Handle edge cases (empty strings, malformed addresses)
  - Test: Verify error messages are clear and helpful

- [ ] **TODO 7: Create unit tests for first tool**
  - Create `internal/tools/email_test.go`
  - Write test cases for valid email addresses
  - Write test cases for invalid email addresses
  - Write test cases for edge cases
  - Test: `go test ./internal/tools/` should pass all tests

## Phase 3: Testing Infrastructure

- [ ] **TODO 8: Set up integration test framework**
  - Create `test/` directory for integration tests
  - Set up test helper functions for starting/stopping MCP server
  - Create basic integration test that connects to server
  - Test: Integration test should connect and verify server is responding

- [ ] **TODO 9: Create manual testing script/documentation**
  - Document how to run the server locally
  - Create example MCP client connection commands
  - Document expected tool signatures and usage
  - Test: Follow documentation to manually test server with MCP client

## Phase 4: Additional Tools

- [ ] **TODO 10: Implement "validate_email_address" tool**
  - Create tool that validates email address format
  - Check syntax, domain existence (optional placements), common patterns
  - Return validation result with detailed feedback
  - Register and test the new tool

- [ ] **TODO 11: Implement "extract_email_headers" tool**
  - Create tool that parses raw email message string
  - Extract headers (From, To, Subject, Date, etc.)
  - Return structured header information
  - Handle malformed emails gracefully
  - Add unit tests

- [ ] **TODO 12: Implement "categorize_email" tool**
  - Create tool that categorizes email content
  - Take subject, body, and sender as inputs
  - Return category (e.g., "promotional", "personal", "important", "spam")
  - Use simple rule-based logic (can be enhanced later)
  - Add comprehensive tests

## Phase 5: Resources (Optional - Advanced)

- [ ] **TODO 13: Create resource package structure**
  - Create `internal/resources/email.go` file
  - Set up package structure for resource definitions
  - Understand MCP resource patterns

- [ ] **TODO 14: Implement "email_templates" resource**
  - Create a resource that exposes email templates
  - Allow listing available templates
  - Allow reading template content
  - Register resource with MCP server
  - Test resource discovery and reading

## Phase 6: Error Handling & Polish

- [ ] **TODO 15: Improve error handling across all tools**
  - Standardize error response format
  - Add error codes for different error types
  - Ensure all errors are properly logged
  - Test: Verify errors are returned in expected format

- [ ] **TODO 16: Add structured logging**
  - Replace print statements with structured logger
  - Add log levels (INFO, DEBUG, ERROR)
  - Log tool invocations and responses
  - Test: Verify logs are helpful for debugging

- [ ] **TODO 17: Add tool descriptions and metadata**
  - Add comprehensive descriptions for each tool
  - Document input/output schemas
  - Add examples to tool metadata
  - Test: Verify metadata is accessible via MCP protocol

## Phase 7: Advanced Features

- [ ] **TODO 18: Implement tool result caching (if applicable)**
  - Identify tools that could benefit from caching
  - Implement simple in-memory cache with TTL
  - Add cache invalidation logic
  - Test: Verify cache works and improves performance

- [ ] **TODO 19: Add rate limiting or throttling**
  - Implement basic rate limiting for tool calls
  - Prevent server overload
  - Return appropriate errors when rate limit exceeded
  - Test: Verify rate limiting works as expected

- [ ] **TODO 20: Add configuration for tool behavior**
  - Make categorization rules configurable
  - Allow configuration of validation strictness
  - Support runtime configuration updates
  - Test: Verify configuration affects tool behavior

## Phase 8: Documentation & Cleanup

- [ ] **TODO 21: Write comprehensive README**
  - Document server capabilities
  - Include installation and setup instructions
  - Provide usage examples
  - Document all available tools and resources

- [ ] **TODO 22: Add code comments and documentation**
  - Add godoc comments to all exported functions
  - Document package-level functionality
  - Add inline comments for complex logic
  - Generate documentation with `go doc`

- [ ] **TODO 23: Code cleanup and refactoring**
  - Review code for duplication
  - Extract common patterns into helper functions
  - Ensure consistent code style
  - Run `gofmt` and `golint` to verify code quality
