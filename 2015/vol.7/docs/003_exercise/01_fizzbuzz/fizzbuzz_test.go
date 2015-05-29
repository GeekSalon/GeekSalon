package fizzbuzz

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestPrintFizzBuzz(t *testing.T) {
	backup := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	opChan := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		opChan <- buf.String()
	}()

	PrintFizzBuzz(100)

	w.Close()
	os.Stdout = backup

	output := <-opChan
	if output != expectedString {
		// t.Errorf("output does not equal\n expected=\n%s\n\n actual=\n%s\n\n", expectedString, output)
		t.Errorf("output does not equal\n output=\n%s\n", output)
	}
}

func TestFizzBuzz(t *testing.T) {
	result := FizzBuzz(100)
	if !reflect.DeepEqual(result, expectedList) {
		t.Errorf("result does not equal.\n expected=%#v\n\n actual=%#v\n", expectedList, result)
	}
}

var expectedList = []interface{}{1, 2, "Fizz", 4, "Buzz", "Fizz", 7, 8, "Fizz", "Buzz", 11, "Fizz", 13, 14, "FizzBuzz", 16, 17, "Fizz", 19, "Buzz", "Fizz", 22, 23, "Fizz", "Buzz", 26, "Fizz", 28, 29, "FizzBuzz", 31, 32, "Fizz", 34, "Buzz", "Fizz", 37, 38, "Fizz", "Buzz", 41, "Fizz", 43, 44, "FizzBuzz", 46, 47, "Fizz", 49, "Buzz", "Fizz", 52, 53, "Fizz", "Buzz", 56, "Fizz", 58, 59, "FizzBuzz", 61, 62, "Fizz", 64, "Buzz", "Fizz", 67, 68, "Fizz", "Buzz", 71, "Fizz", 73, 74, "FizzBuzz", 76, 77, "Fizz", 79, "Buzz", "Fizz", 82, 83, "Fizz", "Buzz", 86, "Fizz", 88, 89, "FizzBuzz", 91, 92, "Fizz", 94, "Buzz", "Fizz", 97, 98, "Fizz", "Buzz"}

var expectedString = `1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
16
17
Fizz
19
Buzz
Fizz
22
23
Fizz
Buzz
26
Fizz
28
29
FizzBuzz
31
32
Fizz
34
Buzz
Fizz
37
38
Fizz
Buzz
41
Fizz
43
44
FizzBuzz
46
47
Fizz
49
Buzz
Fizz
52
53
Fizz
Buzz
56
Fizz
58
59
FizzBuzz
61
62
Fizz
64
Buzz
Fizz
67
68
Fizz
Buzz
71
Fizz
73
74
FizzBuzz
76
77
Fizz
79
Buzz
Fizz
82
83
Fizz
Buzz
86
Fizz
88
89
FizzBuzz
91
92
Fizz
94
Buzz
Fizz
97
98
Fizz
Buzz
`
