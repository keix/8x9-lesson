package entity

// Client represents an OAuth 2.0 client
type Client struct {
	ID                      string   `json:"client_id"`
	Secret                  string   `json:"client_secret,omitempty"`
	RedirectURIs            []string `json:"redirect_uris"`
	Name                    string   `json:"client_name,omitempty"`
	TokenEndpointAuthMethod string   `json:"token_endpoint_auth_method,omitempty"`
	ClientType              string   `json:"client_type,omitempty"` // "public" or "confidential"
}

// IsPublicClient returns true if the client is a public client
func (c *Client) IsPublicClient() bool {
	return c.ClientType == "public"
}

// IsConfidentialClient returns true if the client is a confidential client
// Default to confidential for backward compatibility
func (c *Client) IsConfidentialClient() bool {
	return c.ClientType == "confidential" || c.ClientType == ""
}
