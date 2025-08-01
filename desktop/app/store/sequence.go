package store

import "slices"

type Sequence []string

func (seq *Sequence) Add(s string) {
	if !slices.Contains(*seq, s) {
		*seq = append(*seq, s)
	}
}

func (seq *Sequence) Remove(s string) {
	idx := slices.Index(*seq, s)
	if idx != -1 {
		*seq = append((*seq)[:idx], (*seq)[idx+1:]...)
	}
}

func (seq *Sequence) Move(s string, pos int) {
	if pos < 0 || pos == len(*seq) {
		return
	}
	idx := slices.Index(*seq, s)
	(*seq)[idx], (*seq)[pos] = (*seq)[pos], (*seq)[idx]
}

func sort[T any](seq Sequence, entries []Entry[T]) {
	slices.SortFunc(entries, func(a, b Entry[T]) int {
		idx1 := slices.Index(seq, a.Key)
		idx2 := slices.Index(seq, b.Key)
		return idx1 - idx2
	})
}
