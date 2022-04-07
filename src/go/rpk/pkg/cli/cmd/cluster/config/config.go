// Copyright 2021 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package config

import (
	"github.com/redpanda-data/redpanda/src/go/rpk/pkg/cli/cmd/common"
	"github.com/redpanda-data/redpanda/src/go/rpk/pkg/config"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func NewConfigCommand(fs afero.Fs) *cobra.Command {
	var (
		all            bool
		adminURL       string
		adminEnableTLS bool
		adminCertFile  string
		adminKeyFile   string
		adminCAFile    string
	)

	command := &cobra.Command{
		Use:   "config",
		Args:  cobra.ExactArgs(0),
		Short: "Interact with cluster configuration properties.",
		Long: `Interact with cluster configuration properties.

Cluster properties are redpanda settings that apply to all nodes in
the cluster.  These are separate from node properties, which are set with
'rpk redpanda config'.

Use the 'edit' subcommand to interactively modify the cluster configuration, or
'export' and 'import' to write a configuration to a file that can be edited and
read back later.

These commands take an optional '--all' flag to include all properties. This includes
low-level tunables such as internal buffer sizes, which do not usually need
to be changed during normal operations. These properties generally require
some expertise to set safely, so when in doubt, avoid using '--all'.

Modified properties are propagated immediately to all nodes.  Use the 'status'
subcommand to verify that all nodes are up to date, and to identify
any settings that were rejected by a node. For example, a node might reject a setting if that node is running a redpanda version that does not recognize certain properties.`,
	}

	command.PersistentFlags().StringVar(
		&adminURL,
		config.FlagAdminHosts2,
		"",
		"Comma-separated list of admin API addresses (<IP>:<port>")

	common.AddAdminAPITLSFlags(command,
		&adminEnableTLS,
		&adminCertFile,
		&adminKeyFile,
		&adminCAFile,
	)

	command.PersistentFlags().BoolVar(
		&all,
		"all",
		false,
		"Include all properties, including tunables.",
	)

	command.AddCommand(
		newImportCommand(fs, &all),
		newExportCommand(fs, &all),
		newEditCommand(fs, &all),
		newStatusCommand(fs),
		newForceResetCommand(fs),
		newLintCommand(fs),
		newSetCommand(fs),
		newGetCommand(fs),
	)

	return command
}
