package api // import "github.com/JustCup/go-xenforo-api/api"
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	goxenforoapi "github.com/JustCup/go-xenforo-api"
	"github.com/JustCup/go-xenforo-api/object"
)

// UserAgent for requests.
const UserAgent string = "goXenForoAPI/" + goxenforoapi.Version + ".xf2.1 (+https://github.com/JustCup/go-xenforo-api)"

// XF struct.
type XF struct {
	APIEndpoint string
	XFAPIKey    string
}

// Params type.
type Params map[string]interface{}

// Init XF API.
//
// Insert APIEndpoint(full url without, example: https://example.com/forum/api) and XenForo API key.
func Init(url string, apiKey string) *XF {
	var xf XF
	xf.APIEndpoint = url
	xf.XFAPIKey = apiKey

	// TODO(JustCup): check data

	return &xf
}

// Request func.
func (xf *XF) Request(reqType string, method string, params Params) (response []byte, err error) {

	query := url.Values{}

	FormatParams(&query, params)

	reqRaw := bytes.NewBufferString(query.Encode())

	req, err := http.NewRequest(reqType, fmt.Sprintf("%s/%s", xf.APIEndpoint, method), reqRaw)
	if err != nil {
		return
	}

	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("XF-Api-Key", xf.XFAPIKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		var e object.Error
		err := json.Unmarshal(body, &e)
		if err != nil {
			return response, err
		}

		return response, fmt.Errorf("code: %s, msg: %s", e.Errors[0].Code, e.Errors[0].Message)
	}

	return body, nil
}

// RequestUnmarshal decodes the JSON from the response.
func (xf *XF) RequestUnmarshal(reqType string, method string, params Params, obj interface{}) error {

	resp, err := xf.Request(reqType, method, params)
	if err != nil {
		return err
	}

	return json.Unmarshal(resp, &obj)
}

// FormatParams function.
func FormatParams(q *url.Values, params Params) {
	for k, v := range params {
		key, val := FormatValue(v, k)
		if key != "" && val != "" {
			q.Set(key, val)
		}
	}
}

// FormatValue function.
func FormatValue(value interface{}, key string) (k, v string) {
	if value == nil || key == "" {
		return
	}

	switch iface := value.(type) {
	case bool:
		return key, fmt.Sprintf("%t", iface)
	case Params:
		for kI, vI := range iface {
			keyAns, valAns := FormatValue(vI, kI)
			return fmt.Sprintf("%s[%s]", key, keyAns), valAns
		}
	default:
		return key, fmt.Sprintf("%v", iface)
	}

	return
}
