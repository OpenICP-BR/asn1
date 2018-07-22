package asn1_test

import (
	"testing"

	"github.com/gjvnq/asn1"
)

type unexported struct {
	X int
	y int
	Z int
}

type unexported2 struct {
	X int
	Y int `asn1:"-"`
	Z int
}

type exported struct {
	X int
	Y int
	Z int
}

func TestUnexportedStructField(t *testing.T) {
	_, err := asn1.Marshal(unexported{X: 5, y: 1, Z: 3})
	if err != nil {
		t.Fatal(err)
	}

	b := unexported{X: 5, y: 1, Z: 3}
	bs, err := asn1.Marshal(b)
	if err != nil {
		t.Fatal(err)
	}
	var u unexported
	_, err = asn1.Unmarshal(bs, &u)
	if err != nil {
		t.Fatal(err)
	}
	if b.X != u.X {
		t.Errorf("got X=%d, want X=%d (DER: %x)", u.X, b.X, bs)
	}
	if 0 != u.y {
		t.Errorf("got y=%d, want y=%d (DER: %x)", u.y, 0, bs)
	}
	if b.Z != u.Z {
		t.Errorf("got Z=%d, want Z=%d (DER: %x)", u.Z, b.Z, bs)
	}
}

func TestUnexportedStructField2(t *testing.T) {
	_, err := asn1.Marshal(unexported2{X: 5, Y: 1, Z: 3})
	if err != nil {
		t.Fatal(err)
	}

	b := unexported2{X: 5, Y: 1, Z: 3}
	bs, err := asn1.Marshal(b)
	if err != nil {
		t.Fatal(err)
	}
	var u unexported2
	_, err = asn1.Unmarshal(bs, &u)
	if err != nil {
		t.Fatal(err)
	}
	if b.X != u.X {
		t.Errorf("got X=%d, want X=%d (DER: %x)", u.X, b.X, bs)
	}
	if 0 != u.Y {
		t.Errorf("got Y=%d, want Y=%d (DER: %x)", u.Y, 0, bs)
	}
	if b.Z != u.Z {
		t.Errorf("got Z=%d, want Z=%d (DER: %x)", u.Z, b.Z, bs)
	}
}
