package main

import "testing"

func TestPriceInVNDConvertsDemoUSDPrice(t *testing.T) {
	got := priceInVND(1499)
	const want int64 = 38_970_000
	if got != want {
		t.Fatalf("priceInVND(1499) = %d, want %d", got, want)
	}
}

func TestSelectedROGUsesAuthoritativeDemoPrice(t *testing.T) {
	if got := normalizedSeedPrice("ASUS-0001", 1499); got != 42_000_000 {
		t.Fatalf("normalizedSeedPrice() = %d, want 42000000", got)
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
