package atm

var values = [5]int{20, 50, 100, 200, 500}

var notPossibleValue = []int{-1}

type ATM struct {
	counter [5]int
}

func Constructor() ATM {
	return ATM{}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i, c := range banknotesCount {
		this.counter[i] += c
	}
}

func (this *ATM) Withdraw(amount int) []int {
	ret := make([]int, 5)
	for i := 4; i >= 0; i-- {
		if this.counter[i] != 0 {
			value := values[i]
			count := amount / value
			if count > this.counter[i] {
				count = this.counter[i]
			}
			ret[i] = count
			amount -= value * count
			if amount == 0 {
				for j := 4; j >= i; j-- {
					if ret[j] != 0 {
						this.counter[j] -= ret[j]
					}
				}
				return ret
			}
		}
	}
	return notPossibleValue
}
