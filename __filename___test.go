package __package___test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"os"
	"log"
	"github.com/jmcvetta/randutil"
)

func Setup() {

}

func Teardown() {

}

func Test__funcname__(t *testing.T) {
  assert.True(t, true, "true should be true")
}

func TestMain(m *testing.M) {
  Setup()
  result := m.Run()
  Teardown()
  os.Exit(result)
}
