package main

import (
	"net/http"

	"github.com/coltoneshaw/mattermost-plugin-poc-userid/server/api"
	"github.com/mattermost/mattermost/server/public/plugin"
	"github.com/sirupsen/logrus"

	pluginapi "github.com/mattermost/mattermost/server/public/pluginapi"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.

type Plugin struct {
	plugin.MattermostPlugin

	handler   *api.Handler
	pluginAPI *pluginapi.Client
}

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

// ServeHTTP routes incoming HTTP requests to the plugin's REST API.
func (p *Plugin) ServeHTTP(_ *plugin.Context, w http.ResponseWriter, r *http.Request) {
	p.handler.ServeHTTP(w, r)
}

func (p *Plugin) OnActivate() error {
	pluginAPIClient := pluginapi.NewClient(p.API, p.Driver)
	p.pluginAPI = pluginAPIClient

	logger := logrus.StandardLogger()
	pluginapi.ConfigureLogrus(logger, pluginAPIClient)

	p.handler = api.NewHandler(pluginAPIClient)

	api.NewTestHandler(
		p.handler.APIRouter,
		pluginAPIClient,
	)

	return nil
}
