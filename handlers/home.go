package handlers

import "net/http"

const homePage = `
<html>
<head>
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

      .sign-in-button {
          color: #fff;
          text-decoration: none;
          background-color: #27ae60;
          padding: 8px 12px;
          border-radius: 4px;
      }
      .sign-in-button:hover {
        background-color: #2ecc71;
      }

	  .navigation {
		max-width: 640px;
		margin-left: auto;
		margin-right: auto;
		padding: 1rem;
	  }
    </style>
</head>
<body>
    <div class="content">
        <h1 class="Header--title">Master of Malt Contacts Manager</h1>
        <a class="sign-in-button" href="/auth/redirect" title="Sign in">Sign in with Google</a>
    </div>
	<div class="navigation">
	  <nav>
	    <a href="/contacts">All contacts</a>
	  </nav>
	</div>
</body>
</html>
`

// Home is the handler that sits on the home page
func Home(rw http.ResponseWriter, r *http.Request) {
	body := []byte(homePage)
	rw.Write(body)
}
