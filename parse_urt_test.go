package parseurt

import (
	"fmt"
	"testing"
)

func Test_ParseUpstreamResponseTime(t *testing.T) {
	tests := []struct {
		name   string
		field  string
		uRT    uint32
		hasErr bool
	}{
		// Good
		{
			field:  "-",
			uRT:    0,
			hasErr: false,
		},
		{
			field:  "0.032",
			uRT:    32,
			hasErr: false,
		},
		{
			field:  "0.032 : 0.021",
			uRT:    53,
			hasErr: false,
		},
		{
			field:  "0.037 , 0.021",
			uRT:    37,
			hasErr: false,
		},
		{
			field:  "0.037, 0.021",
			uRT:    37,
			hasErr: false,
		},
		{
			field:  "0.032, 0.021, 0.008 : 0.109",
			uRT:    141,
			hasErr: false,
		},
		{
			field:  "0.032 : 0.021, 0.008 : 0.109",
			uRT:    162,
			hasErr: false,
		},
		{
			field:  "0.032 : 0.021 , 0.008 : 0.109",
			uRT:    162,
			hasErr: false,
		},
		// Bad
		{
			field:  "0.032,",
			uRT:    0,
			hasErr: true,
		},
		{
			field:  "0.032 ,",
			uRT:    0,
			hasErr: true,
		},
		{
			field:  ",0.032",
			uRT:    0,
			hasErr: true,
		},
		{
			field:  ", 0.032",
			uRT:    0,
			hasErr: true,
		},
		{
			field:  "0.032 :",
			uRT:    0,
			hasErr: true,
		},
		{
			field:  ": 0.032",
			uRT:    0,
			hasErr: true,
		},
	}
	for i, tt := range tests {
		if tt.name == "" {
			tt.name = fmt.Sprintf("Test%d", i+1)
		}

		t.Run(tt.name, func(t *testing.T) {
			uRT, err := ParseUpstreamResponseTime(tt.field)
			if (err != nil) != tt.hasErr || uRT != tt.uRT {
				t.Errorf(
					"parseUpstreamResponseTime(): got: uRT: %v, hasErr: %v, expected: uRT: %v, hasErr: %v,",
					uRT, err != nil, tt.hasErr, tt.hasErr,
				)
			}
		})
	}
}
