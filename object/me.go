package object

// MeResponse struct.
type MeResponse struct {
	Success              bool `json:"success"`
	Me                   User `json:"me"`
	ConfirmationRequired bool `json:"confirmation_required"` // True if email confirmation is required for this change
}
