package main

import (
	"fmt"
	"strings"
	"testing"

	"gke/troubleshooting"

	"github.com/stretchr/testify/assert"
)

func TestThis(t *testing.T) {
	arr := troubleshooting.GetNodeFailureCmds()
	//fmt.Println(arr[1])

	t1 := arr[1]
	arr[1] = arr[len(arr)-1]
	arr[len(arr)-1] = t1

	fmt.Println(arr[1])

	//if !strings.Contains(arr[1], "kubectl") {
	//	t.Errorf("blah error '%s'", arr[1])
	//}

	assert.Equal(t, false, strings.Contains(arr[1], "kubectl"), "The two words should be the same.")

}
