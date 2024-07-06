package svg2png

import (
	"bytes"
	"fmt"
	"os/exec"
)

const (
	/*
		Default inkscape binary path
	*/
	BINARY = "/usr/bin/inkscape"
)

type Converter struct {
	bin string
}

/*
	Return a new instance of the converter
*/
func New() *Converter {
	var c Converter
	c.bin = BINARY
	return &c
}

/*
	Set custom inkscape binary path instead of the default
*/
func (c *Converter) SetBinary(b string) error {
	if len(b) == 0 {
		return fmt.Errorf("Empty binary path")
	}
	c.bin = b
	return nil
}

/*
	Try to convert the input SVG to the PNG image
*/
func (c *Converter) Convert(in []byte) (out []byte, err error) {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command(c.bin, "--pipe", "--export-type=png", "--export-filename=-")
	cmd.Stdin = bytes.NewBuffer(in)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if e := cmd.Run(); e != nil {
		err = fmt.Errorf("%s\nSTDERR:\n%s", e.Error(), stderr.String())
		return
	}

	if stdout.Len() == 0 {
		err = fmt.Errorf("Got no data from inkscape")
		return
	}

	out = stdout.Bytes()

	return
}
