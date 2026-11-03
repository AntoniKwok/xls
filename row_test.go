package xls

import "testing"

// A MULBLANK record whose last column precedes its first holds no cells.
// Row.Col used to index into the empty result and panic.
func TestRowColEmptyRun(t *testing.T) {
	empty := &MulBlankCol{Col: Col{RowB: 4, FirstColB: 34}, LastColB: 33}
	row := &Row{cols: map[uint16]contentHandler{34: empty}}
	if got := row.Col(34); got != "" {
		t.Fatalf("empty run cell = %q, want blank", got)
	}

	// a run that does cover the column must still resolve through the scan path
	run := &MulBlankCol{Col: Col{RowB: 4, FirstColB: 10}, Xfs: []uint16{1, 2, 3}, LastColB: 12}
	row = &Row{cols: map[uint16]contentHandler{10: run}}
	if got := row.Col(11); got != "" {
		t.Fatalf("blank run cell = %q, want blank", got)
	}
	// a run whose declared range exceeds its cells must not panic either
	short := &MulBlankCol{Col: Col{RowB: 4, FirstColB: 20}, Xfs: []uint16{1}, LastColB: 25}
	row = &Row{cols: map[uint16]contentHandler{20: short}}
	if got := row.Col(24); got != "" {
		t.Fatalf("short run cell = %q, want blank", got)
	}
}

// Rows without any record are normal in a sparse sheet; Row must report them
// as nil so callers can skip them, not dereference a missing map entry.
func TestWorkSheetRowSparse(t *testing.T) {
	ws := &WorkSheet{wb: &WorkBook{}, rows: map[uint16]*Row{
		0: {cols: map[uint16]contentHandler{}},
		3: {cols: map[uint16]contentHandler{}},
	}, MaxRow: 3}
	for i := 0; i <= 3; i++ {
		got := ws.Row(i)
		present := i == 0 || i == 3
		if present && got == nil {
			t.Fatalf("row %d is populated but Row returned nil", i)
		}
		if !present && got != nil {
			t.Fatalf("row %d has no record but Row returned %v", i, got)
		}
	}
	if ws.Row(60000) != nil {
		t.Fatal("row far past MaxRow should be nil")
	}
}
