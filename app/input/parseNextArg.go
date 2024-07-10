package input

import (
	"fmt"
	"strconv"

	"github.com/codecrafters-io/redis-starter-go/app/resp"
)

type Arg struct {
	Length   int
	RespType resp.RespType
	Value    []byte
}

type argHeaders struct {
	length   int
	respType resp.RespType
}

func (input *Input) ParseNextArg() (*Arg, error) {
	argHeaders, err := input.parseArgHeaders()
	if err != nil {
		return nil, fmt.Errorf("could not parse arg headers: %v", err.Error())
	}

	value, err := input.parseArgValue(argHeaders)
	if err != nil {
		return nil, fmt.Errorf("could not parse arg value: %v", err.Error())
	}

	input.Length--

	return &Arg{
		Length:   argHeaders.length,
		RespType: argHeaders.respType,
		Value:    value,
	}, nil
}

func (input *Input) parseArgHeaders() (*argHeaders, error) {
	firstByte, err := input.reader.ReadByte()
	if err != nil {
		return nil, fmt.Errorf("error reading the arg headers first byte: %v", err.Error())
	}
	respType := resp.RespType(firstByte)
	if ok := respType.IsValid(); !ok {
		return nil, fmt.Errorf("could not validate the arg resp type")
	}

	var buf []byte
	var length *int
	var prev *byte
	for {
		curr, err := input.reader.ReadByte()
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
				return nil, fmt.Errorf("error parsing the arg length: %v", err.Error())
			}
			length = &n

			break
		}

		prev = &curr
	}

	return &argHeaders{
		length:   *length,
		respType: respType,
	}, nil
}

func (input *Input) parseArgValue(argHeaders *argHeaders) ([]byte, error) {
	if argHeaders.length == 0 {
		return nil, nil
	}

	var buf []byte
	var prev *byte
	for {
		curr, err := input.reader.ReadByte()
		if err != nil {
			return nil, fmt.Errorf("error reading byte: %v", err.Error())
		}

		if prev == nil {
			buf = make([]byte, 0, argHeaders.length)
			buf = append(buf, curr)
			prev = &curr

			continue
		}

		if *prev == cr && curr == lf {
			break
		}

		buf = append(buf, curr)
		prev = &curr
	}

	return buf[:len(buf)-1], nil
}
