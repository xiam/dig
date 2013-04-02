package dig

import (
	"encoding/json"
	"testing"
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

	list := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	var i int

	err = Get(&list, &i, 5)

	if err != nil {
		t.Errorf("Test failed")
	}

	if i != 5 {
		t.Errorf("Test failed")
	}
}

func TestMatrix(t *testing.T) {
	var err error

	m33 := [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	var i int

	err = Get(&m33, &i, 1, 1)

	if err != nil {
		t.Errorf("Test failed")
	}

	if i != 4 {
		t.Errorf("Test failed")
	}

	err = Get(&m33, &i, 2, 0)

	if err != nil {
		t.Errorf("Test failed")
	}

	if i != 6 {
		t.Errorf("Test failed")
	}

	err = Get(&m33, &i, 9, 9)

	if i != 0 {
		t.Errorf("Test failed")
	}

	if err == nil {
		t.Errorf("Test failed")
	}

	// Non assignable
	err = Get(&m33, &i, 1)

	if i != 0 {
		t.Errorf("Test failed")
	}

	if err == nil {
		t.Errorf("Test failed")
	}
}

func TestFloatMatrix(t *testing.T) {
	var err error

	mf33 := [][]float64{
		[]float64{0.0, 1.0, 2.0},
		[]float64{3.0, 4.0, 5.0},
		[]float64{6.0, 7.0, 8.0},
	}
	var f float64

	err = Get(&mf33, &f, 1, 2)

	if err != nil {
		t.Errorf("Test failed")
	}

	if f != 5.0 {
		t.Errorf("Test failed")
	}

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
	json.Unmarshal([]byte(s), &m)

	var children []interface{}

	err := Get(&m, &children, "data", "children")

	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	for i, _ := range children {
		child := children[i].(map[string]interface{})
		val := Float64(&child, "data", "created")
		if val == 0.0 {
			t.Fatalf("Failed conversion or traversal.")
		}
	}

}

func TestMap(t *testing.T) {
	var err error

	m := map[string]string{
		"Hello": "World",
	}

	var s string

	// Simple assignment
	err = Get(&m, &s, "Hello")

	if err != nil {
		t.Errorf("Test failed")
	}

	if s != "World" {
		t.Errorf("Test failed")
	}

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

	if err != nil {
		t.Errorf("Test failed")
	}

	if s != "d" {
		t.Errorf("Test failed")
	}

	// Non existent key
	err = Get(&m2, &s, "third", "doest", "not", "exists")

	if err == nil {
		t.Errorf("Test failed")
	}

	if s != "" {
		t.Errorf("Test failed")
	}

}

func TestJSON(t *testing.T) {

	var m map[string]interface{}

	json.Unmarshal([]byte(jsonTest), &m)

	var s string
	Get(&m, &s, "PMap", "17", "Tag")

	if s != "tag17" {
		t.Errorf("Test failed.")
	}

	var f32 float32
	Get(&m, &f32, "PFloat32")

	if f32 != float32(14.1) {
		t.Errorf("Test failed.")
	}

	var f64 float64
	Get(&m, &f64, "PFloat32")

	if f64 != float64(14.1) {
		t.Errorf("Test failed.")
	}

	var b bool
	Get(&m, &b, "PBool")

	if b != true {
		t.Errorf("Test failed.")
	}

	var ui64 uint64
	Get(&m, &ui64, "PUint64")

	if ui64 != uint64(11) {
		t.Errorf("Test failed.")
	}

	var i int
	Get(&m, &ui64, "String")

	if i != 0 {
		t.Errorf("Test failed.")
	}

	Get(&m, &s, "PSlice", 1, "Tag")

	if s != "tag21" {
		t.Errorf("Test failed.")
	}

}

func TestJSON2(t *testing.T) {

	var m map[string]interface{}

	json.Unmarshal([]byte(jsonTest), &m)

	foo := map[string]string{
		"test": String(&m, "PMap", "17", "Tag"),
	}

	if foo["test"] != "tag17" {
		t.Errorf("Test failed.")
	}

}

func TestSet(t *testing.T) {

	var m = map[string]interface{}{
		"path": map[string]interface{}{
			"to": map[string]interface{}{
				"variable": 2,
			},
		},
	}

	var i int

	err := Set(&m, 1, "path", "to", "variable")

	if err != nil {
		t.Errorf("ER: %v\n", err.Error())
	}

	Get(&m, &i, "path", "to", "variable")

	if i != 1 {
		t.Errorf("Test failed.")
	}

	err = Set(&m, 42, "path", "to", "non", "existent", "key")

	if err == nil {
		t.Errorf("Expecting an error.")
	}

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

	if err != nil {
		t.Errorf("ER: %v\n", err.Error())
	}

	Get(&m, &i, "path", "to", "variable")

	if i != 42 {
		t.Errorf("Test failed.")
	}

}

func TestSetSimpleMap(t *testing.T) {

	var m = map[string]int{
		"key": 2,
	}

	var i int

	err := Set(&m, 42, "key")

	if err != nil {
		t.Errorf("ER: %v\n", err.Error())
	}

	Get(&m, &i, "key")

	if i != 42 {
		t.Errorf("Test failed.")
	}

	err = Set(&m, 43, "does", "not-exists")

	if err == nil {
		t.Errorf("Expecting an error.")
	}

	err = Set(&m, 44, "new-key")

	if err != nil {
		t.Errorf("Test failed.")
	}

	Get(&m, &i, "new-key")

	if i != 44 {
		t.Errorf("Test failed.")
	}

	err = Set(&m, 44)

	if err == nil {
		t.Errorf("Expecting an error.")
	}

}

func TestDig(t *testing.T) {
	var i int

	m := map[string]interface{}{}

	Dig(&m, "foo", "bar", "baz")
	Set(&m, 42, "foo", "bar", "baz")
	Get(&m, &i, "foo", "bar", "baz")

	if i != 42 {
		t.Errorf("Test failed.")
	}

}
