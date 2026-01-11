package dsp

import (
	"errors"
	"fmt"
	"math"
	"testing"

	algofft "github.com/MeKo-Christian/algo-fft"
)

// BenchmarkFastPlanVsPlan compares FastPlan to the safe Plan API for complex FFT.
func BenchmarkFastPlanVsPlan(b *testing.B) {
	fftSizes := []int{64, 128, 256, 384, 512}

	for _, size := range fftSizes {
		b.Run(fmt.Sprintf("FastPlan_%d", size), func(b *testing.B) {
			plan, err := algofft.NewFastPlan[complex64](size)
			if err != nil {
				if errors.Is(err, algofft.ErrNotImplemented) {
					b.Skipf("FastPlan not available for size %d", size)
				}
				b.Fatalf("failed to create FastPlan: %v", err)
			}

			input := make([]complex64, size)
			output := make([]complex64, size)
			for i := range size {
				input[i] = complex(float32(math.Sin(float64(i)*0.1)), 0)
			}

			b.SetBytes(int64(size * 8))
			b.ResetTimer()

			for range b.N {
				plan.Forward(output, input)
			}
		})

		b.Run(fmt.Sprintf("Plan_%d", size), func(b *testing.B) {
			plan, err := algofft.NewPlan32(size)
			if err != nil {
				b.Fatalf("failed to create Plan: %v", err)
			}

			input := make([]complex64, size)
			output := make([]complex64, size)
			for i := range size {
				input[i] = complex(float32(math.Sin(float64(i)*0.1)), 0)
			}

			b.SetBytes(int64(size * 8))
			b.ResetTimer()

			for range b.N {
				if err := plan.Forward(output, input); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// BenchmarkFastPlanRealVsPlan compares FastPlanReal32 to the safe PlanReal32 API.
func BenchmarkFastPlanRealVsPlan(b *testing.B) {
	fftSizes := []int{64, 128, 256, 384, 512}

	for _, size := range fftSizes {
		b.Run(fmt.Sprintf("FastPlanReal32_%d", size), func(b *testing.B) {
			plan, err := algofft.NewFastPlanReal32(size)
			if err != nil {
				if errors.Is(err, algofft.ErrNotImplemented) {
					b.Skipf("FastPlanReal32 not available for size %d", size)
				}
				b.Fatalf("failed to create FastPlanReal32: %v", err)
			}

			input := make([]float32, size)
			output := make([]complex64, size/2+1)
			for i := range size {
				input[i] = float32(math.Sin(float64(i) * 0.1))
			}

			b.SetBytes(int64(size * 4))
			b.ResetTimer()

			for range b.N {
				plan.Forward(output, input)
			}
		})

		b.Run(fmt.Sprintf("PlanReal32_%d", size), func(b *testing.B) {
			plan, err := algofft.NewPlanReal32(size)
			if err != nil {
				b.Fatalf("failed to create PlanReal32: %v", err)
			}

			input := make([]float32, size)
			output := make([]complex64, size/2+1)
			for i := range size {
				input[i] = float32(math.Sin(float64(i) * 0.1))
			}

			b.SetBytes(int64(size * 4))
			b.ResetTimer()

			for range b.N {
				if err := plan.Forward(output, input); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
