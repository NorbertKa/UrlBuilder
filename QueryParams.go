package UrlBuilder

import (
	"errors"
)

func (u *Url) SetQueryParams(params map[string][]byte) error {
	u.Lock()
	defer u.Unlock()
	if u.hasSettings {
		for _, p := range params {
			if !u.settings.CheckQueryParams(p) {
				return errors.New(ErrAcceptedQueryParamAlreadyInSettings)
			}
		}
	}
	u.queryParams = params
	return nil
}

func (u Url) GetQueryParams() map[string][]byte {
	u.Lock()
	defer u.Unlock()
	return u.queryParams
}
