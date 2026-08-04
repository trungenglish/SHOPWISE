package main

import "testing"

func TestPriceInVNDConvertsDemoUSDPrice(t *testing.T) {
	got := priceInVND(1499)
	const want int64 = 37_475_000
	if got != want {
		t.Fatalf("priceInVND(1499) = %d, want %d", got, want)
	}
}

func TestSeedConflictUpdatesExistingPrice(t *testing.T) {
	conflict := seedConflictClause()
	for _, assignment := range conflict.DoUpdates {
		if assignment.Column.Name == "price" {
			return
		}
	}
	t.Fatal("seed conflict clause must update price")
}
