package core_test

import (
	// "fmt"
	"fmt"
	"testing"

	"github.com/ra0van/ricedb/core"
)

func TestSimpleStringDecode(t *testing.T) {
    cases := map[string]string {
        "+OK\r\n": "OK",
    }

    for k,v := range cases {
        value, _ := core.Decode([]byte(k))
        if v != value {
            t.Fail()
        }
    }
}

func TestError(t *testing.T){
    cases := map[string]string{
        "-Error message\r\n" : "Error message",
    }

    for k,v := range cases {
        value, _ := core.Decode([]byte(k))
        if v != value {
            t.Fail()
        }
    }
}

func TestInt64(t *testing.T){
    cases := map[string]int64{
        ":0\r\n" : 0,
        ":10001\r\n" : 10001,
    }

    for k,v := range cases{
        value, _ := core.Decode([]byte(k))
        // ivalue := value.(int64)
        if v != value {
            t.Fail()
        }
    }
}

func TestBulkString(t *testing.T){
    cases := map[string]string {
        "$5\r\nhello\r\n": "hello",
        "$0\r\n\r\n" : "",
    }

    for k,v := range cases {
        value, _ := core.Decode([]byte(k))
        if v != value {
            t.Fail()
        }
    }
}

func TestArrayDecode(t *testing.T){
    cases := map[string][]interface{} {
        "*0\r\n" :                                                  {},
        "*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n":                     {"hello", "world"},
        "*3\r\n:1\r\n:2\r\n:3\r\n":                                 {int64(1), int64(2), int64(3)},
        "*5\r\n:1\r\n:2\r\n:3\r\n:4\r\n$5\r\nhello\r\n":            {int64(1), int64(2), int64(3), int64(4), "hello"},
        "*2\r\n*3\r\n:1\r\n:2\r\n:3\r\n*2\r\n+Hello\r\n-World\r\n": {[]int64{int64(1), int64(2), int64(3)}, []interface{}{"Hello", "World"}},

    }

    for k,v := range cases{
        value, _ := core.Decode([]byte(k))
        array := value.([]interface{})
        if len(array) != len(v) {
            t.Logf("Mismatch in length")
            t.Fail()
        }

        for i := range array {
            t.Logf("values " + fmt.Sprintf("%v ", v[i]) + fmt.Sprintf("%v", array[i]))

            if fmt.Sprintf("%v", v[i]) != fmt.Sprintf("%v", array[i]){
                t.Logf("Mismatch in values" + fmt.Sprintf("%v", v[i]) + fmt.Sprintf("%v", array[i]))
                t.Logf("Full array " + fmt.Sprintf("%v ", array))
                t.Fail()
            }
        }
    }
}


