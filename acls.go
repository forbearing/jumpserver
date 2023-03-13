package jumpserver

import "net/url"

type ACLNode struct {
	client *Client
	api    string
}

func (n *ACLNode) LoginACL() *LoginACLOperator {
	api, _ := url.JoinPath(n.api, "/login-acls/")
	return &LoginACLOperator{client: n.client, api: api}
}
func (n *ACLNode) LoginAssetACL() *LoginAssetACLOperator {
	api, _ := url.JoinPath(n.api, "/login-asset-acls/")
	return &LoginAssetACLOperator{client: n.client, api: api}
}
func (n *ACLNode) LoginAsset() *LoginAssetOperator {
	api, _ := url.JoinPath(n.api, "/login-asset/")
	return &LoginAssetOperator{client: n.client, api: api}
}
