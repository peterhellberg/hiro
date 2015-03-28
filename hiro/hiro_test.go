package hiro

import "testing"

func TestReadTemplate_WithBlankFn(t *testing.T) {
	tmpl, err := readTemplate("")
	if err != nil {
		t.Fatalf("unexpected error", err)
	}

	if string(tmpl) != DefaultTemplate {
		t.Fatalf("unexpected template")
	}
}
