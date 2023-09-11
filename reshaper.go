package go_arabic_reshaper

/*
Go Arabic Reshaper is project to reshaping Arabic text to be used in rendering images and pdf files
it inspired by python arabic reshaper, but it is not a port of it
usage example:
reshaper := go_arabic_reshaper.NewArabicReshaper()
reshaper.Reshape("السلام عليكم")
that will return reshaped text
*/

// Define ArabicReshaper struct to hold reshaper options
type ArabicReshaper struct {
	Language                         string           // `Arabic` is default and recommended to work in most of the cases and supports (Arabic, Urdu and Farsi)
	Letters                          map[rune][4]rune //leave it be/ you can load your own letters array that match your case type map[rune][4]rune
	Delete_harakat                   bool             // Whether to shift the Harakat (Tashkeel) one position so they appear correctly when string is reversed
	Shift_harakat_position           bool             // Whether to delete the Tatweel (U+0640) before reshaping or not.
	Delete_tatweel                   bool             // Whether to support ZWJ (U+200D) or not.
	Support_zwj                      bool             //# Use unshaped form instead of isolated form.
	Use_unshaped_instead_of_isolated bool             //# Use unshaped form instead of isolated form.

	Support_ligatures bool //# Whether to use ligatures or not. # Serves as a shortcut to disable all ligatures.

}

// NewArabicReshaper initializes a new ArabicReshaper instance
func NewArabicReshaper(arabicReshaper ...ArabicReshaper) *ArabicReshaper {

	//var letters map[rune][4]rune
	if arabicReshaper == nil {
		return &ArabicReshaper{

			Language:                         "Arabic",
			Letters:                          LETTERS_ARABIC,
			Delete_harakat:                   true,
			Shift_harakat_position:           false,
			Delete_tatweel:                   true,
			Support_zwj:                      true,
			Use_unshaped_instead_of_isolated: false,
			Support_ligatures:                false,
		}

	}

	switch arabicReshaper[0].Language {
	case "ArabicV2":
		arabicReshaper[0].Letters = LETTERS_ARABIC_V2
	case "Kurdish":
		arabicReshaper[0].Letters = LETTERS_KURDISH
	default:
		arabicReshaper[0].Letters = LETTERS_ARABIC
	}

	return &ArabicReshaper{
		Language:                         arabicReshaper[0].Language,
		Letters:                          arabicReshaper[0].Letters,
		Delete_harakat:                   arabicReshaper[0].Delete_harakat,
		Shift_harakat_position:           arabicReshaper[0].Shift_harakat_position,
		Delete_tatweel:                   arabicReshaper[0].Delete_tatweel,
		Support_zwj:                      arabicReshaper[0].Support_zwj,
		Use_unshaped_instead_of_isolated: arabicReshaper[0].Use_unshaped_instead_of_isolated,
		Support_ligatures:                arabicReshaper[0].Support_ligatures,
	}
}

// struck to hold output of reshaper before convert it to string
type OutputStuck struct {
	FORM   int  // save shape form
	LETTER rune // save letter
}

