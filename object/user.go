package object // import "github.com/JustCup/go-xenforo-api/object"

// User struct.
//
// Information about the user. Different information will be included based on permissions and verbosity.
type User struct {
	About                         string            `json:"about"`                            // Returned only if permissions are met. Verbose results only.
	ActivityVisible               bool              `json:"activity_visible"`                 // Returned only if permissions are met.
	Age                           int               `json:"age,omitempty"`                    // The user's current age. Only included if available.
	AlertOptout                   []string          `json:"alert_optout"`                     // Returned only if permissions are met. Verbose results only.
	AllowPostProfile              string            `json:"allow_post_profile"`               // Returned only if permissions are met. Verbose results only.
	AllowReceiveNewsFeed          string            `json:"allow_receive_news_feed"`          // Returned only if permissions are met. Verbose results only.
	AllowSendPersonalConversation string            `json:"allow_send_personal_conversation"` // Returned only if permissions are met. Verbose results only.
	AllowViewIdentities           string            `json:"allow_view_identities"`            // Returned only if permissions are met. Verbose results only.
	AllowViewProfile              string            `json:"allow_view_profile"`               // Returned only if permissions are met. Verbose results only.
	AvatarURLs                    map[string]string `json:"avatar_urls"`                      // Maps from size types to URL.
	CanBan                        bool              `json:"can_ban"`
	CanEdit                       bool              `json:"can_edit"`
	CanFollow                     bool              `json:"can_follow"`
	CanIgnore                     bool              `json:"can_ignore"`
	CanPostProfile                bool              `json:"can_post_profile"`
	CanViewProfile                bool              `json:"can_view_profile"`
	CanViewProfilePosts           bool              `json:"can_view_profile_posts"`
	CanWarn                       bool              `json:"can_warn"`
	ContentShowSignature          bool              `json:"content_show_signature"` // Returned only if permissions are met. Verbose results only.
	CreationWatchState            string            `json:"creation_watch_state"`   // Returned only if permissions are met. Verbose results only.
	CustomFields                  map[string]string `json:"custom_fields"`          // Returned only if permissions are met. Map of custom field keys and values.
	CustomTitle                   string            `json:"custom_title"`           // Returned only if permissions are met. Will have a value if a custom title has been specifically set; prefer user_title instead.
	DOB                           map[string]string `json:"dob"`                    // Returned only if permissions are met. Date of birth with year, month and day keys.
	Email                         string            `json:"email"`                  // Returned only if permissions are met.
	EmailOnConversation           string            `json:"email_on_conversation"`  // Returned only if permissions are met.
	Gravatar                      string            `json:"gravatar"`
	InteractionWatchState         bool              `json:"interaction_watch_state"`
	IsAdmin                       bool              `json:"is_admin"`
	IsBanned                      bool              `json:"is_banned"`
	IsDiscouraged                 bool              `json:"is_discouraged"`
	IsFollowed                    bool              `json:"is_followed"` // True if the visitor is following this user. Only included if visitor is not a guest.
	IsIgnored                     bool              `json:"is_ignored"`  // True if the visitor is ignoring this user. Only included if visitor is not a guest.
	IsModerator                   bool              `json:"is_moderator"`
	IsSuperAdmin                  bool              `json:"is_super_admin"`
	LastActivity                  int               `json:"last_activity"` // Unix timestamp of user's last activity, if available.
	Location                      string            `json:"location"`
	PushOnConversation            bool              `json:"push_on_conversation"`
	PushOptout                    []string          `json:"push_optout"`
	ReceiveAdminEmail             bool              `json:"receive_admin_email"`
	SecondaryGroupIDs             []int             `json:"secondary_group_ids"`
	ShowDOBDate                   bool              `json:"show_dob_date"`
	ShowDOBYear                   bool              `json:"show_dob_year"`
	Signature                     string            `json:"signature"`
	Timezone                      string            `json:"timezone"`
	UsaTFA                        bool              `json:"usa_tfa"` // (?) use_tfa in endpoints
	UserGroupID                   int               `json:"user_group_id"`
	UserState                     string            `json:"user_state"`
	UserTitle                     string            `json:"user_title"`
	Visible                       bool              `json:"visible"`
	WarningPoints                 int               `json:"warning_points"` // Current warning points.
	Website                       string            `json:"website"`
	UserID                        int               `json:"user_id"`
	Username                      string            `json:"username"`
	MessageCount                  int               `json:"message_count"`
	RegisterDate                  int               `json:"register_date"`
	TrophyPoints                  int               `json:"trophy_points"`
	IsStaff                       bool              `json:"is_staff"`
	ReactionScore                 int               `json:"reaction_score"`
}

// UserFindName struct.
type UserFindName struct {
	Exact           *User  `json:"exact,omitempty"` // The user that matched the given username exactly. Empty struct or user info. TODO: refactor(make NULL great again)
	Recommendations []User `json:"recommendations"` // A list of users that match the prefix of the username (but not exactly)
}
