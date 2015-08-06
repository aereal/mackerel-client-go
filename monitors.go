package mackerel

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
{
  "monitors": [
    {
      "id": "2cSZzK3XfmG",
      "type": "passive",
      "name": "connectivity",
      "duration": 5
    },
    {
      "id"  : "2cSZzK3XfmG",
      "type": "host",
      "name": "disk.aa-00.writes.delta",
      "duration": 3,
      "metric": "disk.aa-00.writes.delta",
      "operator": ">",
      "warning": 20000.0,
      "critical": 400000.0,
      "scopes": [
        "SomeService"
      ],
      "excludeScopes": [
        "SomeService: db-slave-backup"
      ]
    },
    {
      "id"  : "2cSZzK3XfmG",
      "type": "service",
      "name": "SomeService - custom.access_num.4xx_count",
      "service": "SomeService",
      "duration": 1,
      "metric": "custom.access_num.4xx_count",
      "operator": ">",
      "warning": 50.0,
      "critical": 100.0
    },
    {
      "id"  : "2cSZzK3XfmG",
      "type": "external",
      "name": "example.com",
      "url": "http://www.example.com"
    }
  ]
}
*/

// Monitor information
type Monitor struct {
	ID            string   `json:"id,omitempty"`
	Name          string   `json:"name,omitempty"`
	Type          string   `json:"type,omitempty"`
	Duration      uint64   `json:"duration,omitempty"`
	Scopes        []string `json:"scopes,omitempty"`
	ExcludeScopes []string `json:"excludeScopes,omitempty"`
}

// FindMonitors find monitors
func (c *Client) FindMonitors() ([]*Monitor, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s", c.urlFor("/api/v0/monitors").String()), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("status code is not 200")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data struct {
		Monitors []*(Monitor) `json:"monitors"`
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data.Monitors, err
}

// CreateMonitor creating monitor
func (c *Client) CreateMonitor(param *Monitor) (*Monitor, error) {
	requestJSON, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		c.urlFor("/api/v0/monitors").String(),
		bytes.NewReader(requestJSON),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data Monitor

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// UpdateMonitor update monitor
func (c *Client) UpdateMonitor(monitorID string, param *Monitor) (*Monitor, error) {
	requestJSON, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"PUT",
		c.urlFor(fmt.Sprintf("/api/v0/monitor/%s", monitorID)).String(),
		bytes.NewReader(requestJSON),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data Monitor

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// DeleteMonitor update monitor
func (c *Client) DeleteMonitor(monitorID string, param *Monitor) (*Monitor, error) {
	requestJSON, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"PUT",
		c.urlFor(fmt.Sprintf("/api/v0/monitor/%s", monitorID)).String(),
		bytes.NewReader(requestJSON),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data Monitor

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
