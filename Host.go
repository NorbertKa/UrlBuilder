package UrlBuilder

import (
	"errors"
)

func (u *Url) SetHost(host []byte) error {
	u.Lock()
	defer u.Unlock()
	if u.hasSettings {
		if !u.settings.CheckHost(host) {
			return errors.New(ErrAcceptedHostAlreadyInSettings)
		}
	}
	u.host = host
	return nil
}

func (u Url) GetHost() []byte {
	u.Lock()
	defer u.Unlock()
	return u.host
}