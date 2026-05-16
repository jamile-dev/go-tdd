package integer

import "testing"

func TestAdd(t *testing.T) {
	sum_result := Add(2, 2)
	expected := 4
	if sum_result != expected {
		t.Errorf("expected '%d', result '%d'", expected, sum_result)
	}
}
