package main

import (
	"fmt"

	"github.com/msoap/raphanus/client"
	"github.com/msoap/raphanus/common"
)

func main() {
	// with default address:
	raph := raphanusclient.New()
	// or with another address:
	// raph := raphanusclient.New(raphanusclient.Cfg{Address: "http://localhost:8771"})
	// or with authentication:
	// raph := raphanusclient.New(raphanusclient.Cfg{User: "uname", Password: "pass"})

	printStat(raph)

	saveIntKey(raph, "k1", 123, 5)
	saveIntKey(raph, "k2", 777, 10)
	saveIntKey(raph, "k3", 771, 7)
	incrDecrIntKey(raph, "k2")
	printIntKey(raph, "k1")
	updateIntKey(raph, "k1", 321)
	printIntKey(raph, "k1")

	testStringValues(raph, "k3")
	testListValues(raph)
	testDictValues(raph)

	printKeys(raph)
	printLength(raph)

	removeKey(raph, "k1")
}

func printStat(raph raphanusclient.Client) {
	stat, err := raph.Stat()
	if err != nil {
		fmt.Printf("Stat got error: %s\n", err)
		return
	}

	fmt.Printf("Stat (version: %s):\n  MemAlloc: %d b\n  MemTotalAlloc: %d b\n  MemMallocs: %d\n  MemFrees: %d\n  MemHeapObjects: %d\n  GCPauseTotalNs: %d\n",
		stat.Version,
		stat.MemAlloc,
		stat.MemTotalAlloc,
		stat.MemMallocs,
		stat.MemFrees,
		stat.MemHeapObjects,
		stat.GCPauseTotalNs,
	)
}

func printKeys(raph raphanusclient.Client) {
	allKeys, err := raph.Keys()
	if err != nil {
		fmt.Printf("Keys got error: %s\n", err)
		return
	}

	fmt.Printf("all keys: %d\n", len(allKeys))
}

func printLength(raph raphanusclient.Client) {
	length, err := raph.Length()
	if err != nil {
		fmt.Printf("Length got error: %s\n", err)
		return
	}

	fmt.Printf("Count of keys: %d\n", length)
}

func saveIntKey(raph raphanusclient.Client, key string, value int64, ttl int) {
	err := raph.SetInt(key, value, ttl)
	if err != nil {
		fmt.Printf("SetInt got error: %s\n", err)
		return
	}

	fmt.Printf("Int value (%s: %d) saved\n", key, value)
}

func updateIntKey(raph raphanusclient.Client, key string, value int64) {
	err := raph.UpdateInt(key, value)
	if err != nil {
		fmt.Printf("UpdateInt got error: %s\n", err)
		return
	}

	fmt.Printf("Int value (%s: %d) updated\n", key, value)
}

func removeKey(raph raphanusclient.Client, key string) {
	err := raph.Remove(key)
	if err != nil {
		fmt.Printf("Remove got error: %s\n", err)
		return
	}

	fmt.Printf("Key %s removed\n", key)
}

func printIntKey(raph raphanusclient.Client, key string) {
	intVal, err := raph.GetInt(key)
	if err != nil {
		fmt.Printf("GetInt got error: %s\n", err)
		return
	}

	fmt.Printf("Key %s, integer value: %d\n", key, intVal)
}

func incrDecrIntKey(raph raphanusclient.Client, key string) {
	if err := raph.IncrInt(key); err != nil {
		fmt.Printf("IncrInt got error: %s\n", err)
		return
	}
	printIntKey(raph, key)

	if err := raph.DecrInt(key); err != nil {
		fmt.Printf("DecrInt got error: %s\n", err)
		return
	}
	printIntKey(raph, key)
}

func printStrKey(raph raphanusclient.Client, key string) {
	strVal, err := raph.GetStr(key)
	if err != nil {
		fmt.Printf("GetStr got error: %s\n", err)
		return
	}

	fmt.Printf("Key %s, string value: %s\n", key, strVal)
}

func testStringValues(raph raphanusclient.Client, key string) {
	if err := raph.SetStr(key, "str val 1", 7); err != nil {
		fmt.Printf("SetStr got error: %s\n", err)
		return
	}
	printStrKey(raph, key)

	if err := raph.UpdateStr(key, "str val new"); err != nil {
		fmt.Printf("SetStr got error: %s\n", err)
		return
	}
	printStrKey(raph, key)
}

func testListValues(raph raphanusclient.Client) {
	if err := raph.SetList("key_list_01", raphanuscommon.ListValue{"l1", "l2", "l3"}, 10); err != nil {
		fmt.Printf("SetList got error: %s\n", err)
		return
	}

	listVal, err := raph.GetList("key_list_01")
	if err != nil {
		fmt.Printf("GetList got error: %s\n", err)
		return
	}
	fmt.Printf("List value: %v\n", listVal)

	if err = raph.UpdateList("key_list_01", raphanuscommon.ListValue{"l1", "l2", "l3_new"}); err != nil {
		fmt.Printf("UpdateList got error: %s\n", err)
		return
	}

	if err = raph.SetListItem("key_list_01", 1, "l2_new"); err != nil {
		fmt.Printf("SetListItem got error: %s\n", err)
		return
	}

	strVal, err := raph.GetListItem("key_list_01", 1)
	if err != nil {
		fmt.Printf("GetListItem got error: %s\n", err)
		return
	}
	fmt.Printf("GetListItem(1): %s\n", strVal)

	listVal, err = raph.GetList("key_list_01")
	if err != nil {
		fmt.Printf("GetList got error: %s\n", err)
		return
	}
	fmt.Printf("List new value: %v\n", listVal)
}

func testDictValues(raph raphanusclient.Client) {
	if err := raph.SetDict("key_dict_01", raphanuscommon.DictValue{"dk1": "d1", "dk2": "d2"}, 10); err != nil {
		fmt.Printf("SetDict got error: %s\n", err)
		return
	}

	if err := raph.UpdateDict("key_dict_01", raphanuscommon.DictValue{"dk1": "d1", "dk2": "d2", "dk3": "d3"}); err != nil {
		fmt.Printf("UpdateDict got error: %s\n", err)
		return
	}

	dictVal, err := raph.GetDict("key_dict_01")
	if err != nil {
		fmt.Printf("GetDict got error: %s\n", err)
		return
	}
	fmt.Printf("DictValue value: %v\n", dictVal)

	if err = raph.SetDictItem("key_dict_01", "dk2", "d2_new"); err != nil {
		fmt.Printf("SetDictItem got error: %s\n", err)
		return
	}

	strVal, err := raph.GetDictItem("key_dict_01", "dk2")
	if err != nil {
		fmt.Printf("GetDictItem got error: %s\n", err)
		return
	}
	fmt.Printf("GetListItem('dk2'): %s\n", strVal)

	if err = raph.RemoveDictItem("key_dict_01", "dk1"); err != nil {
		fmt.Printf("RemoveDictItem got error: %s\n", err)
		return
	}

	dictVal, err = raph.GetDict("key_dict_01")
	if err != nil {
		fmt.Printf("GetDict got error: %s\n", err)
		return
	}
	fmt.Printf("DictValue new value: %v\n", dictVal)
}
