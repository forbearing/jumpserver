package jumpserver

// TempService handles communication with the user related method of the Jumpserver API.
type TempService struct {
	api    string
	client *Client
}

func (u *TempService) Create(p *CreateTempParam) (*User, error) { return nil, nil }
func (u *TempService) Delete(p *DeleteTempParam) (*User, error) { return nil, nil }
func (u *TempService) Update(p *UpdateTempParam) (*User, error) { return nil, nil }
func (u *TempService) List(p *ListTempParam) ([]*User, error)   { return nil, nil }
func (u *TempService) Get(p *GetTempParam) (*User, error)       { return nil, nil }

type CreateTempParam struct{}
type DeleteTempParam struct{}
type UpdateTempParam struct{}
type ListTempParam struct{}
type GetTempParam struct{}
