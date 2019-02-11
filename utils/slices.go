/*

fennel - nintendo network utility library for golang
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

package utils

import "strconv"

// ApplyByteSliceAtOffset applies a byte slice to another byte slice at the specified offset, overwriting any existing indexes already in the slice
func ApplyByteSliceAtOffset(src, dest []byte, offset int) []byte {

	for i, byt := range src {

		dest[offset+i] = byt

	}
	return dest

}

// ConvertInt64SliceToStringSlice converts a slice of int64s to a slice of strings
func ConvertInt64SliceToStringSlice(il []int64) (sl []string) {

	sl = []string{}
	for _, i := range il {

		sl = append(sl, strconv.Itoa(int(i)))

	}

	return

}
