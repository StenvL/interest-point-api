package responses

//ServiceStatusResponse response model
type ServiceStatusResponse struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Details string `json:"details"`
}
