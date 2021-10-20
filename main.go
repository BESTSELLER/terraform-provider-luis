package main

import (
	"github.com/BESTSELLER/terraform-provider-luis/api"
	"github.com/BESTSELLER/terraform-provider-luis/config"
	"github.com/BESTSELLER/terraform-provider-luis/logger"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	
	tracer.Start(tracer.WithAgentAddr("datadog-agent.datadog:8125"))
	defer tracer.Stop()
	

	config.LoadEnvConfig()
	logger.Init()


	api.SetupRouter()
}
