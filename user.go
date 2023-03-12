package jumpserver

import (
	"net/url"
)

// UserNode
type UserNode struct {
	client *Client
	api    string
}

func (n *UserNode) ConnectionToken() *ConnectionTokenOperator {
	api, _ := url.JoinPath(n.api, "connection-token", "/")
	return &ConnectionTokenOperator{client: n.client, api: api}
}
func (n *UserNode) UserGroup() *UserGroupOperator {
	api, _ := url.JoinPath(n.api, "groups", "/")
	return &UserGroupOperator{client: n.client, api: api}
}
func (n *UserNode) UserProfile() *UserProfileOperator {
	api, _ := url.JoinPath(n.api, "profile", "/")
	return &UserProfileOperator{client: n.client, api: api}
}
func (n *UserNode) UserGroupRelation() *UserGroupRelationOperator {
	api, _ := url.JoinPath(n.api, "users-groups-relations", "/")
	return &UserGroupRelationOperator{client: n.client, api: api}
}
func (n *UserNode) User() *UserOperator {
	api, _ := url.JoinPath(n.api, "users", "/")
	return &UserOperator{client: n.client, api: api}
}
