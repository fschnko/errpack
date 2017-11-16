/*
Package errpack declares the Pack type for errors collecting.
This is useful for cyclic or unrelated operations
when journaling is not allowed (library packages, etc.).

	errp := New("pack")
	for i := 0; i < 5; i++ {
		errp.Add(someCallWithError(i))
	}

	return errp.ErrorOrNil()


NOTICE: Currently not being thread safe!
*/
package errpack
