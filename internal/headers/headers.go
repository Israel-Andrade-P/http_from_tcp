package headers

import "bytes"

type Headers map[string]string

var rn = []byte("\r\n")

func NewHeaders() Headers {
	return map[string]string{}
}

func parseHeader(fieldLine []byte)

func (h Headers) Parse(data []byte) (int, bool, error) {

	read := 0
	done := false
	for {
		idx := bytes.Index(data, rn)
		if idx == -1 {
			break
		}
		//found the empty header that separates headers from body, we are done parsing
		if idx == 0 {
			done = true
			break
		}

		header := data[:idx]

	}
	return n, done, nil
}
