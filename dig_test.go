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
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var jsonTest = `{
	"Bool": false,
	"Int": 0,
	"Int8": 0,
	"Int16": 0,
	"Int32": 0,
	"Int64": 0,
	"Uint": 0,
	"Uint8": 0,
	"Uint16": 0,
	"Uint32": 0,
	"Uint64": 0,
	"Uintptr": 0,
	"Float32": 0,
	"Float64": 0,
	"bar": "",
	"bar2": "",
				"IntStr": "0",
	"PBool": true,
	"PInt": 2,
	"PInt8": 3,
	"PInt16": 4,
	"PInt32": 5,
	"PInt64": 6,
	"PUint": 7,
	"PUint8": 8,
	"PUint16": 9,
	"PUint32": 10,
	"PUint64": 11,
	"PUintptr": 12,
	"PFloat32": 14.1,
	"PFloat64": 15.1,
	"String": "",
	"PString": "16",
	"Map": null,
	"MapP": null,
	"PMap": {
		"17": {
			"Tag": "tag17"
		},
		"18": {
			"Tag": "tag18"
		}
	},
	"PMapP": {
		"19": {
			"Tag": "tag19"
		},
		"20": null
	},
	"EmptyMap": null,
	"NilMap": null,
	"Slice": null,
	"SliceP": null,
	"PSlice": [
		{
			"Tag": "tag20"
		},
		{
			"Tag": "tag21"
		}
	],
	"PSliceP": [
		{
			"Tag": "tag22"
		},
		null,
		{
			"Tag": "tag23"
		}
	],
	"EmptySlice": null,
	"NilSlice": null,
	"StringSlice": null,
	"ByteSlice": null,
	"Small": {
		"Tag": ""
	},
	"PSmall": null,
	"PPSmall": {
		"Tag": "tag31"
	},
	"Interface": null,
	"PInterface": 5.2
}`

func TestList(t *testing.T) {
	var err error
	var i int

	list := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	err = Get(&list, &i, 5)
	assert.NoError(t, err)
	assert.Equal(t, 5, i)
}

func TestMatrix(t *testing.T) {
	var err error
	var i int

	m33 := [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}

	err = Get(&m33, &i, 1, 1)
	assert.NoError(t, err)
	assert.Equal(t, 4, i)

	err = Get(&m33, &i, 2, 0)
	assert.NoError(t, err)
	assert.Equal(t, 6, i)

	// Path does not exist
	err = Get(&m33, &i, 9, 9)
	assert.Error(t, err)
	assert.Equal(t, 0, i)

	// Not assignable
	err = Get(&m33, &i, 1)
	assert.Error(t, err)
	assert.Equal(t, 0, i)
}

func TestFloatMatrix(t *testing.T) {
	var err error
	var f float64

	mf33 := [][]float64{
		[]float64{0.0, 1.0, 2.0},
		[]float64{3.0, 4.0, 5.0},
		[]float64{6.0, 7.0, 8.0},
	}

	err = Get(&mf33, &f, 1, 2)
	assert.NoError(t, err)
	assert.Equal(t, 5.0, f)
}

func TestReddit(t *testing.T) {
	var s = `
{
   "kind":"Listing",
   "data":{
      "modhash":"",
      "children":[
         {
            "kind":"t3",
            "data":{
               "domain":"github.com",
               "banned_by":null,
               "media_embed":{

               },
               "subreddit":"golang",
               "selftext_html":null,
               "selftext":"",
               "likes":null,
               "link_flair_text":null,
               "id":"1bep1x",
               "clicked":false,
               "title":"Easy GOPATH management for Go(lang) projects",
               "media":null,
               "score":9,
               "approved_by":null,
               "over_18":false,
               "hidden":false,
               "thumbnail":"",
               "subreddit_id":"t5_2rc7j",
               "edited":false,
               "link_flair_css_class":null,
               "author_flair_css_class":null,
               "downs":4,
               "saved":false,
               "is_self":false,
               "permalink":"/r/golang/comments/1bep1x/easy_gopath_management_for_golang_projects/",
               "name":"t3_1bep1x",
               "created":1364825519.0,
               "url":"https://github.com/divoxx/goproj",
               "author_flair_text":null,
               "author":"divoxx",
               "created_utc":1364796719.0,
               "ups":13,
               "num_comments":17,
               "num_reports":null,
               "distinguished":null
            }
         },
         {
            "kind":"t3",
            "data":{
               "domain":"xaprb.com",
               "banned_by":null,
               "media_embed":{

               },
               "subreddit":"golang",
               "selftext_html":null,
               "selftext":"",
               "likes":null,
               "link_flair_text":null,
               "id":"1anxb9",
               "clicked":false,
               "title":"Building MySQL Database Applications with Go",
               "media":null,
               "score":18,
               "approved_by":null,
               "over_18":false,
               "hidden":false,
               "thumbnail":"",
               "subreddit_id":"t5_2rc7j",
               "edited":false,
               "link_flair_css_class":null,
               "author_flair_css_class":null,
               "downs":3,
               "saved":false,
               "is_self":false,
               "permalink":"/r/golang/comments/1anxb9/building_mysql_database_applications_with_go/",
               "name":"t3_1anxb9",
               "created":1363819085.0,
               "url":"http://www.xaprb.com/blog/2013/03/20/building-mysql-database-applications-with-go/",
               "author_flair_text":null,
               "author":"franciscosouza",
               "created_utc":1363790285.0,
               "ups":21,
               "num_comments":6,
               "num_reports":null,
               "distinguished":null
            }
         }
      ],
      "after":"t3_1aewpb",
      "before":null
   }
}
	`
	var m map[string]interface{}
	err := json.Unmarshal([]byte(s), &m)
	assert.NoError(t, err)

	var children []interface{}

	err = Get(&m, &children, "data", "children")
	assert.NoError(t, err)

	for i := range children {
		child := children[i].(map[string]interface{})
		val := Float64(&child, "data", "created")
		assert.NotZero(t, val)
	}
}

