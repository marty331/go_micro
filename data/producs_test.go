package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Marty",
		Price: 10.5,
		SKU:   "abc-acb-kdj",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
