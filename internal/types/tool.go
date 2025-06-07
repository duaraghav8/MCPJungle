package types

// ToolInvokeResult represents the result of a Tool call.
// It is designed to be passed down to the end user.
type ToolInvokeResult struct {
	Meta    map[string]any   `json:"_meta,omitempty"`
	IsError bool             `json:"isError,omitempty"`
	Content []map[string]any `json:"content"`
}
