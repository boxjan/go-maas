package maas

import (
	"errors"
	"net/url"
)

func FilterToUrlParams(filters ...string) (url.Values, error) {
	if len(filters)%2 != 0 {
		return nil, errors.New("errors machines filter")
	}

	params := url.Values{}
	for i := 0; i < len(filters); i += 2 {
		if _, ok := params[filters[i]]; !ok {
			params[filters[i]] = make([]string, 0, 2)
		}
		params[filters[i]] = append(params[filters[i]], filters[i+1])
	}
	return params, nil
}
