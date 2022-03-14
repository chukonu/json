package json

import (
	"testing"
)

func TestTraverse(t *testing.T) {
	j := make(JsonObject)
	j["size"] = 1
	j["age"] = 1
	j["name"] = map[string]interface{}{
		"forename":   "george",
		"surname":    "osborne",
		"middlename": 100,
		"usedname": map[string]interface{}{
			"forename": "nigel",
		},
	}

	c := Traverse(j)
	entities := make([]JsonEntity, 0)
	for entity := range c {
		entities = append(entities, entity)
	}
	if len(entities) != 2 {
		t.Log(entities)
		t.Fatalf("less entities received. got %v, want %v", len(entities), 2)
	}
	// if entities[0].Path != "/size" {
	// 	t.Fatalf("incorrect path. got: %s want: /size", entities[0].Path)
	// }
}
