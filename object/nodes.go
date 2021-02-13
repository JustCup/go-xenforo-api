package object // import "github.com/JustCup/xenforo-sdk/object"

// Node struct.
type Node struct {
	Breadcrumbs []struct {
		NodeID     uint   `json:"node_id"`
		Title      string `json:"title"`
		NodeTypeID string `json:"node_type_id"`
	} `json:"breadcrumbs"` // A list of breadcrumbs for this node, including the node_id, title, and node_type_id
	TypeData      map[string]interface{} `json:"type_data"` // Data related to the specific node type this represents. Contents will vary significantly.
	ViewURL       string                 `json:"view_url"`
	NodeID        uint                   `json:"node_id"`
	Title         string                 `json:"title"`
	NodeName      string                 `json:"node_name"`
	Description   string                 `json:"description"`
	NodeTypeID    string                 `json:"node_type_id"`
	ParentNodeID  uint                   `json:"parent_node_id"`
	DisplayOrder  uint                   `json:"display_order"`
	DisplayInList bool                   `json:"display_in_list"`
}

// NodeFlat struct.
type NodeFlat struct {
	Node  Node `json:"node"`
	Depth uint `json:"depth"`
}

// NodeResponse struct.
type NodeResponse struct {
	TreeMap   map[uint][]uint `json:"tree_map"` // A mapping that connects node parent IDs to a list of their child node IDs.
	Nodes     []Node          `json:"nodes"`
	Node      Node            `json:"node"`
	NodesFlat []NodeFlat      `json:"nodes_flat"`
	Success   bool            `json:"success"`
}
