package jumpserver

import (
	"encoding/json"
)

// UserOperator handles communication with the user related method of the Jumpserver API.
type UserOperator struct {
	api     string
	client  *Client
	rawData []byte
}

// Create creates Jumpserver user.
func (o *UserOperator) Create() (*User, error) { return nil, nil }

// Delete delete Jumpserver user.
func (o *UserOperator) Delete() (*User, error) { return nil, nil }

// Update update Jumpserver user.
func (o *UserOperator) Update() (*User, error) { return nil, nil }

// List list Jumpserver all users.
func (o *UserOperator) List() ([]*User, error) {
	data, err := o.client.get(o.api)
	o.rawData = data
	if err != nil {
		return nil, err
	}
	var users []*User
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}
	return users, nil
}

// Get gets Jumpserver user.
func (o *UserOperator) Get() (*User, error) { return nil, nil }

// User represents a Jumpserver user.
type User struct {
	ID                      string   `json:"id,omitempty"`
	Name                    string   `json:"name,omitempty"`
	Username                string   `json:"username,omitempty"`
	Password                string   `json:"password,omitempty"`
	PublicKey               string   `json:"public_key,omitempty"`
	Email                   string   `json:"email,omitempty"`
	Wechat                  string   `json:"wechat,omitempty"`
	Phone                   string   `json:"phone,omitempty"`
	MfaLevel                int      `json:"mfa_level,omitempty"`
	Source                  string   `json:"source,omitempty"`
	SourceDisplay           string   `json:"source_display,omitempty"`
	CanPublicKeyAuth        bool     `json:"can_public_key_auth,omitempty"`
	NeedUpdatePassword      bool     `json:"need_update_password,omitempty"`
	MfaEnabled              bool     `json:"mfa_enabled,omitempty"`
	IsServiceAccount        bool     `json:"is_service_account,omitempty"`
	IsValid                 bool     `json:"is_valid,omitempty"`
	IsExpired               bool     `json:"is_expired,omitempty"`
	IsActive                bool     `json:"is_active,omitempty"`
	DateExpired             string   `json:"date_expired,omitempty"`
	DateJoined              string   `json:"date_joined,omitempty"`
	LastLogin               string   `json:"last_login,omitempty"`
	CreatedBy               string   `json:"created_by,omitempty"`
	Comment                 string   `json:"comment,omitempty"`
	IsWecomBound            bool     `json:"is_wecom_bound,omitempty"`
	IsDingtalkBound         bool     `json:"is_dingtalk_bound,omitempty"`
	IsFeishuBound           bool     `json:"is_feishu_bound,omitempty"`
	IsOtpSecretKeyBound     bool     `json:"is_otp_secret_key_bound,omitempty"`
	WecomID                 string   `json:"wecom_id,omitempty"`
	DingtalkID              string   `json:"dingtalk_id,omitempty"`
	FeishuID                string   `json:"feishu_id,omitempty"`
	MfaLevelDisplay         string   `json:"mfa_level_display,omitempty"`
	MfaForceEnabled         bool     `json:"mfa_force_enabled,omitempty"`
	IsFirstLogin            bool     `json:"is_first_login,omitempty"`
	DatePasswordLastUpdated string   `json:"date_password_last_updated,omitempty"`
	AvatarUrl               string   `json:"avatar_url,omitempty"`
	Groups                  []string `json:"groups,omitempty"`
	GroupsDisplay           string   `json:"groups_display,omitempty"`
	SystemRoles             []string `json:"system_roles,omitempty"`
	OrgRoles                []string `json:"org_roles,omitempty"`
	SystemRolesDisplay      string   `json:"system_roles_display,omitempty"`
	OrgRolesDisplay         string   `json:"org_roles_display,omitempty"`
	LoginBlocked            bool     `json:"login_blocked,omitempty"`
	PasswordStrategy        string   `json:"password_strategy,omitempty"`
}

// String
func (u *User) String() string {
	data, err := marshal(u)
	if err != nil {
		return err.Error()
	}
	return String(data)
}

// PrettyString
func (u *User) PrettyString() string {
	data, err := marshal(u, true)
	if err != nil {
		return err.Error()
	}
	return String(data)
}
func (u *User) GetName() string {
	return "User"
}
func (u *User) GetID() string {
	return u.ID
}

// DeleteUserParam implements Parameter interface.
type DeleteUserParam struct {
	ID          string
	Username    string
	Email       string
	Name        string
	Source      string
	OrgRoles    string
	SystemRoles string
	IsActive    bool
	Search      string
	Order       string
	Spm         string
	Ids         string
	Limit       string
	Offset      string
}

// GetUserParam implements Parameter interface.
type GetUserParam DeleteUserParam

func (p *DeleteUserParam) Query() (string, error)         { return concatQuery(p) }
func (p *DeleteUserParam) URL(api string) (string, error) { return concatURL(api, p) }
func (p *DeleteUserParam) GetID() string                  { return p.ID }
func (p *GetUserParam) Query() (string, error)            { return concatQuery(p) }
func (p *GetUserParam) URL(api string) (string, error)    { return concatURL(api, p) }
func (p *GetUserParam) GetID() string                     { return p.ID }
