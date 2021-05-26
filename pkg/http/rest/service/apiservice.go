package service

import (
	"log"

	"com.m3d.dev/nearestprimes/pkg/http/rest/model"
	"github.com/kavehmz/prime"
)

var RestService model.NearestPrimeModel

// InitializePrimes godoc
// This is where the prime numbers pre-computed at application start up
func InitializePrimes(num int) {
	log.Println("Initializing primes")
	// The argument of Primes function indicate the highest prime
	RestService.Primes = prime.Primes(uint64(num))
	log.Println("Done prime initialization")
}

//getNearestPrime
func GetNearestPrime(num int) uint64 {
	return performBinarySearch(0, len(RestService.Primes)-1, num)
}

//performBinarySearch
func performBinarySearch(min, max, num int) uint64 {
	// for numbers less than 3 we return 0
	if num <= 2 {
		return 0
	}
	// logic for the binary search for the closest number
	if min <= max {
		mid := (min + max) / 2
		if mid == 0 || mid == len(RestService.Primes)-1 {
			return RestService.Primes[mid]
		}
		if RestService.Primes[mid] == uint64(num) {
			return RestService.Primes[mid-1]
		}
		if RestService.Primes[mid] < uint64(num) && RestService.Primes[mid+1] > uint64(num) {
			return RestService.Primes[mid]
		}
		if uint64(num) < RestService.Primes[mid] {
			return performBinarySearch(min, mid-1, num)
		} else {
			return performBinarySearch(mid+1, max, num)
		}
	}
	return 0
}
