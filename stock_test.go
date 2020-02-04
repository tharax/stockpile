package stockpile

import (
	"testing"
)

func Test_NewStockpile(t *testing.T) {
	s := NewStockpile()
	if s == nil {
		t.Error("expected s not to be nil")
	}
}

func Test_Stockpile_String(t *testing.T) {
	s := NewStockpile()
	str := s.String()
	if len(str) != 0 {
		t.Error("string of empty stockpile should be zero length")
	}

	s = append(s, Stock{"Example", 1})
	got := s.String()
	expected := `Example                            1
`
	if got != expected {
		t.Errorf("got: %s, expected: %s\n", got, expected)
	}
	s = append(s, Stock{"Example_Two", 3})
	got = s.String()
	expected = `Example                            1
Example_Two                        3
`
	if got != expected {
		t.Errorf("got: %s\n expected: %s\n", got, expected)
	}
}

func Test_Stockpile_Add(t *testing.T) {
	s := NewStockpile()
	err := s.Add(Stock{"Example", 1})

	got := s.String()
	expected := `Example                            1
`
	if err != nil {
		t.Errorf("got error: %v, expected error: %v", err, nil)
	}
	if got != expected {
		t.Errorf("got: %s, expected: %s\n", got, expected)
	}
	err = s.Add(Stock{"Example", 1})
	got = s.String()
	expected = `Example                            2
`
	if err != nil {
		t.Errorf("got error: %v, expected error: %v", err, nil)
	}
	if got != expected {
		t.Errorf("got: %s, expected: %s\n", got, expected)
	}

	err = s.Add(Stock{"Example", 0})
	got = s.String()
	expected = `Example                            2
`
	if err == nil {
		t.Errorf("got error: %v, expected error: %v", nil, err)
	}
	if got != expected {
		t.Errorf("got: %s, expected: %s\n", got, expected)
	}
	err = s.Add(Stock{"Example", -1})
	got = s.String()
	expected = `Example                            2
`
	expectedErrorString := "trying to add a stock with 0 or less quantity"
	if err == nil {
		t.Errorf("got error: %v, expected error: %v", nil, err)
	}
	if got != expected {
		t.Errorf("got: %s, expected: %s\n", got, expected)
	}
	err = s.Add(Stock{"", 1})
	got = s.String()
	expected = `Example                            2
`
	expectedErrorString = "trying to add a stock with no name"
	if err == nil {
		t.Errorf("got error: %v, expected error: %v", nil, err)
	}
	if err.Error() != expectedErrorString {
		t.Errorf("got error: %v, expected error: %v", expectedErrorString, err)
	}
	if got != expected {
		t.Errorf("got: %s, expected: %s\n", got, expected)
	}
}
