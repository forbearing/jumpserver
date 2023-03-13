package jumpserver

import "net/http"

type FavoriteAssetOperator struct {
	client *Client
	api    string
}

type FavoriteAsset struct {
	Asset string `json:"asset,omitempty"`
}

func (o *FavoriteAssetOperator) Create(vl ...*FavoriteAsset) ([]*FavoriteAsset, error) {
	return createFactory(o.client, o.api, vl)
}
func (o *FavoriteAssetOperator) Delete(p *FavoriteAssetDeleteParam) error {
	return deleteFactory[*FavoriteAsset](o.client, o.api, p)
}
func (o *FavoriteAssetOperator) Update(vl ...*FavoriteAsset) ([]*FavoriteAsset, error) {
	return updateFactory(o.client, o.api, http.MethodPut, vl)
}
func (o *FavoriteAssetOperator) UpdatePartial(vl ...*FavoriteAsset) ([]*FavoriteAsset, error) {
	return updateFactory(o.client, o.api, http.MethodPatch, vl)
}
func (o *FavoriteAssetOperator) List(p *FavoriteAssetListParam) ([]*FavoriteAsset, error) {
	return listFactory[*FavoriteAsset](o.client, o.api, p)
}

func (f *FavoriteAsset) String() string       { return jsonAny(f) }
func (f *FavoriteAsset) PrettyString() string { return prettyJsonAny(f) }
func (f *FavoriteAsset) GetName() string      { return "FavoriteAsset" }
func (f *FavoriteAsset) GetID() string        { return "" }

type FavoriteAssetDeleteParam struct {
	// list operation not support query by id.
	ID     string `url:"id,omitempty"`
	Asset  string `url:"asset,omitempty"`
	Search string `url:"search,omitempty"`
	Order  string `url:"order,omitempty"`
	Limit  string `url:"limit,omitempty"`
	Offset string `url:"offset,omitempty"`
}
type FavoriteAssetListParam FavoriteAssetDeleteParam

func (p *FavoriteAssetDeleteParam) Query() (string, error)         { return concatQuery(p) }
func (p *FavoriteAssetDeleteParam) URL(api string) (string, error) { return concatURL(api, p) }
func (p *FavoriteAssetDeleteParam) GetID() string                  { return p.ID }
func (p *FavoriteAssetListParam) Query() (string, error)           { return concatQuery(p) }
func (p *FavoriteAssetListParam) URL(api string) (string, error)   { return concatURL(api, p) }
func (p *FavoriteAssetListParam) GetID() string                    { return p.ID }
