package object // import "github.com/JustCup/go-xenforo-api/object"

// Thread struct.
type Thread struct {
	Username               string            `json:"username"`
	IsWatching             bool              `json:"is_watching,omitempty"`        // If accessing as a user, true if they are watching this thread
	VisitorPostCount       int               `json:"visitor_post_count,omitempty"` // If accessing as a user, the number of posts they have made in this thread
	CustomFields           map[string]string `json:"custom_fields"`
	Tags                   []string          `json:"tags"`
	Prefix                 string            `json:"prefix,omitempty"` // Present if this thread has a prefix. Printable name of the prefix.
	CanEdit                bool              `json:"can_edit"`
	CanEditTags            bool              `json:"can_edit_tags"`
	CanReply               bool              `json:"can_reply"`
	CanSoftDelete          bool              `json:"can_soft_delete"`
	CanHardDelete          bool              `json:"can_hard_delete"`
	CanViewAttachments     bool              `json:"can_view_attachments"`
	Forum                  Node              `json:"Forum"`
	ThreadID               int               `json:"thread_id"`
	NodeID                 int               `json:"node_id"`
	Title                  string            `json:"title"`
	ReplyCount             int               `json:"reply_count"`
	ViewCount              int               `json:"view_count"`
	UserID                 int               `json:"user_id"`
	Sticky                 bool              `json:"sticky"`
	DiscussionState        string            `json:"discussion_state"`
	DiscussionOpen         bool              `json:"discussion_open"`
	DiscussionType         string            `json:"discussion_type"`
	FirstPostID            int               `json:"first_post_id"`
	LastPostID             int               `json:"last_post_id"`
	LastPostDate           int               `json:"last_post_date"`
	LastPostUserID         int               `json:"last_post_user_id"`
	LastPostUsername       string            `json:"last_post_username"`
	FirstPostReactionScore int               `json:"first_post_reaction_score"`
	PrefixID               int               `json:"prefix_id"`
	User                   User              `json:"user"`
}
