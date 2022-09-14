module github.com/dynatrace-oss/terraform-provider-dynatrace

go 1.15

require (
	github.com/dtcookie/dynatrace/api/cluster/v1/groups v1.0.0
	github.com/dtcookie/dynatrace/api/cluster/v1/users v1.0.0
	github.com/dtcookie/dynatrace/api/cluster/v2/envs v1.0.3
	github.com/dtcookie/dynatrace/api/config/alerting v1.0.22
	github.com/dtcookie/dynatrace/api/config/anomalies/applications v1.0.5
	github.com/dtcookie/dynatrace/api/config/anomalies/databaseservices v1.0.5
	github.com/dtcookie/dynatrace/api/config/anomalies/diskevents v1.0.10
	github.com/dtcookie/dynatrace/api/config/anomalies/hosts v1.0.6
	github.com/dtcookie/dynatrace/api/config/anomalies/metricevents v1.0.10
	github.com/dtcookie/dynatrace/api/config/anomalies/services v1.0.11
	github.com/dtcookie/dynatrace/api/config/applications/mobile v1.0.1
	github.com/dtcookie/dynatrace/api/config/applications/web v1.0.11
	github.com/dtcookie/dynatrace/api/config/autotags v1.0.23
	github.com/dtcookie/dynatrace/api/config/credentials/aws v1.0.12
	github.com/dtcookie/dynatrace/api/config/credentials/azure v1.0.12
	github.com/dtcookie/dynatrace/api/config/credentials/cloudfoundry v1.0.1
	github.com/dtcookie/dynatrace/api/config/credentials/kubernetes v1.0.14
	github.com/dtcookie/dynatrace/api/config/credentials/vault v1.0.1
	github.com/dtcookie/dynatrace/api/config/customservices v1.0.16
	github.com/dtcookie/dynatrace/api/config/dashboards v1.0.17
	github.com/dtcookie/dynatrace/api/config/dashboards/sharing v1.0.1
	github.com/dtcookie/dynatrace/api/config/maintenance v1.0.9
	github.com/dtcookie/dynatrace/api/config/managementzones v1.0.18
	github.com/dtcookie/dynatrace/api/config/metrics/calculated/service v1.0.5
	github.com/dtcookie/dynatrace/api/config/naming/hosts v1.0.3
	github.com/dtcookie/dynatrace/api/config/naming/processgroups v1.0.3
	github.com/dtcookie/dynatrace/api/config/naming/services v1.0.3
	github.com/dtcookie/dynatrace/api/config/notifications v1.0.12
	github.com/dtcookie/dynatrace/api/config/requestattributes v1.0.10
	github.com/dtcookie/dynatrace/api/config/requestnaming v1.0.0
	github.com/dtcookie/dynatrace/api/config/synthetic/monitors v1.0.8
	github.com/dtcookie/dynatrace/api/config/topology/application v1.0.1
	github.com/dtcookie/dynatrace/api/config/topology/host v1.0.1
	github.com/dtcookie/dynatrace/api/config/topology/process v1.0.2
	github.com/dtcookie/dynatrace/api/config/topology/processgroup v1.0.1
	github.com/dtcookie/dynatrace/api/config/topology/service v1.0.2
	github.com/dtcookie/dynatrace/api/config/topology/tag v1.0.0
	github.com/dtcookie/dynatrace/api/config/v2/ibmmq/filters v1.0.1
	github.com/dtcookie/dynatrace/api/config/v2/ibmmq/imsbridges v1.0.1
	github.com/dtcookie/dynatrace/api/config/v2/ibmmq/queuemanagers v1.0.1
	github.com/dtcookie/dynatrace/api/config/v2/ibmmq/queuesharinggroups v1.0.1
	github.com/dtcookie/dynatrace/api/config/v2/keyrequests v1.0.2
	github.com/dtcookie/dynatrace/api/config/v2/slo v1.0.6
	github.com/dtcookie/dynatrace/api/config/v2/spans/attributes v1.0.1
	github.com/dtcookie/dynatrace/api/config/v2/spans/capture v1.0.1
	github.com/dtcookie/dynatrace/api/config/v2/spans/ctxprop v1.0.1
	github.com/dtcookie/dynatrace/api/config/v2/spans/entrypoints v1.0.2
	github.com/dtcookie/dynatrace/api/config/v2/spans/resattr v1.0.2
	github.com/dtcookie/dynatrace/rest v1.0.15
	github.com/dtcookie/dynatrace/terraform v1.0.5
	github.com/dtcookie/hcl v0.0.16
	github.com/dtcookie/opt v1.0.0
	github.com/google/uuid v1.3.0
	github.com/hashicorp/terraform-plugin-docs v0.13.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.20.0
)
