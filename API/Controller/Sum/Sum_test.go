package Sum

import "testing"


func TestHitung(t *testing.T)  {
	
	emptyhidup, emptyMati := Sum("osama")
	if emptyhidup == 2 && emptyMati == 2 {
		t.Logf("succes text osama have %v vowel and %v consonant", emptyhidup, emptyMati)
	}else{
		t.Error("faild text osama must have 2 vowel and 2 constant")
	}

	emptylife, emptydead := Sum("omama")
	if emptylife == 2 && emptydead == 2 {
		t.Logf("succes text omama have %v vowel and %v consonant", emptyhidup, emptyMati)
	}else{
		t.Error("faild text omama must have 2 vowel and 2 constant")
	}
}