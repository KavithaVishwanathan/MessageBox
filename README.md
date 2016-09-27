# MessageBox

A single page application (https://infinite-chamber-34739.herokuapp.com/) is built to represent a basic message box functionality. It gets it’s data from a Rest API. The aim of this application is to display the list of messages and the message details. Each message will have the mandatory e-mail components like sender, receiver, subject, body and pre-defined labels.

Stack: Go (Rest API), jQuery (Ajax), SQLite

#### API
Utilizing the simplicity of Go, a rest API was built. Here the endpoints of this API, gives the data of JSON Form. The following are the few endpoints (it displays when you perform the “GET”)

* /message -> displays all the messages
* /label -> displays all the labels
* /user -> displays all the users
* /message/{id} - displays the message details of that particular id
* /message/{id}/labels - displays the labels tagged to that message
* /user/{id} - displays the user details of that particular id

#### FrontEnd:

When you access the root of the url, the jquery sends the Ajax request to get data from /message endpoint and displays it as a table.

When you click on “More Details”, we send multiple ajax requests to display the message details

/message/{id}

/user/{from-id}  (twice one for from and one for to)

/message/{id}/labels

#### Steps to run in local:

* Go get github.com/KavithaVishwanathan/MessageBox
* Go run main.go

If you want to create your own data.

* Insert the mock data in “db/rundb.go”
* Uncomment the db.Run() line in main.go 
* Go run main.go


