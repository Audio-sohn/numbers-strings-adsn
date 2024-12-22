package parse_string

import "testing"

//testing in GO:
// -jede test routine muss im identifier mit "Test..." starten
// -nur ein argument, pointer zu einer testing struct, die verschiedene methoden und daten umfasst

// t.Log(args ...)	Logs informational messages during the test.
// t.Error(args ...)	Logs an error and marks the test as failed, but continues execution.
// t.Errorf(format, ...)	Logs a formatted error message and marks the test as failed, but continues execution.
// t.Fail()	Marks the test as failed, but does not stop execution.
// t.FailNow()	Marks the test as failed and stops execution immediately.
// t.Fatal(args ...)	Logs an error, marks the test as failed, and stops execution immediately.
// t.Fatalf(format, ...)	Logs a formatted error message, marks the test as failed, and stops execution immediately.
// t.Skip(args ...)	Skips the test and logs a message (useful for conditionally skipping tests).
// t.Helper()	Marks the calling function as a helper, improving error stack traces (useful for custom helpers).

func TestParseString(t *testing.T) {

	if ParseString("0") != 0 {

		t.Errorf("ParseString(\"0\") == %d, expected 0", ParseStringHexadecimal("0"))

	}

	if ParseString("A") != 10 {

		t.Errorf("ParseString(\"A\") == %d, expected 10", ParseStringHexadecimal("A"))

	}

	if ParseString("ABCD") != 43981 {

		t.Errorf("ParseString(\"ABCD\") == %d, expected 43981", ParseStringHexadecimal("ABCD"))

	}

	if ParseString("3152") != 3152 {

		t.Errorf("ParseString(\"3152\") == %d, expected 3152", ParseStringHexadecimal("3152"))

	}

	if ParseString("AF21") != 44833 {

		t.Errorf("ParseString(\"AF21\") == %d, expected 44833", ParseStringHexadecimal("AF21"))

	}

}
