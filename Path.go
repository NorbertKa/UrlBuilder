package UrlBuilder

import (
	"errors"
)

func (u *Url) SetPath(path [][]byte) error {
	u.Lock()
	defer u.Unlock()
	if u.hasSettings {
		for _, p := range path {
			if !u.settings.CheckPaths(p) {
				return errors.New(ErrAcceptedPathAlreadyInSettings)
			}
		}
	}
	u.path = path
	return nil
}

func (u Url) GetPath() [][]byte {
	u.Lock()
	defer u.Unlock()
	return u.path
}
