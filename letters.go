package go_arabic_reshaper

// Define constants for letter forms
const (
	UNSHAPED     = 255
	ISOLATED     = 0
	INITIAL      = 1
	MEDIAL       = 2
	FINAL        = 3
	HARAKAT      = 5
	NotSupported = -1
)

// Tatweel and  zero width joiner is common in arabic text
const (
	TATWEEL = '\u0640'
	ZWJ     = '\u200D'
)

// harakat map is used to remove harakat from text, is faster than regex, for loop and contain for this size
var HARAKATMAP = map[rune]rune{
	'\u064B': '\u064B', // ARABIC FATHATAN
	'\u064C': '\u064C', // ARABIC DAMMATAN
	'\u064D': '\u064D', // ARABIC KASRATAN
	'\u064E': '\u064E', // ARABIC FATHA
	'\u064F': '\u064F', // ARABIC DAMMA
	'\u0650': '\u0650', // ARABIC KASRA
	'\u0651': '\u0651', // ARABIC SHADDA
	'\u0652': '\u0652', // ARABIC SUKUN
	'\u0653': '\u0653', // ARABIC MADDAH ABOVE
	'\u0654': '\u0654', // ARABIC HAMZA ABOVE
	'\u0655': '\u0655', // ARABIC HAMZA BELOW
	'\u0656': '\u0656', // ARABIC SUBSCRIPT ALEF
	'\u0657': '\u0657', // ARABIC INVERTED DAMMA
	'\u0658': '\u0658', // ARABIC MARK NOON GHUNNA
	'\u0659': '\u0659', // ARABIC ZWARAKAY
	'\u065A': '\u065A', // ARABIC VOWEL SIGN SMALL V ABOVE
	'\u065B': '\u065B', // ARABIC VOWEL SIGN INVERTED SMALL V ABOVE
	'\u065C': '\u065C', // ARABIC VOWEL SIGN DOT BELOW
	'\u065D': '\u065D', // ARABIC REVERSED DAMMA
	'\u065E': '\u065E', // ARABIC FATHA WITH TWO DOTS
	'\u065F': '\u065F', // ARABIC WAVY HAMZA BELOW
	'\u0670': '\u0670', // ARABIC LETTER SUPERSCRIPT ALEF
	'\u06D6': '\u06D6', // ARABIC SMALL HIGH LIGATURE SAD WITH LAM WITH ALEF MAKSURA
	'\u06D7': '\u06D7', // ARABIC SMALL HIGH LIGATURE QAF WITH LAM WITH ALEF MAKSURA
	'\u06D8': '\u06D8', // ARABIC SMALL HIGH MEEM INITIAL FORM
	'\u06D9': '\u06D9', // ARABIC SMALL HIGH LAM ALEF
	'\u06DA': '\u06DA', // ARABIC SMALL HIGH JEEM
	'\u06DB': '\u06DB', // ARABIC SMALL HIGH THREE DOTS
	'\u06DC': '\u06DC', // ARABIC SMALL HIGH SEEN
	'\u06DF': '\u06DF', // ARABIC SMALL HIGH ROUNDED ZERO
	'\u06E0': '\u06E0', // ARABIC SMALL HIGH UPRIGHT RECTANGULAR ZERO
	'\u06E1': '\u06E1', // ARABIC SMALL HIGH DOTLESS HEAD OF KHAH
	'\u06E2': '\u06E2', // ARABIC SMALL HIGH MEEM ISOLATED FORM
	'\u06E3': '\u06E3', // ARABIC SMALL LOW SEEN
	'\u06E4': '\u06E4', // ARABIC SMALL HIGH MADDA
	'\u06E7': '\u06E7', // ARABIC SMALL HIGH YEH
	'\u06E8': '\u06E8', // ARABIC SMALL HIGH NOON
	'\u06EA': '\u06EA', // ARABIC EMPTY CENTRE LOW STOP
	'\u06EB': '\u06EB', // ARABIC EMPTY CENTRE HIGH STOP
	'\u06EC': '\u06EC', // ARABIC ROUNDED HIGH STOP WITH FILLED CENTRE
	'\u06ED': '\u06ED', // ARABIC SMALL LOW MEEM
	'\u08D3': '\u08D3', // ARABIC SMALL LOW WAW
	'\u08D4': '\u08D4', // ARABIC SMALL HIGH WORD AR-RUB
	'\u08D5': '\u08D5', // ARABIC SMALL HIGH SAD
	'\u08D6': '\u08D6', // ARABIC SMALL HIGH AIN
	'\u08D7': '\u08D7', // ARABIC SMALL HIGH QAF
	'\u08D8': '\u08D8', // ARABIC SMALL HIGH NOON WITH KASRA
	'\u08D9': '\u08D9', // ARABIC SMALL LOW NOON WITH KASRA
	'\u08DA': '\u08DA', // ARABIC SMALL HIGH WORD ATH-THALATHA
	'\u08DB': '\u08DB', // ARABIC SMALL HIGH WORD AS-SAJDA
	'\u08DC': '\u08DC', // ARABIC SMALL HIGH WORD AN-NISF
	'\u08DD': '\u08DD', // ARABIC SMALL HIGH WORD SAKTA
	'\u08DE': '\u08DE', // ARABIC SMALL HIGH WORD QIF
	'\u08DF': '\u08DF', // ARABIC SMALL HIGH WORD WAQFA
	'\u08E0': '\u08E0', // ARABIC SMALL HIGH FOOTNOTE MARKER
	'\u08E1': '\u08E1', // ARABIC SMALL HIGH SIGN SAFHA
	'\u08E3': '\u08E3', // ARABIC TURNED DAMMA BELOW
	'\u08E4': '\u08E4', // ARABIC CURLY FATHA
	'\u08E5': '\u08E5', // ARABIC CURLY DAMMA
	'\u08E6': '\u08E6', // ARABIC CURLY KASRA
	'\u08E7': '\u08E7', // ARABIC CURLY FATHATAN
	'\u08E8': '\u08E8', // ARABIC CURLY DAMMATAN
	'\u08E9': '\u08E9', // ARABIC CURLY KASRATAN
	'\u08EA': '\u08EA', // ARABIC TONE ONE DOT ABOVE
	'\u08EB': '\u08EB', // ARABIC TONE TWO DOTS ABOVE
	'\u08EC': '\u08EC', // ARABIC TONE LOOP ABOVE
	'\u08ED': '\u08ED', // ARABIC TONE ONE DOT BELOW
	'\u08EE': '\u08EE', // ARABIC TONE TWO DOTS BELOW
	'\u08EF': '\u08EF', // ARABIC TONE LOOP BELOW
	'\u08F0': '\u08F0', // ARABIC OPEN FATHATAN
	'\u08F1': '\u08F1', // ARABIC OPEN DAMMATAN
	'\u08F2': '\u08F2', // ARABIC OPEN KASRATAN
	'\u08F3': '\u08F3', // ARABIC SMALL HIGH WAW
	'\u08F4': '\u08F4', // ARABIC FATHA WITH RING
	'\u08F5': '\u08F5', // ARABIC FATHA WITH DOT ABOVE
	'\u08F6': '\u08F6', // ARABIC KASRA WITH DOT BELOW
	'\u08F7': '\u08F7', // ARABIC LEFT ARROWHEAD ABOVE
	'\u08F8': '\u08F8', // ARABIC RIGHT ARROWHEAD ABOVE
	'\u08F9': '\u08F9', // ARABIC LEFT ARROWHEAD BELOW
	'\u08FA': '\u08FA', // ARABIC RIGHT ARROWHEAD BELOW
	'\u08FB': '\u08FB', // ARABIC DOUBLE RIGHT ARROWHEAD ABOVE
	'\u08FC': '\u08FC', // ARABIC DOUBLE RIGHT ARROWHEAD ABOVE WITH DOT
	'\u08FD': '\u08FD', // ARABIC RIGHT ARROWHEAD ABOVE WITH DOT
	'\u08FE': '\u08FE', // ARABIC DAMMA WITH DOT
	'\u08FF': '\u08FF', // ARABIC MARK SIDEWAYS NOON GHUNNA

}

