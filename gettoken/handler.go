package gettoken

import (
	"encoding/json"
	"net/http"

	"github.com/bmf-san/akashigo/api"
	"github.com/pkg/errors"
)

func (gt *GetToken) GetToken(params GetTokenParams) (*GetTokenResponse, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	header := make(http.Header)
	header["Content-Type"] = []string{"application/json"}
	req, err := gt.Client.NewRequest(http.MethodPost, header, body, api.URL(api.GetTokenEndpoint, "", gt.Client.CompanyID))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	body, err = gt.Client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	gtr := &GetTokenResponse{}
	if err = json.Unmarshal(body, &gtr); err != nil {
		return nil, errors.WithStack(err)
	}

	return gtr, nil
}
