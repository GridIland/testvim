package main

func getMonthlyPrice(tier string) int {
	dollarinpenny := 100
	if tier == "basic" {
		return 100 * dollarinpenny
	}
	if tier == "premium" {
		return 150 * dollarinpenny
	}
	if tier == "enterprise" {
		return 500 * dollarinpenny
	}
	return 0
}

