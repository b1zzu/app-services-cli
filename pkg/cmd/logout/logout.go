// Package cluster contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package logout

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
)

type Options struct {
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	localizer  localize.Localizer
}

// NewLogoutCommand gets the command that's logs the current logged in user
func NewLogoutCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:   opts.localizer.MustLocalize("logout.cmd.use"),
		Short: opts.localizer.MustLocalize("logout.cmd.shortDescription"),
		Long:  opts.localizer.MustLocalize("logout.cmd.longDescription"),
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runLogout(opts)
		},
	}
	return cmd
}

func runLogout(opts *Options) error {
	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	err = connection.Logout(context.TODO())

	if err != nil {
		return fmt.Errorf("%v: %w", opts.localizer.MustLocalize("logout.error.unableToLogout"), err)
	}

	logger.Info(opts.localizer.MustLocalize("logout.log.info.logoutSuccess"))

	return nil
}
