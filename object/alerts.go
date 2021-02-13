package object

// Alert struct.
type Alert struct {
	Action        string `json:"action"`
	AlertID       uint   `json:"alert_id"`
	AlertText     string `json:"alert_text"`
	AlertURL      string `json:"alert_url"`
	AlertedUserID uint   `json:"alerted_user_id"`
	AutoRead      bool   `json:"auto_read"`
	ContentID     uint   `json:"content_id"`
	ContentType   string `json:"content_type"`
	EventDate     uint32 `json:"event_date"`
	ReadDate      uint32 `json:"read_date"`
	User          User   `json:"User"`
	UserID        uint   `json:"user_id"`
	Username      string `json:"username"`
	ViewDate      uint32 `json:"view_date"`
}

// AlertsResponse struct.
type AlertsResponse struct {
	Alerts     []Alert    `json:"alerts"`
	Alert      Alert      `json:"alert"`
	Pagination Pagination `json:"pagination"`
	Success    bool       `json:"success"`
}
