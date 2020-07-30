package paperspace

import (
	"encoding/json"
	"fmt"
)

type PaperspaceErrorResponse struct {
	Error *PaperspaceError `json:"error"`
}

type PaperspaceError struct {
	Name    string      `json:"name"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Details interface{} `json:"details,omitempty"`
}

func (e PaperspaceError) Error() string {
	if e.Details == nil {
		return e.Message
	}

	// The experiments api doesn't always serialize the details object nicely
	switch details := e.Details.(type) {
	case string:
		return fmt.Sprint(e.Message, ": ", details)
	case map[string]interface{}:
		if detailMessage, ok := details["message"]; ok {
			if msg, ok := detailMessage.(string); ok {
				return fmt.Sprint(e.Message, ": ", msg)
			}
		}
	}
	return e.Message
}

func (e *PaperspaceError) UnmarshalJSON(raw []byte) error {
	var rawError struct {
		Name    string      `json:"name"`
		Message string      `json:"message"`
		Status  int         `json:"status"`
		Code    *int        `json:"code,omitempty"` // for ps-experiments
		Details interface{} `json:"details,omitempty"`
	}

	if err := json.Unmarshal(raw, &rawError); err != nil {
		return err
	}

	e.Name = rawError.Name
	e.Status = rawError.Status
	if rawError.Code != nil {
		e.Status = *rawError.Code
	}
	e.Message = rawError.Message
	e.Name = rawError.Name

	if e.Message != "" {
		return nil
	}

	// The experiments api doesn't always serialize the details object nicely
	switch details := rawError.Details.(type) {
	case string:
		e.Message = details
	case map[string]interface{}:
		if detailMessage, ok := details["message"]; ok {
			e.Message, _ = detailMessage.(string)
		}
	}
	return nil
}
