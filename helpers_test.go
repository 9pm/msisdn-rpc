package main

import "testing"

func TestGetAlpha(t *testing.T) {

	india := GetAlpha("91")
	if india != "IN" {
		t.Errorf("Country incorrect, got: %s, want: %s.", india, "IN")
	}

	russia := GetAlpha("7")
	if russia != "RU" {
		t.Errorf("Country incorrect, got: %s, want: %s.", russia, "RU")
	}

	as := GetAlpha("1684")
	if as != "AS" {
		t.Errorf("Country incorrect, got: %s, want: %s.", as, "AS")
	}
}
