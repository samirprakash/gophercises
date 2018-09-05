# Exercise details

Create a `http.Handler` package that checks the path of any incoming request and redirects the user to a new page. This is how a URL shortner would work where on trying to access a shortened URL say `https://localhost:8080/google`, the user would get redirected to `https://www.google.com`.

We should be creating this as a package so that the logic can be imported and reused in the calling applications.