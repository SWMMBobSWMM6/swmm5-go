// Generated from EPA SWMM 5.2.4 (hash.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
)

// hash.c:24
func samestr(s1 R.CPtr, s2 R.CPtr) int32 {
	var (
		i int32
	)
	_ = i
	i = 0
	for i = 0; (func() int32 {
		if ((s1.G(i)) >= 97) && ((s1.G(i)) <= 122) {
			return ((s1.G(i)) & (^32))
		}
		return (s1.G(i))
	}()) == (func() int32 {
		if ((s2.G(i)) >= 97) && ((s2.G(i)) <= 122) {
			return ((s2.G(i)) & (^32))
		}
		return (s2.G(i))
	}()); i++ {
		if (!(s1.G((i + 1)) != 0)) && (!(s2.G((i + 1)) != 0)) {
			return (1)
		}
	}
	return (0)
}

// hash.c:33
func hash(str R.CPtr) uint32 {
	var (
		sum1   uint32
		check1 uint32
		sum2   uint64
	)
	_, _, _ = sum1, check1, sum2
	sum1 = 0
	check1 = 0
	sum2 = 0
	for 0 != str.G(0) {
		sum1 += (uint32((func() int32 {
			if ((str.G(0)) >= 97) && ((str.G(0)) <= 122) {
				return ((str.G(0)) & (^32))
			}
			return (str.G(0))
		}())))
		str = str.Add(1)
		if 255 <= sum1 {
			sum1 -= 255
		}
		sum2 += (uint64(sum1))
	}
	check1 = uint32(sum2)
	check1 %= 255
	check1 = (255 - ((sum1 + check1) % 255))
	sum1 = (255 - ((sum1 + check1) % 255))
	return (((check1 << 8) | sum1) % 1999)
}

// hash.c:51
func HTcreate() []*HTentry {
	var (
		i  int32
		ht []*HTentry
	)
	_, _ = i, ht
	i = 0
	ht = make([]*HTentry, 1999)
	if ht != nil {
		for i = 0; i < 1999; i++ {
			ht[i] = nil
		}
	}
	return (ht)
}

// hash.c:59
func HTinsert(ht []*HTentry, key R.CPtr, data int32) int32 {
	var (
		i     uint32
		entry *HTentry
	)
	_, _ = i, entry
	i = hash(key)
	entry = nil
	if i >= 1999 {
		return (0)
	}
	entry = new(HTentry)
	if entry == nil {
		return (0)
	}
	entry.key = key
	entry.data = data
	entry.next = ht[i]
	ht[i] = entry
	return (1)
}

// hash.c:73
func HTfind(ht []*HTentry, key R.CPtr) int32 {
	var (
		i     uint32
		entry *HTentry
	)
	_, _ = i, entry
	i = hash(key)
	entry = nil
	if i >= 1999 {
		return (-1)
	}
	entry = ht[i]
	for entry != nil {
		if samestr(entry.key, key) != 0 {
			return (entry.data)
		}
		entry = entry.next
	}
	return (-1)
}

// hash.c:87
func HTfindKey(ht []*HTentry, key R.CPtr) R.CPtr {
	var (
		i     uint32
		entry *HTentry
	)
	_, _ = i, entry
	i = hash(key)
	entry = nil
	if i >= 1999 {
		return (R.CPtr{})
	}
	entry = ht[i]
	for entry != nil {
		if samestr(entry.key, key) != 0 {
			return (entry.key)
		}
		entry = entry.next
	}
	return (R.CPtr{})
}

// hash.c:101
func HTfree(ht []*HTentry) {
	var (
		entry     *HTentry
		nextentry *HTentry
		i         int32
	)
	_, _, _ = entry, nextentry, i
	entry = nil
	nextentry = nil
	i = 0
	for i = 0; i < 1999; i++ {
		entry = ht[i]
		for entry != nil {
			nextentry = entry.next
			_ = entry
			entry = nextentry
		}
	}
	_ = ht
}
