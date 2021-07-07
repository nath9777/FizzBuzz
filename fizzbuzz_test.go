package fizzbuzz_test

import (
	"testing"

	fizzbuzz "github.com/atulnath977/FizzBuzz"
)

//If i want to test unexported members of fizzbuzz, then use fizzbuzz package.
//If i want to use only exported members of fizzbuzz, then use fizzbuzz_test as package name.

//we'll use fizzbuzz_test since we only have exported function

//Function must not return anything
//function name must begin with Test
func TestFizz(t *testing.T) {
	_, ok := fizzbuzz.Fizzbuzz(1)
	if ok {
		t.Logf("The Value %v should not have returned true", 1)
		t.Fail()
	}

	result, ok := fizzbuzz.Fizzbuzz(3)
	if !ok {
		t.Logf("The Value %v should not have returned true", 1)
		t.Fail()
	}

	if result != "fizz" {
		t.Log("Result should have been fizz")
		t.Fail()
	}
}

// go test path_to_package/current_directory
// go test -v :verbose /more detail
// go test -cover : Tells you how much of your code has been tested(covered)
// go tets -coverprofile=coverage: this will run the test and profile is stored in coverage
// go tool cover -html=coverage: for line coverage
