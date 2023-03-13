package jumpserver

// AccountBackupPlanExecutionOperator implements Operator interface
// Supporte operation are: Create, List, Get
type AccountBackupPlanExecutionOperator struct {
	client *Client
	api    string
}

func (o *AccountBackupPlanExecutionOperator) Create(vl ...*AccountBackupPlanExecution) ([]*AccountBackupPlanExecution, error) {
	return createFactory(o.client, o.api, vl)
}
func (o *AccountBackupPlanExecutionOperator) List(p *AccountBackupPlanExecutionListParam) ([]*AccountBackupPlanExecution, error) {
	return listFactory[*AccountBackupPlanExecution](o.client, o.api, p)
}
func (o *AccountBackupPlanExecutionOperator) Get(id string) (*AccountBackupPlanExecution, error) {
	return getFactory[*AccountBackupPlanExecution](o.client, o.api, id)
}

// AccountBackupPlanExecution implements Object interface.
type AccountBackupPlanExecution struct {
	ID             string
	DateStart      string
	Timedelta      string
	PlanSnapshot   any
	Trigger        string // Enum [ manual, timing ]
	Reason         string
	IsSuccess      bool
	Plan           string
	OrgID          string
	Recipients     string
	TriggerDisplay string
}

// String implements fmt.Stringer interface
func (a *AccountBackupPlanExecution) String() string       { return jsonAny(a) }
func (a *AccountBackupPlanExecution) PrettyString() string { return prettyJsonAny(a) }
func (a AccountBackupPlanExecution) GetName() string       { return "AccountBackupPlanExecution" }
func (a *AccountBackupPlanExecution) GetID() string        { return a.ID }

// AccountBackupPlanExecutionListParam implements Parameter interface
type AccountBackupPlanExecutionListParam struct {
	ID      string `url:"id,omitempty"`
	Trigger string `url:"trigger,omitempty"`
	PlanID  string `url:"plan_id,omitempty"`
	Search  string `url:"search,omitempty"`
	Order   string `url:"order,omitempty"`
	Limit   int    `url:"limit,omitempty"`
	Offset  int    `url:"offset,omitempty"`
}

func (p *AccountBackupPlanExecutionListParam) Query() (string, error) { return concatQuery(p) }
func (p *AccountBackupPlanExecutionListParam) URL(api string) (string, error) {
	return concatURL(api, p)
}
func (p *AccountBackupPlanExecutionListParam) GetID() string { return p.ID }
