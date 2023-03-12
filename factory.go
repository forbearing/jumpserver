package jumpserver

import (
	"encoding/json"
	"net/http"
	gourl "net/url"
)

// createFactory product create method for jumpserver Object Executor.
// create method used to creates jumpserver object. eg: asset, user or node.
//
// URL: api, eg: "/api/v1/assets/"
// Request: list of objects in json format.
// Response: list of objects in json format.
// Status Code: 201
func createFactory[T *Asset | *User | *Node](client *Client, api string, tl []T) ([]T, error) {
	if tl == nil {
		return make([]T, 0), nil
	}
	payload, err := marshal(tl)
	if err != nil {
		return nil, err
	}
	data, code, err := client.request(http.MethodPost, api, payload)
	if err != nil {
		return nil, err
	}
	// successfully created response 201 instead of 200.
	if code != http.StatusCreated {
		return nil, Error(data)
	}
	objs := make([]T, len(tl))
	if err := json.Unmarshal(data, &objs); err != nil {
		return nil, err
	}
	return objs, nil
}

// deleteFactory product delete method for jumpserver Object Executor.
// delete method used to deletes jumpserver object. eg: asset, user or node.
//
// delete jumpserver object by id.
//   - URL: api + id, eg: /api/v1/assets/82ac218c-bed7-4a69-8fef-2f8520bdfb1b/
//   - Request: null
//   - Response: null
//   - Status Code: 204
//
// delete jumpserver object by hostname and/or ip.
//   - URL: api + query parameters
//     eg: /api/v1/assets/assets/?hostname=server
//     eg: /api/v1/assets/assets/?ip=1.1.1.1
//     eg: /api/v1/assets/assets/?hostname=server&ip=1.1.1.1
//   - Request: null
//   - Response: null
//   - Status Code: 204
func deleteFactory[T *Asset | *User | *Node](client *Client, api string, p Parameter) error {
	if p == nil {
		return nil
	}
	var (
		data []byte
		code int
		url  string
		err  error
	)
	if len(p.GetID()) != 0 {
		// 通过 ID 删除对象总是成功, 如果 ID 对应的对象不存在也不会报错,只是没有删除
		if url, err = gourl.JoinPath(api, p.GetID(), "/"); err != nil {
			return err
		}
		data, code, err = client.request(http.MethodDelete, url, nil)
	} else {
		if url, err = p.URL(api); err != nil {
			return err
		}
		data, code, err = client.request(http.MethodDelete, url, nil)
	}
	if err != nil {
		return err
	}
	if code != http.StatusNoContent {
		return Error(data)
	}
	return nil
}

