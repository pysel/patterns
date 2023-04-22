package cor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRequestHandling(t *testing.T) {
	server := NewServer()
	tests := map[string]struct {
		reqString        string
		expectedResponse string
		expectingErr     bool
	}{
		"happy case: split last": {
			reqString:        "split-last|arbitrary data",
			expectedResponse: "data",
			expectingErr:     false,
		},
		"happy case: split first": {
			reqString:        "split-first|arbitrary data",
			expectedResponse: "arbitrary",
			expectingErr:     false,
		},
		"invalid request string": {
			reqString:        "split-last|arbitrary data|extra",
			expectedResponse: "",
			expectingErr:     true,
		},
		"invalid handler type": {
			reqString:        "split-unknown|arbitrary data",
			expectedResponse: "",
			expectingErr:     true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			response, err := server.HandleRequest(test.reqString)
			if test.expectingErr {
				require.Error(t, err)
				require.Equal(t, "", response, "unexpected response") // assert that response is not returned
				return
			}

			require.Equal(t, response, test.expectedResponse, "unexpected response")
		})
	}
}
