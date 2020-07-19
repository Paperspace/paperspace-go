package paperspace

import (
	"fmt"
	"time"
)

type MachineState string

const (
	MachineStateOff          MachineState = "off"
	MachineStateProvisioning MachineState = "provisioning"
	MachineStateRunning      MachineState = "running"
)

type Machine struct {
	ID                     string       `json:"id"`
	Name                   string       `json:"name"`
	OS                     string       `json:"os"`
	RAM                    string       `json:"ram"`
	CPUs                   int          `json:"cpus"`
	GPU                    string       `json:"gpu"`
	State                  MachineState `json:"state"`
	Region                 string       `json:"region"`
	StorageTotal           int          `json:"storageTotal"`
	StorageUsed            int          `json:"storageUsed"`
	UsageRate              float64      `json:"usageRate"`
	ShutdownTimeoutInHours int          `json:"shutdownTimeoutInHours"`
	ShutdownTimeoutForces  bool         `json:"shutdownTimeoutForces"`
	AutoSnapshotFrequency  int          `json:"autoSnapshotFrequency"`
	AutoSnapshotSaveCount  int          `json:"autoSnapshotSaveCount"`
	AgentType              string       `json:"agentType"`
	NetworkID              string       `json:"networkId"`
	PrivateIpAddress       string       `json:"privateIpAddress"`
	PublicIpAddress        string       `json:"publicIpAddress"`
	DtCreated              time.Time    `json:"dtCreated"`
	DtDeleted              time.Time    `json:"dtDeleted"`
	UserID                 string       `json:"userId"`
	TeamID                 string       `json:"teamId"`
	ScriptID               string       `json:"scriptId"`
	DtLastRun              string       `json:"dtLastRun"`
	IsManaged              bool         `json:"isManaged"`
}

type MachineCreateParams struct {
	ClientParams

	Name                   string `json:"name"`
	Region                 string `json:"region"`
	MachineType            string `json:"machineType"`
	Size                   int    `json:"size"`
	BillingType            string `json:"billingType"`
	TemplateID             string `json:"templateId"`
	UserID                 string `json:"userId,omitempty"`
	TeamID                 string `json:"teamId,omitempty"`
	ScriptID               string `json:"scriptId,omitempty"`
	NetworkID              string `json:"networkId,omitempty"`
	ShutdownTimeoutInHours bool   `json:"shutdownTimeoutInHours,omitempty"`
	AssignPublicIP         bool   `json:"assignPublicIP,omitempty"`
	IsManaged              bool   `json:"isManaged,omitempty"`
}

type MachineListParams struct {
	ClientParams

	Filter map[string]string `json:"filter"`
}

type MachineUpdateAttributeParams struct {
	ClientParams

	Name string `json:"name,omitempty" yaml:"name"`
}

type MachineUpdateParams struct {
	ClientParams

	ID                     string `json:"machineId"`
	Name                   string `json:"machineName,omitempty"`
	ShutdownTimeoutInHours int    `json:"shutdownTimeoutInHours,omitempty"`
	ShutdownTimeoutForces  bool   `json:"shutdownTimeoutForces,omitempty"`
	AutoSnapshotFrequency  string `json:"autoSnapshotFrequency,omitempty"`
	AutoSnapshotSaveCount  int    `json:"autoSnapshotSaveCount,omitempty"`
	PerformAutoSnapshot    bool   `json:"performAutoSnapshot,omitempty"`
	DynamicPublicIP        bool   `json:"dynamicPublicIp,omitempty"`
}

func NewMachineListParams() *MachineListParams {
	machineListParams := MachineListParams{
		Filter: make(map[string]string),
	}

	return &machineListParams
}

func (c Client) CreateMachine(params MachineCreateParams) (Machine, error) {
	machine := Machine{}

	url := fmt.Sprintf("/machines/createSingleMachinePublic")
	_, err := c.Request("POST", url, params, &machine, params.ClientParams)

	return machine, err
}

func (c Client) GetMachine(id string, p ...ClientParams) (Machine, error) {
	var params ClientParams
	machine := Machine{}

	if len(p) > 0 {
		params = p[0]
	}

	url := fmt.Sprintf("/machines/getMachinePublic?machineId=%s", id)
	_, err := c.Request("GET", url, nil, &machine, params)

	return machine, err
}

func (c Client) GetMachines(p ...MachineListParams) ([]Machine, error) {
	var machines []Machine
	params := NewMachineListParams()

	if len(p) > 0 {
		params = &p[0]
	}

	url := fmt.Sprintf("/machines/getMachines")
	_, err := c.Request("GET", url, params, &machines, params.ClientParams)

	return machines, err
}

func (c Client) UpdateMachine(p MachineUpdateParams) (Machine, error) {
	machine := Machine{}

	url := fmt.Sprintf("/machines/updateMachine")
	_, err := c.Request("POST", url, p, &machine, p.ClientParams)

	return machine, err
}

func (c Client) DeleteMachine(id string, p ...ClientParams) error {
	var params ClientParams

	if len(p) > 0 {
		params = p[0]
	}

	url := fmt.Sprintf("/machines/%s/destroyMachine", id)
	_, err := c.Request("POST", url, nil, nil, params)

	return err
}