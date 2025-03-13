package apperr

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppErr(t *testing.T) {
	errUserNotFound := New("user not found").WithStatus(http.StatusNotFound)
	warpUserNotFound := fmt.Errorf("warp err : %w", errUserNotFound)

	assert.True(t, errors.Is(warpUserNotFound, errUserNotFound))
}
