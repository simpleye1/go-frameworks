package github

import (
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v41/github"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
)

type Options struct {
	IsEnterprise   bool
	URL            string
	KeyFile        string
	AppId          int64
	InstallationID int64
}

func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("github", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal github option error")
	}
	logger.Info("load github options success", zap.String("url", o.URL))
	return o, err
}

func NewGithub(v *viper.Viper, o *Options, logger *zap.Logger) (*github.Client, error) {
	// Shared transport to reuse TCP connections.
	tr := http.DefaultTransport
	// Wrap the shared transport for use with the app ID 1 authenticating with installation ID 99.
	itr, err := ghinstallation.NewKeyFromFile(tr, o.AppId, o.InstallationID, v.GetString("resources_path")+o.KeyFile)
	if err != nil {
		logger.Error("create github key fail", zap.Error(err))
		return nil, err
	}
	if o.IsEnterprise {
		itr.BaseURL = o.URL
		// Use installation transport with github.com/google/go-github
		client, err := github.NewEnterpriseClient(o.URL, o.URL, &http.Client{Transport: itr})
		if err != nil {
			logger.Error("create github client fail", zap.Error(err))
			return nil, err
		}
		return client, nil
	} else {
		client := github.NewClient(&http.Client{Transport: itr})
		return client, nil
	}
}

var ProviderSet = wire.NewSet(NewGithub, NewOptions)
