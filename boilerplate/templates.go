package boilerplate

var LayoutTemplFileData string = `
package templates

templ Layout(content templ.Component) {
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8"/>
			<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
			<title>Document</title>
			<link
				rel="stylesheet"
				href="/static/output.css"
			/>
			<script nonce={ templ.GetNonce(ctx) }>
				(function() {
					// Get current theme preference (system, light, or dark)
					function getThemePreference() {
						return localStorage.getItem('theme') || 'system';
					}

					const preference = getThemePreference();
					let isDark = false;

					if (preference === 'system') {
						// Use system preference
						isDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
					} else {
						// Use explicit preference
						isDark = preference === 'dark';
					}

					// Apply theme immediately to prevent flash
					if (isDark) {
						document.documentElement.classList.add('dark');
					}
				})();
			</script>
			<script defer type="application/javascript" src="/static/app.js"></script>
		</head>
		<body x-data="{ count: 0 }">
			@content
		</body>
	</html>
}
`

var IndexTemplFileData string = `

package templates
import "%s/components/button"
import "%s/components/themeswitcher"


templ Home() {
  <div class="w-screen h-screen flex flex-col justify-center items-center gap-3">
    @themeswitcher.ThemeSwitcher()
    <h1 class="text-6xl font-black tracking-tight">Hello, Emogo!</h1>
    @button.Button(button.Props{
      Size: button.SizeLg,
      Attributes: templ.Attributes{
        "hx-get": "/htmx-test",
      },
    }){
      Try HTMX
    }
		<div class="flex items-center gap-1">
			@button.Button(button.Props{
				Variant: button.VariantOutline,
				Size: button.SizeLg,
				Attributes: templ.Attributes{
					"@click": "count++",
				},
			}){
				Increment Count with Alpine.js
			}
			<span x-text="count"></span>
		</div>
  </div>
}
`
