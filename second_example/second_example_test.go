package second_example

import (
	"io"
	"os"
	"strings"
	"testing"
)

//func Test_updateMessage(t *testing.T) {
//	msg = "Hello, world!"
//
//	wg.Add(2)
//	go updateMessage("Xanadu")
//	go updateMessage("Hello, universe!")
//	wg.Wait()
//
//	if msg != "Hello, universe!" {
//		t.Errorf("Expected msg to be 'Hello, universe!' but got '%s'", msg)
//	}
//}

func Test_AnotherSecondExample(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	AnotherSecondExample()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "3744000.00") {
		t.Errorf("Expected output to contain '3744000.00' but got '%s'", output)
	}

}
