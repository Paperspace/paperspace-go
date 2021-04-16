package paperspace

import (
	"fmt"
)

type MLModel struct {
	Name              string                 `json:"name,omitempty"`
	IsPublic          *bool                  `json:"isPublic,omitempty"`
	Tag               string                 `json:"tag,omitempty"`
	ProjectID         string                 `json:"projectId,omitempt"`
	ExperimentID      string                 `json:"experimentId,omitempty"`
	URL               string                 `json:"url,omitempty"`
	ModelType         string                 `json:"modelType,omitempty"`
	ModelPath         string                 `json:"modelPath,omitempty"`
	DatasetVersionRef string                 `json:"datasetVersionRef,omitempty"`
	Summary           map[string]interface{} `json:"summary,omitempty"`
	Notes             string                 `json:"notes,omitempty"`
	Params            map[string]interface{} `json:"params,omitempty"`
	Detail            map[string]interface{} `json:"detail,omitempty"`
}

type UpdateModelParams struct {
	RequestParams

	ID         string  `json:"id"`
	Attributes MLModel `json:"attributes"`
}

func (c Client) UpdateModel(params UpdateModelParams) (*MLModel, error) {
	model := &MLModel{}
	url := fmt.Sprintf("/mlModels/updateModel")
	_, err := c.Request("POST", url, params, model, params.RequestParams)
	if err != nil {
		return nil, err
	}
	return model, err
}
