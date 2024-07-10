package input

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/codecrafters-io/redis-starter-go/app/resp"
)

const (
	cr = '\r'
	lf = '\n'
)

type Input struct {
	Length int
	reader *bytes.Reader
	Size   int
}

func NewInput(buffer []byte) (*Input, error) {
	reader := bytes.NewReader(buffer)

	if err := guardAgaintInvalidBuffer(reader); err != nil {
		return nil, fmt.Errorf("invalid buffer")
	}

	size, err := parseInputArgsSize(reader)
	if err != nil {
		return nil, fmt.Errorf("error parsing input args size: %v", err.Error())
	}

	return &Input{
		Length: *size,
		reader: reader,
		Size:   *size,
	}, nil
}

func guardAgaintInvalidBuffer(reader *bytes.Reader) error {
	if reader.Size() <= 0 {
		return fmt.Errorf("empty packet")
	}

	respType, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("error reading the first byte: %v", err.Error())
	}
	if resp.RespType(respType) != resp.ArraysRespType {
		return fmt.Errorf("invalid first byte")
	}

	return nil
}

func parseInputArgsSize(reader *bytes.Reader) (*int, error) {
	var buf []byte
	var prev *byte
	var size *int
	for {
		curr, err := reader.ReadByte()
		if err != nil {
			return nil, fmt.Errorf("error reading byte: %v", err.Error())
		}

		if prev == nil {
			buf = make([]byte, 0)
			buf = append(buf, curr)
			prev = &curr

			continue
		}

		if *prev == cr && curr == lf {
			n, err := strconv.Atoi(string(buf))
			if err != nil {
				return nil, fmt.Errorf("error parsing the input args size: %v", err.Error())
			}
			size = &n

			break
		}

		prev = &curr
	}

	return size, nil
}
