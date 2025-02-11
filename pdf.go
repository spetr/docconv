// +build !ocr

package docconv

import (
	"fmt"
	"io"
)

// ConvertPDF converts an Adobe Acrobat PDF file to text.
func ConvertPDF(r io.Reader) (string, map[string]string, error) {

	f, err := NewLocalFile(r, "/tmp", "sajari-convert-")
	if err != nil {
		return "", nil, fmt.Errorf("error creating local file: %v", err)
	}
	defer f.Done()

	bodyResult, metaResult, convertErr := ConvertPDFText(f.Name())
	if convertErr != nil {
		return "", nil, convertErr
	}
	if bodyResult.err != nil {
		return "", nil, bodyResult.err
	}
	if metaResult.err != nil {
		return "", nil, metaResult.err
	}
	return bodyResult.body, metaResult.meta, nil

}
