package test

import (
	"fmt"
	"github.com/shopspring/decimal"
	"paractice/pkg"
	"testing"
)

func TestBigF(t *testing.T) {
	str := "10000000.000000000001"
	amountInEth, _ := decimal.NewFromString(str)
	got := pkg.EthToWei(&amountInEth)
	fmt.Println(got)
	fmt.Println(got.String())
}

//func DecimalsToWei(amount *decimal.Decimal, decimals *big.Int) *big.Int {
//	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromBigInt(decimals, 0))
//	result := amount.Mul(mul)
//	wei := new(big.Int)
//	wei.SetString(result.String(), 10)
//	return wei
//}
////
////func EthToWei(amount *decimal.Decimal) *big.Int {
////	wei := DecimalsToWei(amount, big.NewInt(18))
////	return wei
////}

//// 左移相当于乘以 10^n
//func leftShiftBig(f *big.Float, n int) *big.Float {
//	shiftFactor := new(big.Float).SetFloat64(math.Pow(10, float64(n)))
//	return new(big.Float).Mul(f, shiftFactor)
//}
//
//// 右移相当于除以 10^n
//func rightShiftBig(f *big.Float, n int) *big.Float {
//	shiftFactor := new(big.Float).SetFloat64(math.Pow(10, float64(n)))
//	return new(big.Float).Quo(f, shiftFactor)
//}
