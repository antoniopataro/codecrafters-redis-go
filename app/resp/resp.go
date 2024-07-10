package resp

import "fmt"

type RespType byte

const (
	ArraysRespType       RespType = '*'
	BulkStringRespType   RespType = '$'
	IntegerRespType      RespType = ':'
	SimpleStringRespType RespType = '+'
)

const (
	crlf = "\r\n"
)

func (respType *RespType) IsValid() bool {
	switch *respType {
	case ArraysRespType, BulkStringRespType, IntegerRespType, SimpleStringRespType:
		return true
	default:
		return false
	}
}

func BulkString(bulkString string) []byte {
	return []byte(fmt.Sprintf("%s%d%s%s%s", string(BulkStringRespType), len(bulkString), crlf, bulkString, crlf))
}

func SimpleString(simpleString string) []byte {
	return []byte(fmt.Sprintf("%s%s%s", string(SimpleStringRespType), simpleString, crlf))
}
