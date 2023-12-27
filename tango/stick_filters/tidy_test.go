package stick_filters

import (
	"bytes"
	"fmt"
	"github.com/tyler-sommer/stick"
	"strings"
	"testing"
)

func TestTidyUpText(t *testing.T) {
	text := `hello
	  world
	  foo

  bar`
	expected := `hello
        world
        foo

    bar
`
	//LogLines(expected)
	//LogLines(tidyUpText(text))
	if tidyUpText(text) != expected {
		t.Error("Expected", expected, "got", tidyUpText(text))
	}
}

func TestTidyFilter(t *testing.T) {
	loader := stick.StringLoader{}
	var st = stick.New(&loader)
	st.Filters = map[string]stick.Filter{
		"tidy": tidyFilter,
	}
	input := `hello


    world`
	expected := `hello

    world
`
	buf := new(bytes.Buffer)
	st.Execute(
		"{{ input|tidy }}",
		buf,
		map[string]stick.Value{"input": input},
	)
	result := buf.String()
	if result != expected {
		t.Errorf("tidy returned an unexpected result: %v expected %v", result, expected)
	}

}

func LogLines(t string) {
	lines := strings.Split(t, "\n")
	for i, line := range lines {
		fmt.Println(i, " ", len(line), ":", line)
	}
}
