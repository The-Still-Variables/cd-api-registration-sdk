package api_registration

var A APIRegistration

type APIRegistration struct {
	HostName    string `json:"hostName"`
	NameSpace   string `json:"nameSpace"`
	Category    string `json:"category"`
	Description string `json:"description"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Domain      string `json:"domain"`
	Environment string `json:"environment"`
	Status      string `json:"status"`
	APIList     []struct {
		URL          string `json:"url"`
		Description  string `json:"description"`
		Method       string `json:"method"`
		Dependencies []struct {
			Category    string `json:"category"`
			Name        string `json:"name"`
			Path        string `json:"path"`
			Method      string `json:"method"`
			Description string `json:"description"`
		} `json:"dependencies"`
	} `json:"apiList"`
	HealthCheck  string `json:"healthCheck"`
	SoftwareID   string `json:"softwareId"`
	SoftwareName string `json:"softwareName"`
	Version      string `json:"version"`
	RegisterAPI  string `json:"registerApi"`
}
