package stamp

import (
	"github.com/bmf-san/akashigo/types"
)

// StampResponse is a response for stamp.
type StampResponse struct {
	Success  bool             `json:"success"`
	Response response         `json:"response"`
	Errors   []errorsresponse `json:"errors"`
}

type response struct {
	LoginCompanyCode string          `json:"login_company_code"`
	StaffID          int             `json:"staff_id"`
	Type             types.TypeStamp `json:"type"`
	StampedAt        string          `json:"stampedAt"`
}

type errorsresponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
