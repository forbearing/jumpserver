package jumpserver

import (
	"encoding/json"
	"errors"
	"fmt"
	gourl "net/url"
	"unsafe"

	"github.com/google/go-querystring/query"
)

// String format anything to string.
func String(x any) string {
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

// Error
func Error(x any) error {
	return errors.New(String(x))
}

func marshal(v any, isPretty ...bool) (data []byte, err error) {
	if len(isPretty) >= 1 {
		if isPretty[0] {
			return json.MarshalIndent(v, "", "    ")
		}
	}
	return json.Marshal(v)
}

func concatPath(x any) (string, error) {
	v, err := query.Values(x)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func concatURL(api string, x any) (string, error) {
	path, err := concatPath(x)
	if err != nil {
		return "", err
	}

	url := api + "?" + path
	if url, err = gourl.PathUnescape(url); err != nil {
		return "", err
	}
	return url, nil
}
