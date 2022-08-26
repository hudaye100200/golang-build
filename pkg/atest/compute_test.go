package test

import "testing"

func TestAdd(t *testing.T) {
	a := 10
	b := 20
	want := 30
	actual := Add(a, b)
	if want != actual {
		t.Errorf("Add函数参数:%d %d, 期望: %d, 实际: %d", a, b, want, actual)
	}
}

func TestMul(t *testing.T) {
	a := 10
	b := 20
	want := 300
	actual := Mul(a, b)
	if want != actual {
		t.Errorf("Mul函数参数:%d %d, 期望: %d, 实际: %d", a, b, want, actual)
	}
}

func TestDiv(t *testing.T) {
	a := 10
	b := 20
	want := 2
	actual := Div(a, b)
	if want != actual {
		t.Errorf("Div函数参数:%d %d, 期望: %d, 实际: %d", a, b, want, actual)
	}
}

func TestSkip(t *testing.T) {
	a := 10
	b := 20
	want := 30
	got := Add(a, b)
	if want == got {
	    t.Skip("Skipping testing")	
	}
}