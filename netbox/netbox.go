package netbox

type NetboxConfigContext struct {
	Monitoring NetboxConfigContextMonitoring
}
type NetboxConfigContextMonitoring struct {
	Monitoring struct {
		Services NetboxConfigContextMonitoringServices
	} `json:"monitoring,omitempty"`
}

type NetboxConfigContextMonitoringServices []struct {
	Service struct {
		Name     string `json:"name,omitempty"`
		Port     int    `json:"port,omitempty"`
		Disabled bool   `json:"disabled,omitempty"`
	} `json:"service,omitempty"`
}
