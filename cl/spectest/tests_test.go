package spectest

import (
	"os"
	"testing"

	"github.com/optimism-java/erigon/spectest"

	"github.com/optimism-java/erigon/cl/transition"

	"github.com/optimism-java/erigon/cl/spectest/consensus_tests"
)

func Test(t *testing.T) {
	spectest.RunCases(t, consensus_tests.TestFormats, transition.ValidatingMachine, os.DirFS("./tests"))
}
