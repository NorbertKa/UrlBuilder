package UrlBuilder

import (
	"sync"
	"errors"
)

type Settings struct {
	sync.Mutex

	acceptedProtocols map[string]bool
	acceptedHosts map[string]bool
	acceptedPorts map[uint16]bool
	acceptedPaths map[string]bool
	acceptedQueryParams map[string]bool
}

const (
	ErrAcceptedProtocolAlreadyInSettings string = "Accepted protocol exists in settings"
	ErrAcceptedHostAlreadyInSettings string = "Accepted host exists in settings"
	ErrAcceptedPortAlreadyInSettings string = "Accepted port exists in settings"
	ErrAcceptedPathAlreadyInSettings string = "Accepted path exists in settings"
	ErrAcceptedQueryParamAlreadyInSettings string = "Accepted query param exists in settings"
)

func (s Settings) CheckProtocol(protocol []byte) bool {
	s.Lock()
	_, check := s.acceptedProtocols[string(protocol)]
	s.Unlock()
	return check
}

func (s *Settings) AddProtocol(protocol []byte) error {
	s.Lock()
	defer s.Unlock()
	check := s.CheckProtocol(protocol)
	if check {
		return errors.New(ErrAcceptedProtocolAlreadyInSettings)
	}
	s.acceptedProtocols[string(protocol)] = true
	return nil
}

func (s *Settings) RemoveProtocol(protocol []byte) {
	s.Lock()
	delete(s.acceptedProtocols, string(protocol))
	s.Unlock()
}

func (s Settings) CheckHost(host []byte) bool {
	s.Lock()
	_, check := s.acceptedHosts[string(host)]
	s.Unlock()
	return check
}

func (s *Settings) AddHost(host []byte) error {
	s.Lock()
	defer s.Unlock()
	check := s.CheckHost(host)
	if check {
		return errors.New(ErrAcceptedHostAlreadyInSettings)
	}
	s.acceptedHosts[string(host)] = true
	return nil
}

func (s *Settings) RemoveHost(host []byte) {
	s.Lock()
	delete(s.acceptedHosts, string(host))
	s.Unlock()
}

func (s Settings) CheckPorts(port uint16) bool {
	s.Lock()
	_, check := s.acceptedPorts[port]
	s.Unlock()
	return check
}

func (s *Settings) AddPort(port uint16) error {
	s.Lock()
	defer s.Unlock()
	check := s.CheckPorts(port)
	if check {
		return errors.New(ErrAcceptedPortAlreadyInSettings)
	}
	s.acceptedPorts[port] = true
	return nil
}

func (s *Settings) RemovePort(port uint16) {
	s.Lock()
	delete(s.acceptedPorts, port)
	s.Unlock()
}

func (s Settings) CheckPaths(path []byte) bool {
	s.Lock()
	_, check := s.acceptedPaths[string(path)]
	s.Unlock()
	return check
}

func (s *Settings) AddPath(path []byte) error {
	s.Lock()
	defer s.Unlock()
	check := s.CheckPaths(path)
	if check {
		return errors.New(ErrAcceptedPathAlreadyInSettings)
	}
	s.acceptedPaths[string(path)] = true
	return nil
}

func (s *Settings) RemovePath(path []byte) {
	s.Lock()
	delete(s.acceptedPaths, string(path))
	s.Unlock()
}

func (s Settings) CheckQueryParams(queryParam []byte) bool {
	s.Lock()
	_, check := s.acceptedQueryParams[string(queryParam)]
	s.Unlock()
	return check
}

func (s *Settings) AddQueryParam(queryParam []byte) error {
	s.Lock()
	defer s.Unlock()
	check := s.CheckQueryParams(queryParam)
	if check {
		return errors.New(ErrAcceptedQueryParamAlreadyInSettings)
	}
	s.acceptedQueryParams[string(queryParam)] = true
	return nil
}

func (s *Settings) RemoveQueryParam(queryParam []byte) {
	s.Lock()
	delete(s.acceptedQueryParams, string(queryParam))
	s.Unlock()
}

func (u *Url) AddSettings(settings *Settings) {
	u.Lock()

	if settings != nil {
		u.settings = settings
		u.hasSettings = true
		u.Unlock()
	} else {
		u.settings = nil
		u.hasSettings = false
		u.Unlock()
	}
}

