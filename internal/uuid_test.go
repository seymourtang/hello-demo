package internal

import (
	"strings"
	"testing"
)

func Test_GetUuid(t *testing.T) {
	uuid := strings.Replace(GetUuid(), "-", "", -1)
	if len(uuid) != 32 {
		t.Error("the length is not 32")
	}
}
