package pyfmt

import (
	"testing"
)

func TestFormatEmpty(t *testing.T) {
	actual := Sprintf("", make(map[string]string))
	expected := ""
	if actual != expected {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestFormatNoSubs(t *testing.T) {
	actual := Sprintf("no substitutions", make(map[string]string))
	expected := "no substitutions"
	if actual != expected {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestFormatAllSubsFound(t *testing.T) {
	dict := map[string]string{"name": "John", "age": "99"}
	actual := Sprintf("hey {name}, are you {age} years old?", dict)
	expected := "hey John, are you 99 years old?"
	if actual != expected {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestFormatUnbalancedOpenCurly(t *testing.T) {
	actual := Sprintf("a{hey", make(map[string]string))
	expected := "a{hey"
	if actual != expected {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestFormatUnbalancedCloseCurly(t *testing.T) {
	actual := Sprintf("a}hey", make(map[string]string))
	expected := "a}hey"
	if actual != expected {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestFormatSwappedCurly(t *testing.T) {
	actual := Sprintf("a}hey{b", make(map[string]string))
	expected := "a}hey{b"
	if actual != expected {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestFormatSomeSubsNotFound(t *testing.T) {
	dict := map[string]string{"name": "John"}
	actual := Sprintf("hey {name} are you {age} years old?", dict)
	expected := "hey John are you {age} years old?"
	if actual != expected {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestFormatNoEmbeddedSubs(t *testing.T) {
	template := "hey {name}, {are} you {age} years old?"
	dict := map[string]string{"name": "John", "age": "99"}
	expected := "hey John, {are} you 99 years old?"

	actual := Sprintf(template, dict)
	if actual != expected {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestFormatEmbeddedSubs(t *testing.T) {
	template := "{hey {name}, {are} you {age} years old?}"
	dict := map[string]string{"name": "John", "age": "99"}
	expected := "{hey John, {are} you 99 years old?}"

	actual := Sprintf(template, dict)
	if actual != expected {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestFormatExtraEmbeddedSubs(t *testing.T) {
	template := "{{hey {name}, {are} you {age} years old?}}"
	dict := map[string]string{"name": "John", "age": "99"}
	expected := "{{hey John, {are} you 99 years old?}}"

	actual := Sprintf(template, dict)
	if actual != expected {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestFormatExtraExtraEmbeddedSubs(t *testing.T) {
	template := "{{name}}"
	var dict map[string]string = nil
	expected := "{{name}}"

	actual := Sprintf(template, dict)
	if actual != expected {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestSprintfNumericalParameters(t *testing.T) {
	actual := Sprintf("-->{0}*{1}<--", 3.5, "silly", "ignored")
	expected := "-->3.5*silly<--"
	if actual != expected {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestSprintfMissingNumericalParameter(t *testing.T) {
	actual := Sprintf("-->{0}*{1}<--", 3.5)
	expected := "-->3.5*{1}<--"
	if actual != expected {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}
