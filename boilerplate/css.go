package boilerplate

var InputCSSFileData string = `
@import "tailwindcss";

@font-face {
  font-family: "Geist";
  font-style: normal;
  font-weight: 100 900;
  font-display: swap;
  src: url("/static/files/geist-latin-wght-normal.woff2") format("woff2");
}

@theme {
  --font-sans: "Geist", ui-sans-serif, system-ui, sans-serif;
}

@layer base {
  body {
    font-family: var(--font-sans);
  }
}
`
