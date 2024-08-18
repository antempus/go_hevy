package main

import (
	"context"
	"github.com/kelseyhightower/envconfig"
	hvy "go_hevy/client"
	"go_hevy/support"
	"log/slog"
	"net/http"
	"net/url"
)

// TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.
const UrlHost = "<TBD>"
const Token = "<TBD>"

func main() {
	var hvyConfig HevyConfig
	var gblConfig GlobalConfig
	err := envconfig.Process("HEVY", &hvyConfig)
	if err != nil {
		panic(err)
	}

	err = envconfig.Process("", &gblConfig)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	httpClient := &http.Client{}
	logger := slog.Logger{}

	obsvr := support.Observer{
		LogLevel: gblConfig.LogLevel,
		Logger:   logger,
	}

	requester := hvy.Requester{
		Url: &url.URL{RawPath: hvyConfig.ApiHost + "/v1/"},
		Headers: map[string]string{
			"api-key": hvyConfig.ApiKey,
		},
		HttpClient: httpClient,
		UserAgent:  "antempus/go_hevy@v0.0.1",
		Context:    ctx,
	}

	hevyClient := &hvy.HevyClient{
		Requester: requester,
		Context:   ctx,
		Client:    httpClient,
	}
	params := hvy.PaginationParams{Page: 2}

	result, err := hevyClient.GetExerciseTemplates(params)
	if err != nil {
		panic(err)
	}
	obsvr.LogJson(result)

	// TODO: @antempus - Setup Tests for client
}
