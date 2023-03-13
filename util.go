package jumpserver

import (
	"encoding/json"
	"errors"
	"fmt"
	gourl "net/url"
	"unsafe"

	"github.com/google/go-querystring/query"
)

// stringAny format anything to string.
func stringAny(x any) string {
	if x == nil {
		return ""
	}
	if v, ok := x.(fmt.Stringer); ok {
		return v.String()
	}

	switch v := x.(type) {
	case []byte:
		return *(*string)(unsafe.Pointer(&v))
	case string:
		return v
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", x)
	default:
		return fmt.Sprintf("%v", x)
	}
}
func errorAny(x any) error {
	return errors.New(stringAny(x))
}
func jsonAny(v any) string {
	data, err := marshal(v)
	if err != nil {
		return err.Error()
	}
	return stringAny(data)
}
func prettyJsonAny(v any) string {
	data, err := marshal(v, true)
	if err != nil {
		return err.Error()
	}
	return stringAny(data)
}
func marshal(v any, isPretty ...bool) (data []byte, err error) {
	if len(isPretty) >= 1 {
		if isPretty[0] {
			return json.MarshalIndent(v, "", "    ")
		}
	}
	return json.Marshal(v)
}
func concatQuery(x any) (string, error) {
	v, err := query.Values(x)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func concatURL(api string, x any) (string, error) {
	query, err := concatQuery(x)
	if err != nil {
		return "", err
	}

	if len(query) == 0 {
		return api, nil
	}
	l := gourl.URL{Path: api, RawQuery: query}
	url, err := gourl.PathUnescape(l.String())
	if err != nil {
		return "", err
	}
	return url, nil
}
