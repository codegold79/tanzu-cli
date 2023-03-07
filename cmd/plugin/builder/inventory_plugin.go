// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/spf13/cobra"

	"github.com/vmware-tanzu/tanzu-cli/cmd/plugin/builder/imgpkg"
	"github.com/vmware-tanzu/tanzu-cli/cmd/plugin/builder/inventory"
	"github.com/vmware-tanzu/tanzu-cli/pkg/cli"
)

// newInventoryPluginCmd creates a new command for plugin inventory operations.
func newInventoryPluginCmd() *cobra.Command {
	var inventoryPluginCmd = &cobra.Command{
		Use:   "plugin",
		Short: "Plugin Inventory Operations",
	}

	inventoryPluginCmd.SetUsageFunc(cli.SubCmdUsageFunc)

	inventoryPluginCmd.AddCommand(
		newInventoryPluginInsertCmd(),
		newInventoryPluginActivateCmd(),
		newInventoryPluginDeactivateCmd(),
	)

	return inventoryPluginCmd
}

type inventoryPluginInsertFlags struct {
	Repository        string
	InventoryImageTag string
	ManifestFile      string
	Publisher         string
	Vendor            string
	DeactivatePlugins bool
}

func newInventoryPluginInsertCmd() *cobra.Command {
	var ipiFlags = &inventoryPluginInsertFlags{}

	var pluginInsertCmd = &cobra.Command{
		Use:     "insert",
		Short:   "Insert the plugin to the inventory database available on the remote repository",
		Example: ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			piOptions := inventory.InventoryPluginUpdateOptions{
				Repository:        ipiFlags.Repository,
				InventoryImageTag: ipiFlags.InventoryImageTag,
				ManifestFile:      ipiFlags.ManifestFile,
				Vendor:            ipiFlags.Vendor,
				Publisher:         ipiFlags.Publisher,
				DeactivatePlugins: ipiFlags.DeactivatePlugins,
				ImgpkgOptions:     imgpkg.NewImgpkgCLIWrapper(),
			}
			return piOptions.PluginInsert()
		},
	}

	pluginInsertCmd.Flags().StringVarP(&ipiFlags.Repository, "repository", "", "", "repository to publish plugin inventory image")
	pluginInsertCmd.Flags().StringVarP(&ipiFlags.InventoryImageTag, "plugin-inventory-image-tag", "", "latest", "tag to which plugin inventory image needs to be published")
	pluginInsertCmd.Flags().StringVarP(&ipiFlags.ManifestFile, "manifest", "", "", "manifest file specifying plugin details that needs to be processed")
	pluginInsertCmd.Flags().StringVarP(&ipiFlags.Vendor, "vendor", "", "", "name of the vendor")
	pluginInsertCmd.Flags().StringVarP(&ipiFlags.Publisher, "publisher", "", "", "name of the publisher")
	pluginInsertCmd.Flags().BoolVarP(&ipiFlags.DeactivatePlugins, "deactivate", "", false, "mark plugins as deactivated")

	_ = pluginInsertCmd.MarkFlagRequired("repository")
	_ = pluginInsertCmd.MarkFlagRequired("vendor")
	_ = pluginInsertCmd.MarkFlagRequired("publisher")
	_ = pluginInsertCmd.MarkFlagRequired("manifest")

	return pluginInsertCmd
}

type inventoryPluginActivateDeactivateFlags struct {
	Repository        string
	InventoryImageTag string
	ManifestFile      string
	Publisher         string
	Vendor            string
}

func newInventoryPluginActivateCmd() *cobra.Command {
	pluginDeactivateCmd, flags := getActivateDeactivateBaseCmd()
	pluginDeactivateCmd.Use = "activate"
	pluginDeactivateCmd.Short = "Activate the existing plugin in the inventory database available on the remote repository"
	pluginDeactivateCmd.Example = ""
	pluginDeactivateCmd.RunE = func(cmd *cobra.Command, args []string) error {
		piOptions := inventory.InventoryPluginUpdateOptions{
			Repository:        flags.Repository,
			InventoryImageTag: flags.InventoryImageTag,
			ManifestFile:      flags.ManifestFile,
			Vendor:            flags.Vendor,
			Publisher:         flags.Publisher,
			DeactivatePlugins: false,
			ImgpkgOptions:     imgpkg.NewImgpkgCLIWrapper(),
		}
		return piOptions.UpdatePluginActivationState()
	}
	return pluginDeactivateCmd
}

func newInventoryPluginDeactivateCmd() *cobra.Command {
	pluginDeactivateCmd, flags := getActivateDeactivateBaseCmd()
	pluginDeactivateCmd.Use = "deactivate"
	pluginDeactivateCmd.Short = "Deactivate the existing plugin in the inventory database available on the remote repository"
	pluginDeactivateCmd.Example = ""
	pluginDeactivateCmd.RunE = func(cmd *cobra.Command, args []string) error {
		piOptions := inventory.InventoryPluginUpdateOptions{
			Repository:        flags.Repository,
			InventoryImageTag: flags.InventoryImageTag,
			ManifestFile:      flags.ManifestFile,
			Vendor:            flags.Vendor,
			Publisher:         flags.Publisher,
			DeactivatePlugins: true,
			ImgpkgOptions:     imgpkg.NewImgpkgCLIWrapper(),
		}
		return piOptions.UpdatePluginActivationState()
	}
	return pluginDeactivateCmd
}

func getActivateDeactivateBaseCmd() (*cobra.Command, *inventoryPluginActivateDeactivateFlags) {
	var flags = &inventoryPluginActivateDeactivateFlags{}

	var activateDeactivateCmd = &cobra.Command{}

	activateDeactivateCmd.Flags().StringVarP(&flags.Repository, "repository", "", "", "repository to publish plugin inventory image")
	activateDeactivateCmd.Flags().StringVarP(&flags.InventoryImageTag, "plugin-inventory-image-tag", "", "latest", "tag to which plugin inventory image needs to be published")
	activateDeactivateCmd.Flags().StringVarP(&flags.ManifestFile, "manifest", "", "", "manifest file specifying plugin details that needs to be processed")
	activateDeactivateCmd.Flags().StringVarP(&flags.Vendor, "vendor", "", "", "name of the vendor")
	activateDeactivateCmd.Flags().StringVarP(&flags.Publisher, "publisher", "", "", "name of the publisher")

	_ = activateDeactivateCmd.MarkFlagRequired("repository")
	_ = activateDeactivateCmd.MarkFlagRequired("vendor")
	_ = activateDeactivateCmd.MarkFlagRequired("publisher")
	_ = activateDeactivateCmd.MarkFlagRequired("manifest")

	return activateDeactivateCmd, flags
}