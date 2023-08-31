package dating_test

import (
	"fmt"
	dating "forum/dating"
	"testing"
)

func Test_date_for_database(t *testing.T) {
	dating.Parse_date_from_database(dating.That_moment())
	dating.Parse_date_from_database(dating.That_moment())
	dating.Parse_date_from_database(dating.Rdm_date())
	dating.Parse_date_from_database(dating.Rdm_date())
}

func Test_rdm_date(t *testing.T) {
	fmt.Println("getting random date for database")
	fmt.Println(dating.Rdm_date())
	fmt.Println(dating.Rdm_date())
	fmt.Println(dating.Rdm_date())
	fmt.Println(dating.Rdm_date())
	fmt.Println(dating.Rdm_date())
	fmt.Println()
}
func Test_that_moment(t *testing.T) {
	fmt.Println("getting time now for database")
	fmt.Println(dating.That_moment())
	fmt.Println(dating.That_moment())
	fmt.Println(dating.That_moment())
	fmt.Println(dating.That_moment())
	fmt.Println(dating.That_moment())
	fmt.Println()
}

func Test_date_only_to_display_in_thread(t *testing.T) {
	fmt.Println("create a date as string, parse it to an time object, sring it into handy date format:")
	fmt.Println(dating.Just_the_date(dating.Parse_date_from_database(dating.That_moment())))
	fmt.Println(dating.Just_the_date(dating.Parse_date_from_database(dating.Rdm_date())))
	fmt.Println(dating.Just_the_date(dating.Parse_date_from_database(dating.Rdm_date())))
}