func TestMap(t *testing.T) {
	var err error
	var s string

	m := map[string]string{
		"Hello": "World",
	}

	// Simple assignment
	err = Get(&m, &s, "Hello")
	assert.NoError(t, err)
	assert.Equal(t, "World", s)

	m2 := map[string]map[string]map[string]string{
		"first": map[string]map[string]string{
			"first.1": map[string]string{
				"col.1": "a",
				"col.2": "b",
				"col.3": "c",
			},
		},
		"second": map[string]map[string]string{
			"second.2": map[string]string{
				"col.4": "d",
				"col.5": "e",
				"col.6": "f",
			},
		},
	}

	// Nested keys
	err = Get(&m2, &s, "second", "second.2", "col.4")
	assert.NoError(t, err)
	assert.Equal(t, "d", s)

	// Non existent key
	err = Get(&m2, &s, "third", "doest", "not", "exists")
	assert.Error(t, err)
	assert.Equal(t, "", s)
}

func TestJSON(t *testing.T) {
	var m map[string]interface{}

	err := json.Unmarshal([]byte(jsonTest), &m)
	assert.NoError(t, err)

	var s string
	err = Get(&m, &s, "PMap", "17", "Tag")
	assert.NoError(t, err)
	assert.Equal(t, "tag17", s)

	var f32 float32
	err = Get(&m, &f32, "PFloat32")
	assert.NoError(t, err)
	assert.Equal(t, float32(14.1), f32)

	var f64 float64
	err = Get(&m, &f64, "PFloat32")
	assert.NoError(t, err)
	assert.Equal(t, float64(14.1), f64)

	var b bool
	err = Get(&m, &b, "PBool")
	assert.NoError(t, err)
	assert.Equal(t, true, b)

	var ui64 uint64
	err = Get(&m, &ui64, "PUint64")
	assert.NoError(t, err)
	assert.Equal(t, uint64(11), ui64)

	var i int
	err = Get(&m, &ui64, "String")
	assert.NoError(t, err)
	assert.Zero(t, i)

	err = Get(&m, &s, "PSlice", 1, "Tag")
	assert.NoError(t, err)
	assert.Equal(t, "tag21", s)
}

func TestJSON2(t *testing.T) {
	var m map[string]interface{}

	err := json.Unmarshal([]byte(jsonTest), &m)
	assert.NoError(t, err)

	foo := map[string]string{
		"test": String(&m, "PMap", "17", "Tag"),
	}
	assert.NoError(t, err)
	assert.Equal(t, "tag17", foo["test"])
}

func TestSet(t *testing.T) {
	var i int
	var m = map[string]interface{}{
		"path": map[string]interface{}{
			"to": map[string]interface{}{
				"variable": 2,
			},
		},
	}

	err := Set(&m, 1, "path", "to", "variable")
	assert.NoError(t, err)

	err = Get(&m, &i, "path", "to", "variable")
	assert.NoError(t, err)
	assert.Equal(t, 1, i)

	err = Set(&m, 42, "path", "to", "non", "existent", "key")
	assert.Error(t, err)
}

func TestSetN(t *testing.T) {
	var m = map[string]map[string]map[string]int{
		"path": map[string]map[string]int{
			"to": map[string]int{
				"variable": 2,
			},
		},
	}
	var i int

	err := Set(&m, 42, "path", "to", "variable")
	assert.NoError(t, err)

	err = Get(&m, &i, "path", "to", "variable")
	assert.NoError(t, err)
	assert.Equal(t, 42, i)
}

func TestSetSimpleMap(t *testing.T) {
	var m = map[string]int{
		"key": 2,
	}
	var i int

	err := Set(&m, 42, "key")
	assert.NoError(t, err)

	err = Get(&m, &i, "key")
	assert.NoError(t, err)
	assert.Equal(t, 42, i)

	err = Set(&m, 43, "does", "not-exists")
	assert.Error(t, err)

	err = Set(&m, 44, "new-key")
	assert.NoError(t, err)

	err = Get(&m, &i, "new-key")
	assert.NoError(t, err)
	assert.Equal(t, 44, i)

	err = Set(&m, 44)
	assert.Error(t, err)
}

func TestDig(t *testing.T) {
	var i int

	m := map[string]interface{}{}

	err := Dig(&m, "foo", "bar", "baz")
	assert.NoError(t, err)

	err = Set(&m, 42, "foo", "bar", "baz")
	assert.NoError(t, err)

	err = Get(&m, &i, "foo", "bar", "baz")
	assert.NoError(t, err)
	assert.Equal(t, 42, i)
}
