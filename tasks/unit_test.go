package tasks

import (
	"testing"
)

func TestRepeatStr(t *testing.T) {

	TestTable := []struct {
		count    int
		stroka   string
		expected string
	}{
		{
			count:    5,
			stroka:   "lol",
			expected: "lollollollollol",
		},
		{
			count:    2,
			stroka:   "qwe",
			expected: "qweqwe",
		},
		{
			count:    5,
			stroka:   "",
			expected: "",
		},
		{
			count:    0,
			stroka:   "qwert",
			expected: "",
		},
		{
			count:    1,
			stroka:   "qwert",
			expected: "testBug",
		},
	}

	for i, testCase := range TestTable {

		var result string = RepeatStr(testCase.count, testCase.stroka)

		t.Logf("OKEY. №%d, get %s", i, result)

		if result != testCase.expected {
			t.Errorf("ERROR. Expect %s, get %s", testCase.expected, result)
		}

	}

}
