package teams

import (
	"testing"

	"github.com/keybase/client/go/libkb"
)

func SetupTest(tb testing.TB, name string, depth int) (tc libkb.TestContext) {
	ret := libkb.SetupTest(tb, name, depth+1)
	NewTeamLoaderAndInstall(ret.G)
	return ret
}
