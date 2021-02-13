package object // import "github.com/JustCup/xenforo-sdk/object"

// User struct.
//
// Information about the user. Different information will be included based on permissions and verbosity.
type User struct {
	About                         string            `json:"about"`                                      // Returned only if permissions are met. Verbose results only.
	ActivityVisible               bool              `json:"activity_visible,omitempty"`                 // Returned only if permissions are met.
	Age                           int               `json:"age,omitempty"`                              // The user's current age. Only included if available.
	AlertOptout                   []string          `json:"alert_optout,omitempty"`                     // Returned only if permissions are met. Verbose results only.
	AllowPostProfile              string            `json:"allow_post_profile,omitempty"`               // Returned only if permissions are met. Verbose results only.
	AllowReceiveNewsFeed          string            `json:"allow_receive_news_feed,omitempty"`          // Returned only if permissions are met. Verbose results only.
	AllowSendPersonalConversation string            `json:"allow_send_personal_conversation,omitempty"` // Returned only if permissions are met. Verbose results only.
	AllowViewIdentities           string            `json:"allow_view_identities,omitempty"`            // Returned only if permissions are met. Verbose results only.
	AllowViewProfile              string            `json:"allow_view_profile,omitempty"`               // Returned only if permissions are met. Verbose results only.
	AvatarURLs                    map[string]string `json:"avatar_urls"`                                // Maps from size types to URL.
	CanBan                        bool              `json:"can_ban"`
	CanConverse                   bool              `json:"can_converse"`
	CanEdit                       bool              `json:"can_edit"`
	CanFollow                     bool              `json:"can_follow"`
	CanIgnore                     bool              `json:"can_ignore"`
	CanPostProfile                bool              `json:"can_post_profile"`
	CanViewProfile                bool              `json:"can_view_profile"`
	CanViewProfilePosts           bool              `json:"can_view_profile_posts"`
	CanWarn                       bool              `json:"can_warn"`
	ContentShowSignature          bool              `json:"content_show_signature,omitempty"` // Returned only if permissions are met. Verbose results only.
	CreationWatchState            string            `json:"creation_watch_state,omitempty"`   // Returned only if permissions are met. Verbose results only.
	CustomFields                  map[string]string `json:"custom_fields"`                    // Returned only if permissions are met. Map of custom field keys and values.
	CustomTitle                   string            `json:"custom_title,omitempty"`           // Returned only if permissions are met. Will have a value if a custom title has been specifically set; prefer user_title instead.
	DOB                           DOB               `json:"dob,omitempty"`                    // Returned only if permissions are met. Date of birth with year, month and day keys.
	Email                         string            `json:"email,omitempty"`                  // Returned only if permissions are met.
	EmailOnConversation           bool              `json:"email_on_conversation,omitempty"`  // Returned only if permissions are met.
	Gravatar                      string            `json:"gravatar,omitempty"`
	InteractionWatchState         string            `json:"interaction_watch_state,omitempty"`
	IsAdmin                       bool              `json:"is_admin,omitempty"`
	IsBanned                      bool              `json:"is_banned,omitempty"`
	IsDiscouraged                 bool              `json:"is_discouraged,omitempty"`
	IsFollowed                    bool              `json:"is_followed,omitempty"` // True if the visitor is following this user. Only included if visitor is not a guest.
	IsIgnored                     bool              `json:"is_ignored,omitempty"`  // True if the visitor is ignoring this user. Only included if visitor is not a guest.
	IsModerator                   bool              `json:"is_moderator,omitempty"`
	IsSuperAdmin                  bool              `json:"is_super_admin,omitempty"`
	LastActivity                  int               `json:"last_activity,omitempty"` // Unix timestamp of user's last activity, if available.
	Location                      string            `json:"location,omitempty"`
	PushOnConversation            bool              `json:"push_on_conversation,omitempty"`
	PushOptout                    []string          `json:"push_optout,omitempty"`
	ReceiveAdminEmail             bool              `json:"receive_admin_email,omitempty"`
	SecondaryGroupIDs             []int             `json:"secondary_group_ids,omitempty"`
	ShowDOBDate                   bool              `json:"show_dob_date,omitempty"`
	ShowDOBYear                   bool              `json:"show_dob_year,omitempty"`
	Signature                     string            `json:"signature,omitempty"`
	Timezone                      string            `json:"timezone,omitempty"`
	UsaTFA                        bool              `json:"usa_tfa,omitempty"` // (?) use_tfa in endpoints
	UserGroupID                   int               `json:"user_group_id,omitempty"`
	UserState                     string            `json:"user_state,omitempty"`
	UserTitle                     string            `json:"user_title,omitempty"`
	Visible                       bool              `json:"visible,omitempty"`
	WarningPoints                 int               `json:"warning_points,omitempty"` // Current warning points.
	Website                       string            `json:"website"`
	UserID                        int               `json:"user_id"`
	Username                      string            `json:"username"`
	MessageCount                  int               `json:"message_count"`
	RegisterDate                  int               `json:"register_date"`
	TrophyPoints                  int               `json:"trophy_points,omitempty"`
	IsStaff                       bool              `json:"is_staff"`
	ReactionScore                 int               `json:"reaction_score"`
}

// UserGetFindName struct.
type UserGetFindName struct {
	Exact           User   `json:"exact,omitempty"` // The user that matched the given username exactly. Empty struct or user info. TODO: refactor(make NULL great again)
	Recommendations []User `json:"recommendations"` // A list of users that match the prefix of the username (but not exactly)
}

// UsersGet struct.
type UsersGet struct {
	Users      []User     `json:"users"`
	Pagination Pagination `json:"pagination,omitempty"`
}

// UsersGetID stuct.
type UsersGetID struct {
	User         User          `json:"user"`
	ProfilePosts []ProfilePost `json:"profile_posts,omitempty"` // List of profile posts on the requested page
	Pagination   Pagination    `json:"pagination,omitempty"`    // Pagination details
}

// UserCreate struct.
type UserCreate struct {
	Success bool `json:"success"`
	User    User `json:"user"`
}
