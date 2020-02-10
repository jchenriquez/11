package mostwater

import (
	"io"
	"encoding/json"
	"bufio"
	"os"
	"testing"
)

type Test struct {
	Input []int `json:"input"`
	Output int `json:"output"`
}

func TestMaxArea (test *testing.T) {
	f, err := os.Open("./tests.json")

	if err != nil {
		test.Error(err)
		return
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	decoder := json.NewDecoder(reader)

	var tests map[string]Test

	for {
		err = decoder.Decode(&tests)

		if err == nil {

			for name, tst := range tests {
				test.Run(name, func (st *testing.T) {

					if MaxArea(tst.Input) != tst.Output {
						st.Errorf("test data %v\n", tst)
					}

				})
			}

		} else if err == io.EOF {
			break
		} else {
			test.Error(err)
			break
		}
	}

}
