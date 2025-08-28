package layer

import (
	"encoding/json"
	"fmt"

	"github.com/dudubtw/figma/app"
)

func UnmarshalJSON(data []byte, wf *app.WorkplaceFile) error {
	type Alias app.WorkplaceFile
	aux := struct {
		Layers []json.RawMessage
		*Alias
	}{Alias: (*Alias)(wf)}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	for _, raw := range aux.Layers {
		// Peek at the "type" field
		var kind struct {
			Type string
		}
		if err := json.Unmarshal(raw, &kind); err != nil {
			return err
		}

		switch kind.Type {
		case "rectangle":
			var rect Rectangle
			if err := json.Unmarshal(raw, &rect); err != nil {
				return err
			}

			wf.Layers = append(wf.Layers, &rect)
		case "text":
			var txt Text
			if err := json.Unmarshal(raw, &txt); err != nil {
				return err
			}
			wf.Layers = append(wf.Layers, &txt)

		case "image":
			var image Image
			if err := json.Unmarshal(raw, &image); err != nil {
				return err
			}
			wf.Layers = append(wf.Layers, &image)

		default:
			return fmt.Errorf("unknown layer type: %s", kind.Type)
		}
	}

	return nil
}