// arabic letters
var LETTERS_ARABIC = map[rune][4]rune{
	'\u0621': {'\uFE80', 0, 0, 0},                      // ARABIC LETTER HAMZA
	'\u0622': {'\uFE81', 0, 0, '\uFE82'},               // ARABIC LETTER ALEF WITH MADDA ABOVE
	'\u0623': {'\uFE83', 0, 0, '\uFE84'},               // ARABIC LETTER ALEF WITH HAMZA ABOVE
	'\u0624': {'\uFE85', 0, 0, '\uFE86'},               // ARABIC LETTER WAW WITH HAMZA ABOVE
	'\u0625': {'\uFE87', 0, 0, '\uFE88'},               // ARABIC LETTER ALEF WITH HAMZA BELOW
	'\u0626': {'\uFE89', '\uFE8B', '\uFE8C', '\uFE8A'}, // ARABIC LETTER YEH WITH HAMZA ABOVE
	'\u0627': {'\uFE8D', 0, 0, '\uFE8E'},               // ARABIC LETTER ALEF
	'\u0628': {'\uFE8F', '\uFE91', '\uFE92', '\uFE90'}, // ARABIC LETTER BEH
	'\u0629': {'\uFE93', 0, 0, '\uFE94'},               // ARABIC LETTER TEH MARBUTA
	'\u062A': {'\uFE95', '\uFE97', '\uFE98', '\uFE96'}, // ARABIC LETTER TEH
	'\u062B': {'\uFE99', '\uFE9B', '\uFE9C', '\uFE9A'}, // ARABIC LETTER THEH
	'\u062C': {'\uFE9D', '\uFE9F', '\uFEA0', '\uFE9E'}, // ARABIC LETTER JEEM
	'\u062D': {'\uFEA1', '\uFEA3', '\uFEA4', '\uFEA2'}, // ARABIC LETTER HAH
	'\u062E': {'\uFEA5', '\uFEA7', '\uFEA8', '\uFEA6'}, // ARABIC LETTER KHAH
	'\u062F': {'\uFEA9', 0, 0, '\uFEAA'},               // ARABIC LETTER DAL
	'\u0630': {'\uFEAB', 0, 0, '\uFEAC'},               // ARABIC LETTER THAL
	'\u0631': {'\uFEAD', 0, 0, '\uFEAE'},               // ARABIC LETTER REH
	'\u0632': {'\uFEAF', 0, 0, '\uFEB0'},               // ARABIC LETTER ZAIN
	'\u0633': {'\uFEB1', '\uFEB3', '\uFEB4', '\uFEB2'}, // ARABIC LETTER SEEN
	'\u0634': {'\uFEB5', '\uFEB7', '\uFEB8', '\uFEB6'}, // ARABIC LETTER SHEEN
	'\u0635': {'\uFEB9', '\uFEBB', '\uFEBC', '\uFEBA'}, // ARABIC LETTER SAD
	'\u0636': {'\uFEBD', '\uFEBF', '\uFEC0', '\uFEBE'}, // ARABIC LETTER DAD
	'\u0637': {'\uFEC1', '\uFEC3', '\uFEC4', '\uFEC2'}, // ARABIC LETTER TAH
	'\u0638': {'\uFEC5', '\uFEC7', '\uFEC8', '\uFEC6'}, // ARABIC LETTER ZAH
	'\u0639': {'\uFEC9', '\uFECB', '\uFECC', '\uFECA'}, // ARABIC LETTER AIN
	'\u063A': {'\uFECD', '\uFECF', '\uFED0', '\uFECE'}, // ARABIC LETTER GHAIN
	TATWEEL:  {TATWEEL, TATWEEL, TATWEEL, TATWEEL},     // ARABIC TATWEEL
	'\u0641': {'\uFED1', '\uFED3', '\uFED4', '\uFED2'}, // ARABIC LETTER FEH
	'\u0642': {'\uFED5', '\uFED7', '\uFED8', '\uFED6'}, // ARABIC LETTER QAF
	'\u0643': {'\uFED9', '\uFEDB', '\uFEDC', '\uFEDA'}, // ARABIC LETTER KAF
	'\u0644': {'\uFEDD', '\uFEDF', '\uFEE0', '\uFEDE'}, // ARABIC LETTER LAM
	'\u0645': {'\uFEE1', '\uFEE3', '\uFEE4', '\uFEE2'}, // ARABIC LETTER MEEM
	'\u0646': {'\uFEE5', '\uFEE7', '\uFEE8', '\uFEE6'}, // ARABIC LETTER NOON
	'\u0647': {'\uFEE9', '\uFEEB', '\uFEEC', '\uFEEA'}, // ARABIC LETTER HEH
	'\u0648': {'\uFEED', 0, 0, '\uFEEE'},               // ARABIC LETTER WAW
	'\u0649': {'\uFEEF', '\uFBE8', '\uFBE9', '\uFEF0'}, // ARABIC LETTER (UIGHUR KAZAKH KIRGHIZ)? ALEF MAKSURA
	'\u064A': {'\uFEF1', '\uFEF3', '\uFEF4', '\uFEF2'}, // ARABIC LETTER YEH
	'\u0671': {'\uFB50', 0, 0, '\uFB51'},               // ARABIC LETTER ALEF WASLA
	'\u0677': {'\uFBDD', 0, 0, 0},                      // ARABIC LETTER U WITH HAMZA ABOVE
	'\u0679': {'\uFB66', '\uFB68', '\uFB69', '\uFB67'}, // ARABIC LETTER TTEH
	'\u067A': {'\uFB5E', '\uFB60', '\uFB61', '\uFB5F'}, // ARABIC LETTER TTEHEH
	'\u067B': {'\uFB52', '\uFB54', '\uFB55', '\uFB53'}, // ARABIC LETTER BEEH
	'\u067E': {'\uFB56', '\uFB58', '\uFB59', '\uFB57'}, // ARABIC LETTER PEH
	'\u067F': {'\uFB62', '\uFB64', '\uFB65', '\uFB63'}, // ARABIC LETTER TEHEH
	'\u0680': {'\uFB5A', '\uFB5C', '\uFB5D', '\uFB5B'}, // ARABIC LETTER BEHEH
	'\u0683': {'\uFB76', '\uFB78', '\uFB79', '\uFB77'}, // ARABIC LETTER NYEH
	'\u0684': {'\uFB72', '\uFB74', '\uFB75', '\uFB73'}, // ARABIC LETTER DYEH
	'\u0686': {'\uFB7A', '\uFB7C', '\uFB7D', '\uFB7B'}, // ARABIC LETTER TCHEH
	'\u0687': {'\uFB7E', '\uFB80', '\uFB81', '\uFB7F'}, // ARABIC LETTER TCHEHEH
	'\u0688': {'\uFB88', 0, 0, '\uFB89'},               // ARABIC LETTER DDAL
	'\u068C': {'\uFB84', 0, 0, '\uFB85'},               // ARABIC LETTER DAHAL
	'\u068D': {'\uFB82', 0, 0, '\uFB83'},               // ARABIC LETTER DDAHAL
	'\u068E': {'\uFB86', 0, 0, '\uFB87'},               // ARABIC LETTER DUL
	'\u0691': {'\uFB8C', 0, 0, '\uFB8D'},               // ARABIC LETTER RREH
	'\u0698': {'\uFB8A', 0, 0, '\uFB8B'},               // ARABIC LETTER JEH
	'\u06A4': {'\uFB6A', '\uFB6C', '\uFB6D', '\uFB6B'}, // ARABIC LETTER VEH
	'\u06A6': {'\uFB6E', '\uFB70', '\uFB71', '\uFB6F'}, // ARABIC LETTER PEHEH
	'\u06A9': {'\uFB8E', '\uFB90', '\uFB91', '\uFB8F'}, // ARABIC LETTER KEHEH
	'\u06AD': {'\uFBD3', '\uFBD5', '\uFBD6', '\uFBD4'}, // ARABIC LETTER NG
	'\u06AF': {'\uFB92', '\uFB94', '\uFB95', '\uFB93'}, // ARABIC LETTER GAF
	'\u06B1': {'\uFB9A', '\uFB9C', '\uFB9D', '\uFB9B'}, // ARABIC LETTER NGOEH
	'\u06B3': {'\uFB96', '\uFB98', '\uFB99', '\uFB97'}, // ARABIC LETTER GUEH
	'\u06BA': {'\uFB9E', 0, 0, '\uFB9F'},               // ARABIC LETTER NOON GHUNNA
	'\u06BB': {'\uFBA0', '\uFBA2', '\uFBA3', '\uFBA1'}, // ARABIC LETTER RNOON
	'\u06BE': {'\uFBAA', '\uFBAC', '\uFBAD', '\uFBAB'}, // ARABIC LETTER HEH DOACHASHMEE
	'\u06C0': {'\uFBA4', 0, 0, '\uFBA5'},               // ARABIC LETTER HEH WITH YEH ABOVE
	'\u06C1': {'\uFBA6', '\uFBA8', '\uFBA9', '\uFBA7'}, // ARABIC LETTER HEH GOAL
	'\u06C5': {'\uFBE0', 0, 0, '\uFBE1'},               // ARABIC LETTER KIRGHIZ OE
	'\u06C6': {'\uFBD9', 0, 0, '\uFBDA'},               // ARABIC LETTER OE
	'\u06C7': {'\uFBD7', 0, 0, '\uFBD8'},               // ARABIC LETTER U
	'\u06C8': {'\uFBDB', 0, 0, '\uFBDC'},               // ARABIC LETTER YU
	'\u06C9': {'\uFBE2', 0, 0, '\uFBE3'},               // ARABIC LETTER KIRGHIZ YU
	'\u06CB': {'\uFBDE', 0, 0, '\uFBDF'},               // ARABIC LETTER VE
	'\u06CC': {'\uFBFC', '\uFBFE', '\uFBFF', '\uFBFD'}, // ARABIC LETTER FARSI YEH
	'\u06D0': {'\uFBE4', '\uFBE6', '\uFBE7', '\uFBE5'}, // ARABIC LETTER E
	'\u06D2': {'\uFBAE', 0, 0, '\uFBAF'},               // ARABIC LETTER YEH BARREE
	'\u06D3': {'\uFBB0', 0, 0, '\uFBB1'},               // ARABIC LETTER YEH BARREE WITH HAMZA ABOVE
	ZWJ:      {ZWJ, ZWJ, ZWJ, ZWJ},                     // ZERO WIDTH JOINER

}

