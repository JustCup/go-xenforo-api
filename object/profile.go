package object // import "github.com/JustCup/xenforo-sdk/object"

// ProfilePost struct.
type ProfilePost struct {
	Username          string               `json:"username"`
	CanEdit           bool                 `json:"can_edit"`
	CanSoftDelete     bool                 `json:"can_soft_delete"`
	CanHardDelete     bool                 `json:"can_hard_delete"`
	CanReact          bool                 `json:"can_react"`
	ProfileUser       User                 `json:"ProfileUser"`         // If requested by context, the user this profile post was left for.
	LatestComments    []ProfilePostComment `json:"LatestComments"`      // If requested, the most recent comments on this profile post.
	IsReactedTo       bool                 `json:"is_reacted_to"`       // True if the viewing user has reacted to this content.
	VisitorReactionID int                  `json:"visitor_reaction_id"` // If the viewer reacted, the ID of the reaction they used.
	ProfilePostID     int                  `json:"profile_post_id"`
	ProfileUserID     int                  `json:"profile_user_id"`
	UserID            int                  `json:"user_id"`
	PostDate          int                  `json:"post_date"`
	Message           string               `json:"message"`
	MessageState      string               `json:"message_state"`
	WarningMessage    string               `json:"warning_message"`
	CommentCount      int                  `json:"comment_count"`
	FirstCommentDate  int                  `json:"first_comment_date"`
	LastCommentDate   int                  `json:"last_comment_date"`
	ReactionScore     int                  `json:"reaction_score"`
	User              User                 `json:"User"`
}

// ProfilePostComment struct.
type ProfilePostComment struct {
	Username             string      `json:"username"`
	CanEdit              bool        `json:"can_edit"`
	CanSoftDelete        bool        `json:"can_soft_delete"`
	CanHardDelete        bool        `json:"can_hard_delete"`
	CanReact             bool        `json:"can_react"`
	ProfilePost          ProfilePost `json:"ProfilePost"`         // If requested by context, the profile post this comment relates to.
	IsReactedTo          bool        `json:"is_reacted_to"`       // True if the viewing user has reacted to this content.
	VisitorReactionID    int         `json:"visitor_reaction_id"` // If the viewer reacted, the ID of the reaction they used.
	ProfilePostCommentID int         `json:"profile_post_comment_id"`
	ProfilePostID        int         `json:"profile_post_id"`
	UserID               int         `json:"user_id"`
	CommentDate          int         `json:"comment_date"`
	Message              string      `json:"message"`
	MessageState         string      `json:"message_state"`
	WarningMessage       string      `json:"warning_message"`
	ReactionScore        int         `json:"reaction_score"`
	User                 User        `json:"User"`
}
