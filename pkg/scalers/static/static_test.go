package static

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewStaticScaler(t *testing.T) {
	scale := scaleResponse{
		Name:     "foo",
		Replicas: 0,
		Region:   "east",
		Primary:  false,
	}
	response, err := json.Marshal(scale)
	if err != nil {
		return
	}
	fmt.Println(string(response))
}
