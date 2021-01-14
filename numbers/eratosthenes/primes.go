package primes

func primes(num int) []int {
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
