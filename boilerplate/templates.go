package boilerplate

var LayoutTemplFileData string = `
package templates

templ Layout(content templ.Component) {
  <!DOCTYPE html>
  <html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <link
      rel="stylesheet"
      href="/static/output.css"
    >
		<script defer type="application/javascript" src="/static/app.js"></script>
  </head>
  <body class="">
    @content
  </body>
  </html>
}
`

var IndexTemplFileData string = `
package templates

templ Home() {
  <div class="w-screen h-screen flex flex-col justify-center items-center gap-3">
    <h1 class="text-6xl font-black tracking-tight">Hello, Emogo!</h1>
    <button hx-get="/htmx-test" class="bg-black text-white px-3 py-2 rounded-lg tracking-wider">Use HTMX</button>
  </div>
}
`
