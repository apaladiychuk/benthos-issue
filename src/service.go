package main

import (
	"benthos-issue/plugin"
	_ "embed"
	"github.com/redpanda-data/benthos/v4/public/service"
	_ "github.com/redpanda-data/connect/v4/public/components/io"
	_ "github.com/redpanda-data/connect/v4/public/components/jaeger"
	_ "github.com/redpanda-data/connect/v4/public/components/nats"
	_ "github.com/redpanda-data/connect/v4/public/components/pure"
	_ "plugin"
)

//go:embed stream.yaml
var streamSrc string

//go:embed stream-input.yaml
var streamSrcInput string

func BuildStream(src string) (service.MessageHandlerFunc, *service.Stream, error) {

	builder := service.NewStreamBuilder()
	// Read the YAML configuration file

	customPlugin := plugin.NewCustomPlugin()
	err := service.RegisterProcessor("custom", customPlugin.GetConfig(), customPlugin.Constructor)
	if err != nil {
		return nil, nil, err
	}

	// Set the configuration from the YAML file
	err = builder.SetYAML(src)
	if err != nil {
		return nil, nil, err
	}

	err = builder.SetLoggerYAML(`level: off`)
	if err != nil {
		return nil, nil, err
	}

	sendFn, err := builder.AddProducerFunc()

	if err != nil {
		return nil, nil, err
	}
	err = builder.AddOutputYAML(`drop: {}`)

	stream, err := builder.Build()

	if err != nil {
		return nil, nil, err
	}

	return sendFn, stream, nil
}
