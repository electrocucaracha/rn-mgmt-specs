package services

import (
	"fmt"
	"math"

	"rental-property-mgmt/internal/models"
)

// CalculationService handles financial metric calculations
type CalculationService struct{}

// NewCalculationService creates a new calculation service
func NewCalculationService() *CalculationService {
	return &CalculationService{}
}

// CalculateMetrics calculates all financial metrics for a property
func (cs *CalculationService) CalculateMetrics(property *models.Property) (*models.FinancialMetrics, error) {
	if !property.HasRequiredFieldsForMetrics() {
		return nil, fmt.Errorf("property missing required fields for metric calculations")
	}

	metrics := &models.FinancialMetrics{
		PropertyID: property.ID,
		IsCurrent:  true,
	}

	// Calculate monthly mortgage payment
	monthlyPayment, err := cs.calculateMonthlyMortgagePayment(property)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate monthly mortgage payment: %w", err)
	}
	metrics.MonthlyMortgagePayment = &monthlyPayment

	// Calculate Net Operating Income (NOI)
	noi, err := cs.calculateNOI(property)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate NOI: %w", err)
	}
	metrics.NetOperatingIncome = &noi

	// Calculate Cap Rate
	capRate := cs.calculateCapRate(noi, property.PurchasePrice)
	metrics.CapRate = &capRate

	// Calculate Cash to Close
	cashToClose := cs.calculateCashToClose(property)
	metrics.CashToClose = &cashToClose

	// Calculate Cash-on-Cash Return
	cocReturn, err := cs.calculateCashOnCashReturn(property, noi, monthlyPayment, cashToClose)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate cash-on-cash return: %w", err)
	}
	metrics.CashOnCashReturn = &cocReturn

	// Calculate Rent-to-Value Ratio
	rtv := cs.calculateRentToValueRatio(property)
	metrics.RentToValueRatio = &rtv

	// Calculate Gross Rent Multiplier
	grm := cs.calculateGrossRentMultiplier(property)
	metrics.GrossRentMultiplier = &grm

	return metrics, nil
}

// calculateMonthlyMortgagePayment calculates monthly mortgage payment
// Formula: P = L[c(1 + c)^n]/[(1 + c)^n - 1]
// Where: P = payment, L = loan amount, c = monthly interest rate, n = number of payments
func (cs *CalculationService) calculateMonthlyMortgagePayment(property *models.Property) (float64, error) {
	interestRate := property.GetFinancingTerm("interest_rate")
	loanTerm := property.GetFinancingTerm("loan_term")
	downPaymentPercent := property.GetFinancingTerm("down_payment_percent")

	if interestRate <= 0 || loanTerm <= 0 {
		return 0, fmt.Errorf("invalid financing terms: interest_rate=%f, loan_term=%f", interestRate, loanTerm)
	}

	// Calculate loan amount
	downPaymentAmount := property.PurchasePrice * (downPaymentPercent / 100)
	loanAmount := property.PurchasePrice - downPaymentAmount

	if loanAmount <= 0 {
		return 0, nil // No loan needed
	}

	// Convert annual rate to monthly and term to months
	monthlyRate := (interestRate / 100) / 12
	numberOfPayments := loanTerm * 12

	// Calculate monthly payment using standard amortization formula
	if monthlyRate == 0 {
		// Special case: 0% interest rate
		return loanAmount / numberOfPayments, nil
	}

	numerator := loanAmount * monthlyRate * math.Pow(1+monthlyRate, numberOfPayments)
	denominator := math.Pow(1+monthlyRate, numberOfPayments) - 1

	return numerator / denominator, nil
}

