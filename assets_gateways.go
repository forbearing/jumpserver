package jumpserver

import "net/http"

type GatewayOperator struct {
	client *Client
	api    string
}

func (o *GatewayOperator) Create(vl ...*Gateway) ([]*Gateway, error) {
	return createFactory(o.client, o.api, vl)
}
func (o *GatewayOperator) Delete(p *GatewayDeleteParam) error {
	return deleteFactory[*Gateway](o.client, o.api, p)
}
func (o *GatewayOperator) Update(vl ...*Gateway) ([]*Gateway, error) {
	return updateFactory(o.client, o.api, http.MethodPut, vl)
}
func (o *GatewayOperator) UpdatePartial(vl ...*Gateway) ([]*Gateway, error) {
	return updateFactory(o.client, o.api, http.MethodPatch, vl)
}
func (o *GatewayOperator) List(p *GatewayListParam) ([]*Gateway, error) {
	return listFactory[*Gateway](o.client, o.api, p)
}
func (o *GatewayOperator) Get(id string) (*Gateway, error) {
	return getFactory[*Gateway](o.client, o.api, id)
}

type Gateway struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Password     string `json:"password,omitempty"`
	PrivateKey   string `json:"private_key,omitempty"`
	PublicKey    string `json:"public_key,omitempty"`
	Passphrase   string `json:"passphrase,omitempty"`
	Username     string `json:"username,omitempty"`
	IP           string `json:"ip,omitempty"`
	Port         int    `json:"port,omitempty"`
	Protocol     string `json:"protocol,omitempty"`
	IsActive     bool   `json:"is_active,omitempty"`
	IsConnective bool   `json:"is_connective,omitempty"`
	DateCreated  string `json:"date_created,omitempty"`
	DateUpdated  string `json:"date_updated,omitempty"`
	CreatedBy    string `json:"created_by,omitempty"`
	Comment      string `json:"comment,omitempty"`
	Domain       string `json:"domain,omitempty"`
	OrgID        string `json:"org_id,omitempty"`
	OrgName      string `json:"org_name,omitempty"`
}

func (g *Gateway) String() string       { return jsonAny(g) }
func (g *Gateway) PrettyString() string { return prettyJsonAny(g) }
func (g *Gateway) GetName() string      { return "Gateway" }
func (g *Gateway) GetID() string        { return g.ID }

type GatewayDeleteParam struct {
	// list operator not support query by id
	ID         string `url:"id,omitempty"`
	DomainName string `url:"domain__name,omitempty"`
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
type GatewayListParam GatewayDeleteParam

func (p *GatewayDeleteParam) Query() (string, error)         { return concatQuery(p) }
func (p *GatewayDeleteParam) URL(api string) (string, error) { return concatURL(api, p) }
func (p *GatewayDeleteParam) GetID() string                  { return p.ID }
func (p *GatewayListParam) Query() (string, error)           { return concatQuery(p) }
func (p *GatewayListParam) URL(api string) (string, error)   { return concatURL(api, p) }
func (p *GatewayListParam) GetID() string                    { return p.ID }
