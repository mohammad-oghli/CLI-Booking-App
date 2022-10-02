# CLI Booking APP
Command-line interface Booking program using `Golang + Mysql + Gorm`.

The application main functionalities:
* Register a user 
* Login to the app
* List conferences to book with total available tickets for each one
* Number of booked users for each conference
* Book conference tickets for registered user with their required info 
* Admin user can add new conference to the app

**Note**: All Create, Read, Update program operations will be executed on the connected Mysql Database using `Gorm` ORM for Golang.

The app interface look like:

<code>

PS C:\Users\mhdsh\Go Workpsace\CLI-Booking-App> go run app.go        

WELCOME TO CLI BOOKING APP!

###########################

Select option:
1. Register to the App
2. Login if you already a registered user

1

Enter UserName: oghli

Enter Password: 123456

Account Successfuly Created!

PS C:\Users\mhdsh\Go Workpsace\CLI-Booking-App> go run app.go

WELCOME TO CLI BOOKING APP!

###########################

Select option:
1. Register to the App
2. Login if you already a registered user

2

Login to your account

UserName: oghli

Password: 123456 

Welcome oghli!

List of of currently available Conferences

##########################################
1. Cisco Live | Total Tickets: 60 | Total booked users: 0
2. Microsoft MVP Global Summit | Total Tickets: 120 | Total booked users: 2
3. Samsung Unpacked | Total Tickets: 80 | Total booked users: 1
4. Red Hat Summit | Total Tickets: 23 | Total booked users: 1
5. Dell Technologies World | Total Tickets: 165 | Total booked users: 0    
6. Apple Unleashed | Total Tickets: 250 | Total booked users: 0
7. IGN Entertainment | Total Tickets: 50 | Total booked users: 0

##########################################

Please select conference number to book: 2

Enter your first name: Mohamad

Enter your last name: Oghli

Enter your email address: oghli@gmail.com

Enter number of tickets: 25

Thank you Mohamad Oghli for booking 25 tickets. You will recive confirmation email at oghli@gmail.com.

95 tickets remaining for Microsoft MVP Global Summit.

</code>

To connect to your local Mysql database you should change this line of code in `db_conn.go` file inside `db_config` directory:

`dsn := "user:pass@tcp(localhost:3306)/booking_app_db?charset=utf8mb4&parseTime=True&loc=Local"`

add here your **user** and **pass** for Mysql.

In order to build the app run this command in your terminal inside application directory:

`go build`

After that run the app build using:

`.\cli-booking-app.exe`

You can star the repository to follow the latest updates of this version of CLI-Booking App.