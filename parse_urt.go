package parseurt

import (
	"fmt"
	"strconv"
	"strings"
)

// https://nginx.org/en/docs/http/ngx_http_upstream_module.html#var_upstream_response_time
func ParseUpstreamResponseTime(responseTimeField string) (responseTime uint32, err error) {
	// Special case: no value for upstream response time
	if responseTimeField == "-" {
		return 0, nil
	}

	var groupURT []uint32 = []uint32{0}
	var fields []string = strings.Split(responseTimeField, " ")
	var parseError error = fmt.Errorf("bad field value: '%s'", responseTimeField)

	const stateGroupStartedNoValues = 1
	const stateGroupWithValues = 2
	const stateGroupWithValuesAndSep = 3
	parseState := stateGroupStartedNoValues

	for _, v := range fields {

		if v == "," {
			// Same group
			if parseState != stateGroupWithValues {
				return 0, parseError
			}

			parseState = stateGroupWithValuesAndSep

		} else if v == ":" {
			// New group started
			if parseState != stateGroupWithValues {
				return 0, parseError
			}

			groupURT = append(groupURT, 0)
			parseState = stateGroupStartedNoValues

		} else if strings.HasSuffix(v, ",") {
			// Same group, next value and comma
			if parseState != stateGroupStartedNoValues && parseState != stateGroupWithValuesAndSep {
				return 0, parseError
			}

			uv, err := strconv.ParseFloat(v[:len(v)-1], 64)
			if err != nil {
				return 0, err
			}

			urt := uint32(uv * 1000)
			gidx := len(groupURT) - 1
			if groupURT[gidx] < urt {
				groupURT[gidx] = urt
			}

			parseState = stateGroupWithValuesAndSep

		} else {
			// Same group, next value
			if parseState != stateGroupStartedNoValues && parseState != stateGroupWithValuesAndSep {
				return 0, parseError
			}

			uv, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return 0, err
			}

			urt := uint32(uv * 1000)
			gidx := len(groupURT) - 1
			if groupURT[gidx] < urt {
				groupURT[gidx] = urt
			}

			parseState = stateGroupWithValues
		}
	}

	if parseState != stateGroupWithValues {
		return 0, parseError
	}

	// Sum max group values
	for _, v := range groupURT {
		responseTime += v
	}

	return responseTime, nil
}