// updateFactory product update method for jumpserver Object Executor.
// update method used to update or partial update jumpserver object. eg: asset, user or node.
//
// http method `put` for update.
// http method `patch` for partial update.
//
// 1.If `al` length greater than 1
//   - URL: api, eg: "/api/v1/assets/".
//     Request: list
//     Response: list
//     Status Code: 200
//
// 2.If `al` length is 1, but not set ID.
//   - URL: api, eg: "/api/v1/assets/".
//     Request: list
//     Response: list
//     Status Code: 200
//
// 3.If `al` length is 1 and set id.
//   - URL: api + id, eg: "/api/v1/assets/82ac218c-bed7-4a69-8fef-2f8520bdfb1b/".
//     Request: dict
//     Response: dict
//     Status Code: 200
func updateFactory[T *Asset | *User | *Node](client *Client, api, method string, tl []T) ([]T, error) {
	if len(tl) == 0 {
		return make([]T, 0), nil
	}

	var (
		bulkUpdate bool
		payload    []byte
		data       []byte
		code       int
		url        string
		err        error
	)
	if len(tl) != 1 {
		// URL: api, eg: "/api/v1/assets/".
		// Data send to jumpserver is a list.
		// The bulk update response data from jumpserver is a list.
		bulkUpdate = true
		if payload, err = marshal(tl); err != nil {
			return nil, err
		}
		data, code, err = client.request(method, api, payload)
	} else {
		obj := any(tl[0]).(Object)
		if len(obj.GetID()) == 0 {
			// URL: api, eg: "/api/v1/assets/"
			// Data send to jumpserver is a list
			// The bulk update response data from jumpserver is a list.
			bulkUpdate = true
			if payload, err = marshal(tl); err != nil {
				return nil, err
			}
			data, code, err = client.request(method, api, payload)
		} else {
			// URL: api + id, eg: "/api/v1/assets/82ac218c-bed7-4a69-8fef-2f8520bdfb1b/".
			// Data send to jumpserver is a dict.
			// The update response data from jumpserver is a dict.
			if payload, err = marshal(tl[0]); err != nil {
				return nil, err
			}
			if url, err = gourl.JoinPath(api, obj.GetID(), "/"); err != nil {
				return nil, err
			}
			data, code, err = client.request(method, url, payload)
		}
	}
	if err != nil {
		return nil, err
	}
	// successfully update response 200 instead of 201.
	if code != http.StatusOK {
		return nil, Error(data)
	}

	objs := make([]T, len(tl))
	if bulkUpdate {
		// The bulk update response data from jumpserver is a list.
		if err = json.Unmarshal(data, &objs); err != nil {
			return nil, err
		}
	} else {
		// The update response data from jumpserver is a dict.
		obj := new(T)
		if err = json.Unmarshal(data, obj); err != nil {
			return nil, err
		}
		objs = append(objs, *obj)
	}
	return objs, nil
}

// listFactory product list method for jumpserver Object Executor.
// list uesed to lists jumpserver objects, eg: assets, users, nodes.
//
// URL: api, eg: "/api/v1/assets/"
// Request: null
// Response: list of object in json format
// Status Code: 200
func listFactory[T *Asset | *User | *Node](client *Client, api string) ([]T, error) {
	data, code, err := client.request(http.MethodGet, api, nil)
	if err != nil {
		return nil, err
	}
	if code != http.StatusOK {
		return nil, Error(data)
	}
	objs := make([]T, 0)
	if err := json.Unmarshal(data, &objs); err != nil {
		return nil, err
	}
	return objs, nil
}

// getFactory product get method for jumpserver Object Executor.
// get method used to get or query jumpserver object, eg: assets, users or nodes.
//
// get jumpserver object by id.
//   - URL: api + id, eg: /api/v1/assets/82ac218c-bed7-4a69-8fef-2f8520bdfb1b/
//   - Request: null
//   - Response: dict
//   - Status Code: 200
//
// get jumpserver object by hostname and/or ip.
//   - URL: api + query parameters
//     eg: /api/v1/assets/assets/?hostname=server
//     eg: /api/v1/assets/assets/?ip=1.1.1.1
//     eg: /api/v1/assets/assets/?hostname=server&ip=1.1.1.1
//   - Request: null
//   - Response: list
//   - Status Code: 200
func getFactory[T *Asset | *User | *Node](client *Client, api string, p Parameter) ([]T, error) {
	if p == nil {
		return make([]T, 0), nil
	}

	var (
		getByID bool
		data    []byte
		code    int
		url     string
		err     error
	)
	if len(p.GetID()) != 0 {
		if url, err = gourl.JoinPath(api, p.GetID(), "/"); err != nil {
			return nil, err
		}
		data, code, err = client.request(http.MethodGet, url, nil)
		getByID = true
	} else {
		if url, err = p.URL(api); err != nil {
			return nil, err
		}
		data, code, err = client.request(http.MethodGet, url, nil)
	}
	if err != nil {
		return nil, err
	}
	if code != http.StatusOK {
		return nil, Error(data)
	}

	// 通过 ID 来获取 Jumpserver 对象, 返回的结果是一个对象,而不是多个对象
	if getByID {
		obj := new(T)
		if err = json.Unmarshal(data, obj); err != nil {
			return nil, err
		}
		return []T{*obj}, nil
	} else {
		objs := make([]T, 0)
		if err = json.Unmarshal(data, &objs); err != nil {
			return nil, err
		}
		return objs, nil
	}
}
