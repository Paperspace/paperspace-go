package paperspace

import (
	"fmt"
	"time"
)

type Network struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Network   string    `json:"network"`
	Netmask   string    `json:"netmask"`
	TeamID    string    `json:"teamId"`
	IsManaged bool      `json:"isManaged"`
	DtCreated time.Time `json:"dtCreated"`
	DtDeleted time.Time `json:"dtDeleted"`
}

type NetworkCreateParams struct {
	RequestParams

	Name   string `json:"name"`
	Region string `json:"region"`
}

type networkCreateParamsInternal struct {
	RequestParams

	Name     string `json:"name"`
	RegionID int    `json:"regionId"`
}

type NetworkDeleteParams struct {
	RequestParams
}

type NetworkGetParams struct {
	RequestParams
}

type NetworkListParams struct {
	RequestParams
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Region  string `json:"region,omitempty"`
	Network string `json:"network,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	TeamID  string `json:"teamId,omitempty"`
}

func (c Client) CreateNetwork(params NetworkCreateParams) (Network, error) {
	regionID, ok := RegionMap[params.Region]
	if !ok {
		return Network{}, fmt.Errorf("no region found for %s", params.Region)
	}

	intParams := networkCreateParamsInternal{Name: params.Name, RegionID: regionID, RequestParams: params.RequestParams}

	network := Network{}
	url := fmt.Sprintf("/networks")
	_, err := c.Request("POST", url, intParams, &network, params.RequestParams)

	return network, err
}

func (c Client) GetNetwork(id string, params NetworkGetParams) (Network, error) {
	networks, err := c.GetNetworks(NetworkListParams{ID: id, RequestParams: params.RequestParams})
	if err != nil {
		return Network{}, err
	}
	if len(networks) == 0 {
		return Network{}, fmt.Errorf("no network found for ID %s", id)
	}
	if len(networks) > 1 {
		return Network{}, fmt.Errorf("found more than one network for ID %s", id)
	}
	return networks[0], nil
}

func (c Client) GetNetworks(params NetworkListParams) ([]Network, error) {
	var networks []Network

	url := fmt.Sprintf("/networks")
	_, err := c.Request("GET", url, params, &networks, params.RequestParams)

	return networks, err
}

func (c Client) DeleteNetwork(id string, params NetworkDeleteParams) error {
	url := fmt.Sprintf("/networks/%s", id)
	_, err := c.Request("DELETE", url, nil, nil, params.RequestParams)

	return err
}
