// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package airgapped

import "github.com/vmware-tanzu/tanzu-cli/test/e2e/framework"

// TODO(anujc): Remove the hardcoding of the plugins within the plugin-groups here once the `tanzu plugin group get`
// command is implemented. We can use the output of `tanzu plugin group get` command with the original repository
// and use the created plugin map to do the validation of plugins availability after plugin gets migrated to new repository.

var pluginsForPG_TKG_001 = []*framework.PluginInfo{
	{Name: "cluster", Target: "kubernetes", Version: "v0.0.1", Description: "Desc for cluster"},
	{Name: "feature", Target: "kubernetes", Version: "v0.0.1", Description: "Desc for feature"},
	{Name: "kubernetes-release", Target: "kubernetes", Version: "v0.0.1", Description: "Desc for kubernetes-release"},
	{Name: "management-cluster", Target: "kubernetes", Version: "v0.0.1", Description: "Desc for management-cluster"},
	{Name: "package", Target: "kubernetes", Version: "v0.0.1", Description: "Desc for package"},
	{Name: "secret", Target: "kubernetes", Version: "v0.0.1", Description: "Desc for secret"},
	{Name: "telemetry", Target: "kubernetes", Version: "v0.0.1", Description: "Desc for telemetry"},
}

var pluginsForPG_TKG_999 = []*framework.PluginInfo{
	{Name: "cluster", Target: "kubernetes", Version: "v9.9.9", Description: "Desc for cluster"},
	{Name: "feature", Target: "kubernetes", Version: "v9.9.9", Description: "Desc for feature"},
	{Name: "kubernetes-release", Target: "kubernetes", Version: "v9.9.9", Description: "Desc for kubernetes-release"},
	{Name: "management-cluster", Target: "kubernetes", Version: "v9.9.9", Description: "Desc for management-cluster"},
	{Name: "package", Target: "kubernetes", Version: "v9.9.9", Description: "Desc for package"},
	{Name: "secret", Target: "kubernetes", Version: "v9.9.9", Description: "Desc for secret"},
	{Name: "telemetry", Target: "kubernetes", Version: "v9.9.9", Description: "Desc for telemetry"},
}

var pluginsForPG_TMC_001 = []*framework.PluginInfo{
	{Name: "account", Target: "mission-control", Version: "v0.0.1", Description: "Desc for account"},
	{Name: "apply", Target: "mission-control", Version: "v0.0.1", Description: "Desc for apply"},
	{Name: "audit", Target: "mission-control", Version: "v0.0.1", Description: "Desc for audit"},
	{Name: "cluster", Target: "mission-control", Version: "v0.0.1", Description: "Desc for cluster"},
	{Name: "clustergroup", Target: "mission-control", Version: "v0.0.1", Description: "Desc for clustergroup"},
	{Name: "data-protection", Target: "mission-control", Version: "v0.0.1", Description: "Desc for data-protection"},
	{Name: "ekscluster", Target: "mission-control", Version: "v0.0.1", Description: "Desc for ekscluster"},
	{Name: "events", Target: "mission-control", Version: "v0.0.1", Description: "Desc for events"},
	{Name: "iam", Target: "mission-control", Version: "v0.0.1", Description: "Desc for iam"},
	{Name: "inspection", Target: "mission-control", Version: "v0.0.1", Description: "Desc for inspection"},
	{Name: "integration", Target: "mission-control", Version: "v0.0.1", Description: "Desc for integration"},
	{Name: "management-cluster", Target: "mission-control", Version: "v0.0.1", Description: "Desc for management-cluster"},
	{Name: "policy", Target: "mission-control", Version: "v0.0.1", Description: "Desc for policy"},
	{Name: "workspace", Target: "mission-control", Version: "v0.0.1", Description: "Desc for workspace"},
	{Name: "helm", Target: "mission-control", Version: "v0.0.1", Description: "Desc for helm"},
	{Name: "secret", Target: "mission-control", Version: "v0.0.1", Description: "Desc for secret"},
	{Name: "continuousdelivery", Target: "mission-control", Version: "v0.0.1", Description: "Desc for continuousdelivery"},
	{Name: "tanzupackage", Target: "mission-control", Version: "v0.0.1", Description: "Desc for tanzupackage"},
}

var pluginsForPG_TMC_999 = []*framework.PluginInfo{
	{Name: "account", Target: "mission-control", Version: "v9.9.9", Description: "Desc for account"},
	{Name: "apply", Target: "mission-control", Version: "v9.9.9", Description: "Desc for apply"},
	{Name: "audit", Target: "mission-control", Version: "v9.9.9", Description: "Desc for audit"},
	{Name: "cluster", Target: "mission-control", Version: "v9.9.9", Description: "Desc for cluster"},
	{Name: "clustergroup", Target: "mission-control", Version: "v9.9.9", Description: "Desc for clustergroup"},
	{Name: "data-protection", Target: "mission-control", Version: "v9.9.9", Description: "Desc for data-protection"},
	{Name: "ekscluster", Target: "mission-control", Version: "v9.9.9", Description: "Desc for ekscluster"},
	{Name: "events", Target: "mission-control", Version: "v9.9.9", Description: "Desc for events"},
	{Name: "iam", Target: "mission-control", Version: "v9.9.9", Description: "Desc for iam"},
	{Name: "inspection", Target: "mission-control", Version: "v9.9.9", Description: "Desc for inspection"},
	{Name: "integration", Target: "mission-control", Version: "v9.9.9", Description: "Desc for integration"},
	{Name: "management-cluster", Target: "mission-control", Version: "v9.9.9", Description: "Desc for management-cluster"},
	{Name: "policy", Target: "mission-control", Version: "v9.9.9", Description: "Desc for policy"},
	{Name: "workspace", Target: "mission-control", Version: "v9.9.9", Description: "Desc for workspace"},
	{Name: "helm", Target: "mission-control", Version: "v9.9.9", Description: "Desc for helm"},
	{Name: "secret", Target: "mission-control", Version: "v9.9.9", Description: "Desc for secret"},
	{Name: "continuousdelivery", Target: "mission-control", Version: "v9.9.9", Description: "Desc for continuousdelivery"},
	{Name: "tanzupackage", Target: "mission-control", Version: "v9.9.9", Description: "Desc for tanzupackage"},
}

var pluginsNotInAnyPG_001 = []*framework.PluginInfo{
	{Name: "isolated-cluster", Target: "global", Version: "v0.0.1", Description: "Desc for isolated-cluster"},
	{Name: "pinniped-auth", Target: "global", Version: "v0.0.1", Description: "Desc for pinniped-auth"},
}

var pluginsNotInAnyPG_999 = []*framework.PluginInfo{
	{Name: "isolated-cluster", Target: "global", Version: "v9.9.9", Description: "Desc for isolated-cluster"},
	{Name: "pinniped-auth", Target: "global", Version: "v9.9.9", Description: "Desc for pinniped-auth"},
}
