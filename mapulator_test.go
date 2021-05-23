package mapulator

import (
	"fmt"
	"testing"
)

func TestMapulatorSetSingleValue(t *testing.T) {
	name := "Jhon"
	want := map[string]interface{}{
		"user": map[string]interface{}{
			"name": name,
		},
	}
	if got := Set(map[string]interface{}{}, "user.name", name); fmt.Sprint(got) != fmt.Sprint(want) {
		t.Error("\nExpect:\n", fmt.Sprint(want))
		t.Error("\nInstead:\n", fmt.Sprint(got))
	}
}

func TestMapulatorSetMultiValue(t *testing.T) {
	fullname := "Jhon"
	want := map[string]interface{}{
		"user": map[string]interface{}{
			"name": map[string]interface{}{
				"full": map[string]interface{}{
					"value":  fullname,
					"length": len(fullname),
				},
			},
		},
	}
	if got := Set(map[string]interface{}{}, "user.name.full", map[string]interface{}{
		"value":  fullname,
		"length": len(fullname),
	}); fmt.Sprint(got) != fmt.Sprint(want) {
		t.Error("\nExpect:\n", fmt.Sprint(want))
		t.Error("\nInstead:\n", fmt.Sprint(got))
	}
}
