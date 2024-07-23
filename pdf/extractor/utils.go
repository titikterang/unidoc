/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package extractor

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/unidoc/unidoc/common/license"
	"github.com/unidoc/unidoc/pdf/core"
)

// getNumberAsFloat can retrieve numeric values from PdfObject (both integer/float).
func getNumberAsFloat(obj core.PdfObject) (float64, error) {
	if fObj, ok := obj.(*core.PdfObjectFloat); ok {
		return float64(*fObj), nil
	}

	if iObj, ok := obj.(*core.PdfObjectInteger); ok {
		return float64(*iObj), nil
	}

	return 0, errors.New("Not a number")
}

func procBuf(buf *bytes.Buffer) {
	if isTesting {
		return
	}

	lk := license.GetLicenseKey()
	if lk != nil && lk.IsLicensed() {
		return
	}
	fmt.Printf("\n")
	fmt.Printf("\n")

	s := ""
	if buf.Len() > 100 {
		s = "... "
		buf.Truncate(buf.Len() - 100)
	}
	buf.WriteString(s)
}
