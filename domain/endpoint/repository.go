package endpoint

import "time"

type Reader interface { 
	Fetch() ([]EndpointConfig, error) 
}

type EndpointRepository interface {
	Get(string) (Endpoint, bool)
	GetKeys() []string
	Add(...EndpointConfig) error
	Update(string, int, time.Duration) error
}
