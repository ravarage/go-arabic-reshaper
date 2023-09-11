package go_arabic_reshaper

import (
	"testing"
)

func TestArabic(t *testing.T) {
	reshaper := NewArabicReshaper()
	if reshaper.Reshape("السلام عليكم") == "ﺍﻟﺴﻠﺎﻣ ﻋﻠﻴﻜﻡ" {
		t.Log("OK")
	} else {
		t.Error("Not OK")
	}
}
func TestKurdish(t *testing.T) {
	Kurdish := NewArabicReshaper(ArabicReshaper{
		Language: "Kurdish",
	})
	if Kurdish.Reshape("ڕاڤیار سەربەست طاهر") == "ڕاﭬﯿﺎر ﺳ\uE000رﺑ\uE000ﺳﺖ ﻃﺎھر" {
		t.Log("OK")
	} else {
		t.Error("Not OK")
	}
}
func TestRemoveHarakat(t *testing.T) {
	RemoveHarakat := NewArabicReshaper(ArabicReshaper{
		Delete_harakat: true,
	})
	if RemoveHarakat.Reshape("فَتْحَة") == "ﻓﺘﺤﺓ" {
		t.Log("OK")
	} else {
		t.Error("Not OK")
	}

}
