/*
  Copyright (c) 2013 José Carlos Nieto, http://xiam.menteslibres.org/

  Permission is hereby granted, free of charge, to any person obtaining
  a copy of this software and associated documentation files (the
  "Software"), to deal in the Software without restriction, including
  without limitation the rights to use, copy, modify, merge, publish,
  distribute, sublicense, and/or sell copies of the Software, and to
  permit persons to whom the Software is furnished to do so, subject to
  the following conditions:

  The above copyright notice and this permission notice shall be
  included in all copies or substantial portions of the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
  LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
  OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
  WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package dig

import (
	"fmt"
	"github.com/gosexy/to"
	"reflect"
)

/*
	Returns a boolean starting from a Slice or Map.
*/
func Bool(src interface{}, route ...interface{}) bool {
	var b bool
	err := Pick(src, &b, route...)
	if err != nil {
		return false
	}
	return b
}

/*
	Returns an uint64 starting from a Slice or Map.
*/
func Uint64(src interface{}, route ...interface{}) uint64 {
	var i uint64
	err := Pick(src, &i, route...)
	if err != nil {
		return uint64(0)
	}
	return i
}

/*
	Returns an int64 starting from a Slice or Map.
*/
func Int64(src interface{}, route ...interface{}) int64 {
	var i int64
	err := Pick(src, &i, route...)
	if err != nil {
		return int64(0)
	}
	return i
}

/*
	Returns a float32 starting from a Slice or Map.
*/
func Float32(src interface{}, route ...interface{}) float32 {
	var f float32
	err := Pick(src, &f, route...)
	if err != nil {
		return float32(0)
	}
	return f
}

/*
	Returns a float64 starting from a Slice or Map.
*/
func Float64(src interface{}, route ...interface{}) float64 {
	var f float64
	err := Pick(src, &f, route...)
	if err != nil {
		return float64(0)
	}
	return f
}

/*
	Returns a string starting from a Slice or Map.
*/
func String(src interface{}, route ...interface{}) string {
	var s string
	err := Pick(src, &s, route...)
	if err != nil {
		return ""
	}
	return s
}

/*
	Starts with src (pointer to Slice or Map) and follows the given route, if the
	route is met, tries to copy or convert the found node into the value pointed
	by dst.
*/
func Pick(src interface{}, dst interface{}, route ...interface{}) error {

	var err error = nil

	dv := reflect.ValueOf(dst)

	if dv.Kind() != reflect.Ptr || dv.IsNil() {
		return fmt.Errorf("Destination is not a pointer.")
	}

	sv := reflect.ValueOf(src)

	if sv.Kind() != reflect.Ptr || sv.IsNil() {
		return fmt.Errorf("Source is not a pointer.")
	}

	var p reflect.Value

	p = sv.Elem()

	for _, curr := range route {
		if err != nil {
			break
		}
		switch p.Kind() {
		// Current p is a slice
		case reflect.Slice:
			switch curr.(type) {
			case int:
				i := curr.(int)
				if i < p.Len() {
					p = p.Index(i)
				} else {
					err = fmt.Errorf("Undefined index %d.", i)
				}
			}
		case reflect.Map:
			p = p.MapIndex(reflect.ValueOf(curr))
		case reflect.Interface:
			err = fmt.Errorf("Got %v.", p.Kind())
		default:
			err = fmt.Errorf("Expecting slice or map. Got %v.", p.Kind())
		}
		if p.IsValid() == true {
			if p.CanInterface() == true {
				p = reflect.ValueOf(p.Interface())
			}
		}
	}

	if err == nil {
		if p.IsValid() == true {
			if dv.Elem().Type() != p.Type() {
				// Trying conversion
				if p.CanInterface() == true {
					var t interface{}
					t, err = to.Convert(p.Interface(), dv.Elem().Kind())
					if err == nil {
						tv := reflect.ValueOf(t)
						if dv.Elem().Type() == tv.Type() {
							p = tv
						}
					}
				}
			}
		} else {
			err = fmt.Errorf("Could not find path.")
		}
	}

	if err != nil {
		dv.Elem().Set(reflect.Zero(dv.Elem().Type()))
	} else {
		if dv.Elem().Type() == p.Type() {
			dv.Elem().Set(p)
		} else {
			err = fmt.Errorf("Could not assign %s to %s.", p.Type(), dv.Elem().Type())
		}
	}

	return err
}
