package tidytext

import (
	"fmt"
	"strings"
	"testing"
)

func TestTidyText_RemoveDoubleNewlines(t *testing.T) {
	tidyText := NewTidyText("Hello\n\n\nWorld")
	tidyText.RemoveDoubleNewlines()
	if tidyText.GetText() != "Hello\n\nWorld" {
		t.Error("Expected 'Hello\nWorld', got", tidyText.GetText())
	}

	text := `foo

bar`
	expected := text
	actual := NewTidyText(text).RemoveDoubleNewlines().GetText()
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestTidyText_TabsToSpaces(t *testing.T) {
	tidyText := NewTidyText("Hello\tWorld").SetIndentation(4)
	tidyText.TabsToSpaces()
	if tidyText.GetText() != "Hello    World" {
		t.Error("Expected 'Hello    World', got", tidyText.GetText())
	}
}

func TestTidyText_TrimEmptyLines(t *testing.T) {
	tidyText := NewTidyText("Hello\n\nWorld")
	tidyText.TrimEmptyLines()
	if tidyText.GetText() != "Hello\n\nWorld" {
		t.Error("Expected 'Hello\n\nWorld', got", tidyText.GetText())
	}

	tidyText = NewTidyText("Hello\n \nWorld")
	tidyText.TrimEmptyLines()
	if tidyText.GetText() != "Hello\n\nWorld" {
		t.Error("Expected 'Hello\n\nWorld', got", tidyText.GetText())
	}
	tidyText = NewTidyText("Hello\n\t\n \n\nWorld")
	tidyText.TrimEmptyLines()
	if tidyText.GetText() != "Hello\n\n\n\nWorld" {
		t.Error("Expected 'Hello\n\n\n\nWorld', got", tidyText.GetText())
	}

}

func TestTidyText_ReformatIndentation(t *testing.T) {
	tidyText := NewTidyText(`
hello
  world
  foo
    bar
 foobar
`).SetIndentation(4).ReformatIndentation()
	expected := `
hello
        world
        foo
            bar
    foobar
`
	if tidyText.GetText() != expected {
		t.Error("Expected `"+expected+"`, got", tidyText.GetText())
	}
}

func LogLines(t string) {
	lines := strings.Split(t, "\n")
	for i, line := range lines {
		fmt.Println(i, " ", len(line), ":", line)
	}
}