// currently same as v1 but its place holder before converting it from python
var LETTERS_ARABIC_V2 = map[rune][4]rune{
	'\u0621': {'\uFE80', 0, 0, 0},                      // ARABIC LETTER HAMZA
	'\u0622': {'\uFE81', 0, 0, '\uFE82'},               // ARABIC LETTER ALEF WITH MADDA ABOVE
	'\u0623': {'\uFE83', 0, 0, '\uFE84'},               // ARABIC LETTER ALEF WITH HAMZA ABOVE
	'\u0624': {'\uFE85', 0, 0, '\uFE86'},               // ARABIC LETTER WAW WITH HAMZA ABOVE
	'\u0625': {'\uFE87', 0, 0, '\uFE88'},               // ARABIC LETTER ALEF WITH HAMZA BELOW
	'\u0626': {'\uFE89', '\uFE8B', '\uFE8C', '\uFE8A'}, // ARABIC LETTER YEH WITH HAMZA ABOVE
	'\u0627': {'\uFE8D', 0, 0, '\uFE8E'},               // ARABIC LETTER ALEF
	'\u0628': {'\uFE8F', '\uFE91', '\uFE92', '\uFE90'}, // ARABIC LETTER BEH
	'\u0629': {'\uFE93', 0, 0, '\uFE94'},               // ARABIC LETTER TEH MARBUTA
	'\u062A': {'\uFE95', '\uFE97', '\uFE98', '\uFE96'}, // ARABIC LETTER TEH
	'\u062B': {'\uFE99', '\uFE9B', '\uFE9C', '\uFE9A'}, // ARABIC LETTER THEH
	'\u062C': {'\uFE9D', '\uFE9F', '\uFEA0', '\uFE9E'}, // ARABIC LETTER JEEM
	'\u062D': {'\uFEA1', '\uFEA3', '\uFEA4', '\uFEA2'}, // ARABIC LETTER HAH
	'\u062E': {'\uFEA5', '\uFEA7', '\uFEA8', '\uFEA6'}, // ARABIC LETTER KHAH
	'\u062F': {'\uFEA9', 0, 0, '\uFEAA'},               // ARABIC LETTER DAL
	'\u0630': {'\uFEAB', 0, 0, '\uFEAC'},               // ARABIC LETTER THAL
	'\u0631': {'\uFEAD', 0, 0, '\uFEAE'},               // ARABIC LETTER REH
	'\u0632': {'\uFEAF', 0, 0, '\uFEB0'},               // ARABIC LETTER ZAIN
	'\u0633': {'\uFEB1', '\uFEB3', '\uFEB4', '\uFEB2'}, // ARABIC LETTER SEEN
	'\u0634': {'\uFEB5', '\uFEB7', '\uFEB8', '\uFEB6'}, // ARABIC LETTER SHEEN
	'\u0635': {'\uFEB9', '\uFEBB', '\uFEBC', '\uFEBA'}, // ARABIC LETTER SAD
	'\u0636': {'\uFEBD', '\uFEBF', '\uFEC0', '\uFEBE'}, // ARABIC LETTER DAD
	'\u0637': {'\uFEC1', '\uFEC3', '\uFEC4', '\uFEC2'}, // ARABIC LETTER TAH
	'\u0638': {'\uFEC5', '\uFEC7', '\uFEC8', '\uFEC6'}, // ARABIC LETTER ZAH
	'\u0639': {'\uFEC9', '\uFECB', '\uFECC', '\uFECA'}, // ARABIC LETTER AIN
	'\u063A': {'\uFECD', '\uFECF', '\uFED0', '\uFECE'}, // ARABIC LETTER GHAIN
	TATWEEL:  {TATWEEL, TATWEEL, TATWEEL, TATWEEL},     // ARABIC TATWEEL
	'\u0641': {'\uFED1', '\uFED3', '\uFED4', '\uFED2'}, // ARABIC LETTER FEH
	'\u0642': {'\uFED5', '\uFED7', '\uFED8', '\uFED6'}, // ARABIC LETTER QAF
	'\u0643': {'\uFED9', '\uFEDB', '\uFEDC', '\uFEDA'}, // ARABIC LETTER KAF
	'\u0644': {'\uFEDD', '\uFEDF', '\uFEE0', '\uFEDE'}, // ARABIC LETTER LAM
	'\u0645': {'\uFEE1', '\uFEE3', '\uFEE4', '\uFEE2'}, // ARABIC LETTER MEEM
	'\u0646': {'\uFEE5', '\uFEE7', '\uFEE8', '\uFEE6'}, // ARABIC LETTER NOON
	'\u0647': {'\uFEE9', '\uFEEB', '\uFEEC', '\uFEEA'}, // ARABIC LETTER HEH
	'\u0648': {'\uFEED', 0, 0, '\uFEEE'},               // ARABIC LETTER WAW
	'\u0649': {'\uFEEF', '\uFBE8', '\uFBE9', '\uFEF0'}, // ARABIC LETTER (UIGHUR KAZAKH KIRGHIZ)? ALEF MAKSURA
	'\u064A': {'\uFEF1', '\uFEF3', '\uFEF4', '\uFEF2'}, // ARABIC LETTER YEH
	'\u0671': {'\uFB50', 0, 0, '\uFB51'},               // ARABIC LETTER ALEF WASLA
	'\u0677': {'\uFBDD', 0, 0, 0},                      // ARABIC LETTER U WITH HAMZA ABOVE
	'\u0679': {'\uFB66', '\uFB68', '\uFB69', '\uFB67'}, // ARABIC LETTER TTEH
	'\u067A': {'\uFB5E', '\uFB60', '\uFB61', '\uFB5F'}, // ARABIC LETTER TTEHEH
	'\u067B': {'\uFB52', '\uFB54', '\uFB55', '\uFB53'}, // ARABIC LETTER BEEH
	'\u067E': {'\uFB56', '\uFB58', '\uFB59', '\uFB57'}, // ARABIC LETTER PEH
	'\u067F': {'\uFB62', '\uFB64', '\uFB65', '\uFB63'}, // ARABIC LETTER TEHEH
	'\u0680': {'\uFB5A', '\uFB5C', '\uFB5D', '\uFB5B'}, // ARABIC LETTER BEHEH
	'\u0683': {'\uFB76', '\uFB78', '\uFB79', '\uFB77'}, // ARABIC LETTER NYEH
	'\u0684': {'\uFB72', '\uFB74', '\uFB75', '\uFB73'}, // ARABIC LETTER DYEH
	'\u0686': {'\uFB7A', '\uFB7C', '\uFB7D', '\uFB7B'}, // ARABIC LETTER TCHEH
	'\u0687': {'\uFB7E', '\uFB80', '\uFB81', '\uFB7F'}, // ARABIC LETTER TCHEHEH
	'\u0688': {'\uFB88', 0, 0, '\uFB89'},               // ARABIC LETTER DDAL
	'\u068C': {'\uFB84', 0, 0, '\uFB85'},               // ARABIC LETTER DAHAL
	'\u068D': {'\uFB82', 0, 0, '\uFB83'},               // ARABIC LETTER DDAHAL
	'\u068E': {'\uFB86', 0, 0, '\uFB87'},               // ARABIC LETTER DUL
	'\u0691': {'\uFB8C', 0, 0, '\uFB8D'},               // ARABIC LETTER RREH
	'\u0698': {'\uFB8A', 0, 0, '\uFB8B'},               // ARABIC LETTER JEH
	'\u06A4': {'\uFB6A', '\uFB6C', '\uFB6D', '\uFB6B'}, // ARABIC LETTER VEH
	'\u06A6': {'\uFB6E', '\uFB70', '\uFB71', '\uFB6F'}, // ARABIC LETTER PEHEH
	'\u06A9': {'\uFB8E', '\uFB90', '\uFB91', '\uFB8F'}, // ARABIC LETTER KEHEH
	'\u06AD': {'\uFBD3', '\uFBD5', '\uFBD6', '\uFBD4'}, // ARABIC LETTER NG
	'\u06AF': {'\uFB92', '\uFB94', '\uFB95', '\uFB93'}, // ARABIC LETTER GAF
	'\u06B1': {'\uFB9A', '\uFB9C', '\uFB9D', '\uFB9B'}, // ARABIC LETTER NGOEH
	'\u06B3': {'\uFB96', '\uFB98', '\uFB99', '\uFB97'}, // ARABIC LETTER GUEH
	'\u06BA': {'\uFB9E', 0, 0, '\uFB9F'},               // ARABIC LETTER NOON GHUNNA
	'\u06BB': {'\uFBA0', '\uFBA2', '\uFBA3', '\uFBA1'}, // ARABIC LETTER RNOON
	'\u06BE': {'\uFBAA', '\uFBAC', '\uFBAD', '\uFBAB'}, // ARABIC LETTER HEH DOACHASHMEE
	'\u06C0': {'\uFBA4', 0, 0, '\uFBA5'},               // ARABIC LETTER HEH WITH YEH ABOVE
	'\u06C1': {'\uFBA6', '\uFBA8', '\uFBA9', '\uFBA7'}, // ARABIC LETTER HEH GOAL
	'\u06C5': {'\uFBE0', 0, 0, '\uFBE1'},               // ARABIC LETTER KIRGHIZ OE
	'\u06C6': {'\uFBD9', 0, 0, '\uFBDA'},               // ARABIC LETTER OE
	'\u06C7': {'\uFBD7', 0, 0, '\uFBD8'},               // ARABIC LETTER U
	'\u06C8': {'\uFBDB', 0, 0, '\uFBDC'},               // ARABIC LETTER YU
	'\u06C9': {'\uFBE2', 0, 0, '\uFBE3'},               // ARABIC LETTER KIRGHIZ YU
	'\u06CB': {'\uFBDE', 0, 0, '\uFBDF'},               // ARABIC LETTER VE
	'\u06CC': {'\uFBFC', '\uFBFE', '\uFBFF', '\uFBFD'}, // ARABIC LETTER FARSI YEH
	'\u06D0': {'\uFBE4', '\uFBE6', '\uFBE7', '\uFBE5'}, // ARABIC LETTER E
	'\u06D2': {'\uFBAE', 0, 0, '\uFBAF'},               // ARABIC LETTER YEH BARREE
	'\u06D3': {'\uFBB0', 0, 0, '\uFBB1'},               // ARABIC LETTER YEH BARREE WITH HAMZA ABOVE
	ZWJ:      {ZWJ, ZWJ, ZWJ, ZWJ},                     // ZERO WIDTH JOINER

}

