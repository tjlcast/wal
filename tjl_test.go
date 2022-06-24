package wal

import "testing"



func Test_demo(t *testing.T) {
	logPath:="1655903359.dat"

	opts := DefaultOptions
	opts.SegmentSize = int(1<<20)
	opts.SegmentCacheSize = 1
	l, err := Open(logPath, opts)
	if err != nil {
		t.Fatalf("expected %v, got %v", ErrNotFound, err)
	}

	oldNum := 0
	newNum := 0

	// Read -- read back all entries
	//lastIndex, _ := l.LastIndex()
	//markTime1 := MarkTime()
	//t.Logf("First: %d last: %d segments: %d.", int(l.firstIndex), int(lastIndex), len(l.segments))
	//for i := int(l.firstIndex); i<=int(lastIndex); i++ {
	//	//t.Logf("Read idx: %d.", i)
	//	_, err := l.Read(uint64(i))
	//	oldNum++
	//	if err != nil {
	//		t.Fatalf("error while getting %d", i)
	//	}
	//}
	//t.Logf("Read cost: %d.", markTime1.Gap())

	it := l.Iterator()
	for !l.ItEmpry(it) {
		_, err = l.ItNext(it)
		newNum++
		if err != nil {
			t.Fatalf("error while getting.")
		}
	}

	t.Logf("old %d new %d.\n", oldNum, newNum)

	err = l.TruncateBack(500000)
	if err != nil {
		panic(err)
	}

	newNum1 := 0
	it = l.Iterator()
	for !l.ItEmpry(it) {
		_, err = l.ItNext(it)
		newNum1++
		if err != nil {
			t.Fatalf("error while getting.")
		}
	}

	t.Logf("old %d new %d.\n", oldNum, newNum1)
}

