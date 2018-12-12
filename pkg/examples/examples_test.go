package examples

import (
	"github.com/elek/flekszible/pkg/processor"
	"testing"
)

func TestGettingStarted(t *testing.T) {
	processor.TestExample(t, "gettingstarted")
}

func TestGettingEnvs(t *testing.T) {
	processor.TestExample(t, "envs/dev")
}