package jumpserver

import "net/http"

type GatheredUserOperator struct {
	client *Client
	api    string
}

func (o *GatheredUserOperator) Create(vl ...*GatheredUser) ([]*GatheredUser, error) {
	return createFactory(o.client, o.api, vl)
}
func (o *GatheredUserOperator) Delete(p *GatheredUserDeleteParam) error {
	return deleteFactory[*GatheredUser](o.client, o.api, p)
}
func (o *GatheredUserOperator) Update(vl ...*GatheredUser) ([]*GatheredUser, error) {
	return updateFactory(o.client, o.api, http.MethodPut, vl)
}
func (o *GatheredUserOperator) UpdatePartial(vl ...*GatheredUser) ([]*GatheredUser, error) {
	return updateFactory(o.client, o.api, http.MethodPatch, vl)
}
func (o *GatheredUserOperator) List(p *GatheredUserListParam) ([]*GatheredUser, error) {
	return listFactory[*GatheredUser](o.client, o.api, p)
}
func (o *GatheredUserOperator) Get(id string) (*GatheredUser, error) {
	return getFactory[*GatheredUser](o.client, o.api, id)
}

type GatheredUser struct {
	ID            string `json:"id,omitempty"`
	Username      string `json:"username,omitempty"`
	IpLastLogin   string `json:"ip_last_login,omitempty"`
	Present       bool   `json:"present,omitempty"`
	DateLastLogin string `json:"date_last_login,omitempty"`
	DateCreated   string `json:"date_created,omitempty"`
	DateUpdated   string `json:"date_updated,omitempty"`
	Asset         string `json:"asset,omitempty"`
	Hostname      string `json:"hostname,omitempty"`
	IP            string `json:"ip,omitempty"`
	OrgID         string `json:"org_id,omitempty"`
	OrgName       string `json:"org_name,omitempty"`
}

func (g *GatheredUser) String() string       { return jsonAny(g) }
func (g *GatheredUser) PrettyString() string { return prettyJsonAny(g) }
func (g *GatheredUser) GetName() string      { return "GatheredUser" }
func (g *GatheredUser) GetID() string        { return g.ID }

type GatheredUserDeleteParam struct {
	// list operation not support query by id.
	ID         string `url:"id,omitempty"`
	DomainName string `url:"domain_name,omitempty"`
	Name       string `url:"name,omitempty"`
	Username   string `url:"username,omitempty"`
	IP         string `url:"ip,omitempty"`
	Domain     string `url:"domain,omitempty"`
	Search     string `url:"search,omitempty"`
	Order      string `url:"order,omitempty"`
	Spm        string `url:"spm,omitempty"`
	Ids        string `url:"ids,omitempty"`
	Limit      int    `url:"limit,omitempty"`
	Offset     int    `url:"offset,omitempty"`
}
type GatheredUserListParam GatheredUserDeleteParam

func (p *GatheredUserDeleteParam) Query() (string, error)         { return concatQuery(p) }
func (p *GatheredUserDeleteParam) URL(api string) (string, error) { return concatURL(api, p) }
func (p *GatheredUserDeleteParam) GetID() string                  { return p.ID }
func (p *GatheredUserListParam) Query() (string, error)           { return concatQuery(p) }
func (p *GatheredUserListParam) URL(api string) (string, error)   { return concatURL(api, p) }
func (p *GatheredUserListParam) GetID() string                    { return p.ID }
