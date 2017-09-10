package pkg

import (
	"testing"
	"fmt"
)

func Test_CreateChart(t *testing.T)()  {
	err := Build(map[string]string{
		"path":"/Users/andrepinto/Documents/workspace/golang/src/github.com/andrepinto/helmsman/demo",
		"destination":".",
		"repository":"http://localhost:8000/charts/upload/",
		"image":"andrepinto/server",
		"tag":"v1.0.0",
	})

	fmt.Println(err)
}


