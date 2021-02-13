package object // import "github.com/JustCup/xenforo-sdk/object"

// Node struct.
type Node struct {
	Breadcrumbs []struct {
		NodeID     int    `json:"node_id"`
		Title      string `json:"title"`
		NodeTypeID string `json:"node_type_id"`
	} `json:"breadcrumbs"` // A list of breadcrumbs for this node, including the node_id, title, and node_type_id
	TypeData      map[string]string `json:"type_data"` // Data related to the specific node type this represents. Contents will vary significantly.
	NodeID        int               `json:"node_id"`
	Title         string            `json:"title"`
	NodeName      string            `json:"node_name"`
	Description   string            `json:"description"`
	NodeTypeID    string            `json:"node_type_id"`
	ParentNodeID  int               `json:"parent_node_id"`
	DisplayOrder  int               `json:"display_order"`
	DisplayInList bool              `json:"display_in_list"`
}
