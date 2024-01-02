package stick_filters

import (
	"fmt"
	"testing"
)

func TestObjectsFromPrefix(t *testing.T) {
	fmt.Println("!!!")
	vars := map[string]interface{}{
		"proxy_host":  "foo",
		"proxy_port":  "bar",
		"proxy0_host": "foo1",
		"proxy0_port": "bar1",
		"proxy2_host": "foo2",
		"proxy2_port": "bar2",
		"proxy3_port": "bar3",
	}
	defaultObj := map[string]interface{}{
		"host":      "default_host",
		"generator": "tango",
	}
	expected := []map[string]interface{}{
		map[string]interface{}{
			"host":      "foo",
			"port":      "bar",
			"generator": "tango",
		},
		map[string]interface{}{
			"host":      "foo1",
			"port":      "bar1",
			"generator": "tango",
		},
		map[string]interface{}{
			"host":      "foo2",
			"port":      "bar2",
			"generator": "tango",
		},
		map[string]interface{}{
			"host":      "default_host",
			"port":      "bar3",
			"generator": "tango",
		},
	}
	results := objectsFromPrefix(vars, "proxy", 10, defaultObj)
	if len(results) != len(expected) {
		t.Errorf("Expected %v results, got %v", len(expected), len(results))
	}
	fmt.Println(results)
	for i, result := range results {
		expectResult := expected[i]
		for key, val := range expectResult {
			if result.(map[string]interface{})[key] != val {
				t.Errorf("Expected %v, got %v", val, result.(map[string]interface{})[key])
			}
		}
	}

}
