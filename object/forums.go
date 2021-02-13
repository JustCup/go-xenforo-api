package object // import "github.com/JustCup/xenforo-sdk/object"

// Forum struct.
type Forum struct {
	AllowPosting  bool `json:"allow_posting"`
	AllowPool     bool `json:"allow_pool"`
	RequirePrefix bool `json:"require_prefix"`
	MinTags       int  `json:"min_tags"`
}

// ForumsGet struct.
type ForumsGet struct {
	Forum      Forum      `json:"forum"`
	Threads    []Thread   `json:"threads"`    // Threads on this page
	Pagination Pagination `json:"Pagination"` // Pagination information
	Sticky     []Thread   `json:"sticky"`     // If on page 1, a list of sticky threads in this forum. Does not count towards the per page limit.
}

// ForumsGetThreads struct.
type ForumsGetThreads struct {
	Threads    []Thread   `json:"threads"`    // Threads on this page
	Pagination Pagination `json:"pagination"` // Pagination information
	Sticky     []Thread   `json:"sticky"`     // If on page 1, a list of sticky threads in this forum. Does not count towards the per page limit.
}
