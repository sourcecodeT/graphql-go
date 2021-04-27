package altair

import (
	"html/template"
	"net/http"
)

var page = template.Must(template.New("graphiql").Parse(`
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <title>Altair</title>
    <base href="https://cdn.jsdelivr.net/npm/altair-static@4.0.2/build/dist/" />
    <meta name="viewport" content="width=device-width,initial-scale=1" />
    <link rel="icon" type="image/x-icon" href="favicon.ico" />
    <link href="styles.css" rel="stylesheet" />
  </head>

  <body>
    <app-root>
      <style>
        .loading-screen {
          /*Prevents the loading screen from showing until CSS is downloaded*/
          display: none;
        }
      </style>
      <div class="loading-screen styled">
        <div class="loading-screen-inner">
          <div class="loading-screen-logo-container">
            <img src="assets/img/logo_350.svg" alt="Altair" />
          </div>
          <div class="loading-screen-loading-indicator">
            <span class="loading-indicator-dot"></span>
            <span class="loading-indicator-dot"></span>
            <span class="loading-indicator-dot"></span>
          </div>
        </div>
      </div>
    </app-root>
    <script
      rel="preload"
      as="script"
      type="text/javascript"
      src="runtime-es2018.js"
    ></script>
    <script
      rel="preload"
      as="script"
      type="text/javascript"
      src="polyfills-es2018.js"
    ></script>
    <script
      rel="preload"
      as="script"
      type="text/javascript"
      src="main-es2018.js"
    ></script>
    <script>
      document.addEventListener("DOMContentLoaded", () => {
        AltairGraphQL.init({
					endpointURL: 'http://localhost:8000/graphql'
				});
      });
    </script>
  </body>
</html>
`))

func Handler(title string, endpoint string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		err := page.Execute(w, map[string]string{
			"endpoint": endpoint,
		})
		if err != nil {
			panic(err)
		}
	}
}
