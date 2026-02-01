package repository

import (
	"encoding/json"
	"os"
	"stat_by_sites/domain"
	"time"
)

type ResourceSchema struct{
	Url string `json:"url"`
	Interval string `json:"interval"`
}

type resourcesDataSchema struct{
	Resources []ResourceSchema `json:"resources"`
}

func Fetch(path string) ([]domain.EndpointConfig, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var schema resourcesDataSchema;
	err = json.Unmarshal(data, &schema)

	if err != nil {
		return nil, err
	}

	endpoints := make([]domain.EndpointConfig, len(schema.Resources) )

	for i, el := range schema.Resources{
		if parsedInterval, err := time.ParseDuration(el.Interval); err == nil{
			endpoints[i] = domain.EndpointConfig{
				URL: el.Url,
				Interval: parsedInterval,
			}
		}
	}

	return endpoints, nil
}