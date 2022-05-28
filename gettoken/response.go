package gettoken

// GetTokenResponse is a response for getting a token.
type GetTokenResponse struct {
	Success  bool             `json:"success"`
	Response response         `json:"response"`
	Errors   []errorsresponse `json:"errors"`
}

type response struct {
	LoginCompanyCode string `json:"login_company_code"`
	StaffID          int    `json:"staff_id"`
	AgencyManagerID  int    `json:"agency_manager_id"`
	Token            string `json:"token"`
	ExpiredAt        string `json:"expired_at"`
}

type errorsresponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
