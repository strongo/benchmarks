package gocontrib

import (
	"testing"
	"strconv"
)

func TestAtoiImproved(t *testing.T) {
	v, err := AtoiImproved("+12")
	t.Logf("v=%v, err=%v", v, err)

	v, err = AtoiImproved("-12")
	t.Logf("v=%v, err=%v", v, err)
}

func BenchmarkAtoi(b *testing.B) {

	// unsigned := [10]string{
	// 	"0",
	// 	"1",
	// 	"12",
	// 	"123",
	// 	"1234",
	// 	"12345",
	// 	"123456",
	// 	"1234567",
	// 	"123456789",
	// }

	// pos := [10]string{
	// 	"+0",
	// 	"+1",
	// 	"+12",
	// 	"+123",
	// 	"+1234",
	// 	"+12345",
	// 	"+123456",
	// 	"+1234567",
	// 	"+123456789",
	// }
	//
	// neg := [10]string{
	// 	"-0",
	// 	"-1",
	// 	"-12",
	// 	"-123",
	// 	"-1234",
	// 	"-12345",
	// 	"-123456",
	// 	"-1234567",
	// 	"-123456789",
	// }

	// wrongSigned := [10]string{
	// 	"-",
	// 	"-a",
	// 	"-1b",
	// 	"-12c",
	// 	"-123d",
	// 	"-1234e",
	// 	"-12345f",
	// 	"-123456g",
	// 	"-1234567h",
	// 	"-12345678i",
	// }

	b.Run("original_unsigned", func(b *testing.B) {
		const v = "12"
		for i := 0; i < b.N; i++ {
			i, _ := strconv.Atoi(v)
			BenchSink += i
		}
	})

	b.Run("improved_unsigned", func(b *testing.B) {
		const v = "12"
		for i := 0; i < b.N; i++ {
			i, _ := AtoiImproved(v)
			BenchSink += i
		}
	})

	b.Run("original_positive", func(b *testing.B) {
		const v = "+12"
		for i := 0; i < b.N; i++ {
			i, _ := strconv.Atoi(v)
			BenchSink += i
		}
	})
	b.Run("improved_positive", func(b *testing.B) {
		const v = "+12"
		for i := 0; i < b.N; i++ {
			i, _ := AtoiImproved(v)
			BenchSink += i
		}
	})

	b.Run("original_negative", func(b *testing.B) {
		const v = "-12"
		for i := 0; i < b.N; i++ {
			i, _ := strconv.Atoi(v)
			BenchSink += i
		}
	})
	b.Run("improved_negative", func(b *testing.B) {
		const v = "-12"
		for i := 0; i < b.N; i++ {
			i, _ := AtoiImproved(v)
			BenchSink += i
		}
	})
	// b.Run("original_wrong_signed", func(b *testing.B) {
	// 	for i := 0; i < b.N; i++ {
	// 		for _, v := range wrongSigned {
	// 			i, err := strconv.Atoi(v)
	// 			BenchSink += i
	// 			lastErr = err
	// 		}
	// 	}
	// })
	// b.Run("improved_unsigned", func(b *testing.B) {
	// 	for i := 0; i < b.N; i++ {
	// 		for _, v := range unsigned {
	// 			i, _ := AtoiImproved(v)
	// 			BenchSink += i
	// 		}
	// 	}
	// })


	b.Run("original_minus_only", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := strconv.Atoi("-")
			lastErr = err
		}
	})
	b.Run("improved_minus_only", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := AtoiImproved("-")
			lastErr = err
		}
	})
	// b.Run("improve_wrong_signed", func(b *testing.B) {
	// 	for i := 0; i < b.N; i++ {
	// 		for _, v := range wrongSigned {
	// 			i, err := AtoiImproved(v)
	// 			BenchSink += i
	// 			lastErr = err
	// 		}
	// 	}
	// })
	b.Log(BenchSink)
}

var BenchSink int
var lastErr error
