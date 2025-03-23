package command

import (
	"log"
	"strings"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	labelCommand = "command"
)

type Runner struct {
	prometheusNamespace string
	metricKeys          sync.Map
	counterVectors      map[string]*prometheus.CounterVec
}

// DefaultRunner does not record Prometheus metrics
var DefaultRunner = func() *Runner {
	result, err := NewRunner()
	if err != nil {
		log.Printf("Failed to create default runner: %v\n", err)
	}
	return result
}()

type RunnerOption func(*Runner) error

func NewRunner(options ...RunnerOption) (*Runner, error) {
	result := &Runner{}

	for _, option := range options {
		err := option(result)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// WithPrometheusNamespace sets the namespace for the prometheus metrics.
// Prometheus metrics will be exposed only if this is called.
func WithPrometheusNamespace(namespace string) RunnerOption {
	return func(r *Runner) error {
		r.prometheusNamespace = namespace
		// TODO: Validate / normalize namespace

		r.counterVectors = make(map[string]*prometheus.CounterVec)
		return nil
	}
}

func (r *Runner) prometheusEnabled() bool {
	return r.prometheusNamespace != ""
}

// Run executes the Command with the given name and arguments,
// injecting observability calls as required by the Runner configuration
func (r *Runner) Run(subsystem string, cmdName string, cmd Command, args ...string) error {
	if !r.prometheusEnabled() {
		return cmd(args)
	}

	// TODO: Prometheus...
	result := cmd(args)
	return result
}

func (r *Runner) counterVec(subsystem string) (*prometheus.CounterVec, error) {
	key := r.prometheusNamespace + "_" + subsystem
	if _, ok := r.counterVectors[key]; !ok {
		r.counterVectors[key] = prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: r.prometheusNamespace,
			Subsystem: subsystem,
			Name:      "command_run_total",
			Help:      "Command execution count",
		}, []string{labelCommand})
	}
	return r.counterVectors[key], nil
}

func (r *Runner) metricKey(subsystem string) (string, error) {
	return "blah", nil
}

// validCommandName returns true if the command name is valid, false otherwise.
// A valid command name is non-empty and trimmed.
func (r *Runner) validCommandName(cmdName string) bool {
	trimmedCmdName := strings.TrimSpace(cmdName)
	return trimmedCmdName != "" && trimmedCmdName == cmdName
}
