package main

import "testing"

func TestUnpack1(t *testing.T) {
	var ps string
	var got, want string

	ps = "a4bc2d5e"
	want = "aaaabccddddde"
	got = Unpack([]rune(ps))

	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}

func TestUnpack2(t *testing.T) {
	var ps string
	var got, want string

	ps = "abcd"
	want = "abcd"
	got = Unpack([]rune(ps))

	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}

func TestUnpack3(t *testing.T) {
	var ps string
	var got, want string

	ps = "45"
	want = ""
	got = Unpack([]rune(ps))

	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}

func TestUnpack4(t *testing.T) {
	var ps string
	var got, want string

	ps = `qwe\4\5`
	want = `qwe45`
	got = Unpack([]rune(ps))

	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}

func TestUnpack5(t *testing.T) {
	var ps string
	var got, want string

	ps = `qwe\45`
	want = `qwe44444`
	got = Unpack([]rune(ps))

	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}

func TestUnpack6(t *testing.T) {
	var ps string
	var got, want string

	ps = `qwe\\5`
	want = `qwe\\\\\`
	got = Unpack([]rune(ps))

	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}

func TestUnpackRunes(t *testing.T) {
	var ps string
	var got, want string

	ps = `а11пр2т5н`
	want = `ааааааааааапрртттттн`
	got = Unpack([]rune(ps))

	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}
