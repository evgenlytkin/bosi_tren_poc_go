package main

import "testing"

// func BenchmarkOneCallWithOneBigResponse(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		one_call_with_one_big_response()
// 	}
// }

// func BenchmarkOneCallWithOneStubbedBigResponse(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		one_call_with_one_stubbed_big_response()
// 	}
// }

func BenchmarkMultipleCallsWithResponseInParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		multiple_calls_with_response_in_parallel()
	}
}

// func BenchmarkMultipleCallsWithMultipleSmallStubbedResponses(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		multiple_calls_with_multiple_small_stubbed_responses()
// 	}
// }
