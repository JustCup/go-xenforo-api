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
	ProfileBannerURLs             map[string]string `json:"profile_banner_urls"`                        // Maps from size types to URL.
	CanBan                        bool              `json:"can_ban"`
	CanConverse                   bool              `json:"can_converse"`
	CanEdit                       bool              `json:"can_edit"`
	CanFollow                     bool              `json:"can_follow"`
	CanIgnore                     bool              `json:"can_ignore"`
	CanPostProfile                bool              `json:"can_post_profile"`
	CanViewProfile                bool              `json:"can_view_profile"`
	CanViewProfilePosts           bool              `json:"can_view_profile_posts"`
	CanWarn                       bool              `json:"can_warn"`
	ContentShowSignature          bool              `json:"content_show_signature,omitempty"`  // Returned only if permissions are met. Verbose results only.
	CreationWatchState            string            `json:"creation_watch_state,omitempty"`    // Returned only if permissions are met. Verbose results only.
	CustomFields                  map[string]string `json:"custom_fields"`                     // Returned only if permissions are met. Map of custom field keys and values.
	CustomTitle                   string            `json:"custom_title,omitempty"`            // Returned only if permissions are met. Will have a value if a custom title has been specifically set; prefer user_title instead.
	DOB                           DOB               `json:"dob,omitempty"`                     // Returned only if permissions are met. Date of birth with year, month and day keys.
	Email                         string            `json:"email,omitempty"`                   // Returned only if permissions are met. Verbose results only.
	EmailOnConversation           bool              `json:"email_on_conversation,omitempty"`   // Returned only if permissions are met. Verbose results only.
	Gravatar                      string            `json:"gravatar,omitempty"`                // Returned only if permissions are met. Verbose results only.
	InteractionWatchState         string            `json:"interaction_watch_state,omitempty"` // Returned only if permissions are met. Verbose results only.
	IsAdmin                       bool              `json:"is_admin,omitempty"`                // Returned only if permissions are met.
	IsBanned                      bool              `json:"is_banned,omitempty"`               // Returned only if permissions are met.
	IsDiscouraged                 bool              `json:"is_discouraged,omitempty"`          // Returned only if permissions are met.
	IsFollowed                    bool              `json:"is_followed,omitempty"`             // True if the visitor is following this user. Only included if visitor is not a guest.
	IsIgnored                     bool              `json:"is_ignored,omitempty"`              // True if the visitor is ignoring this user. Only included if visitor is not a guest.
	IsModerator                   bool              `json:"is_moderator,omitempty"`            // Returned only if permissions are met.
	IsSuperAdmin                  bool              `json:"is_super_admin,omitempty"`          // Returned only if permissions are met.
	LastActivity                  uint32            `json:"last_activity,omitempty"`           // Returned only if permissions are met. Unix timestamp of user's last activity, if available.
	Location                      string            `json:"location,omitempty"`
	PushOnConversation            bool              `json:"push_on_conversation,omitempty"` // Returned only if permissions are met. Verbose results only.
	PushOptout                    []string          `json:"push_optout,omitempty"`          // Returned only if permissions are met.
	ReceiveAdminEmail             bool              `json:"receive_admin_email,omitempty"`  // Returned only if permissions are met. Verbose results only.
	SecondaryGroupIDs             []uint            `json:"secondary_group_ids,omitempty"`  // Returned only if permissions are met.
	ShowDOBDate                   bool              `json:"show_dob_date,omitempty"`        // Returned only if permissions are met. Verbose results only.
	ShowDOBYear                   bool              `json:"show_dob_year,omitempty"`        // Returned only if permissions are met. Verbose results only.
	Signature                     string            `json:"signature,omitempty"`
	Timezone                      string            `json:"timezone,omitempty"`      // Returned only if permissions are met. Verbose results only.
	UseTFA                        bool              `json:"usa_tfa,omitempty"`       // Returned only if permissions are met. Verbose results only.
	UserGroupID                   uint              `json:"user_group_id,omitempty"` // Returned only if permissions are met.
	UserState                     string            `json:"user_state,omitempty"`    // Returned only if permissions are met.
	UserTitle                     string            `json:"user_title,omitempty"`
	Visible                       bool              `json:"visible,omitempty"`        // Returned only if permissions are met.
	WarningPoints                 uint              `json:"warning_points,omitempty"` // Returned only if permissions are met. Current warning points.
	Website                       string            `json:"website"`                  // Returned only if permissions are met.
	ViewURL                       string            `json:"view_url"`
	UserID                        uint              `json:"user_id"`
	Username                      string            `json:"username"`
	MessageCount                  uint              `json:"message_count"`
	QuestionSolutionCount         uint              `json:"question_solution_count"`
	RegisterDate                  uint32            `json:"register_date"`
	TrophyPoints                  uint              `json:"trophy_points,omitempty"`
	IsStaff                       bool              `json:"is_staff"`
	ReactionScore                 uint              `json:"reaction_score"`
	VoteScore                     uint              `json:"vote_score"`
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
