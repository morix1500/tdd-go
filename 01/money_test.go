package money

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultipliction(t *testing.T) {
	five := NewDollar(5)
	assert.Equal(t, five.Times(2), NewDollar(10), "times function have problem")
	assert.Equal(t, five.Times(3), NewDollar(15), "times function have problem")
}

func TestEquality(t *testing.T) {
	assert.True(t, NewDollar(5).Equals(NewDollar(5)), "amount is not equals")
	assert.False(t, NewDollar(6).Equals(NewDollar(5)), "amount is equals")
	assert.False(t, NewDollar(5).Equals(NewFranc(5)), "amount is equals")
}

func TestCurrency(t *testing.T) {
	assert.Equal(t, "USD", NewDollar(5).currency, "currency is not USD")
	assert.Equal(t, "CHF", NewFranc(5).currency, "currency is not CHF")
}

func TestSimpleAddition(t *testing.T) {
	five := NewDollar(5)
	sum := five.Plus(five)
	bank := NewBank()
	reduced := bank.Reduce(sum, "USD")
	assert.Equal(t, NewDollar(10), reduced)
}

func TestPlusReturnsSum(t *testing.T) {
	five := NewDollar(5)
	res := five.Plus(five)
	sum := res.(Sum)
	assert.Equal(t, five, sum.augend)
	assert.Equal(t, five, sum.addend)
}

func TestReduceSum(t *testing.T) {
	sum := Sum{
		augend: NewDollar(3),
		addend: NewDollar(4),
	}
	bank := NewBank()
	res := bank.Reduce(sum, "USD")
	assert.Equal(t, NewDollar(7), res)
}

func TestReduceMoney(t *testing.T) {
	bank := NewBank()
	res := bank.Reduce(NewDollar(1), "USD")
	assert.Equal(t, NewDollar(1), res)
}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	res := bank.Reduce(NewFranc(2), "USD")
	assert.Equal(t, NewDollar(1), res)
}

func TestIdentityRate(t *testing.T) {
	assert.Equal(t, 1, NewBank().Rate("USD", "USD"))
}

func TestMixedAddition(t *testing.T) {
	fiveBucks := NewDollar(5)
	tenFrancs := NewFranc(10)
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	res := bank.Reduce(fiveBucks.Plus(tenFrancs), "USD")
	assert.Equal(t, NewDollar(10), res)
}