// LETTERS_KURDISH is a map of Kurdish letters to their Unicode representations.
var LETTERS_KURDISH = map[rune][4]rune{
	'\u0621': {'\uFE80', 0, 0, 0},                      // ARABIC LETTER HAMZA
	'\u0622': {'\u0622', 0, 0, '\uFE82'},               // ARABIC LETTER ALEF WITH MADDA ABOVE
	'\u0623': {'\u0623', 0, 0, '\uFE84'},               // ARABIC LETTER ALEF WITH HAMZA ABOVE
	'\u0624': {'\u0624', 0, 0, '\uFE86'},               // ARABIC LETTER WAW WITH HAMZA ABOVE
	'\u0625': {'\u0625', 0, 0, '\uFE88'},               // ARABIC LETTER ALEF WITH HAMZA BELOW
	'\u0626': {'\u0626', '\uFE8B', '\uFE8C', '\uFE8A'}, // ARABIC LETTER YEH WITH HAMZA ABOVE
	'\u0627': {'\u0627', 0, 0, '\uFE8E'},               // ARABIC LETTER ALEF
	'\u0628': {'\u0628', '\uFE91', '\uFE92', '\uFE90'}, // ARABIC LETTER BEH
	'\u0629': {'\u0629', 0, 0, '\uFE94'},               // ARABIC LETTER TEH MARBUTA
	'\u062A': {'\u062A', '\uFE97', '\uFE98', '\uFE96'}, // ARABIC LETTER TEH
	'\u062B': {'\u062B', '\uFE9B', '\uFE9C', '\uFE9A'}, // ARABIC LETTER THEH
	'\u062C': {'\u062C', '\uFE9F', '\uFEA0', '\uFE9E'}, // ARABIC LETTER JEEM
	'\u062D': {'\uFEA1', '\uFEA3', '\uFEA4', '\uFEA2'}, // ARABIC LETTER HAH
	'\u062E': {'\u062E', '\uFEA7', '\uFEA8', '\uFEA6'}, // ARABIC LETTER KHAH
	'\u062F': {'\u062F', 0, 0, '\uFEAA'},               // ARABIC LETTER DAL
	'\u0630': {'\u0630', 0, 0, '\uFEAC'},               // ARABIC LETTER THAL
	'\u0631': {'\u0631', 0, 0, '\uFEAE'},               // ARABIC LETTER REH
	'\u0632': {'\u0632', 0, 0, '\uFEB0'},               // ARABIC LETTER ZAIN
	'\u0633': {'\u0633', '\uFEB3', '\uFEB4', '\uFEB2'}, // ARABIC LETTER SEEN
	'\u0634': {'\u0634', '\uFEB7', '\uFEB8', '\uFEB6'}, // ARABIC LETTER SHEEN
	'\u0635': {'\u0635', '\uFEBB', '\uFEBC', '\uFEBA'}, // ARABIC LETTER SAD
	'\u0636': {'\u0636', '\uFEBF', '\uFEC0', '\uFEBE'}, // ARABIC LETTER DAD
	'\u0637': {'\u0637', '\uFEC3', '\uFEC4', '\uFEC2'}, // ARABIC LETTER TAH
	'\u0638': {'\u0638', '\uFEC7', '\uFEC8', '\uFEC6'}, // ARABIC LETTER ZAH
	'\u0639': {'\u0639', '\uFECB', '\uFECC', '\uFECA'}, // ARABIC LETTER AIN
	'\u063A': {'\u063A', '\uFECF', '\uFED0', '\uFECE'}, // ARABIC LETTER GHAIN
	TATWEEL:  {'\u0640', '\u0640', '\u0640', '\u0640'}, // ARABIC TATWEEL
	'\u0641': {'\u0641', '\uFED3', '\uFED4', '\uFED2'}, // ARABIC LETTER FEH
	'\u0642': {'\u0642', '\uFED7', '\uFED8', '\uFED6'}, // ARABIC LETTER QAF
	'\u0643': {'\u0643', '\uFEDB', '\uFEDC', '\uFEDA'}, // ARABIC LETTER KAF
	'\u0644': {'\u0644', '\uFEDF', '\uFEE0', '\uFEDE'}, // ARABIC LETTER LAM
	'\u0645': {'\u0645', '\uFEE3', '\uFEE4', '\uFEE2'}, // ARABIC LETTER MEEM
	'\u0646': {'\u0646', '\uFEE7', '\uFEE8', '\uFEE6'}, // ARABIC LETTER NOON
	'\uFBAB': {'\uFBAB', '\uFBAB', '\uFBAB', '\uFBAB'}, // ARABIC LETTER HEH
	'\u0648': {'\u0648', 0, 0, '\uFEEE'},               // ARABIC LETTER WAW
	'\u0649': {'\u0649', '\uFBE8', '\uFBE9', '\uFEF0'}, // ARABIC LETTER (UIGHUR KAZAKH KIRGHIZ)? ALEF MAKSURA
	'\u064A': {'\u064A', '\uFEF3', '\uFEF4', '\uFEF2'}, // ARABIC LETTER YEH
	'\u0671': {'\u0671', 0, 0, '\uFB51'},               // ARABIC LETTER ALEF WASLA
	'\u0677': {'\u0677', 0, 0, 0},                      // ARABIC LETTER U WITH HAMZA ABOVE
	'\u0679': {'\u0679', '\uFB68', '\uFB69', '\uFB67'}, // ARABIC LETTER TTEH
	'\u067A': {'\u067A', '\uFB60', '\uFB61', '\uFB5F'}, // ARABIC LETTER TTEHEH
	'\u067B': {'\u067B', '\uFB54', '\uFB55', '\uFB53'}, // ARABIC LETTER BEEH
	// ARABIC LETTER PEH
	'\u067E': {'\u067E', '\uFB58', '\uFB59', '\uFB57'},
	// ARABIC LETTER TEHEH
	'\u067F': {'\u067F', '\uFB64', '\uFB65', '\uFB63'},
	// ARABIC LETTER BEHEH
	'\u0680': {'\u0680', '\uFB5C', '\uFB5D', '\uFB5B'},
	// ARABIC LETTER NYEH
	'\u0683': {'\u0683', '\uFB78', '\uFB79', '\uFB77'},
	// ARABIC LETTER DYEH
	'\u0684': {'\u0684', '\uFB74', '\uFB75', '\uFB73'},
	// ARABIC LETTER TCHEH
	'\u0686': {'\u0686', '\uFB7C', '\uFB7D', '\uFB7B'},
	// ARABIC LETTER TCHEHEH
	'\u0687': {'\u0687', '\uFB80', '\uFB81', '\uFB7F'},
	// ARABIC LETTER DDAL
	'\u0688': {'\u0688', 0, 0, '\uFB89'},
	// ARABIC LETTER DAHAL
	'\u068C': {'\u068C', 0, 0, '\uFB85'},
	// ARABIC LETTER DDAHAL
	'\u068D': {'\u068D', 0, 0, '\uFB83'},
	// ARABIC LETTER DUL
	'\u068E': {'\u068E', 0, 0, '\uFB87'},
	// ARABIC LETTER RREH
	'\u0691': {'\u0691', 0, 0, '\uFB8D'},
	// ARABIC LETTER JEH
	'\u0698': {'\u0698', 0, 0, '\uFB8B'},
	// ARABIC LETTER VEH
	'\u06A4': {'\u06A4', '\uFB6C', '\uFB6D', '\uFB6B'},
	// ARABIC LETTER PEHEH
	'\u06A6': {'\u06A6', '\uFB70', '\uFB71', '\uFB6F'},
	// ARABIC LETTER KEHEH
	'\u06A9': {'\u06A9', '\uFB90', '\uFB91', '\uFB8F'},
	// ARABIC LETTER NG
	'\u06AD': {'\u06AD', '\uFBD5', '\uFBD6', '\uFBD4'},
	// ARABIC LETTER GAF
	'\u06AF': {'\u06AF', '\uFB94', '\uFB95', '\uFB93'},
	// ARABIC LETTER NGOEH
	'\u06B1': {'\u06B1', '\uFB9C', '\uFB9D', '\uFB9B'},
	// ARABIC LETTER GUEH
	'\u06B3': {'\u06B3', '\uFB98', '\uFB99', '\uFB97'},
	// ARABIC LETTER NOON GHUNNA
	'\u06BA': {'\u06BA', 0, 0, '\uFB9F'},
	// ARABIC LETTER RNOON
	'\u06BB': {'\u06BB', '\uFBA2', '\uFBA3', '\uFBA1'},
	// ARABIC LETTER HEH DOACHASHMEE
	'\u06BE': {'\u06BE', '\uFBAC', '\uFBAD', '\uFBAB'},
	// ARABIC LETTER HEH WITH YEH ABOVE
	'\u06C0': {'\u06C0', 0, 0, '\uFBA5'},
	// ARABIC LETTER HEH GOAL
	'\u06C1': {'\u06C1', '\uFBA8', '\uFBA9', '\uFBA7'},
	// ARABIC LETTER KIRGHIZ OE
	'\u06C5': {'\u06C5', 0, 0, '\uFBE1'},
	// ARABIC LETTER OE
	'\u06C6': {'\u06C6', 0, 0, '\uFBDA'},
	// ARABIC LETTER U
	'\u06C7': {'\u06C7', 0, 0, '\uFBD8'},
	// ARABIC LETTER YU
	'\u06C8': {'\u06C8', 0, 0, '\uFBDC'},
	// ARABIC LETTER KIRGHIZ YU
	'\u06C9': {'\u06C9', 0, 0, '\uFBE3'},
	// ARABIC LETTER VE
	'\u06CB': {'\u06CB', 0, 0, '\uFBDF'},
	// ARABIC LETTER FARSI YEH
	'\u06CC': {'\u06CC', '\uFBFE', '\uFBFF', '\uFBFD'},
	// ARABIC LETTER E
	'\u06D0': {'\u06D0', '\uFBE6', '\uFBE7', '\uFBE5'},
	// ARABIC LETTER YEH BARREE
	'\u06D2': {'\u06D2', 0, 0, '\uFBAF'},
	// ARABIC LETTER YEH BARREE WITH HAMZA ABOVE
	'\u06D3': {'\u06D3', 0, 0, '\uFBB1'},
	// Kurdish letter YEAH
	'\u06ce': {'\uE004', '\uE005', '\uE006', '\uE004'},
	// Kurdish letter Hamza same as arabic Teh without the point
	'\u06d5': {'\u06d5', 0, 0, '\uE000'},
	'\u0695': {'\u0695', 0, 0, '\uE001'},
	'\u0647': {'\u0647', '\u06BE', '\uFBAB', '\uFBAB'}, // ARABIC LETTER HEH
	'\u06b5': {'\u06B5', '\uE008', '\uE009', '\uE007'},
	// ZWJ
	ZWJ: {ZWJ, ZWJ, ZWJ, ZWJ},
	// Add more Kurdish letters here
}

