// This file is auto-generated by internal/cmd/genoptions/main.go. DO NOT EDIT

package jwk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptionIdent(t *testing.T) {
	require.Equal(t, "WithFetchBackoff", identFetchBackoff{}.String())
	require.Equal(t, "WithFetchWhitelist", identFetchWhitelist{}.String())
	require.Equal(t, "WithHTTPClient", identHTTPClient{}.String())
	require.Equal(t, "WithIgnoreParseError", identIgnoreParseError{}.String())
	require.Equal(t, "withLocalRegistry", identLocalRegistry{}.String())
	require.Equal(t, "WithMinRefreshInterval", identMinRefreshInterval{}.String())
	require.Equal(t, "WithPEM", identPEM{}.String())
	require.Equal(t, "WithRefreshInterval", identRefreshInterval{}.String())
	require.Equal(t, "WithThumbprintHash", identThumbprintHash{}.String())
}
