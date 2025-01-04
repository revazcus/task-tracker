package commonTesting

import (
	"fmt"
	"github.com/revazcus/task-tracker/backend/infrastructure/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func AssertErrors(t *testing.T, err error, expectedErrors []error) {
	require.Error(t, err)

	errs, ok := err.(*errors.Errors)
	require.True(t, ok)

	require.Equal(t, len(expectedErrors), errs.Size())
	for _, expectedError := range expectedErrors {
		assert.True(t, errs.Contains(expectedError), fmt.Sprintf("Fail assert - missing expected error = %q", expectedError))
	}
}
