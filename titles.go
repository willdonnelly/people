package people

import (
	"fmt"
	"math/rand"
)

func Title() string {
	title := Titles[rand.Int31n(int32(len(Titles)))]
	return fmt.Sprintf("%s, ExampleCorp", title)
}

var Titles = []string{
	"Recruiter",
	"Human Resources Assistant",
	"Payroll Specialist",
	"Payroll Clerk",
	"Human Resources Specialist",
	"Technical Recruiter",
	"Administrator Payroll",
	"Training Specialist",
	"Human Resources Coordinator",
	"Trainer",
	"Physical Therapist",
	"Occupational Therapist",
	"Registered Nurse",
	"Pharmacist",
	"Speech Language Pathologist",
	"Physical Therapist Thera",
	"Nurse",
	"Speech Pathologist Assistant",
	"Occupational Therapist Occupation",
	"Marketing Manager",
	"Product Manager",
	"Business Development Manager",
	"Marketing Director",
	"Director Business Development",
	"Product Marketing Manager",
	"Senior Product Manager",
	"Account Manager",
	"Marketing",
	"Senior Marketing Manager",
	"Project Manager",
	"Software Engineer",
	"Java Developer",
	"Business Analyst",
	"Systems Engineer",
	"Network Engineer",
	"Senior Software Engineer",
	"Web Developer",
	".net Developer",
	"Systems Administrator",
	"Engineer",
	"Mechanical Engineer",
	"Electrical Engineer",
	"Manufacturing Engineer",
	"Quality Engineer",
	"Project Engineer",
	"Process Engineer",
	"Design Engineer",
	"Technician",
	"Engineering Technician",
	"Accountant",
	"Senior Accountant",
	"Controller",
	"Financial Analyst",
	"Cost Accountant",
	"Accounting Manager",
	"Senior Financial Analyst",
	"Auditor",
	"Assistant Controller",
	"Tax Manager",
}
