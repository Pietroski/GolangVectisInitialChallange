package primes

import (
	"math"
	"strconv"
)

type primeListHelper struct {
	primeList           *[]int64
	primeListLastEleIdx int

	listSum int64

	wIdx int
	rIdx int
	rEle int64
}

func NewPrimeListHelper(number int64) *primeListHelper {
	primeList := make([]int64, number, number)
	primeList[0] = 2
	primeList[1] = 3

	return &primeListHelper{
		primeList:           &primeList,
		primeListLastEleIdx: int(number) - 1,

		wIdx:    2,
		listSum: 5,
	}
}

func GenPrimesUpToPos(number int64) (*[]int64, int64) {
	plh := NewPrimeListHelper(number)
	num := int64(5)
	halfNum := genHalfNum(num)

primeGen:
	for plh.genPrime() {

		// Initialises the prime list iterator
		plh.iter()
		for primeNum, ok := plh.nextUpTo(halfNum); ok; primeNum, ok = plh.nextUpTo(halfNum) {
			if num%primeNum == 0 {
				num = num + 2
				halfNum = genHalfNum(num)
				continue primeGen
			}
		}

		plh.updatePrimeListInfo(num)
		num = num + 2
		halfNum = genHalfNum(num)
	}

	return plh.primeList, plh.listSum
}

func genHalfNum(number int64) int64 {
	halfNum := int64(math.Sqrt(float64(number)))
	return halfNum
}

func (plh *primeListHelper) genPrime() bool {
	if plh.wIdx <= plh.primeListLastEleIdx {
		return true
	}

	return false
}

func (plh *primeListHelper) iter() {
	plh.rIdx = 1
	plh.rEle = (*plh.primeList)[plh.rIdx]
}

func (plh *primeListHelper) nextUpTo(limit int64) (int64, bool) {
	//defer func() (int64, bool) {
	//	if r := recover(); r != nil {
	//		return plh.rEle, false
	//	}
	//
	//	return plh.rEle, false
	//}()

	plh.rEle = (*plh.primeList)[plh.rIdx]
	//// TODO: test it!
	//if r := recover(); r != nil {
	//	return plh.rEle, false
	//}

	if plh.rEle <= limit {
		plh.rIdx++
		return plh.rEle, true
	}

	return plh.rEle, false
}

func (plh *primeListHelper) updatePrimeListInfo(newPrimeNum int64) {
	(*plh.primeList)[plh.wIdx] = newPrimeNum
	plh.wIdx++
	plh.sumWith(newPrimeNum)
}

func (plh *primeListHelper) sumWith(newPrimeNum int64) {
	plh.listSum += newPrimeNum
}

func LargestPrimeFactorOfTheMostRightDigitsOf(primesList *[]int64, number int64, digitsNum int) int64 {
	num := mostRightDigits(number, digitsNum)
	lpf := largestPrimeFactor(primesList, num)

	return lpf
}

func mostRightDigits(number int64, digitsNum int) int64 {
	strNum := strconv.Itoa(int(number))
	strNumLen := len(strNum)
	mrd := strNum[strNumLen-digitsNum:] // strNum[strNumLen-3] + strNum[strNumLen-2] + strNum[strNumLen-1]

	newNum, _ := strconv.ParseInt(mrd, 10, 64)

	return newNum
}

func largestPrimeFactor(primesList *[]int64, num int64) int64 {
	sqrtNum := num / 2 // int64(math.Sqrt(float64(num)))
	primeNum := (*primesList)[0]
	lpf := int64(-1)

	for idx := 0; primeNum <= sqrtNum*3; idx++ {
		if num%primeNum == 0 && primeNum > lpf {
			lpf = primeNum
		}

		primeNum = (*primesList)[idx]
	}

	return lpf
}
