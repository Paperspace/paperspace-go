package paperspace

import (
	"context"
	"time"
)

type Experiment struct {
	DtCreated                   *time.Time                `json:"dtCreated,omitempty"`
	DtDeleted                   *time.Time                `json:"dtDeleted,omitempty"`
	DtFinished                  *time.Time                `json:"dtFinished,omitempty"`
	DtModified                  *time.Time                `json:"dtModified,omitempty"`
	DtProvisioningFinished      *time.Time                `json:"dtProvisioningFinished,omitempty"`
	DtProvisioningStarted       *time.Time                `json:"dtProvisioningStarted,omitempty"`
	DtStarted                   *time.Time                `json:"dtStarted,omitempty"`
	DtTeardownFinished          *time.Time                `json:"dtTeardownFinished,omitempty"`
	DtTeardownStarted           *time.Time                `json:"dtTeardownStarted,omitempty"`
	ExperimentError             *string                   `json:"experimentError,omitempty"`
	ExperimentTemplateHistoryID int64                     `json:"experimentTemplateHistoryId"`
	ExperimentTemplateID        int64                     `json:"experimentTemplateId"`
	ExperimentTypeID            int64                     `json:"experimentTypeId"`
	Handle                      string                    `json:"handle"`
	ID                          int64                     `json:"id"`
	ProjectHandle               string                    `json:"projectHandle"`
	ProjectID                   int64                     `json:"projectId"`
	StartedByUserID             int64                     `json:"started_by_user_id"`
	State                       int64                     `json:"state"`
	Tags                        []string                  `json:"tags"`
	TemplateHistory             ExperimentTemplateHistory `json:"templateHistory,omitempty"`
}

type ExperimentTemplateHistory struct {
	ID                   int64            `json:"id"`
	DtCreated            time.Time        `json:"dtCreated"`
	DtDeleted            time.Time        `json:"dtDeleted"`
	ExperimentTemplateID int64            `json:"experimentTemplateId"`
	TriggerEvent         interface{}      `json:"triggerEvent"`
	TriggerEventID       int64            `json:"triggerEventId"`
	CliCommand           *string          `json:"cliCommand,omitempty"`
	Params               ExperimentParams `json:"params"`
}

type TriggerEvent struct {
	ID        int64                  `json:"id"`
	Type      string                 `json:"type"`
	DtCreated *time.Time             `json:"dtCreated"`
	EventData map[string]interface{} `json:"eventData"`
}

type ExperimentParams struct {
	Name              *string           `json:"name"`
	Ports             *string           `json:"ports"`
	WorkingDirectory  *string           `json:"workingDirectory"`
	ArtifactDirectory *string           `json:"artifactDirectory"`
	ExperimentEnv     map[string]string `json:"experimentEnv"`

	ProjectID      *int64  `json:"projectId"`
	ProjectHandle  *string `json:"projectHandle"`
	ClusterID      *string `json:"clusterId"`
	TriggerEventID *int64  `json:"triggerEventId"`

	Workspace         *string `json:"workspaceUrl"`
	WorkspaceRef      *string `json:"workspaceRef"`
	WorkspaceUsername *string `json:"workspaceUsername"`
	WorkspacePassword *string `json:"workspacePassword"`

	ModelType *string `json:"modelType"`
	ModelPath *string `json:"modelPath"`

	IsPreemptible *bool   `json:"isPreemptible"`
	CustomMetrics *string `json:"customMetrics"`

	Datasets []ExperimentDataset `json:"datasets"`

	SingleNodeExperimentParams

	MultiNodeExperimentWorkerParams
	MultiNodeExperimentParameterServerParams
	MultiNodeMPIExperimentMasterParams
}

type SingleNodeExperimentParams struct {
	Container   *string `json:"container"`
	MachineType *string `json:"machineType"`
	Command     *string `json:"command"`

	ContainerUser    *string `json:"containerUser"`
	RegistryUsername *string `json:"registryUsername"`
	RegistryPassword *string `json:"registryPassword"`

	UseDockerfile  *bool   `json:"useDockerfile"`
	DockerfilePath *string `json:"dockerfilePath"`
}

type MultiNodeExperimentWorkerParams struct {
	WorkerContainer   *string `json:"workerContainer"`
	WorkerMachineType *string `json:"workerMachineType"`
	WorkerCommand     *string `json:"workerCommand"`

	WorkerContainerUser    *string `json:"workerContainerUser"`
	WorkerRegistryUsername *string `json:"workerRegistryUsername"`
	WorkerRegistryPassword *string `json:"workerRegistryPassword"`

	WorkerUseDockerfile  *bool   `json:"workerUseDockerfile"`
	WorkerDockerfilePath *string `json:"workerDockerfilePath"`
}

type MultiNodeExperimentParameterServerParams struct {
	ParameterServerContainer   *string `json:"parameterServerContainer"`
	ParameterServerMachineType *string `json:"parameterServerMachineType"`
	ParameterServerCommand     *string `json:"parameterServerCommand"`

	ParameterServerContainerUser    *string `json:"parameterServerContainerUser"`
	ParameterServerRegistryUsername *string `json:"parameterServerRegistryUsername"`
	ParameterServerRegistryPassword *string `json:"parameterServerRegistryPassword"`

	ParameterServerUseDockerfile  *bool   `json:"parameterServerUseDockerfile"`
	ParameterServerDockerfilePath *string `json:"parameterServerDockerfilePath"`
}

type MultiNodeMPIExperimentMasterParams struct {
	MasterContainer   *string `json:"masterContainer"`
	MasterMachineType *string `json:"masterMachineType"`
	MasterCommand     *string `json:"masterCommand"`

	MasterContainerUser    *string `json:"masterContainerUser"`
	MasterRegistryUsername *string `json:"masterRegistryUsername"`
	MasterRegistryPassword *string `json:"masterRegistryPassword"`

	MasterUseDockerfile  *bool   `json:"masterUseDockerfile"`
	MasterDockerfilePath *string `json:"masterDockerfilePath"`
}

type ExperimentDataset struct {
	URI               string                          `json:"uri"`
	AWSAccessKeyID    *string                         `json:"awsAccessKeyId"`
	AWSSecretAcessKey *string                         `json:"awsSecretAccessKey"`
	ETag              *string                         `json:"etag"`
	VersionID         *string                         `json:"versionId"`
	Name              *string                         `json:"name"`
	VolumeOptions     *ExperimentDatasetVolumeOptions `json:"volumeOptions"`
}

type ExperimentDatasetVolumeOptions struct {
	Kind string  `json:"kind"`
	Size *string `json:"size,omitempty"`
}

func (c Client) RunExperiment(ctx context.Context, params ExperimentParams) (Experiment, error) {
	experiment := struct {
		Data    Experiment `json:"data"`
		Message string     `json:"message"`
	}{}

	url := "/experiments/v2/run"
	_, err := c.Request(ctx, "POST", url, params, &experiment)

	return experiment.Data, err
}
