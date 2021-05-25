package service

import (
	"log"

	"com.m3d.dev/nearestprimes/pkg/http/rest/model"
	"github.com/kavehmz/prime"
)

var RestService model.NearestPrimeModel

// Initializing primes
//
func Init() {
	log.Println("Initializing primes")
	RestService.Primes = prime.Primes(999999999)
	log.Println("Done prime initialization")
}

//getNearestPrime
func GetNearestPrime(num int) uint64 {
	return performBinarySearch(0, len(RestService.Primes)-1, num)
}

//performBinarySearch
func performBinarySearch(min, max, num int) uint64 {
	if num <= 2 {
		return 0
	}
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
