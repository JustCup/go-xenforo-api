package object // import "github.com/JustCup/xenforo-sdk/object"

// Forum struct.
type Forum struct {
	ForumTypeID   string `json:"forum_type_id"`
	AllowPosting  bool   `json:"allow_posting"`
	RequirePrefix bool   `json:"require_prefix"`
	MinTags       uint   `json:"min_tags"`
}

// ForumsResponse struct.
type ForumsResponse struct {
	Forum      Forum      `json:"forum"`
	Threads    []Thread   `json:"threads"`    // Threads on this page
	Pagination Pagination `json:"Pagination"` // Pagination information
	Sticky     []Thread   `json:"sticky"`     // If on page 1, a list of sticky threads in this forum. Does not count towards the per page limit.
	Success    bool       `json:"success"`
}
