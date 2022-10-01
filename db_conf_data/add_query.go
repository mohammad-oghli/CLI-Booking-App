package db_conf_data

import (
	"oghli.pro/cli-booking-app/models"
)

// Adding intial available Conferences for the first time
// You can add as much as you want conferences to DB
func Init_Conferences() {
	cisco_conf := models.Conference{Name: "Cisco Live", TotalTickets: 60}
	microsoft_conf := models.Conference{Name: "Microsoft MVP Global Summit", TotalTickets: 150}
	samsung_conf := models.Conference{Name: "Samsung Unpacked", TotalTickets: 100}
	redhat_conf := models.Conference{Name: "Red Hat Summit", TotalTickets: 25}
	dell_conf := models.Conference{Name: "Dell Technologies World", TotalTickets: 200}
	// Adding objects to booking app DB
	cisco_conf.NewConference()
	microsoft_conf.NewConference()
	samsung_conf.NewConference()
	redhat_conf.NewConference()
	dell_conf.NewConference()
}

// Adding admin user for the app
// Admin can create new conferences and add them to the app
func Set_Admin() {
	admin := models.User{UserName: "admin", Password: "admin@123"}
	// adding object to booking app DB
	admin.NewUser()
}
