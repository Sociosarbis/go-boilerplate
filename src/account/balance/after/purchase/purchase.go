package purchase

func accountBalanceAfterPurchase(purchaseAmount int) int {
	onesDigit := purchaseAmount % 10
	if onesDigit < 5 {
		return 100 - (purchaseAmount - onesDigit)
	}
	return 90 - (purchaseAmount - onesDigit)
}
