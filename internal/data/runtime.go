package data

import (
	"fmt"
	"strconv"
	"strings"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}

func (r *Runtime) UnmarshalJSON(data []byte) error {
	trimmedStr := strings.Trim(string(data), `"`)
	strValue := strings.Split(trimmedStr, " ")[0]
	intValue, err := strconv.ParseInt(strValue, 10, 32)
	if err != nil {
		return err
	}

	*r = Runtime(int32(intValue))
	return nil
}
