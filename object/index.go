package object

// IndexResponse struct.
type IndexResponse struct {
	VersionID uint   `json:"version_id"` // XenForo version ID.
	SiteTitle string `json:"site_title"` // Title of the site this API relates to.
	BaseURL   string `json:"base_url"`   // The base URL of the XenForo install this API relates to.
	ApiURL    string `json:"api_url"`    // The base API URL.
	Key       struct {
		Type           string   `json:"type"`             // Type of the API key accessing the API (guest, user or super).
		UserID         uint     `json:"user_id"`          // If a user key, the ID of the user the key is for; null otherwise.
		AllowAllScopes bool     `json:"allow_all_scopes"` // If true, all scopes can be accessed.
		Scopes         []string `json:"scopes"`           // A list of scopes this key can access (if not allowed to access all scopes).
	} `json:"key"`
}
