/**
 * package jsonops；
 * 	学习使用直接操作json的库：
 * 	2020-09-12: github.com/Jeffail/gabs
 */
package jsonops

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Jeffail/gabs/v2"
)

var (
	demo1 = []byte(`
{
    "info": {
      "name": {
        "first": "lee",
        "last": "darjun"
      },
      "age": 18,
      "hobbies": [
        "game",
        "programming"
      ]
    }
}`)

	demo2 = []byte(`
{
	"outter":{
		"inner":{
			"value1":10,
			"value2":22
		},
		"alsoInner":{
			"value1":20,
			"array1":[
				30, 40
			]
		}
	}
}`)
)

/*-------------------读取json类api---------------------------*/

// 获取某个值得3种api： Path , Search 和 JSONPointer
// 判断是否存在: exists
func GetValueDemo() {
	jsonObj, err := gabs.ParseJSON(demo2)
	if err != nil {
		panic(err)
	}

	var value float64
	var ok bool

	key := "outter.inner.value1"
	value, ok = jsonObj.Path(key).Data().(float64)
	fmt.Printf("outter.inner.value1: value=%v ,ok=%v\n", value, ok)
	// value == 10.0, ok == true

	value, ok = jsonObj.Search("outter", "inner", "value1").Data().(float64)
	// value == 10.0, ok == true
	fmt.Printf("outter.inner.value1: value=%v ,ok=%v\n", value, ok)

	gObj, err := jsonObj.JSONPointer("/outter/alsoInner/array1/1")
	if err != nil {
		panic(err)
	}
	value, ok = gObj.Data().(float64)
	fmt.Printf("/outter/alsoInner/array1/1: value=%v ,ok=%v\n", value, ok)
	// value == 40.0, ok == true

	value, ok = jsonObj.Path("does.not.exist").Data().(float64)
	// value == 0.0, ok == false
	fmt.Printf("does.not.exist: value=%v ,ok=%v\n", value, ok)
	exists := jsonObj.Exists("outter", "inner", "value1")
	fmt.Printf("outter.inner.value1: exists=%v\n", exists)

	// exists == true

	exists = jsonObj.ExistsP("does.not.exist")
	fmt.Printf("does.not.exist: exists=%v\n", exists)

	// exists == false
}

// 通过map遍历: ChildrenMap
func IterMap() {
	jsonParsed, err := gabs.ParseJSON([]byte(`{"object":{"first":1,"second":2,"third":3}}`))
	if err != nil {
		panic(err)
	}

	// S is shorthand for Search
	for key, child := range jsonParsed.S("object").ChildrenMap() {
		fmt.Printf("key: %v, value: %v\n", key, child.Data().(string))
	}
}

// 通过数组遍历: Children
func IterArray() {
	jsonParsed, err := gabs.ParseJSON([]byte(`{"array":["first","second","third"]}`))
	if err != nil {
		panic(err)
	}

	for _, child := range jsonParsed.S("array").Children() {
		fmt.Println(child.Data().(string))
	}
	//	Will print:
	//
	//first
	//second
	//third
}

//
func SearchThroughArrays() {
	jsonParsed, err := gabs.ParseJSON([]byte(`{"array":[{"value":1},{"value":2},{"value":3}]}`))
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonParsed.Path("array.1.value").String())
}

/*-------------------写入json类api-------------------------*/

func Gen() {
	jsonObj := gabs.New()
	// or gabs.Wrap(jsonObject) to work on an existing map[string]interface{}

	jsonObj.Set(10, "outter", "inner", "value")
	jsonObj.SetP(20, "outter.inner.value2")
	jsonObj.Set(30, "outter", "inner2", "value3")

	// To pretty-print:
	//fmt.Println(jsonObj.StringIndent("", "  "))
	fmt.Println(jsonObj.String())
	//{"outter":{"inner":{"value":10,"value2":20},"inner2":{"value3":30}}}
}

func GenArray() {
	jsonObj := gabs.New()

	jsonObj.Array("foo", "array")
	// Or .ArrayP("foo.array")

	jsonObj.ArrayAppend(10, "foo", "array")
	jsonObj.ArrayAppend(20, "foo", "array")
	jsonObj.ArrayAppend(30, "foo", "array")

	fmt.Println(jsonObj.String())
	//{"foo":{"array":[10,20,30]}}
}

func GenArrayByIndex() {
	jsonObj := gabs.New()

	// Create an array with the length of 3
	jsonObj.ArrayOfSize(3, "foo")

	jsonObj.S("foo").SetIndex("test1", 0)
	jsonObj.S("foo").SetIndex("test2", 1)

	// Create an embedded array with the length of 3
	jsonObj.S("foo").ArrayOfSizeI(3, 2)

	jsonObj.S("foo").Index(2).SetIndex(1, 0)
	jsonObj.S("foo").Index(2).SetIndex(2, 1)
	jsonObj.S("foo").Index(2).SetIndex(3, 2)

	fmt.Println(jsonObj.String())
	//{"foo":["test1","test2",[1,2,3]]}
}

func JsonIntParse() {
	sample := []byte(`{"test":{"int":10,"float":6.66}}`)
	dec := json.NewDecoder(bytes.NewReader(sample))
	dec.UseNumber()

	val, err := gabs.ParseJSONDecoder(dec)
	if err != nil {
		panic(fmt.Errorf("Failed to parse:%w ", err))
		return
	}

	intValue, err := val.Path("test.int").Data().(json.Number).Int64()

	fmt.Println(intValue)

}
