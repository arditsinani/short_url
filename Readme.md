# URL shortener

Url shortener writen in **GO** with **Gin Gonic** web server and **MongoDB**. It stores the url's and creates a short path for each. The new path is exposed when created, and redirects to the original url.

## Steps

The application is build with **Docker** and can be run with *docker-compose*, so docker is required in the machine to run the application. A *Makefile* is used to facilitate application comands.

```
> Clone the repository
> cd short_url
> make build
> make up
```

## Routes

The app has the following routes

```
POST => http://localhost:8080/create
payload: {
    url: "url that needs to be shortened"
}
```

```
GET => http://localhost:8080/api/read/(shortened url) => get the info of the short url
```

```
GET => http://localhost:8080/(shortened url) => redirects to the original url
```

```
DELETE => http://localhost:8080/api/delete/(shortened url) => updates entity status to deleted, doest delete it
```

Enjoy!