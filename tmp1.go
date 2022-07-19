package main

type Targets []*Target

type MonitoringObjects []struct {
	Host           string
	OverrideConfig NetboxConfigContext
	Target         Target
}

type Target struct {
	Targets []string     `json:"targets"`
	Labels  TargetLabels `json:"labels"`
	Metrics svcMetrics
}
type NetboxConfigContext struct {
	Monitoring struct {
		Services []struct {
			Service struct {
				Name     string `json:"name,omitempty"`
				Port     int    `json:"port,omitempty"`
				Disabled bool   `json:"disabled,omitempty"`
			} `json:"service,omitempty"`
		} `json:"services,omitempty"`
	} `json:"monitoring,omitempty"`
}
type TargetLabels struct {
	MetricsPath                          string `json:"__metrics_path__,omitempty"`
	MetaNetboxIp                         string `json:"__meta_netbox_ip"`
	MetaNetboxName                       string `json:"__meta_netbox_name" validate:"fqdn"`
	MetaNetboxType                       string `json:"__meta_netbox_type"`
	MetaNetboxExporter                   string `json:"__meta_netbox_exporter"`
	MetaNetboxId                         string `json:"__meta_netbox_id"`
	MetaNetboxTenant                     string `json:"__meta_netbox_tenant"`
	MetaNetboxTenantSlug                 string `json:"__meta_netbox_tenant_slug"`
	MetaNetboxTenantGroup                string `json:"__meta_netbox_tenant_group"`
	MetaNetboxTenantGroupSlug            string `json:"__meta_netbox_tenant_group_slug"`
	MetaNetboxSite                       string `json:"__meta_netbox_site"`
	MetaNetboxSiteSlug                   string `json:"__meta_netbox_site_slug"`
	MetaNetboxRole                       string `json:"__meta_netbox_role"`
	MetaNetboxRoleSlug                   string `json:"__meta_netbox_role_slug"`
	MetaNetboxCluster                    string `json:"__meta_netbox_cluster"`
	MetaNetboxPlatform                   string `json:"__meta_netbox_platform"`
	MetaNetboxPlatformSlug               string `json:"__meta_netbox_platform_slug"`
	MetaNetboxCumstomFieldDeployProfile  string `json:"__meta_netbox_custom_field_deploy_profile"`
	MetaNetboxCumstomFieldEnvironment    string `json:"__meta_netbox_custom_field_environment"`
	MetaNetboxCumstomFieldHypervisorId   string `json:"__meta_netbox_custom_field_hypervisor_id"`
	MetaNetboxCumstomFieldHypervisorName string `json:"__meta_netbox_custom_field_hypervisor_name"`
	MetaNetboxCumstomFieldHypervisorUrl  string `json:"__meta_netbox_custom_field_hypervisor_url"`
	MetaNetboxCumstomFieldMonitoring     string `json:"__meta_netbox_custom_field_monitoring"`
	MetaNetboxTags                       string `json:"__meta_netbox_tags,omitempty"`
	MetaNetboxTagSlugs                   string `json:"__meta_netbox_tag_slugs,omitempty"`
}

type svcMetrics struct {
	Avaliable      *kitprometheus.Gauge
	DnsMismatch    *kitprometheus.Gauge
	NetboxMismatch *kitprometheus.Gauge
}
