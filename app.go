package main

import (
	"bufio"
	"fmt"
	"os"

	"oghli.pro/cli-booking-app/models"
)

func main() {
	fmt.Println("WELCOME TO CLI BOOKING APP!")
	fmt.Println("###########################")
	// Adding admin user to booking app DB
	// Comment it after first run of the app
	//db_conf_data.Set_Admin()
	// Adding intial Conferences data to booking app DB
	// Comment it after first run of the app
	//db_conf_data.Init_Conferences()
	var option uint
	var logged_user string
	var logged bool = false
	fmt.Println("Select option:")
	fmt.Println("1. Register to the App")
	fmt.Println("2. Login if you already a registered user")
	fmt.Scan(&option)
	if option == 1 {
		var username string
		var password string
		fmt.Print("Enter UserName: ")
		fmt.Scan(&username)
		fmt.Print("Enter Password: ")
		fmt.Scan(&password)
		q_user := models.GetUserByUserName(username)
		if q_user.ID == 0 {
			newUser := models.User{UserName: username, Password: password}
			newUser.NewUser()
			fmt.Println("Account Successfuly Created!")
		} else {
			fmt.Printf("%v UserName is already available!", q_user.UserName)
		}
	} else if option == 2 {
		var username string
		var password string
		fmt.Println("Login to your account")
		fmt.Print("UserName: ")
		fmt.Scan(&username)
		fmt.Print("Password: ")
		fmt.Scan(&password)
		q_user := models.GetUserByUserName(username)
		if q_user.ID != 0 {
			if username == "admin" {
				if q_user.Password == password {
					fmt.Println("Administrator logged in!")
					logged = true
					logged_user = username
				} else {
					fmt.Println("Incorrect Password!")
				}
			} else {
				if q_user.Password == password {
					fmt.Printf("Welcome %v!\n", q_user.UserName)
					logged = true
					logged_user = username
				} else {
					fmt.Println("Incorrect Password!")
				}
			}
		} else {
			fmt.Println("Invalid UserName or Password!")
		}
	} else {
		fmt.Println("Invalid Option!")
	}
	if logged {
		if logged_user == "admin" {
			var conf_name string
			var total_tickets uint
			fmt.Println("Create New Conference:")
			fmt.Println("######################")
			fmt.Print("Enter conference name: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			scanner.Scan()
			conf_name = scanner.Text()
			fmt.Print("Enter total number of tickets: ")
			fmt.Scan(&total_tickets)
			q_conf := models.GetConferenceByName(conf_name)
			if q_conf.ID == 0 {
				newConf := models.Conference{Name: conf_name, TotalTickets: total_tickets}
				newConf.NewConference()
				fmt.Printf("%v Successfuly Created!", newConf.Name)
			} else {
				fmt.Printf("%v Conference is already available!", conf_name)
			}
		} else {
			fmt.Println("List of of currently available Conferences")
			fmt.Println("##########################################")
			var option uint
			var firstName string
			var lastName string
			var email string
			var userTickets uint
			q_user := models.GetUserByUserName(logged_user)
			conferences := models.GetAllConferences()
			for index, item := range conferences {
				fmt.Printf("%v. %v | Total Tickets: %v | Total booked users: %v\n", index+1, item.Name, item.TotalTickets, len(item.Users))
			}
			fmt.Println("##########################################")
			if q_user.Conf_ID == 0 {
				fmt.Print("Please select conference number to book: ")
				fmt.Scan(&option)
				fmt.Print("Enter your first name: ")
				fmt.Scan(&firstName)
				fmt.Print("Enter your last name: ")
				fmt.Scan(&lastName)
				fmt.Print("Enter your email address: ")
				fmt.Scan(&email)
				fmt.Print("Enter number of tickets: ")
				fmt.Scan(&userTickets)
				if option >= 1 && option <= uint(len(conferences)) {
					selected_conf := conferences[option-1]
					q_conf := models.GetConferenceByName(selected_conf.Name)
					if int(q_conf.TotalTickets-userTickets) >= 0 {
						q_conf.TotalTickets = q_conf.TotalTickets - userTickets
						q_conf.Users = append(q_conf.Users, *q_user)
						q_user.FirstName = firstName
						q_user.LastName = lastName
						q_user.Email = email
						q_user.BookedTickets = userTickets
						q_conf.UpdateConferenceBooking()
						q_user.UpdateUserBooking()
						fmt.Printf("Thank you %v %v for booking %v tickets. You will recive confirmation email at %v.\n", firstName, lastName, userTickets, email)
						fmt.Printf("%v tickets remaining for %v.\n", q_conf.TotalTickets, q_conf.Name)
					} else {
						fmt.Printf("%v conference is booked out. Come back next year.", q_conf.Name)
					}
				} else {
					fmt.Println("Invalid Option!")
				}
			} else {
				var names []string
				q_conf := models.GetConferenceById(q_user.Conf_ID)
				fmt.Printf("You already booked %v tickets for %v.\n", q_user.BookedTickets, q_conf.Name)
				for _, item := range q_conf.Users {
					names = append(names, item.FirstName)
				}
				fmt.Println("The first names of booked people in this conference:")
				for _, item := range names {
					fmt.Printf("• %v\n", item)
				}
				fmt.Println("Thank you for using our CLI Booking App!")
			}
		}
	}
}
