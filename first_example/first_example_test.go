package first_example

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup

	wg.Add(1)

	go printSomething("Hello, World!", &wg)

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)

	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, World!") {
		t.Errorf("Expected output to contain 'Hello, World!' but got '%s'", output)
	}
}

func Test_Challenge1(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	Challenge1()

	_ = w.Close()

	result, _ := io.ReadAll(r)

	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Errorf("Expected output to contain 'Hello, universe!' but got '%s'", output)
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Errorf("Expected output to contain 'Hello, cosmos!' but got '%s'", output)
	}

	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("Expected output to contain 'Hello, world!' but got '%s'", output)
	}
}
