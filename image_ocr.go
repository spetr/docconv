// +build ocr

package docconv

import (
	"fmt"
	"io"
	"sync"

	"github.com/otiai10/gosseract"
)

var langs = struct {
	sync.RWMutex
	lang string
}{lang: "eng"}

// ConvertImage converts images to text.
// Requires gosseract.
func ConvertImage(r io.Reader) (string, map[string]string, error) {
	f, err := NewLocalFile(r, "/tmp", "sajari-convert-")
	if err != nil {
		return "", nil, fmt.Errorf("error creating local file: %v", err)
	}
	defer f.Done()

	meta := make(map[string]string)

	gs := gosseract.NewClient()
	defer gs.Close()
	gs.SetImage(f.Name())
	body, err := gs.Text()
	if err != nil {
		return "", nil, fmt.Errorf("tesseract error: %v", err)
	}
	fmt.Println(body)

	return body, meta, nil
}

// SetImageLanguages sets the languages parameter passed to gosseract.
func SetImageLanguages(l string) {
	langs.Lock()
	langs.lang = l
	langs.Unlock()
}
