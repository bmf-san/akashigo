package stamp

import (
	"encoding/json"
	"net/http"

	"github.com/bmf-san/akashigo/api"
	"github.com/pkg/errors"
)

func (s *Stamp) Stamp(params StampParams) (*StampResponse, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	header := make(http.Header)
	header["Content-Type"] = []string{"application/json"}
	req, err := s.Client.NewRequest(http.MethodPost, header, body, api.URL(api.StampEndpoint, "", s.Client.CompanyID))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	body, err = s.Client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	sr := &StampResponse{}
	if err = json.Unmarshal(body, &sr); err != nil {
		return nil, errors.WithStack(err)
	}

	return sr, nil
}
