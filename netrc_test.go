package netrc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFile(t *testing.T) {
	wantData := Data{
		Machines: []Machine{
			{
				Name:     "machinegun",
				Login:    "youngboy",
				Password: "oldboy",
			},
			{
				Name:     "double-2",
				Password: "redact",
			},
			{
				Name:  "tripple-3",
				Login: "login",
			},
		},
	}

	data, err := ParseFile("testdata/example.netrc")
	assert.NoError(t, err)
	assert.Equal(t, wantData, data)
}
