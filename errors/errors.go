/*

libninty - nintendo network utility library for golang
Copyright (C) 2018 superwhiskers <whiskerdev@protonmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

*/

package errors

import "fmt"

// LibnintyError implements a custom error type used in libninty
type LibnintyError struct{
	scope string
	error string
}

// Error formats the error held in a LibnintyError as a string
func (e LibnintyError) Error() string {

	return fmt.Sprintf("libninty: %s: %s", e.scope, e.error)

}

var (
	ByteBufferOverreadError = LibnintyError{
		scope: "bytebuffer",
		error: "read exceeds buffer capacity",
	}
	ByteBufferOverwriteError = LibnintyError{
		scope: "bytebuffer",
		error: "write exceeds buffer capacity",
	}
)
