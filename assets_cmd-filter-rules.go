package jumpserver

// CmdFilterRuleOperator implements Operator interface.
// supported operations are: List.
type CmdFilterRuleOperator struct {
	client *Client
	api    string
}

func (o *CmdFilterRuleOperator) List(p *CmdFilterRuleListParam) ([]*CmdFilterRule, error) {
	return listFactory[*CmdFilterRule](o.client, o.api, p)
}

type CmdFilterRule struct {
	ID            string   `json:"id,omitempty"`
	Type          string   `json:"type,omitempty"`
	TypeDisplay   string   `json:"type_display,omitempty"`
	Content       string   `json:"content,omitempty"`
	IgnoreCase    bool     `json:"ignore_case,omitempty"`
	Pattern       string   `json:"pattern,omitempty"`
	Priority      int      `json:"priority,omitempty"`
	Action        int      `json:"action,omitempty"`
	ActionDisplay string   `json:"action_display,omitempty"`
	Reviewers     []string `json:"reviewers,omitempty"`
	DateCreated   string   `json:"date_created,omitempty"`
	DateUpdated   string   `json:"date_updated,omitempty"`
	Comment       string   `json:"comment,omitempty"`
	CreatedBy     string   `json:"created_by,omitempty"`
	Filter        string   `json:"filter,omitempty"`
	OrgID         string   `json:"org_id,omitempty"`
	OrgName       string   `json:"org_name,omitempty"`
}

func (c *CmdFilterRule) String() string       { return jsonAny(c) }
func (c *CmdFilterRule) PrettyString() string { return prettyJsonAny(c) }
func (c *CmdFilterRule) GetName() string      { return "CmdFilter" }
func (c *CmdFilterRule) GetID() string        { return c.ID }

type CmdFilterRuleListParam struct {
	Search string `url:"search,omitempty"`
	Order  string `url:"order,omitempty"`
	Limit  int    `url:"limit,omitempty"`
	Offset int    `url:"offset,omitempty"`
}

func (p *CmdFilterRuleListParam) Query() (string, error)         { return concatQuery(p) }
func (p *CmdFilterRuleListParam) URL(api string) (string, error) { return concatURL(api, p) }
func (p *CmdFilterRuleListParam) GetID() string                  { return "" }
