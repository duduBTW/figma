package main

func minF(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

// Helper function to find the maximum of two float32 values
func maxF(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}
