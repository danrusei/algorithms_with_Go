package primes

import "math"

func segmentedSieve(num int) []int {
	limit := int(math.Floor(math.Sqrt(float64(num))) + 1)
	primes := getPrimes(limit)

	result := []int{}

	low := limit
	high := limit * 2

	for low < num {
		if high >= num {
			high = num
		}

		// mark primes in current range. A value in mark[i]
		// will finally be false if 'i-low' is Not a prime, else true.
		mark := make([]bool, limit+1)
		for i := 0; i < limit+1; i++ {
			mark[i] = true
		}

		for i := 0; i < len(primes); i++ {

			//find the minimum number in [low..high] that is a multiple of prime[i]
			loLim := int(math.Floor(float64(low/primes[i])) * float64(primes[i]))

			if loLim < low {
				loLim += primes[i]
			}
			// mark multiples of prime[i] in [low..high]: marking j - low for j
			for j := loLim; j < high; j += primes[i] {
				mark[j-low] = false
			}
		}

		for i := low; i < high; i++ {
			if mark[i-low] {
				result = append(result, i)
			}
		}

		// update low and high for next segment
		low = low + limit
		high = high + limit

	}
	primes = append(primes, result...)
	return primes
}

func getPrimes(num int) []int {
	numbers := make([]bool, num)

	//create the slice of bools and mark initially all true
	//a value in numbers[i] will finally be false if i is Not a prime, else true.
	for i := 0; i < num; i++ {
		numbers[i] = true
	}

	for p := 2; p*p < num; p++ {
		//if numbers[p] is not changed, then it is a prime
		// this selects only the alreade identified prime numbers
		if numbers[p] {
			// search for multiple of p, and mark them as false
			for i := p * p; i < num; i += p {
				numbers[i] = false
			}
		}
	}

	//create the list that containts only the prime numbers
	primes := []int{}
	for i := 2; i < num; i++ {
		if numbers[i] {
			primes = append(primes, i)
		}
	}
	return primes
}
