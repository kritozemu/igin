package igin

import "testing"

func TestNextedGroup(t *testing.T) {
	r := NewEngine()
	v1 := r.Group("/v1")
	v2 := v1.Group("/v2")
	v3 := v2.Group("/v3")
	if v2.prefix != "/v1/v2" {
		t.Fatalf("v2.prefix != /v1/v2")
	}
	if v3.prefix != "/v1/v2/v3" {
		t.Fatalf("v3.prefix != /v1/v2/v3")
	}
}
