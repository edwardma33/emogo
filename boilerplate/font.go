package boilerplate

import _ "embed"

// GeistLatinVariableFont is the Google Fonts Latin variable face, covering
// font weights 100 through 900.
//
//go:embed assets/geist-latin-wght-normal.woff2
var GeistLatinVariableFont []byte

// GeistOFL is the SIL Open Font License for the bundled Geist font.
//
//go:embed assets/geist-OFL.txt
var GeistOFL []byte
