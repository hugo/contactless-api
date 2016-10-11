package handlers

import "net/http"

const homePage = `
<html>
<head>
    <meta name="google-site-verification" content="kA-AnHBsa5Zyu9r8K81laEYgDy45x1WT54FZplcY7Dc" /> 
    <style>
      body {
          font-family: -apple-system, 'helvetica neue', helvetica, arial, sans-serif;
          font-size: 16px;
      }

      .content {
          max-width: 640px;
          margin-left: auto;
          margin-right: auto;
          padding: 1rem;
          background-color: #2c3e50;
      }

      .Header--title {
        color: #fff;
      }
    </style>
</head>
<body>
  <div class="content">
    <h1 class="Header--title">Contactless</h1>
  </div>
</body>
</html>
`

// Home is the handler that sits on the home page
func Home(rw http.ResponseWriter, r *http.Request) {
	body := []byte(homePage)
	rw.Write(body)
}