// Reshape reshapes the input text based on the Arabic reshaping rules
// need major refactor here, by merging two for loops to one will shave time by half
// removing checking for harakat either removing it, or shifting it will shave time this option is good for Kurdish and Persian
// removing more if mean more performance but arabic writing is complex
// doing this maybe it will match gAarbic performance as current solution is x4 slower than gArabic
func (a *ArabicReshaper) Reshape(text string) string {
	var outputStuck []OutputStuck
	for _, runeValue := range text {

		//check if delete harakat is true
		if a.Delete_harakat {
			if isHarakat(runeValue) {
				continue
			}
		}
		if !a.Delete_harakat {
			if isHarakat(runeValue) {

				outputStuck = append(outputStuck, OutputStuck{
					FORM:   HARAKAT,
					LETTER: runeValue,
				})

				continue
			}
		}
		if a.Delete_tatweel {
			if isTatweel(runeValue, a.Letters) {
				continue
			}
		}
		if !a.Support_zwj {
			if isZWJ(runeValue, a.Letters) {
				continue
			}
		}
		if _, ok := a.Letters[runeValue]; !ok {
			if !isHarakat(runeValue) {

				outputStuck = append(outputStuck, OutputStuck{
					FORM:   NotSupported,
					LETTER: runeValue,
				})
				continue
			}
		}
		if len(outputStuck) == 0 {
			if isInitial(runeValue, a.Letters) {
				outputStuck = append(outputStuck, OutputStuck{
					FORM:   INITIAL,
					LETTER: runeValue,
				})
				continue
			}
			outputStuck = append(outputStuck, OutputStuck{
				FORM:   ISOLATED,
				LETTER: runeValue,
			})
			continue
		}
		previosLetter := outputStuck[len(outputStuck)-1]
		harakatfount := false
		if previosLetter.FORM == HARAKAT {
			if len(outputStuck) > 1 {
				harakatfount = true
				previosLetter = outputStuck[len(outputStuck)-2]

			} else {
				previosLetter = OutputStuck{
					FORM:   ISOLATED,
					LETTER: runeValue,
				}
			}
		}
		if previosLetter.FORM == ISOLATED {
			if isInitial(runeValue, a.Letters) {
				//outputStuck[len(outputStuck)-1].FORM =
				outputStuck = append(outputStuck, OutputStuck{
					FORM:   INITIAL,
					LETTER: runeValue,
				})
				continue
			}

			outputStuck = append(outputStuck, OutputStuck{
				FORM:   ISOLATED,
				LETTER: runeValue,
			})
			continue
		}

		if previosLetter.FORM == NotSupported {
			if !isInitial(runeValue, a.Letters) {
				outputStuck = append(outputStuck, OutputStuck{
					ISOLATED, runeValue,
				})
			}
			outputStuck = append(outputStuck, OutputStuck{
				INITIAL, runeValue,
			})
			continue
		}

		if previosLetter.FORM == ISOLATED {

			if !isInitial(runeValue, a.Letters) {
				outputStuck = append(outputStuck, OutputStuck{
					ISOLATED, runeValue,
				})
				continue
			}
			outputStuck = append(outputStuck, OutputStuck{
				INITIAL, runeValue,
			})
			continue
		}
		if previosLetter.FORM == INITIAL {

			if isFinal(runeValue, a.Letters) {
				outputStuck = append(outputStuck, OutputStuck{
					FINAL, runeValue,
				})
				continue
			}
			outputStuck = append(outputStuck, OutputStuck{
				ISOLATED, runeValue,
			})
			continue

		}
		if previosLetter.FORM == MEDIAL {
			if isMedial(runeValue, a.Letters) {
				outputStuck = append(outputStuck, OutputStuck{
					MEDIAL, runeValue,
				})
				continue
			}
			if isFinal(runeValue, a.Letters) {
				outputStuck = append(outputStuck, OutputStuck{
					FINAL, runeValue,
				})
				continue
			}
			outputStuck = append(outputStuck, OutputStuck{
				ISOLATED, runeValue,
			})
			continue

		}
		if previosLetter.FORM == FINAL {
			if isMedial(previosLetter.LETTER, a.Letters) {
				if harakatfount {
					outputStuck[len(outputStuck)-2].FORM = MEDIAL
				} else {
					outputStuck[len(outputStuck)-1].FORM = MEDIAL
				}

			} else {
				if isInitial(runeValue, a.Letters) {
					outputStuck = append(outputStuck, OutputStuck{
						INITIAL, runeValue,
					})
				} else {
					outputStuck = append(outputStuck, OutputStuck{
						ISOLATED, runeValue,
					})

				}
				continue
			}
			if isFinal(runeValue, a.Letters) {
				outputStuck = append(outputStuck, OutputStuck{
					FINAL, runeValue,
				})
				continue
			}
			if isInitial(runeValue, a.Letters) {
				outputStuck = append(outputStuck, OutputStuck{
					INITIAL, runeValue,
				})
				continue
			}
			outputStuck = append(outputStuck, OutputStuck{
				ISOLATED, runeValue,
			})
			continue
		}
	}
	var output string
	for i := 0; i < len(outputStuck); i++ {
		if a.Use_unshaped_instead_of_isolated {
			if outputStuck[i].FORM == ISOLATED {
				output += string(outputStuck[i].LETTER)
			}
		}
		if outputStuck[i].FORM == NotSupported {
			output += string(outputStuck[i].LETTER)
			continue
		}
		if outputStuck[i].FORM == HARAKAT {

			output += string(outputStuck[i].LETTER)
			if a.Shift_harakat_position {
				oldletter := outputStuck[i-1]
				//remove last letter from output
				output = output[:len(output)-1]
				output += string(outputStuck[i].LETTER)
				output += string(a.Letters[oldletter.LETTER][oldletter.FORM])

			}
			output += string(outputStuck[i].LETTER)
			continue
		}
		if i == len(outputStuck)-1 {
			if isFinal(outputStuck[i].LETTER, a.Letters) { // this is to make sure last letter is final not median
				if len(outputStuck) > 1 {

					//if before final is final then make it isolated
					if isFinal(outputStuck[i-1].LETTER, a.Letters) {
						outputStuck[i].FORM = ISOLATED
					} else {
						outputStuck[i].FORM = FINAL
					}
				} else {
					outputStuck[i].FORM = ISOLATED
				}
			}
		}
		output += string(a.Letters[outputStuck[i].LETTER][outputStuck[i].FORM])

	}

	return output

}
