package jumpserver

import "net/http"

type DomainOperator struct {
	client *Client
	api    string
}

func (o *DomainOperator) Create(vl ...*Domain) ([]*Domain, error) {
	return createFactory(o.client, o.api, vl)
}
func (o *DomainOperator) Delete(p *DomainDeleteParam) error {
	return deleteFactory[*Domain](o.client, o.api, p)
}
func (o *DomainOperator) Update(vl ...*Domain) ([]*Domain, error) {
	return updateFactory(o.client, o.api, http.MethodPut, vl)
}
func (o *DomainOperator) UpdatePartial(vl ...*Domain) ([]*Domain, error) {
	return updateFactory(o.client, o.api, http.MethodPatch, vl)
}
func (o *DomainOperator) List(p *DomainListParam) ([]*Domain, error) {
	return listFactory[*Domain](o.client, o.api, p)
}
func (o *DomainOperator) Get(id string) (*Domain, error) {
	return getFactory[*Domain](o.client, o.api, id)
}

type Domain struct {
	ID               string   `json:"id,omitempty"`
	Name             string   `json:"name,omitempty"`
	Comment          string   `json:"comment,omitempty"`
	DateCreated      string   `json:"date_created,omitempty"`
	AssetCount       int      `json:"asset_count,omitempty"`
	Assets           []string `json:"assets,omitempty"`
	ApplicationCount int      `json:"application_count,omitempty"`
	GatewayCount     int      `json:"gateway_count,omitempty"`
	OrgID            string   `json:"org_id,omitempty"`
	OrgName          string   `json:"org_name,omitempty"`
}

func (d *Domain) String() string       { return jsonAny(d) }
func (d *Domain) PrettyString() string { return prettyJsonAny(d) }
func (d *Domain) GetName() string      { return "Domain" }
func (d *Domain) GetID() string        { return d.ID }

type DomainDeleteParam struct {
	// list operation not support query by id.
	ID     string `url:"id,omitempty"`
	Name   string `url:"name,omitempty"`
	Search string `url:"search" json:"search,omitempty"`
	Order  string `url:"order" json:"order,omitempty"`
	Spm    string `url:"spm" json:"spm,omitempty"`
	Ids    string `url:"ids" json:"ids,omitempty"`
	Limit  int    `url:"limit" json:"limit,omitempty"`
	Offset int    `url:"offset" json:"offset,omitempty"`
}
type DomainListParam DomainDeleteParam

func (p *DomainDeleteParam) Query() (string, error)         { return concatQuery(p) }
func (p *DomainDeleteParam) URL(api string) (string, error) { return concatURL(api, p) }
func (p *DomainDeleteParam) GetID() string                  { return p.ID }
func (p *DomainListParam) Query() (string, error)           { return concatQuery(p) }
func (p *DomainListParam) URL(api string) (string, error)   { return concatURL(api, p) }
func (p *DomainListParam) GetID() string                    { return p.ID }
