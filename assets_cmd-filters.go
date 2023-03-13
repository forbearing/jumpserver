package jumpserver

import "net/http"

type CmdFilterOperator struct {
	client *Client
	api    string
}

func (o *CmdFilterOperator) Create(vl ...*CmdFilter) ([]*CmdFilter, error) {
	return createFactory(o.client, o.api, vl)
}
func (o *CmdFilterOperator) Delete(p *CmdFilterDeleteParam) error {
	return deleteFactory[*CmdFilter](o.client, o.api, p)
}
func (o *CmdFilterOperator) Update(vl ...*CmdFilter) ([]*CmdFilter, error) {
	return updateFactory(o.client, o.api, http.MethodPut, vl)
}
func (o *CmdFilterOperator) UpdatePartial(vl ...*CmdFilter) ([]*CmdFilter, error) {
	return updateFactory(o.client, o.api, http.MethodPatch, vl)
}
func (o *CmdFilterOperator) List(p *CmdFilterListParam) ([]*CmdFilter, error) {
	return listFactory[*CmdFilter](o.client, o.api, p)
}
func (o *CmdFilterOperator) Get(id string) (*CmdFilter, error) {
	return getFactory[*CmdFilter](o.client, o.api, id)
}

type CmdFilter struct {
	ID           string   `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	OrgID        string   `json:"org_id,omitempty"`
	OrgName      string   `json:"org_name,omitempty"`
	IsActive     bool     `json:"is_active,omitempty"`
	DateCreated  string   `json:"date_created,omitempty"`
	DateUpdated  string   `json:"date_updated,omitempty"`
	Comment      string   `json:"comment,omitempty"`
	CreatedBy    string   `json:"created_by,omitempty"`
	Rules        []string `json:"rules,omitempty"`
	Users        []string `json:"users,omitempty"`
	UserGroups   []string `json:"user_groups,omitempty"`
	SystemUsers  []string `json:"system_users,omitempty"`
	Assets       []string `json:"assets,omitempty"`
	Applications []string `json:"applications,omitempty"`
}

func (c *CmdFilter) String() string       { return jsonAny(c) }
func (c *CmdFilter) PrettyString() string { return prettyJsonAny(c) }
func (c *CmdFilter) GetName() string      { return "CmdFilter" }
func (c *CmdFilter) GetID() string        { return c.ID }

type CmdFilterDeleteParam struct {
	// list operation not support query by id.
	ID     string `url:"id,omitempty"`
	Name   string `url:"name,omitempty"`
	Search string `url:"search,omitempty"`
	Order  string `url:"order,omitempty"`
	Spm    string `url:"spm,omitempty"`
	Ids    string `url:"ids,omitempty"`
	Limit  int    `url:"limit,omitempty"`
	Offset int    `url:"offset,omitempty"`
}
type CmdFilterListParam CmdFilterDeleteParam

func (p *CmdFilterDeleteParam) Query() (string, error)         { return concatQuery(p) }
func (p *CmdFilterDeleteParam) URL(api string) (string, error) { return concatURL(api, p) }
func (p *CmdFilterDeleteParam) GetID() (api string)            { return p.ID }
func (p *CmdFilterListParam) Query() (string, error)           { return concatQuery(p) }
func (p *CmdFilterListParam) URL(api string) (string, error)   { return concatURL(api, p) }
func (p *CmdFilterListParam) GetID() (api string)              { return p.ID }
