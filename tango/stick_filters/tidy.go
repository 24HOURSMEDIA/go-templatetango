package stick_filters

import (
	"github.com/tyler-sommer/stick"
	"templatetango/tidytext"
)

func tidyUpText(text string) string {
	tidy := tidytext.NewTidyText(text)
	tidy.TabsToSpaces().
		TrimEmptyLines().
		RemoveDoubleNewlines().
		RemoveLeadingAndTrailingNewlines().
		ReformatIndentation().
		AddTrailingNewLine()
	return tidy.GetText()
}

func tidyFilter(
	ctx stick.Context,
	val stick.Value,
	args ...stick.Value,
) stick.Value {
	text := stick.CoerceString(val)
	strategy := "text"
	if len(args) > 0 {
		strategy = stick.CoerceString(args[0])
	}
	switch strategy {
	case "text":
		return tidyUpText(text)
	default:
		return val
	}
}
