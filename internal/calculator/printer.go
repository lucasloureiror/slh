package calculator

import "fmt"

func print() {
	fmt.Println(downtime.SLA+"%", "availability represents the following maximum allowable downtime to meet your Service Level:")
	fmt.Println(downtime.Daily)
	fmt.Println(downtime.Weekly)
	fmt.Println(downtime.Monthly)
	fmt.Println(downtime.Quartely)
	fmt.Println(downtime.Yearly)

}
