package object // import "github.com/JustCup/go-xenforo-api/object"

// Pagination struct.
//
// Pagination information
type Pagination struct {
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	PerPage     int `json:"per_page"`
	Shown       int `json:"shown"`
	Total       int `json:"total"`
}

// DOB struct.
type DOB struct {
	Year  int `json:"year,omitempty"`
	Day   int `json:"day"`
	Month int `json:"month"`
}

// Error struct.
type Error struct {
	Errors []struct {
		Code    string            `json:"code"`
		Message string            `json:"message"`
		Params  map[string]string `json:"params,omitempty"`
	} `json:"errors"`
}
