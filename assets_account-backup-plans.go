package jumpserver

import "net/http"

// AccountBackupPlanOperator implements Operator interface
type AccountBackupPlanOperator struct {
	client *Client
	api    string
}

func (o *AccountBackupPlanOperator) Create(vl ...*AccountBackupPlan) ([]*AccountBackupPlan, error) {
	return createFactory(o.client, o.api, vl)
}
func (o *AccountBackupPlanOperator) Delete(p *AccountBackupPlanDeleteParam) error {
	return deleteFactory[*AccountBackupPlan](o.client, o.api, p)
}
func (o *AccountBackupPlanOperator) Update(vl ...*AccountBackupPlan) ([]*AccountBackupPlan, error) {
	return updateFactory(o.client, o.api, http.MethodPut, vl)
}
func (o *AccountBackupPlanOperator) UpdatePartial(vl ...*AccountBackupPlan) ([]*AccountBackupPlan, error) {
	return updateFactory(o.client, o.api, http.MethodPatch, vl)
}
func (o *AccountBackupPlanOperator) List(p *AccountBackupPlanListParam) ([]*AccountBackupPlan, error) {
	return listFactory[*AccountBackupPlan](o.client, o.api, p)
}
func (o *AccountBackupPlanOperator) Get(id string) (*AccountBackupPlan, error) {
	return getFactory[*AccountBackupPlan](o.client, o.api, id)
}

// AccountBackupPlan implements Object interface
type AccountBackupPlan struct {
	ID              string   `json:"id,omitempty"`
	Name            string   `json:"name,omitempty"`
	IsPeriodic      bool     `json:"is_periodic,omitempty"`
	Interval        int      `json:"interval,omitempty"`
	Crontab         string   `json:"crontab,omitempty"`
	DateCreated     string   `json:"date_created,omitempty"`
	DateUpdated     string   `json:"date_updated,omitempty"`
	CreatedBy       string   `json:"created_by,omitempty"`
	PeriodicDisplay string   `json:"periodic_display,omitempty"`
	Comment         string   `json:"comment,omitempty"`
	Recipients      []string `json:"recipients,omitempty"`
	Types           []string `json:"types,omitempty"` // Enum [ all, asset, application ]
	OrgID           string   `json:"org_id,omitempty"`
	OrgName         string   `json:"org_name,omitempty"`
}

func (a *AccountBackupPlan) String() string       { return jsonAny(a) }
func (a *AccountBackupPlan) PrettyString() string { return prettyJsonAny(a) }
func (a *AccountBackupPlan) GetName() string {
	return "AccountBackupPlan"
}
func (a *AccountBackupPlan) GetID() string {
	return a.ID
}

// AccountBackupPlanDeleteParam implements Parameter interface.
type AccountBackupPlanDeleteParam struct {
	ID     string `url:"id,omitempty"`
	Name   string `url:"name,omitempty"`
	Search string `url:"search,omitempty"`
	Order  string `url:"order,omitempty"`
	Spm    string `url:"spm,omitempty"`
	Ids    string `url:"ids,omitempty"`
	Limit  int    `url:"limit,omitempty"`
	Offset int    `url:"offset,omitempty"`
}

// AccountBackupPlanListParam implements Parameter interface.
// NOTE: list operation not support query by id.
type AccountBackupPlanListParam AccountBackupPlanDeleteParam

func (p *AccountBackupPlanDeleteParam) Query() (string, error)         { return concatQuery(p) }
func (p *AccountBackupPlanDeleteParam) URL(api string) (string, error) { return concatURL(api, p) }
func (p *AccountBackupPlanDeleteParam) GetID() string                  { return p.ID }
func (p *AccountBackupPlanListParam) Query() (string, error)           { return concatQuery(p) }
func (p *AccountBackupPlanListParam) URL(api string) (string, error)   { return concatURL(api, p) }
func (p *AccountBackupPlanListParam) GetID() string                    { return p.ID }
