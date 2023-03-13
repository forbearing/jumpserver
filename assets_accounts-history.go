package jumpserver

// AccountHistoryOperator implements Operator interface
// supported operations are: List, Get.
type AccountHistoryOperator struct {
	client *Client
	api    string
}

func (o *AccountHistoryOperator) List(p *AccountHistoryListParam) ([]*AccountHistory, error) {
	return listFactory[*AccountHistory](o.client, o.api, p)
}
func (o *AccountHistoryOperator) Get(historyID string) (*AccountHistory, error) {
	return getFactory[*AccountHistory](o.client, o.api, historyID)
}

type AccountHistory struct {
	SystemuserDisplay string `json:"systemuser_display,omitempty"`
	OrgID             string `json:"org_id,omitempty"`
	Protocols         string `json:"protocols,omitempty"`
	Version           int    `json:"version,omitempty"`
	PublicKey         string `json:"public_key,omitempty"`
	HistoryID         int    `json:"history_id,omitempty"`
	ID                string `json:"id,omitempty"`
	Password          string `json:"password,omitempty"`
	Platform          string `json:"platform,omitempty"`
	Username          string `json:"username,omitempty"`
	Passphrase        string `json:"passphrase,omitempty"`
	DateUpdated       string `json:"date_updated,omitempty"`
	Asset             string `json:"asset,omitempty"`
	IP                string `json:"ip,omitempty"`
	Hostname          string `json:"hostname,omitempty"`
	DateCreated       string `json:"date_created,omitempty"`
	PrivateKey        string `json:"private_key,omitempty"`
	Systemuser        string `json:"systemuser,omitempty"`
}

func (a *AccountHistory) String() string       { return jsonAny(a) }
func (a *AccountHistory) PrettyString() string { return prettyJsonAny(a) }
func (a *AccountHistory) GetName() string      { return "AccountHistory" }
func (a *AccountHistory) GetID() string        { return a.ID }

type AccountHistoryListParam struct {
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

func (p *AccountHistoryListParam) Query() (string, error)         { return concatQuery(p) }
func (p *AccountHistoryListParam) URL(api string) (string, error) { return concatURL(api, p) }
func (p *AccountHistoryListParam) GetID() string                  { return "" } // no historyID
