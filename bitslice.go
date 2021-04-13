package bitslice

import (
	"strconv"
)

// BitSlice is a bit array
type BitSlice struct {
	// buf is the underlying byte slice storing the bits
	buf []byte
	// len is the number of bits stored in buf
	len int64
}

// New creates a BitSlice with l zeroed bits with the capacity to store c bits,
// rounded up to the nearest octet. Panics if l > c, l < 0, or c < 0.
func New(l int64, c int64) BitSlice {
	return BitSlice{
		// Round l and c up to nearest octet
		buf: make([]byte, (l+7)/8, (c+7)/8),
		len: l,
	}
}

// Len returns the number of bits stored in the BitSlice
func (b BitSlice) Len() int64 {
	return b.len
}

// Cap returns the capacity of the BitSlice in bits
func (b BitSlice) Cap() int64 {
	return int64(cap(b.buf)) * 8
}

// Get returns the bit at position i. If i is out of range, it panics.
func (b BitSlice) Get(i int64) bool {
	if i >= b.len {
		panic("index of out range [" + strconv.FormatInt(i, 10) +
			"] with length " + strconv.FormatInt(b.len, 10))
	}
	var mask byte = 1 << (i % 8)
	return (b.buf[i/8] & mask) > 0
}

// Set sets the bit at position i to v. If i is out of range, it panics.
func (b BitSlice) Set(i int64, v bool) {
	if i >= b.len {
		panic("index of out range [" + strconv.FormatInt(i, 10) +
			"] with length " + strconv.FormatInt(b.len, 10))
	}
	var mask byte = 1 << (i % 8)
	if v {
		b.buf[i/8] |= mask
	} else {
		b.buf[i/8] &= ^mask
	}
}

// Append adds bits to the end of the BitSlice much like the built-in append for
// slices. If the BitSlice has sufficient capacity, the underlying buffer is
// resliced to accommodate the new bits. If it does not, a new underlying array
// will be allocated. Append returns the updated BitSlice.
func (b BitSlice) Append(bits ...bool) BitSlice {
	bitsI := 0    // Index of next bit to be added
	var mask byte // Use `b.buf[bufI] |= mask` to set current bit

	bufI := int(b.len / 8) // Index of next byte in b to be written
	mask = 1 << (b.len % 8)
	if mask > 1 {
		// Fill last byte in b with bits
		for ; bitsI < len(bits) && mask > 0; bitsI++ {
			if bits[bitsI] {
				// Set bit
				b.buf[bufI] |= mask
			}
			mask <<= 1
		}
	}

	// Iterate through each bit, appending bytes to b.buf until done
	var curByte byte
	if bitsI < len(bits) {
		for mask = 1; bitsI < len(bits); bitsI++ {
			if mask == 0 {
				// Append byte and reset
				b.buf = append(b.buf, curByte)
				curByte = 0
				mask = 1
			}
			if bits[bitsI] {
				curByte |= mask
			}
			mask <<= 1
		}
		// Append partial byte
		b.buf = append(b.buf, curByte)
	}

	// Update BitSlice length and return
	b.len += int64(len(bits))
	return b
}

// AppendBytes pads the last byte of the BitSlice with zeroes and appends bytes
// to the end of the BitSlice
func (b BitSlice) AppendBytes(bytes ...byte) BitSlice {
	b.buf = append(b.buf, bytes...)
	b.len = int64(len(b.buf)) * 8
	return b
}

// Bytes returns the underlying byte slice storing the bits
func (b BitSlice) Bytes() []byte {
	return b.buf
}
