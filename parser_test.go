package axmlParser

import (
	"testing"
)

func TestParser(t *testing.T) {
	var filename = "a.apk"

	listener := new(PlainListener)
	_, err := ParseApk(filename, listener)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v\n", listener)
	//	fmt.Println("Init package is", listener.PackageName,
	//		"Activity is", listener.ActivityName)
}
