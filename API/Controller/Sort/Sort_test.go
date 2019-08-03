package Sort

import "testing"


func TestHitung(t *testing.T)  {
	
	SortOne := Sort("osama")
	if SortOne == "aaoms" {
		t.Logf("succes text osama has been sorted %v", SortOne)
	}else{
		t.Error("faild to sort should be aaoms")
	}

	SortTwo := Sort("omama")
	if SortTwo == "aaomm" {
		t.Logf("succes text omama has been sorted %v", SortTwo )
	}else{
		t.Error("faild to sort should be aaomm")
	}
}