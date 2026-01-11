package dsp

import (
	"errors"

	algofft "github.com/MeKo-Christian/algo-fft"
)

type complexFFTPlan interface {
	Forward(dst, src []complex64) error
	Inverse(dst, src []complex64) error
}

type fastComplexFFTPlan struct {
	plan *algofft.FastPlan[complex64]
}

func (p *fastComplexFFTPlan) Forward(dst, src []complex64) error {
	p.plan.Forward(dst, src)
	return nil
}

func (p *fastComplexFFTPlan) Inverse(dst, src []complex64) error {
	p.plan.Inverse(dst, src)
	return nil
}

type safeComplexFFTPlan struct {
	plan *algofft.Plan[complex64]
}

func (p *safeComplexFFTPlan) Forward(dst, src []complex64) error {
	return p.plan.Forward(dst, src)
}

func (p *safeComplexFFTPlan) Inverse(dst, src []complex64) error {
	return p.plan.Inverse(dst, src)
}

func newComplexFFTPlan(size int) (complexFFTPlan, error) {
	fastPlan, err := algofft.NewFastPlan[complex64](size)
	if err == nil {
		return &fastComplexFFTPlan{plan: fastPlan}, nil
	}

	if !errors.Is(err, algofft.ErrNotImplemented) {
		return nil, err
	}

	plan, err := algofft.NewPlan32(size)
	if err != nil {
		return nil, err
	}

	return &safeComplexFFTPlan{plan: plan}, nil
}

type realFFTPlan32 interface {
	Forward(dst []complex64, src []float32) error
	Inverse(dst []float32, src []complex64) error
}

type fastRealFFTPlan32 struct {
	plan *algofft.FastPlanReal32
}

func (p *fastRealFFTPlan32) Forward(dst []complex64, src []float32) error {
	p.plan.Forward(dst, src)
	return nil
}

func (p *fastRealFFTPlan32) Inverse(dst []float32, src []complex64) error {
	p.plan.Inverse(dst, src)
	return nil
}

type safeRealFFTPlan32 struct {
	plan *algofft.PlanRealT[float32, complex64]
}

func (p *safeRealFFTPlan32) Forward(dst []complex64, src []float32) error {
	return p.plan.Forward(dst, src)
}

func (p *safeRealFFTPlan32) Inverse(dst []float32, src []complex64) error {
	return p.plan.Inverse(dst, src)
}

func newRealFFTPlan32(size int) (realFFTPlan32, error) {
	fastPlan, err := algofft.NewFastPlanReal32(size)
	if err == nil {
		return &fastRealFFTPlan32{plan: fastPlan}, nil
	}

	if !errors.Is(err, algofft.ErrNotImplemented) {
		return nil, err
	}

	plan, err := algofft.NewPlanReal32(size)
	if err != nil {
		return nil, err
	}

	return &safeRealFFTPlan32{plan: plan}, nil
}
