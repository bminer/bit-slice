package bitslice

import (
	"testing"
)

func TestNew(t *testing.T) {
	b := New(0, 0)
	if b.Len() != 0 {
		t.Errorf("length %d, expected 0", b.Len())
	}
	if b.Cap() != 0 {
		t.Errorf("cap %d, expected 0", b.Cap())
	}

	b = New(5, 5)
	if b.Len() != 5 {
		t.Errorf("length %d, expected 5", b.Len())
	}
	if b.Cap() != 8 {
		t.Errorf("cap %d, expected 8", b.Cap())
	}
}

func TestAppend(t *testing.T) {
	b := New(0, 0)
	b = b.Append(true, true, true, true, false, false, false, true)
	b = b.Append(true, true)
	b = b.Append(false, true)
	b = b.Append(true, true)
	b = b.Append(false)
	b = b.Append(false)
	b = b.Append(true)
	b = b.Append(true, true, false)
	b = b.Append(true, true, false, false, false, false, true, true)
	b = b.Append(false)
	b = b.Append(false)
	b = b.Append(false)
	b = b.Append(false)
	b = b.Append(false)
	b = b.Append(false)
	b = b.Append(false)
	b = b.Append(true)
	b = b.Append(true, true, true, true, false, false, false, true)
	b = b.Append(true, true, true, false)
	b = b.Append(true, true, true, true, false, false, false, true, false)
	if b.Len() != 57 {
		t.Errorf("length %d, expected 57", b.Len())
	}
	vals := []bool{
		true, true, true, true, false, false, false, true,
		true, true,
		false, true,
		true, true,
		false,
		false,
		true,
		true, true, false,
		true, true, false, false, false, false, true, true,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		true,
		true, true, true, true, false, false, false, true,
		true, true, true, false,
		true, true, true, true, false, false, false, true, false,
	}
	for i := int64(0); i < int64(len(vals)); i++ {
		if b.Get(i) != vals[i] {
			t.Errorf("bit %d is %v, expected %v", i, b.Get(i), vals[i])
		}
	}
}
