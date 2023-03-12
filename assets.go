package jumpserver

import (
	"net/http"
)

// AssetExecutor handles communication with the user related method of the Jumpserver API.
type AssetExecutor struct {
	api    string
	client *Client
}

// Create creates jumpserver assets.
// more detail see createFactory function document.
func (s *AssetExecutor) Create(al ...*Asset) ([]*Asset, error) {
	return createFactory(s.client, s.api, al)
}

// Delete deletes jumpserver asset.
// more detail see deleteFactory function document.
func (s *AssetExecutor) Delete(p *DeleteAssetParam) (err error) {
	return deleteFactory[*Asset](s.client, s.api, p)
}

// Update updates jumpserver assets.
// more detail see updateFactory function document.
func (s *AssetExecutor) Update(al ...*Asset) ([]*Asset, error) {
	return updateFactory(s.client, s.api, http.MethodPut, al)
}

// UpdatePartial partial updates jumpserver assets.
func (s *AssetExecutor) UpdatePartial(al ...*Asset) ([]*Asset, error) {
	return updateFactory(s.client, s.api, http.MethodPatch, al)
}

// List lists all jumpserver assets.
// more detail see listFactory function document.
func (s *AssetExecutor) List() ([]*Asset, error) {
	return listFactory[*Asset](s.client, s.api)
}

// Get gets jumpserver assets.
// more detail see getFactory function document.
func (s *AssetExecutor) Get(p *GetAssetParam) ([]*Asset, error) {
	return getFactory[*Asset](s.client, s.api, p)
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

// String
func (s *Asset) String() string {
	data, err := marshal(s, false)
	if err != nil {
		return err.Error()
	}
	return String(data)
}

// PrettyString
func (s *Asset) PrettyString() string {
	data, err := marshal(s, true)
	if err != nil {
		return err.Error()
	}
	return String(data)
}

// GetName returns jumpserver asset name.
func (s *Asset) GetName() string {
	return "Asset"
}

// GetID returns jumpserver asset id.
func (s *Asset) GetID() string {
	return s.ID
}

// GetAssetParam implements Parameter interface.
type GetAssetParam struct {
	// ID 不参与 url query 查询
	// 如果提供了 ID, queyr 只有 ID 一个
	ID string

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

// DeleteAssetParam implements Parameter interface.
type DeleteAssetParam GetAssetParam

func (p *GetAssetParam) GetID() string                     { return p.ID }
func (p *GetAssetParam) Path() (string, error)             { return concatPath(p) }
func (p *GetAssetParam) URL(api string) (string, error)    { return concatURL(api, p) }
func (p *DeleteAssetParam) GetID() string                  { return p.ID }
func (p *DeleteAssetParam) Path() (string, error)          { return concatPath(p) }
func (p *DeleteAssetParam) URL(api string) (string, error) { return concatURL(api, p) }
