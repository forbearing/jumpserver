package jumpserver

import (
	"encoding/json"
)

// NodeOperator handles communication with the user related method of the Jumpserver API.
type NodeOperator struct {
	api    string
	client *Client
}

func (o *NodeOperator) Create(p *CreateNodeParam) (*Node, error) { return nil, nil }
func (o *NodeOperator) Delete(p *DeleteNodeParam) (*Node, error) { return nil, nil }
func (o *NodeOperator) Update(p *UpdateNodeParam) (*Node, error) { return nil, nil }
func (o *NodeOperator) List() ([]*Node, error) {
	data, err := o.client.get(o.api)
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
func (o *NodeOperator) Get(p *GetNodeParam) (*Node, error) {
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

func (o *Node) String() string {
	data, err := marshal(o, false)
	if err != nil {
		return ""
	}
	return String(data)
}
func (o *Node) PrettyString() string {
	data, err := marshal(o, true)
	if err != nil {
		return ""
	}
	return String(data)
}
func (o *Node) GetName() string {
	return o.Name
}
func (o *Node) GetID() string {
	return o.ID
}

type CreateNodeParam struct {
}
type DeleteNodeParam struct{}
type UpdateNodeParam struct{}

// GetNodeParam
type GetNodeParam struct {
}
