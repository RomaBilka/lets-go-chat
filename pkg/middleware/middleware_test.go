package middleware

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func logRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("teeeeest")
}

func TestLogRequest(t *testing.T) {

	buf := &bytes.Buffer{}

	// Redirect STDOUT to a buffer
	//stdout := os.Stdout
	r, _, err := os.Pipe()
	if err != nil {
		t.Errorf("Failed to redirect STDOUT")
	}
	//os.Stdout = w
	go func() {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			fmt.Println(scanner.Text()+"1111")
			buf.WriteString(scanner.Text())
		}
	}()

	fmt.Println(55555)

	ts := httptest.NewServer(LogRequest(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", greeting)
//////////////////////////
//	w.Close()
//	os.Stdout = stdout
//
//	// Test output
//	t.Log(buf)
//	if buf.Len() == 0 {
//		t.Error("No information logged to STDOUT")
//	}
//
//	if strings.Count(buf.String(), "\n") > 1 {
//		t.Error("Expected only a single line of log output")
//	}

}
