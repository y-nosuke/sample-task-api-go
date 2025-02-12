package otel

type Config struct {
	ExporterEndpoint string `env:"EXPORTER_ENDPOINT"`
}
