package jumpserver

import (
	"encoding/json"
)

// NodeService handles communication with the user related method of the Jumpserver API.
type NodeService struct {
	api    string
	client *Client
}

func (n *NodeService) Create(p *CreateNodeParam) (*Node, error) { return nil, nil }
func (n *NodeService) Delete(p *DeleteNodeParam) (*Node, error) { return nil, nil }
func (n *NodeService) Update(p *UpdateNodeParam) (*Node, error) { return nil, nil }
func (n *NodeService) List() ([]*Node, error) {
	data, err := n.client.get(n.api)
	if err != nil {
		return nil, err
	}
	var nodes []*Node
	if err := json.Unmarshal(data, &nodes); err != nil {
		return nil, err
	}
	return nodes, nil
}

// Get
func (n *NodeService) Get(p *GetNodeParam) (*Node, error) {
	return nil, nil
}

type Node struct {
	ID        string `json:"id,omitempty"`
	Key       string `json:"key,omitempty"`
	Value     string `json:"value,omitempty"`
	OrgID     string `json:"org_id,omitempty"`
	Name      string `json:"name,omitempty"`
	FullValue string `json:"full_value,omitempty"`
	OrgName   string `json:"org_name,omitempty"`
}

func (n *Node) String() string {
	data, err := marshal(n, false)
	if err != nil {
		return ""
	}
	return String(data)
}
func (n *Node) PrettyString() string {
	data, err := marshal(n, true)
	if err != nil {
		return ""
	}
	return String(data)
}
func (n *Node) GetName() string {
	return n.Name
}
func (n *Node) GetID() string {
	return n.ID
}

type CreateNodeParam struct {
}
type DeleteNodeParam struct{}
type UpdateNodeParam struct{}

// GetNodeParam
type GetNodeParam struct {
}