// majority of the code is from that used to check state of the letter in arabic and the output
func connectsWithLetterBefore(letter rune, LETTERS map[rune][4]rune) bool {
	forms, exists := LETTERS[letter]
	if !exists {
		return false
	}
	return forms[FINAL] != 0 || forms[MEDIAL] != 0
}

func connectsWithLetterAfter(letter rune, LETTERS map[rune][4]rune) bool {
	forms, exists := LETTERS[letter]
	if !exists {
		return false
	}
	return forms[INITIAL] != 0 || forms[MEDIAL] != 0
}

func connectsWithLettersBeforeAndAfter(letter rune, LETTERS map[rune][4]rune) bool {
	_, exists := LETTERS[letter]
	if !exists {
		return false
	}
	return connectsWithLetterBefore(letter, LETTERS) && connectsWithLetterAfter(letter, LETTERS)
}

func isIsolated(letter rune, LETTERS map[rune][4]rune) bool {
	forms, exists := LETTERS[letter]
	if !exists {
		return false
	}
	return forms[ISOLATED] != 0
}

func isFinal(letter rune, LETTERS map[rune][4]rune) bool {
	forms, exists := LETTERS[letter]
	if !exists {
		return false
	}
	return forms[FINAL] != 0
}

func isInitial(letter rune, LETTERS map[rune][4]rune) bool {
	forms, exists := LETTERS[letter]
	if !exists {
		return false
	}
	return forms[INITIAL] != 0
}

func isMedial(letter rune, LETTERS map[rune][4]rune) bool {
	forms, exists := LETTERS[letter]
	if !exists {
		return false
	}
	return forms[MEDIAL] != 0
}

func isLetter(letter rune, LETTERS map[rune][4]rune) bool {
	_, exists := LETTERS[letter]
	return exists
}

func isTatweel(letter rune, LETTERS map[rune][4]rune) bool {
	return letter == TATWEEL
}

func isZWJ(letter rune, LETTERS map[rune][4]rune) bool {
	return letter == ZWJ
}

func isArabicLetter(letter rune) bool {
	return isLetter(letter, LETTERS_ARABIC) || isLetter(letter, LETTERS_ARABIC_V2) || isLetter(letter, LETTERS_KURDISH)
}

func isHarakat(letter rune) bool {
	//check if the key in the harakat map array
	_, exists := HARAKATMAP[letter]
	return exists

}
