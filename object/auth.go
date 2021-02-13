package object // import "github.com/JustCup/xenforo-sdk/object"

// AuthResponse struct.
type AuthResponse struct {
	User       User   `json:"user"`    // If successful, the user record of the matching user.
	Success    bool   `json:"success"` // If false, no session or remember cookie could be found.
	LoginToken string `json:"login_token"`
	LoginURL   string `json:"login_url"`   // Direct user to this URL to trigger a login.
	ExpiryDate uint32 `json:"expiry_date"` // Unix timestamp of when the token expires. An error will be displayed if the token is expired or invalid.
}
