## Project
This is the *Loan Calculator (GO)* project that is part of Hyperskill platform from Jetbrains Academy.

## Project purpose
Learning GO path

Loan calculator project for calculating annuity and differentiated payments.

Sample run:
```bash
Calculate differentiated payment: 

go run cmd/main.go --type=diff --principal=1000000 --periods=10 --interest=10
Month 1: payment is 108334
Month 2: payment is 107500
Month 3: payment is 106667
Month 4: payment is 105834
Month 5: payment is 105000
Month 6: payment is 104167
Month 7: payment is 103334
Month 8: payment is 102500
Month 9: payment is 101667
Month 10: payment is 100834

Overpayment = 45837

```

```bash
Calculate the principal for a user paying 8,722 per month for 120 months (10 years) at 5.6% interest

> go run cmd/main.go --type=annuity --payment=8722 --periods=120 --interest=5.6
Your loan principal = 800018!
Overpayment = 246622

```

```bash
Calculate how long it will take to repay a loan with 500,000 principal, monthly payment of 23,000, and 7.8% interest

go run cmd/main.go --type=annuity --principal=500000 --payment=23000 --interest=7.8
It will take 2 years to repay this loan!
Overpayment = 52000
```