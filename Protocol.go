package UrlBuilder

import (
	"errors"
)

func (u *Url) SetProtocol(protocol []byte) error {
	u.Lock()
	defer u.Unlock()
	if u.hasSettings {
		if !u.settings.CheckProtocol(protocol) {
			return errors.New(ErrAcceptedProtocolAlreadyInSettings)
		}
	}
	u.protocol = protocol
	return nil
}

func (u Url) GetProtocol() []byte {
	u.Lock()
	defer u.Unlock()
	return u.protocol
}