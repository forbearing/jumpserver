package jumpserver

import (
	"net/http"
)

// AssetOperator handles communication with the user related method of the Jumpserver API.
type AssetOperator struct {
	api    string
	client *Client
}

// Create creates jumpserver assets.
// more detail see createFactory function document.
func (o *AssetOperator) Create(vl ...*Asset) ([]*Asset, error) {
	return createFactory(o.client, o.api, vl)
}

// Delete deletes jumpserver asset.
// more detail see deleteFactory function document.
func (o *AssetOperator) Delete(p *DeleteAssetParam) error {
	return deleteFactory[*Asset](o.client, o.api, p)
}

// Update updates jumpserver assets.
// more detail see updateFactory function document.
func (o *AssetOperator) Update(vl ...*Asset) ([]*Asset, error) {
	return updateFactory(o.client, o.api, http.MethodPut, vl)
}

// UpdatePartial partial updates jumpserver assets.
func (o *AssetOperator) UpdatePartial(vl ...*Asset) ([]*Asset, error) {
	return updateFactory(o.client, o.api, http.MethodPatch, vl)
}

// List lists all jumpserver assets.
// more detail see listFactory function document.
func (o *AssetOperator) List(p *ListAssetParam) ([]*Asset, error) {
	return listFactory[*Asset](o.client, o.api, p)
}

// Get gets jumpserver assets.
// more detail see getFactory function document.
func (o *AssetOperator) Get(id string) (*Asset, error) {
	return getFactory[*Asset](o.client, o.api, id)
}

// Asset
type Asset struct {
	ID               string   `json:"id,omitempty"`
	Hostname         string   `json:"hostname,omitempty"`
	IP               string   `json:"ip,omitempty"`
	Platform         Platform `json:"platform,omitempty"`
	Protocols        []string `json:"protocols,omitempty"`
	Protocol         string   `json:"protocol,omitempty"`
	Port             int      `json:"port,omitempty"`
	IsActive         bool     `json:"is_active,omitempty"`
	PublicIP         string   `json:"public_ip,omitempty"`
	Number           string   `json:"number,omitempty"`
	Comment          string   `json:"comment,omitempty"`
	Vendor           string   `json:"vendor,omitempty"`
	Model            string   `json:"model,omitempty"`
	SN               string   `json:"sn,omitempty"`
	CpuModel         string   `json:"cpu_model,omitempty"`
	CpuCount         int      `json:"cpu_count,omitempty"`
	CpuCores         int      `json:"cpu_cores,omitempty"`
	CpuVcpus         int      `json:"cpu_vcpus,omitempty"`
	Memory           string   `json:"memory,omitempty"`
	DiskTotal        string   `json:"disk_total,omitempty"`
	DiskInfo         string   `json:"disk_info,omitempty"`
	OS               string   `json:"os,omitempty"`
	OsVersion        string   `json:"os_version,omitempty"`
	OsArch           string   `json:"os_arch,omitempty"`
	HostnameRaw      string   `json:"hostname_raw,omitempty"`
	CpuInfo          string   `json:"cpu_info,omitempty"`
	HardwareInfo     string   `json:"hardware_info,omitempty"`
	Domain           string   `json:"domain,omitempty"`
	DomainDisplay    string   `json:"domain_display,omitempty"`
	AdminUser        string   `json:"admin_user,omitempty"`
	AdminUserDisplay string   `json:"admin_user_display,omitempty"`
	Nodes            []string `json:"nodes,omitempty"`
	NodesDisplay     []string `json:"nodes_display,omitempty"`
	Labels           []string `json:"labels,omitempty"`
	LabelsDisplay    []string `json:"labels_display,omitempty"`
	Connectivity     string   `json:"connectivity,omitempty"`
	DateVerified     string   `json:"date_verified,omitempty"`
	CreatedBy        string   `json:"created_by,omitempty"`
	DateCreated      string   `json:"date_created,omitempty"`
	OrgID            string   `json:"org_id,omitempty"`
	OrgName          string   `json:"org_name,omitempty"`
}

// String implements fmt.Stringer interface.
func (a *Asset) String() string { return jsonAny(a) }

// PrettyString
func (a *Asset) PrettyString() string { return prettyJsonAny(a) }

// GetName returns jumpserver asset name.
func (a *Asset) GetName() string { return "Asset" }

// GetID returns jumpserver asset id.
func (a *Asset) GetID() string { return a.ID }

// DeleteAssetParam implements Parameter interface.
type DeleteAssetParam struct {
	// List operation not support query by ID
	// Delete operation is supported query by ID.
	ID                 string `url:"id,omitempty"`
	Hostname           string `url:"hostname,omitempty"`
	IP                 string `url:"ip,omitempty"`
	SystemUsersID      string `url:"system_users__id,omitempty"`
	PlatformBase       string `url:"platform__base,omitempty"`
	IsActive           bool   `url:"is_active,omitempty"`
	Protocols          string `url:"protocols,omitempty"`
	ProtocolsIcontains string `url:"protocols__icontains,omitempty"`
	Search             string `url:"search,omitempty"`
	Order              string `url:"order,omitempty"`
	Spm                string `url:"spm,omitempty"`
	Ids                string `url:"ids,omitempty"`
	Node               string `url:"node,omitempty"`
	All                string `url:"all,omitempty"`
	Label              string `url:"label,omitempty"`
	Ips                string `url:"ips,omitempty"`
	Limit              string `url:"limit,omitempty"`
	Offset             int    `url:"offset,omitempty"`
}

// ListAssetParam implements Parameter interface.
type ListAssetParam DeleteAssetParam

func (p *DeleteAssetParam) Query() (string, error)         { return concatQuery(p) }
func (p *DeleteAssetParam) URL(api string) (string, error) { return concatURL(api, p) }
func (p *DeleteAssetParam) GetID() string                  { return p.ID }
func (p *ListAssetParam) Query() (string, error)           { return concatQuery(p) }
func (p *ListAssetParam) URL(api string) (string, error)   { return concatURL(api, p) }
func (p *ListAssetParam) GetID() string                    { return p.ID }
