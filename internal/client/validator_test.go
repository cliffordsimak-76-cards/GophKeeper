package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_notEmpty(t *testing.T) {
	t.Run("error empty", func(t *testing.T) {
		require.NotNil(t, notEmpty(""))
	})
	t.Run("success", func(t *testing.T) {
		require.Nil(t, notEmpty("test"))
	})
}

func Test_any(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		require.Nil(t, any(""))
	})
}
