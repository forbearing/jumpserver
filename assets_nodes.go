package jumpserver

// NodeOperator handles communication with the user related method of the Jumpserver API.
type NodeOperator struct {
	api    string
	client *Client
}

func (o *NodeOperator) Create(vl ...*Node) ([]*Node, error)        { return nil, nil }
func (o *NodeOperator) Delete(p *DeleteNodeParam) error            { return nil }
func (o *NodeOperator) Update(vl ...*Node) ([]*Node, error)        { return nil, nil }
func (o *NodeOperator) UpdatePartial(vl ...*Node) ([]*Node, error) { return nil, nil }
func (o *NodeOperator) List(p *ListNodeParam) ([]*Node, error)     { return nil, nil }
func (o *NodeOperator) Get(id string) (*Node, error)               { return nil, nil }

type Node struct {
	ID        string `json:"id,omitempty"`
	Key       string `json:"key,omitempty"`
	Value     string `json:"value,omitempty"`
	OrgID     string `json:"org_id,omitempty"`
	Name      string `json:"name,omitempty"`
	FullValue string `json:"full_value,omitempty"`
	OrgName   string `json:"org_name,omitempty"`
}

func (n *Node) String() string       { return jsonAny(n) }
func (n *Node) PrettyString() string { return prettyJsonAny(n) }
func (n *Node) GetName() string      { return "Node" }
func (n *Node) GetID() string        { return n.ID }

type DeleteNodeParam struct{}
type ListNodeParam struct{}
