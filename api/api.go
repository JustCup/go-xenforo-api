package api // import "github.com/JustCup/go-xenforo-api/api"

// XF struct.
type XF struct {
	RequestURL string
	XFApiKey   string
}

// Init XF API
//
// Insert RequestURL(full url without, example: https://example.com/forum/api) and XF API Key
func Init(url string, apiKey string) *XF {
	var xf XF

	xf.RequestURL = url
	xf.XFApiKey = apiKey

	return &xf
}
