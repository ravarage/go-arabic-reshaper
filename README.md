# go-arabic-reshaper

A pure go library to handle Arabic text and it variant for rendering



###ToDO
 Markup : - [x] Reshape Arabic,Kurdish,Persian and Urdo
          - [x] Support Harakat
          - [x] Shift Harakat
          - [x] Remove Harakat
          - [x] Support Tatweel
          - [x] Support width JOINER
          - [x] Use Unshaped instead of isolated form
          - [ ] Support Ligatures
          - [ ] Make it one loop.
          - [ ] impove condintion







## Installation
```bash
go get  github.com/ravarage/go-arabic-reshaper
```
### Example

Using reshape arabic with default setting
```go

import (
	"github.com/ravarage/go-arabic-reshaper"
)

func main() {
	Reshaper := go_arabic_reshaper.NewArabicReshaper()
	reshapedtext := Reshaper.Reshape("السلام عليكم")
}

```

Using to reshape Kurdish letters, Recommended for Persian too
```go
import (
	"github.com/ravarage/go-arabic-reshaper"
)

func main() {
	Reshaper := go_arabic_reshaper.NewArabicReshaper(go_arabic_reshaper.ArabicReshaper{
		Language: "Kurdish",
	})
	reshapedtext := Reshaper.Reshape("ڕاڤیار")
}



```

Creating config, defualt are false in all cases
```go
go_arabic_reshaper.ArabicReshaper{
		Language:                         "Arabic", // `Arabic` is default Kurdish, and Arabic_V2 is suppurted
		Letters:                          nil,      //leave it be/ you can load your own letters array that match your case type map[rune][4]rune
		Delete_harakat:                   false,    // Whether to shift the Harakat (Tashkeel) one position so they appear correctly when string is reversed
		Shift_harakat_position:           false,    // Whether to delete the Tatweel (U+0640) before reshaping or not.
		Delete_tatweel:                   false,    // Whether to support ZWJ (U+200D) or not.
		Support_zwj:                      false,    //# Use unshaped form instead of isolated form.
		Use_unshaped_instead_of_isolated: false,    //# Use unshaped form instead of isolated form.

		Support_ligatures: false, //# Whether to use ligatures or not. # Serves as a shortcut to disable all ligatures.not impliment yet

	}
```

