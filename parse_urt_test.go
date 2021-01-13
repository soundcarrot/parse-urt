package parseurt

import (
	"fmt"
	"math"
	"testing"
)

func Test_ParseUpstreamResponseTime(t *testing.T) {
	tests := []struct {
		name   string
		field  string
		uRT    float64
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
			uRT:    0.032,
			hasErr: false,
		},
		{
			field:  "0.032 : 0.021",
			uRT:    0.053,
			hasErr: false,
		},
		{
			field:  "0.037 , 0.021",
			uRT:    0.037,
			hasErr: false,
		},
		{
			field:  "0.037, 0.021",
			uRT:    0.037,
			hasErr: false,
		},
		{
			field:  "0.032, 0.021, 0.008 : 0.109",
			uRT:    0.141,
			hasErr: false,
		},
		{
			field:  "0.032 : 0.021, 0.008 : 0.109",
			uRT:    0.162,
			hasErr: false,
		},
		{
			field:  "0.032 : 0.021 , 0.008 : 0.109",
			uRT:    0.162,
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
			if (err != nil) != tt.hasErr || math.Abs(uRT-tt.uRT) > 1.0e-6 {
				t.Errorf(
					"parseUpstreamResponseTime(): got: uRT: %v, hasErr: %v, expected: uRT: %v, hasErr: %v,",
					uRT, err != nil, tt.uRT, tt.hasErr,
				)
			}
		})
	}
}
