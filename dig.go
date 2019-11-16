// Copyright (c) 2013-today José Nieto, https://xiam.dev
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package dig

import (
	"fmt"
	"reflect"

	"github.com/xiam/to"
)

// Bool returns a bool value following the given src path (map or slice)
func Bool(src interface{}, path ...interface{}) bool {
	var b bool
	err := Get(src, &b, path...)
	if err != nil {
		return false
	}
	return b
}

// Uint64 returns a uint64 value following the given src path (map or slice)
func Uint64(src interface{}, path ...interface{}) uint64 {
	var i uint64
	err := Get(src, &i, path...)
	if err != nil {
		return uint64(0)
	}
	return i
}

// Int64 returns a int64 value following the given src path (map or slice)
func Int64(src interface{}, path ...interface{}) int64 {
	var i int64
	err := Get(src, &i, path...)
	if err != nil {
		return int64(0)
	}
	return i
}

// Float32 returns a float32 value following the given src path (map or slice)
func Float32(src interface{}, path ...interface{}) float32 {
	var f float32
	err := Get(src, &f, path...)
	if err != nil {
		return float32(0)
	}
	return f
}

// Float64 returns a float64 value following the given src path (map or slice)
func Float64(src interface{}, path ...interface{}) float64 {
	var f float64
	err := Get(src, &f, path...)
	if err != nil {
		return float64(0)
	}
	return f
}

// Interface returns an interface{} value following the given src path (map or slice)
func Interface(src interface{}, path ...interface{}) interface{} {
	var i interface{}
	err := Get(src, &i, path...)
	if err != nil {
		return nil
	}
	return i
}

// String returns a string value following the given src path (map or slice)
func String(src interface{}, path ...interface{}) string {
	var s string
	err := Get(src, &s, path...)
	if err != nil {
		return ""
	}
	return s
}

// Set follows the given path on a slice or map and attempts to set it to the
// given value.
func Set(src interface{}, val interface{}, path ...interface{}) error {
	l := len(path)

	if l < 1 {
		return fmt.Errorf("missing path")
	}

	parent := path[0 : l-1]
	last := path[l-1 : l]

	p, err := pick(src, false, parent...)

	if err != nil {
		return err
	}

	if !p.IsValid() {
		return fmt.Errorf("the given path does not exists")
	}

	p.SetMapIndex(reflect.ValueOf(last[0]), reflect.ValueOf(val))

	return nil
}

// Get follows the given path on a slice or map and attempts to set dst to the
// value described by that path.
func Get(src interface{}, dst interface{}, path ...interface{}) error {

	if len(path) < 1 {
		return fmt.Errorf("missing path")
	}

	dv := reflect.ValueOf(dst)

	if dv.Kind() != reflect.Ptr || dv.IsNil() {
		return fmt.Errorf("destination is not a pointer")
	}

	sv := reflect.ValueOf(src)

	if sv.Kind() != reflect.Ptr || sv.IsNil() {
		return fmt.Errorf("source is not a pointer")
	}

	// Setting to zero before setting it again.
	dv.Elem().Set(reflect.Zero(dv.Elem().Type()))

	p, err := pick(src, false, path...)

	if err != nil {
		return err
	}

	if !p.IsValid() {
		return fmt.Errorf("could not find the path: %#v", path)
	}

	p = convert(dv, p)

	if dv.Elem().Type() == p.Type() || dv.Elem().Kind() == reflect.Interface {
		dv.Elem().Set(*p)
	} else {
		return fmt.Errorf("could not assign %s to %s", p.Type(), dv.Elem().Type())
	}

	return nil
}

func convert(dv reflect.Value, p *reflect.Value) *reflect.Value {
	if dv.Elem().Type() != p.Type() {
		// Trying conversion
		if p.CanInterface() {
			t, err := to.Convert(p.Interface(), dv.Elem().Kind())
			if err == nil {
				tv := reflect.ValueOf(t)
				if dv.Elem().Type() == tv.Type() {
					p = &tv
				}
			}
		}
	}
	return p
}

// Dig creates a route to the given path. If the path already exists it
// overwrites it with a zero value.
func Dig(src interface{}, path ...interface{}) error {
	v, err := pick(src, true, path...)
	if !v.IsValid() {
		return fmt.Errorf("could not reach node")
	}
	return err
}

func pick(src interface{}, dig bool, path ...interface{}) (*reflect.Value, error) {
	var err error

	v := reflect.ValueOf(src)

	if v.Kind() != reflect.Ptr || v.IsNil() {
		return nil, fmt.Errorf("source is not a pointer")
	}

	v = v.Elem()

	for _, key := range path {
		u := v
		switch kind := v.Kind(); kind {
		case reflect.Slice:
			v, err = checkSlice(key, v)
			if err != nil {
				return &v, err
			}
		case reflect.Map:
			v = checkMap(key, v, u, dig)
		}
		if v.IsValid() {
			if v.CanInterface() {
				v = reflect.ValueOf(v.Interface())
			}
		}
	}

	return &v, err
}

func checkSlice(key interface{}, v reflect.Value) (reflect.Value, error) {
	switch i := key.(type) {
	case int:
		if i < v.Len() {
			v = v.Index(i)
		} else {
			return v, fmt.Errorf("undefined index: %d", i)
		}
	}
	return v, nil
}

func checkMap(key interface{}, v reflect.Value, u reflect.Value, dig bool) reflect.Value {
	vkey := reflect.ValueOf(key)
	v = v.MapIndex(vkey)
	if dig && !v.IsValid() {
		u.SetMapIndex(vkey, reflect.MakeMap(u.Type()))
		v = u.MapIndex(vkey)
	}
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return v
}
