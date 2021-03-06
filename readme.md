# How to have this repo

1. You can fork it. This means you will have a copy on your github.com profile. When changes are made to either repo (either this or yours), they won't appear on the other. Basically, each repo will be independent.

2. You can clone this repo (recommended)

```sh
cd go/src/
git clone https://github.com/innv8/api-introduction
cd api-introduction
```

This allows you to pull changes when they are made on this repo (on github)

```sh
git pull origin main
```

Then if you want to make it your own changes, I can make you a contributor and then you can have a branch e.g `dev`

```sh
git checkout -b dev
```

This will make a copy of the `main` branch into a new branch called `dev`

If you want to branch from a specific branch e.g `lesson/04` you need to checkout to that branch before creating the `dev` branch

```sh
git checkout lesson/04
git checkout -b dev
```

So lets say you have made changes on the dev branch you can commit with a message and push

```sh
# git add . basically bundles all your changes
# the . means this whole directory
# you need to run this in the root directory of the project

git add . 

# e.g
# git commit -m "fixed bug that was preventing /users endpoint to get users from the db"
git commit -m "the message describing what you have done"

# push the changes to github/ any site that is hosting the repo
git push origin dev
```

3. The other option is to clone this project, and then change its remote origin to be your own.


- On your githu profile, create a repo and call it `api-introduction`

- Clone the repo from `innv8`

```sh
cd go/src/
git clone https://github.com/innv8/api-introduction.git
cd api-introduction
```

- remove the remote. This is basically detaching the offline code from the online repo

```sh
# note: origin is the default name for the remote.
# you can different origins if you want to be pushing your code to different remotes e.g github and gitlab but this usually almost never happens.
git remote remove origin
```


- now connect the offline code to your online repo

```sh
git remote add origin https://github.com/innv8/api-introduction.git
```

- now push to your remote

```sh
# if you do
# git push origin main
# it will only push main
# if you replace the branch name with '--all', then all the branches will be pushed.
git push origin --all
```


---
---

# Web API

A web API is an application that runs on a server and offers resources e.g a database to the clients.

We will create a simple API and access it from the browser.

---

## Lesson 2 : MVC

There is a design paradigm called MVC:
M : Models
V : Views
C : Controller

In Go, we can create packages for each of them.

#### Models

Models are basically functions that deal with exchange of data. Let's say your API reads and writes data in the database, models will be functions that are responsible for this.  They can also be used to exchange data between your API and any other external resource which can be a database, cache, another API etc.

#### View

View are basically routes. These are endpoints that you give the client. So from the client's perspective, they can go to a certain route (view) to view data.

#### Controllers

These are functions that sit between views and models. When a request comes to a view (route), a controller is called. The controller then calls the relevant model(s), and then it processes the output from the model and then sends the response to the client.

If you are writing in a language that has pointers, the controllers can have direct access to resource connections e.g the database and then share the resources with the models as needed.



## Project Directories

Note that : this is a personal preference. You can explore other options online.

- controllers : will have controllers and views (routes) and also connections to outside resources e.g database

- models : functions that deal with data

- entities : structs to define the structure of data that is expected from and by the clients and also any struct you need e.g for models.

- utils : these will be utility functions - functions that can be used anywhere e.g a function that configure our logger, encrypt

---

## Lesson 3 : Middleware/ Gorrila Library

Gorrila is just a collection of libraries that can be used in web APIs. We'll start with mux which helps in creating routers (views)

With mux, you can specify the HTTP method (e.g GET, POST, PUT, PATCH, DELETE ) etc
e.g

```go
var router = mux.NewRouter()
router.HandleFunc("/users", controllers.UsersController).Methods("PUT", "POST")
```

This means, `/users` endpoint will only accept PUT and POST methods.

In the controller, if you want to know the method, you can use r.Method from the *http.Request



### Middleware

A middleware is a middleman between the router and the controller. For example, controller1 returns data as an array [1,2,3] while controller2 returns data as a number e.d 4. However, you want to maintain the data structure between you and the client so that controller1 returns 

```json
{
    "data": [1,2,3]
}
```

and controller2 returns

```json
{
    "data": 4
}
```

Instead of writing the logic for conversion in each controller, you can just create a middleware.

Another common use case for middlewares is authenticating requests. Therefore, before a request reaches a controller, it is authenticated first.

In the first use of middlware (controller -> middleware -> router), we just define the middleware as a function and call it from the controller.

In the second use of middleware (router -> middleware -> controller), we'll need to set the middleware on the router.

---

## Lesson 4 : Middleware 2, Connect to a Database

The second type of middleware receives the request from the router (view) before it forwards it to the controller. this could be to authenticate the request or to check for anything.

We will create a middleware that prints out all the headers sent by the client.

The code is in `middleware/headers.go` and the function is called in the router (`controllers/base.go`)

### Connecting to Mysql

When connecting to MySQL we need a connection URI (Uniform Resource Identifier) its a connection string/ identifier for resources e.g databases, caches, queues.

A MySQL URI is in the following format

```shell
username:password@host:port/db_name
```

E.g if your username is `root`, password is `password`, the database is on your `localhost` running on port `3306` and the database name is `chama`

```shell
root:password@tcp(localhost:3306)/chama
```


So to connect you can write this function that returns the pointer to the connection and an error (if it occures)

```go


import "database/sql"

func ConnectToDB() (connection *sql.DB, err error) {

	var dbURI = "root:password@tcp(localhost:3306)/chama"

	connection, err = sql.Open("mysql", dbURI)
	if err != nil {
		return
	}

	// it is advisable to ping
	err = connection.Ping()
	if err != nil {
		return
	}

	return

}


```

To read from the database, we need write a query and use the db connection pointer to fetch the details


```go

type User struct {
	ID int
	Name string
}

func FetchUsers(db *sql.DB) (users []User, err error) {
	var query = "SELECT id, name FROM users"

	rows, err := db.Query(query)
	if err != nil {
		// deal with error
		return
	}

	// loop through the rows until the end
	// in each loop read into an instance of User (or individual id, name variables)
	for rows.Next() {
		var user User

		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			// deal with error
			return
		}

		users = append(users, user)

	}

	log.Println("got %d users from db", len(users))
	return
}
```

---
## Lesson 5 : Dealing with Databases

We can import the `chama.sql` file into mysql. 

- Connect to mysql and create a database called `chama`

```sh
mysql -u username -ppassword
```

```sh
create database chama
```

- Exit mysql console and import the `chama.sql` file

```sh
mysql -u user -ppassword chama < chama.sql
```

We will change the `GetUsers` to return members.

Refer to `models/users.go`


## Lesson 6 : Create a new member