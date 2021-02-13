package api

import (
	"fmt"
	"github.com/JustCup/xenforo-sdk/object"
)

// Nodes function.
//
// Gets the node tree.
//
// https://xenforo.com/community/pages/api-endpoints/#route_get_nodes_
func (xf *XF) Nodes() (response object.NodeResponse, err error) {
	err = xf.RequestUnmarshal("GET", "nodes", Params{}, &response)
	return
}

// NodesCreate function.
//
// Creates a new node.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_nodes_
func (xf *XF) NodesCreate(params Params) (response object.NodeResponse, err error) {
	err = xf.RequestUnmarshal("POST", "nodes", params, &response)
	return
}

// NodesFlattened function.
//
// Gets a flattened node tree. Traversing this will return a list of nodes in the expected order.
//
// https://xenforo.com/community/pages/api-endpoints/#route_get_nodes_flattened
func (xf *XF) NodesFlattened() (response object.NodeResponse, err error) {
	err = xf.RequestUnmarshal("GET", "nodes/flattened", Params{}, &response)
	return
}

// Node function.
//
// Gets information about the specified node.
//
// https://xenforo.com/community/pages/api-endpoints/#route_get_nodes_id_
func (xf *XF) Node(id int) (response object.NodeResponse, err error) {
	err = xf.RequestUnmarshal("GET", fmt.Sprintf("nodes/%d", id), Params{}, &response)
	return
}

// NodeUpdate function.
//
// Updates the specified node.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_nodes_id_
func (xf *XF) NodeUpdate(id int, params Params) (response object.NodeResponse, err error) {
	err = xf.RequestUnmarshal("GET", fmt.Sprintf("nodes/%d", id), params, &response)
	return
}

// NodeDelete function.
//
// Deletes the specified node.
//
// https://xenforo.com/community/pages/api-endpoints/#route_delete_nodes_id_
func (xf *XF) NodeDelete(id int, params Params) (response object.NodeResponse, err error) {
	err = xf.RequestUnmarshal("DELETE", fmt.Sprintf("nodes/%d", id), params, &response)
	return
}
