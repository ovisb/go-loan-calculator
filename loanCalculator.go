package loanCalculator

import (
	"flag"
	"fmt"
	"math"
	"os"
)

const (
	annuityOpt = "annuity"
	diffOpt    = "diff"
)

func calculateMonthlyInterest(interest float64) float64 {
	return (interest / 100) / 12
}

func calculateNumberOfMonths(interestRate, payment, principal float64) float64 {
	monthlyInterestRate := calculateMonthlyInterest(interestRate)
	return math.Ceil(math.Log(payment/(payment-monthlyInterestRate*principal)) / math.Log(1+monthlyInterestRate))
}

func calculateMonthlyPayment(principal, interestRate, periods float64) float64 {
	i := calculateMonthlyInterest(interestRate)
	if i == 0 {
		return principal / periods
	}
	return principal * (i * math.Pow(1+i, periods)) / (math.Pow(1+i, periods) - 1)
}

func calculateLoanPrincipal(payment, annualInterestRate, periods float64) float64 {
	monthlyInterestRate := calculateMonthlyInterest(annualInterestRate)
	if monthlyInterestRate == 0 {
		return payment * periods
	}
	return payment * ((math.Pow(1+monthlyInterestRate, periods) - 1) / (monthlyInterestRate * math.Pow(1+monthlyInterestRate, periods)))
}

func printNumberOfMonths(numberOfMonths float64) {
	years := numberOfMonths / 12
	months := int(numberOfMonths) % 12

	if years == 1 {
		fmt.Println("It will take 1 year to repay the loan!")
		return
	}

	if years > 1 {
		monthsSuff := "month"
		if months > 1 {
			monthsSuff = "months"
		}
		if months == 0 {
			fmt.Printf("It will take %d years to repay the loan!\n", int(years))
			return
		}
		fmt.Printf("It will take %d years and %d %s to repay the loan!\n", int(years), months, monthsSuff)
		return
	}
}

func calculateDiffPayment(principal, periods, interest, month float64) float64 {
	return math.Ceil(principal/periods + calculateMonthlyInterest(interest)*(principal-(principal*(month-1))/periods))
}

func getDiffOverPayment(principal, periods, interest float64) float64 {
	var sum float64
	for month := 1.0; month <= periods; month++ {
		diff := calculateDiffPayment(principal, periods, interest, month)
		sum += diff
		fmt.Printf("Month %d: payment is %d\n", int(month), int(diff))
	}
	return sum - principal
}

func printAnnuityBasedOnMissingArg(periods *float64, interest *float64, payment *float64, principal *float64) {
	switch {
	case *periods == 0:
		numberOfMonths := calculateNumberOfMonths(*interest, *payment, *principal)
		printNumberOfMonths(numberOfMonths)
		overPayment := getAnnuityOverPayment(*principal, numberOfMonths, *payment)
		fmt.Printf("Overpayment = %d\n", overPayment)

	case *payment == 0:
		annuity := calculateMonthlyPayment(*principal, *interest, *periods)
		monthlyP := math.Ceil(annuity)
		overPayment := getAnnuityOverPayment(*principal, *periods, monthlyP)
		fmt.Printf("Your annuity payment = %d!\n", int(monthlyP))
		fmt.Printf("Overpayment = %d\n", overPayment)

	case *principal == 0:
		princ := calculateLoanPrincipal(*payment, *interest, *periods)
		fmt.Printf("Your loan principal = %d!\n", int(princ))
		overPayment := getAnnuityOverPayment(princ, *periods, *payment)
		fmt.Printf("Overpayment = %d\n", overPayment)
	}
}

func getAnnuityOverPayment(principal, periods, payment float64) int {
	return int(math.Ceil(periods*payment - principal))
}

func checkIfValuesNotPositive(principal, periods, payment, interest float64) bool {
	if principal < 0 || periods < 0 || payment < 0 || interest <= 0 {
		return true
	}
	return false
}

func Start() {

	paymentType := flag.String("type", "", "Payment type. Annuity or diff.")
	principal := flag.Float64("principal", 0.0, "Loan principal")
	periods := flag.Float64("periods", 0.0, "Loan periods")
	interest := flag.Float64("interest", 0.0, "Loan interest")
	payment := flag.Float64("payment", 0.0, "Loan payment")

	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Println("Pass --help for command options.")
		return
	}

	if !(*paymentType == diffOpt || *paymentType == annuityOpt) {
		fmt.Println("Incorrect parameters")
		return
	}

	if *paymentType == diffOpt {

		if *payment != 0 || !(*principal != 0 && *periods != 0 && *interest != 0) {
			fmt.Println("Incorrect parameters")
			return
		}
		if checkNotPositive := checkIfValuesNotPositive(*principal, *periods, *payment, *interest); checkNotPositive {
			fmt.Println("Incorrect parameters")
			return
		}

		overPayment := getDiffOverPayment(*principal, *periods, *interest)
		fmt.Println()
		fmt.Printf("Overpayment = %d\n", int(overPayment))
	}

	if *paymentType == annuityOpt {
		if len(os.Args[1:]) != 4 {
			fmt.Println("Incorrect parameters")
			return
		}

		if checkNotPositive := checkIfValuesNotPositive(*principal, *periods, *payment, *interest); checkNotPositive {
			fmt.Println("Incorrect parameters")
			return
		}

		printAnnuityBasedOnMissingArg(periods, interest, payment, principal)
	}
}
