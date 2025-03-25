package modules_test

import (
	"testing"

	"github.com/mewmewlab/eki/pkg/modules"
	"github.com/mewmewlab/eki/pkg/utils"
)

func init() {
	utils.InitUtils()
}

func TestUnitPull(t *testing.T) {
	u := modules.NewUnit("nginx", "latest")
	if err := u.Pull(nil); err != nil {
		t.Error(err)
	}
}
