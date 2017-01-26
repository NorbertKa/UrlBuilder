package UrlBuilder

import (
	"errors"
)

func (u *Url) SetPort(port uint16) error {
	u.Lock()
	defer u.Unlock()
	if u.hasSettings {
		if !u.settings.CheckPorts(port) {
			return errors.New(ErrAcceptedPortAlreadyInSettings)
		}
	}
	u.port = port
	return nil
}

func (u Url) GetPort() uint16 {
	u.Lock()
	defer u.Unlock()
	return u.port
}
