package status

import (
	"context"
	"fmt"
	"io"
	"reflect"
	"text/tabwriter"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/color"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	serviceapi "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/serviceapi"
	serviceapiclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/serviceapi/client"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
	"github.com/openconfig/goyang/pkg/indent"
)

const tagTitle = "title"

type Status struct {
	Kafka *KafkaStatus `json:"kafka,omitempty" title:"Kafka"`
}

type KafkaStatus struct {
	ID                  string `json:"id,omitempty"`
	Name                string `json:"name,omitempty"`
	Status              string `json:"status,omitempty"`
	BootstrapServerHost string `json:"bootstrap_server_host,omitempty" title:"Bootstrap URL"`
}

type Options struct {
	Config     config.IConfig
	Logger     func() (logging.Logger, error)
	Connection func() (connection.Connection, error)

	// request specific services
	Services []string
}

// Get gets the status of all services currently set in the user config
func Get(ctx context.Context, opts *Options) (status *Status, ok bool, err error) {
	cfg, err := opts.Config.Load()
	if err != nil {
		return nil, false, err
	}
	connection, err := opts.Connection()
	if err != nil {
		return nil, false, err
	}
	logger, err := opts.Logger()
	if err != nil {
		return nil, false, err
	}

	status = &Status{}
	api := connection.API()

	if stringInSlice("kafka", opts.Services) {
		kafkaCfg := cfg.Services.Kafka
		if cfg.HasKafka() {
			// nolint:govet
			kafkaStatus, err := getKafkaStatus(ctx, api.Kafka, kafkaCfg.ClusterID)
			if err != nil {
				if serviceapi.IsNotFoundError(err) {
					logger.Infof("Kafka instance with ID %v not found.", color.Info(kafkaCfg.ClusterID))
					logger.Info("Run", color.CodeSnippet("rhoas kafka use --id=<kafka-instance-id>"), "to use another Kafka instance.")
				}
			} else {
				status.Kafka = kafkaStatus
				ok = true
			}
		} else {
			logger.Debug("No Kafka instance is currently used, skipping status check")
		}
	}

	return status, ok, err
}

// Print prints the status information of all set services
func Print(w io.Writer, status *Status) {
	v := reflect.ValueOf(status).Elem()

	indirectVal := reflect.Indirect(v)
	for i := 0; i < indirectVal.NumField(); i++ {
		fieldType := indirectVal.Type().Field(i)
		fieldVal := indirectVal.Field(i)

		if !fieldVal.IsNil() {
			title := getTitle(&fieldType)
			fmt.Fprintln(w, "")
			printServiceStatus(w, title, fieldVal)
		}
	}
}

// print the status of service v
func printServiceStatus(w io.Writer, name string, v reflect.Value) {
	indentWriter := indent.NewWriter(w, "  ")

	// set table padding
	padding := 5

	// create a new tabwriter
	tw := tabwriter.NewWriter(indentWriter, 0, 0, padding, ' ', tabwriter.TabIndent)

	// tracks the longest row in chars so we can set an equal length divider
	maxRowLen := 0
	// iterate over every field in the type

	indirectV := reflect.Indirect(v)
	for i := 0; i < indirectV.NumField(); i++ {
		// get field type metadata
		fieldType := indirectV.Type().Field(i)

		// get value of the field
		fieldValue := indirectV.Field(i)

		// get the title to use for the field
		title := getTitle(&fieldType)

		// print the row and take note of its character length
		len, _ := fmt.Fprintf(tw, "%v:\t\t%v\n", title, fieldValue)
		if len > maxRowLen {
			maxRowLen = len
		}
	}
	// print the service header
	fmt.Fprintln(indentWriter, name)
	// print the title divider
	fmt.Fprintln(indentWriter, createDivider(maxRowLen+padding))

	tw.Flush()
}

// get title of the field from the "title" tag
// If the tag does not exist, use the name of the field
func getTitle(f *reflect.StructField) string {
	// Get the field tag value
	tag := f.Tag.Get(tagTitle)
	if tag == "" {
		tag = f.Name
	}

	return tag
}

// create a divider for the top of the table of n length
func createDivider(n int) string {
	b := "-"
	for i := 0; i <= n; i++ {
		b += "-"
	}

	return b
}

func getKafkaStatus(ctx context.Context, api serviceapiclient.DefaultApi, id string) (status *KafkaStatus, err error) {
	kafka, _, apiErr := api.GetKafkaById(ctx, id).Execute()
	if apiErr.Error() != "" {
		return nil, apiErr
	}

	status = &KafkaStatus{
		ID:                  kafka.GetId(),
		Name:                kafka.GetName(),
		Status:              kafka.GetStatus(),
		BootstrapServerHost: kafka.GetBootstrapServerHost(),
	}

	return status, err
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}