package jumpserver

import "net/http"

type AccountOperator struct {
	client *Client
	api    string
}

func (o *AccountOperator) Create(vl ...*Account) ([]*Account, error) {
	return createFactory(o.client, o.api, vl)
}
func (o *AccountOperator) Delete(p *AccountDeleteParam) error {
	return deleteFactory[*Account](o.client, o.api, p)
}
func (o *AccountOperator) Update(vl ...*Account) ([]*Account, error) {
	return updateFactory(o.client, o.api, http.MethodPut, vl)
}
func (o *AccountOperator) UpdatePartial(vl ...*Account) ([]*Account, error) {
	return updateFactory(o.client, o.api, http.MethodPatch, vl)
}
func (o *AccountOperator) List(p *AccountListParam) ([]*Account, error) {
	return listFactory[*Account](o.client, o.api, p)
}

type Account struct {
	ID                string `json:"id,omitempty"`
	Username          string `json:"username,omitempty"`
	IP                string `json:"ip,omitempty"`
	Hostname          string `json:"hostname,omitempty"`
	Platform          string `json:"platform,omitempty"`
	Protocols         string `json:"protocols,omitempty"`
	Version           int    `json:"version,omitempty"`
	Password          string `json:"password,omitempty"`
	PrivateKey        string `json:"private_key,omitempty"`
	PublicKey         string `json:"public_key,omitempty"`
	Passphrase        string `json:"passphrase,omitempty"`
	DateCreated       string `json:"date_created,omitempty"`
	DateUpdated       string `json:"date_updated,omitempty"`
	Connectivity      string `json:"connectivity,omitempty"` // Enum [unknown, ok, failed]
	DateVerified      string `json:"date_verified,omitempty"`
	Comment           string `json:"comment,omitempty"`
	Asset             string `json:"asset,omitempty"`
	Systemuser        string `json:"systemuser,omitempty"`
	SystemuserDisplay string `json:"systemuser_display,omitempty"`
	OrgID             string `json:"org_id,omitempty"`
	OrgName           string `json:"org_name,omitempty"`
}

func (a *Account) String() string       { return jsonAny(a) }
func (a *Account) PrettyString() string { return prettyJsonAny(a) }
func (a *Account) GetName() string      { return "Account" }
func (a *Account) GetID() string        { return a.ID }

type AccountDeleteParam struct {
	Asset      string `url:"asset,omitempty"`
	Systemuser string `url:"systemuser,omitempty"`
	ID         string `url:"id,omitempty"`
	Username   string `url:"username,omitempty"`
	IP         string `url:"ip,omitempty"`
	Hostname   string `url:"hostname,omitempty"`
	Node       string `url:"node,omitempty"`
	Search     string `url:"search,omitempty"`
	Order      string `url:"order,omitempty"`
	Spm        string `url:"spm,omitempty"`
	Ids        string `url:"ids,omitempty"`
	Limit      int    `url:"limit,omitempty"`
	Offset     int    `url:"offset,omitempty"`
}
type AccountListParam AccountDeleteParam

func (p *AccountDeleteParam) Query() (string, error)         { return concatQuery(p) }
func (p *AccountDeleteParam) URL(api string) (string, error) { return concatURL(api, p) }
func (p *AccountDeleteParam) GetID() string                  { return "" }
func (p *AccountListParam) Query() (string, error)           { return concatQuery(p) }
func (p *AccountListParam) URL(api string) (string, error)   { return concatURL(api, p) }
func (p *AccountListParam) GetID() string                    { return "" }
