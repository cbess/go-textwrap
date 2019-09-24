package textwrap

import (
	"strings"
	"testing"
)

func joinStrings(strs []string) string {
	return strings.Join(strs, " ")
}

func TestWordWrapSimple(t *testing.T) {
	result, err := WordWrap(" 3 7 ", 1, 0)
	if err != nil {
		t.Error(err)
	}

	if !result.IsValid() {
		t.Error("The result is not valid")
	}

	if len(result.TextGroups) != 2 {
		t.Errorf("Wrong group count, should 2 be, got %d", len(result.TextGroups))
	}
}

func TestWordWrapStringRebuild(t *testing.T) {
	origText := "Jesus is God. He died for sinners and rose from the dead by His own power and is coming back to judge the world in righteousness. Repent and believe the gospel!"

	result, err := WordWrap(origText, 50, 0)
	if err != nil {
		t.Error(err)
	}

	text := joinStrings(result.TextGroups)
	if text != origText {
		t.Logf("orig text: %s", origText)
		t.Logf("join text: %s", text)

		t.Error("The original text does NOT match the wrapped and joined text")
	}
}

func TestWordWrapSingleGroup(t *testing.T) {
	// no need to wrap, text fits in width
	origText := "Jesus is God. He Saves by grace through faith alone."
	result, err := WordWrap(origText, len(origText)+1, 0)
	if err != nil {
		t.Error(err)
	}

	if len(result.TextGroups) != 1 {
		t.Error("Expected one text group")
	}

	// a width smaller than the largest word
	result, err = WordWrap("Propitiation", 1, 0)
	if err != nil {
		t.Error(err)
	}

	if len(result.TextGroups) != 1 {
		t.Error("Expected one text group")
	}

	// no words
	result, err = WordWrap("", 1, 0)
	if err != nil {
		t.Error(err)
	}

	if !result.IsValid() {
		t.Error("Empty result should be valid")
	}

	if len(result.TextGroups) != 0 {
		t.Errorf("Expected no text groups")
	}
}

func TestWordWrapFail(t *testing.T) {
	// text exceeds max word count
	_, err := WordWrap("one two three", 7, 2)
	if err == nil {
		t.Error("Expected to return error")
	}
}