// calculateNOI calculates Net Operating Income
// NOI = (Monthly Rent × 12) - Annual Operating Expenses
func (cs *CalculationService) calculateNOI(property *models.Property) (float64, error) {
	if property.IntendedRent == nil || *property.IntendedRent <= 0 {
		return 0, fmt.Errorf("intended rent not set or invalid")
	}

	annualRent := *property.IntendedRent * 12

	// Calculate operating expenses
	insurance := property.GetOperatingExpense("insurance")
	propertyTaxes := property.GetOperatingExpense("property_taxes")
	hoa := property.GetOperatingExpense("hoa")
	utilities := property.GetOperatingExpense("utilities")

	// Calculate percentage-based expenses
	vacancyRate := property.GetOperatingAssumption("vacancy_rate")
	maintenancePct := property.GetOperatingAssumption("maintenance_pct")
	managementPct := property.GetOperatingAssumption("management_pct")

	vacancyLoss := annualRent * vacancyRate
	maintenanceCost := annualRent * maintenancePct
	managementCost := annualRent * managementPct

	totalOperatingExpenses := insurance + propertyTaxes + hoa + utilities + vacancyLoss + maintenanceCost + managementCost

	return annualRent - totalOperatingExpenses, nil
}

// calculateCapRate calculates Capitalization Rate
// Cap Rate = (NOI / Purchase Price) × 100
func (cs *CalculationService) calculateCapRate(noi, purchasePrice float64) float64 {
	if purchasePrice <= 0 {
		return 0
	}
	return (noi / purchasePrice) * 100
}

// calculateCashToClose calculates total cash needed to close
// Cash to Close = Down Payment + Closing Costs
func (cs *CalculationService) calculateCashToClose(property *models.Property) float64 {
	downPaymentPercent := property.GetFinancingTerm("down_payment_percent")
	closingCosts := property.GetFinancingTerm("closing_costs")

	downPaymentAmount := property.PurchasePrice * (downPaymentPercent / 100)

	return downPaymentAmount + closingCosts
}

// calculateCashOnCashReturn calculates Cash-on-Cash Return
// Annual Cash Flow = NOI - Annual Debt Service
// Cash-on-Cash Return = (Annual Cash Flow / Initial Cash Investment) × 100
func (cs *CalculationService) calculateCashOnCashReturn(property *models.Property, noi, monthlyPayment, cashToClose float64) (float64, error) {
	if cashToClose <= 0 {
		return 0, fmt.Errorf("cash to close must be greater than 0")
	}

	annualDebtService := monthlyPayment * 12
	annualCashFlow := noi - annualDebtService

	return (annualCashFlow / cashToClose) * 100, nil
}

// calculateRentToValueRatio calculates Rent-to-Value Ratio
// RTV = (Monthly Rent × 12 / Purchase Price) × 100
func (cs *CalculationService) calculateRentToValueRatio(property *models.Property) float64 {
	if property.IntendedRent == nil || *property.IntendedRent <= 0 || property.PurchasePrice <= 0 {
		return 0
	}

	annualRent := *property.IntendedRent * 12
	return (annualRent / property.PurchasePrice) * 100
}

// calculateGrossRentMultiplier calculates Gross Rent Multiplier
// GRM = Purchase Price / (Monthly Rent × 12)
func (cs *CalculationService) calculateGrossRentMultiplier(property *models.Property) float64 {
	if property.IntendedRent == nil || *property.IntendedRent <= 0 {
		return 0
	}

	annualRent := *property.IntendedRent * 12
	if annualRent <= 0 {
		return 0
	}

	return property.PurchasePrice / annualRent
}

// RecalculateIfNeeded checks if metrics need recalculation and does so if needed
func (cs *CalculationService) RecalculateIfNeeded(property *models.Property, currentMetrics *models.FinancialMetrics) (*models.FinancialMetrics, bool, error) {
	// Always recalculate if metrics don't exist or are marked as not current
	if currentMetrics == nil || !currentMetrics.IsCurrent {
		newMetrics, err := cs.CalculateMetrics(property)
		if err != nil {
			return nil, false, err
		}
		return newMetrics, true, nil
	}

	// For now, we'll assume metrics are current if they exist and are marked as such
	// In a more sophisticated system, we might check timestamps or property modification dates
	return currentMetrics, false, nil
}
