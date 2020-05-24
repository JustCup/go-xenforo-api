package object // import "github.com/JustCup/go-xenforo-api/object"

// Forum struct.
type Forum struct {
	AllowPosting  bool `json:"allow_posting"`
	AllowPool     bool `json:"allow_pool"`
	RequirePrefix bool `json:"require_prefix"`
	MinTags       int  `json:"min_tags"`
}

// ForumInfo struct.
type ForumInfo struct {
	Forum      Forum      `json:"forum"`
	Threads    []Thread   `json:"threads"`    // Threads on this page
	Pagination Pagination `json:"Pagination"` // Pagination information
	Sticky     []Thread   `json:"sticky"`     // If on page 1, a list of sticky threads in this forum. Does not count towards the per page limit.
}
