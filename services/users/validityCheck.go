package users

import "regexp"

func isValidId(id int) bool {

	if id < 1 {

		return false
	}
	return true
}

func isValidName(name string) bool {

	if name == "" {

		return false
	}
	return true
}

func isValidEmail(email string) bool {
	emailreg := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailreg.MatchString(email)

}

func isValidPhone(phone string) bool {
	phonereg := regexp.MustCompile(`^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\s\./0-9]*$`)
	return phonereg.MatchString(phone)
}

func isValidAge(age int) bool {

	if age < 1 {

		return false
	}
	return true
}
