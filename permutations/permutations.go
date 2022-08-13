package permutations

const digits = 4

type Number [digits]int8

type Permutaions struct {
	numbers []Number
}

func New() (result Permutaions) {
	return generate(Permutaions{})
}

func (this Permutaions) Len() int {
	return len(this.numbers)
}

func (this Permutaions) Get(i int) Number {
	return this.numbers[i]
}

func generate(result Permutaions) Permutaions {
	currentNumber := firstNumber()
	for hasNext := true; hasNext; currentNumber, hasNext = next(currentNumber) {
		result.numbers = append(result.numbers, currentNumber)
	}
	return result
}

func firstNumber() Number {
	var result Number
	var i int8
	for i = 0; i < digits; i++ {
		result[i] = i
	}
	return result
}

func next(number Number) (Number, bool) {
	return inc(number, digits-1)
}

func inc(number Number, position int) (Number, bool) {
	if position < 0 {
		return number, false
	}
	number[position] = nextFree(number, position)
	if number[position] != -1 {
		return number, true
	}
	result, hasNext := inc(number, position-1)
	result[position] = nextFree(result, position)
	return result, hasNext
}

func nextFree(number Number, position int) int8 {
	result := number[position] + 1
	for ; result < 10; result++ {
		if !occupied(number, result, position) {
			return result
		}
	}
	return -1
}

func occupied(number Number, digit int8, position int) bool {
	for i := 0; i < digits; i++ {
		if i != position && number[i] == digit {
			return true
		}
	}
	return false
}
