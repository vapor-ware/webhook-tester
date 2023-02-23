package checkers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gh.tarampamp.am/webhook-tester/internal/checkers"
)

func TestLiveChecker_Check(t *testing.T) {
	assert.NoError(t, checkers.NewLiveChecker().Check())
}
