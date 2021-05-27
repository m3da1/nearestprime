package tests

import (
	"testing"

	"com.m3d.dev/nearestprimes/pkg/http/rest/service"
	"github.com/kavehmz/prime"
)

// TestGeneratePrimes
// Testing Generate Primes; This test is to confirm whether the library does the generation of primes
// Testing with a value of 100 should produces a list of primes below
// [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97]
func TestGeneratedPrimes(t *testing.T) {
	primes := prime.Primes(100)
	expected := 25
	actual := len(primes)
	if expected != actual {
		t.Errorf("TestGeneratedPrimes; Expected: %v, Actual: %v", expected, actual)
	}
}

// TestGeneratedPrimeFirstValue
// Testing this should be produce a result of 2 since 2 is the first prime
func TestGeneratedPrimeFirstValue(t *testing.T) {
	primes := prime.Primes(100)
	expected := 2
	actual := primes[0]
	if expected != int(actual) {
		t.Errorf("TestGeneratedPrimeFirstValue; Expected: %v, Actual: %v", expected, actual)
	}
}

// TestSuccessfulInitialized
// This is to confirm whether the primes we generated and store in primes array
func TestSuccessfulInitialized(t *testing.T) {
	err := service.InitializePrimes(100)
	if err != nil {
		t.Errorf(err.Error())
	}
}

// TestFailedInitialized
// This is to confirm whether the primes we generated and store in primes array
// Testing with a value less than 3
func TestFailedInitialized(t *testing.T) {
	err := service.InitializePrimes(1)
	if err == nil {
		t.Errorf("TestFailedInitialized failed")
	}
}

// TestSuccessfulInitializedPrimes
// This is to confirm whether the primes array has been populated
func TestSuccessfulInitializedPrimes(t *testing.T) {
	err := service.InitializePrimes(100)
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(service.RestService.Primes) < 1 {
		t.Errorf("TestSuccessfulInitializedPrimes; Expected a value greater than 0 but found %v", len(service.RestService.Primes))
	}
}

// TestSuccessfulNearestPrime
// This is to test the nearnest prime service
func TestSuccessfulNearestPrime(t *testing.T) {
	err := service.InitializePrimes(100)
	if err != nil {
		t.Errorf(err.Error())
	}
	testdata := 30
	expected := 29
	actual := service.GetNearestPrime(testdata)
	if expected != int(actual) {
		t.Errorf("TestSuccessfulNearestPrime: Expected: %v, Actual: %v", expected, actual)
	}
}

// TestFailedNearestPrime
func TestFailedNearestPrime(t *testing.T) {
	err := service.InitializePrimes(100)
	if err != nil {
		t.Errorf(err.Error())
	}
	testdata := 30
	expected := 27
	actual := service.GetNearestPrime(testdata)
	if expected == int(actual) {
		t.Errorf("TestFailedNearestPrime: Excepted: %v, Actual: %v", expected, actual)
	}
}
