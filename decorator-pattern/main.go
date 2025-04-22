package main

import (
	"fmt"

	"github.com/GodwinJacobR/design-patterns/decorator-pattern/pkg"
)

type PlanType string

const (
	BasicPlan      PlanType = "basic"
	PremiumPlan    PlanType = "premium"
	EnterprisePlan PlanType = "enterprise"
)

type PlanDecorator struct {
	feeProvider pkg.FeeProvider
}

type PremiumPlanDecorator struct {
	PlanDecorator
}

func (p *PremiumPlanDecorator) ProcessingFee() float64 {
	return p.feeProvider.ProcessingFee() * 0.8
}

func (p *PremiumPlanDecorator) TransferFee() float64 {
	return p.feeProvider.TransferFee() * 0.9
}

type EnterprisePlanDecorator struct {
	PlanDecorator
}

func (e *EnterprisePlanDecorator) ProcessingFee() float64 {
	return 25
}

func (e *EnterprisePlanDecorator) TransferFee() float64 {
	return e.feeProvider.TransferFee() * 0.75
}

func CreateFeeProvider(amount float64, plan PlanType) pkg.FeeProvider {
	transfer := &pkg.Transfer{
		Amount: amount,
	}

	// Apply decorator based on plan
	switch plan {
	case PremiumPlan:
		return &PremiumPlanDecorator{
			PlanDecorator: PlanDecorator{feeProvider: transfer},
		}
	case EnterprisePlan:
		return &EnterprisePlanDecorator{
			PlanDecorator: PlanDecorator{feeProvider: transfer},
		}
	default: // Basic plan uses default implementation
		return transfer
	}
}

func main() {
	transferAmount := 1000.0

	basicProvider := CreateFeeProvider(transferAmount, BasicPlan)
	premiumProvider := CreateFeeProvider(transferAmount, PremiumPlan)
	enterpriseProvider := CreateFeeProvider(transferAmount, EnterprisePlan)

	fmt.Printf("Basic Plan:\n")
	fmt.Printf("  Processing Fee: $%.2f\n", basicProvider.ProcessingFee())
	fmt.Printf("  Transfer Fee: $%.2f\n\n", basicProvider.TransferFee())

	fmt.Printf("Premium Plan:\n")
	fmt.Printf("  Processing Fee: $%.2f\n", premiumProvider.ProcessingFee())
	fmt.Printf("  Transfer Fee: $%.2f\n\n", premiumProvider.TransferFee())

	fmt.Printf("Enterprise Plan:\n")
	fmt.Printf("  Processing Fee: $%.2f\n", enterpriseProvider.ProcessingFee())
	fmt.Printf("  Transfer Fee: $%.2f\n", enterpriseProvider.TransferFee())
}
