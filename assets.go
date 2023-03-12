package jumpserver

import (
	"net/url"
)

type AssetNode struct {
	client *Client
	api    string
}

func (n *AssetNode) AccountBackupPlanExecution() *AccountBackupPlanExecutionOperator {
	api, _ := url.JoinPath(n.api, "/account-backup-plan-executions/")
	return &AccountBackupPlanExecutionOperator{client: n.client, api: api}
}
func (n *AssetNode) AccountBackup() *AccountBackupPlanOperator {
	api, _ := url.JoinPath(n.api, "/account-backup-plans/")
	return &AccountBackupPlanOperator{client: n.client, api: api}
}
func (n *AssetNode) AccountHistory() *AccountHistoryOperator {
	api, _ := url.JoinPath(n.api, "/accounts-history/")
	return &AccountHistoryOperator{client: n.client, api: api}
}
func (n *AssetNode) Account() *AccountOperator {
	api, _ := url.JoinPath(n.api, "/accounts/")
	return &AccountOperator{client: n.client, api: api}
}
func (n *AssetNode) Asset() *AssetOperator {
	api, _ := url.JoinPath(n.api, "/assets/")
	return &AssetOperator{client: n.client, api: api}
}
func (n *AssetNode) CmdFilterRule() *CmdFilterRuleOperator {
	api, _ := url.JoinPath(n.api, "/cmd-filter-rules/")
	return &CmdFilterRuleOperator{client: n.client, api: api}
}
func (n *AssetNode) CmdFilter() *CmdFilterOperator {
	api, _ := url.JoinPath(n.api, "/cmd-filters/")
	return &CmdFilterOperator{client: n.client, api: api}
}
func (n *AssetNode) Domain() *DomainOperator {
	api, _ := url.JoinPath(n.api, "/domains/")
	return &DomainOperator{client: n.client, api: api}
}
func (n *AssetNode) FavoriteAsset() *FavoriteAssetOperator {
	api, _ := url.JoinPath(n.api, "/favorite-assets/")
	return &FavoriteAssetOperator{client: n.client, api: api}
}
func (n *AssetNode) Gateway() *GatewayOperator {
	api, _ := url.JoinPath(n.api, "/favorite-assets/")
	return &GatewayOperator{client: n.client, api: api}
}
func (n *AssetNode) GatheredUser() *GatheredUserOperator {
	api, _ := url.JoinPath(n.api, "/gathered-users/")
	return &GatheredUserOperator{client: n.client, api: api}
}
func (n *AssetNode) Label() *LabelOperator {
	api, _ := url.JoinPath(n.api, "/labels/")
	return &LabelOperator{client: n.client, api: api}
}
func (n *AssetNode) Node() *NodeOperator {
	api, _ := url.JoinPath(n.api, "/nodes/")
	return &NodeOperator{client: n.client, api: api}
}
func (n *AssetNode) Platform() *PlatformOperator {
	api, _ := url.JoinPath(n.api, "/platforms/")
	return &PlatformOperator{client: n.client, api: api}
}
