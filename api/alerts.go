package api

import (
	"fmt"
	"github.com/JustCup/xenforo-sdk/object"
)

// AlertsGet function.
//
// Gets the API user's list of alerts.
//
// https://xenforo.com/community/pages/api-endpoints/#route_get_alerts_
func (xf *XF) AlertsGet(params Params) (response object.AlertsResponse, err error) {
	err = xf.RequestUnmarshal("GET", "alerts", params, &response)
	return
}

// AlertsSend function.
//
// Sends an alert to the specified user. Only available to super user keys.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_alerts_
func (xf *XF) AlertsSend(params Params) (response object.AlertsResponse, err error) {
	err = xf.RequestUnmarshal("POST", "alerts", params, &response)
	return
}

// AlertsMarkAll function.
//
// Marks all of the API user's alerts as read or viewed. Must specify "read" or "viewed" parameters.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_alerts_mark_all
func (xf *XF) AlertsMarkAll(params Params) (response object.AlertsResponse, err error) {
	err = xf.RequestUnmarshal("POST", "alerts/mark-all", params, &response)
	return
}

// AlertsIDGet function.
//
// Gets information about the specified alert.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_alerts_mark_all
func (xf *XF) AlertsIDGet(id uint) (response object.AlertsResponse, err error) {
	err = xf.RequestUnmarshal("POST", fmt.Sprintf("%s/%d", "alerts", id), Params{}, &response)
	return
}

// AlertsIDMark function.
//
// Gets information about the specified alert.
//
// Marks the alert as viewed/read/unread. (Marking as unviewed is not supported.)
func (xf *XF) AlertsIDMark(id uint, params Params) (response object.AlertsResponse, err error) {
	err = xf.RequestUnmarshal("POST", fmt.Sprintf("%s/%d/%s", "alerts", id, "mark"), params, &response)
	return
}
