package mapulator

import (
	"fmt"
	"testing"
)

func Validate(t *testing.T, got interface{}, want interface{}) {
	if fmt.Sprint(got) != fmt.Sprint(want) {
		t.Error("\nExpect:\n", fmt.Sprint(want))
		t.Error("\nInstead:\n", fmt.Sprint(got))
	}
}

//? Test Set()

func TestMapulatorSetSingleValue(t *testing.T) {
	name := "Jhon"
	want := map[string]interface{}{
		"user": map[string]interface{}{
			"name": name,
		},
	}
	got := Set(map[string]interface{}{}, "user.name", name)

	Validate(t, got, want)
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
	got := Set(map[string]interface{}{}, "user.name.full", map[string]interface{}{
		"value":  fullname,
		"length": len(fullname),
	})

	Validate(t, got, want)
}

//? Test SetNew()
func TestMapulatorSetNew(t *testing.T) {
	want := map[string]interface{}{
		"status": "success",
	}
	got := SetNew("status", "success")

	Validate(t, got, want)
}
