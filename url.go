package UrlBuilder

import "sync"

const (
	defaultPort uint16 = 80
	initialProtocolLength = 4
	initialHostLength = 5
)

var (
	slash =  []byte{'/'}
	questionMark = []byte{'?'}
	ampersand = []byte{'&'}
	path = []byte{':','/','/'}
)

type Url struct {
	sync.Mutex

	protocol []byte
	host []byte
	port uint16
	path [][]byte
	queryParams map[string][]byte

	settings *Settings
	hasSettings bool
}

func newUrl() *Url {
	return &Url{
		protocol: make([]byte, initialProtocolLength),
		host: make([]byte, initialHostLength),
		port: defaultPort,
		settings: nil,
		hasSettings: false,
	}
}

func NewUrl() *Url {
	return newUrl()
}