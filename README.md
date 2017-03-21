# Boilerplate

This is my personal boilerplate for a deployment ready web development server. It includes the following:

 - Working web server running on `gin-gonic`
 - Routes for a REST api at `api/v1/users` route
 - Models with a basic user
 - Storage that initialized a database connection and automatically migrates models
 - Config that allows for feature flagging and statically typed environment variables
 - Controllers for serving front end
 - Heroku ready deployment and local testing using `heroku local`
 - Gulp with automatic `go build` and SASS and JS compilation
 - CSRF protection with `justinas/nosurf`
 - Templating engine with Pongo, standard chrome, and partials
 - Standard jQuery library
