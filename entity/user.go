package entity

//User for our system
type User struct {
	email       string
	password    string
	name        string
	dateOfBirth string
}

//validates at a business level
func (u *User) Validate() error {
	return nil
	//   ...
}
