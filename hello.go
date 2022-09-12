package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	studentArray:= [74]string{"Behling, John;",
	"Bell, Dan;",
	"Benson, Cole;",
	"Brockbank, Stephen;",
	"Carlock, Cody;",
	"Choi, Brittany;",
	"Choi, Ye Jin;",
	"Clark, Madison;",
	"Ferolino, Alexis Jane;",
	"Guevara Alvarenga, Stefany;",
	"Hansen, Nathan;",
	"Hoather, Jeff;",
	"Horlacher, Ethan;",
	"Hunt, Brandon;",
	"Jensen, David;",
	"Jung, Euigun;",
	"Kimball, Logan;",
	"Ladle, Dallin;",
	"Lee, SeungEun;",
	"Lei, Sheng;",
	"Leung, Miles;",
	"Lo, Shaun;",
	"Marks, Greg;",
	"Marquis, Caden;",
	"McConkie, Liberty;",
	"McCord, Matthew;",
	"McMillan, Zac;",
	"Monson, Bailey;",
	"Nelson, Sloan;",
	"Peterson, James;",
	"Piscione, Michael;",
	"Prettyman, Samantha;",
	"Ridd, Hayden;",
	"Salvesen, Connor;",
	"Shipley, David;",
	"Stanley, Madison;",
	"Sweeten, Daniela;",
	"Tempest, Jordan;",
	"Trammell,  Mark;",
	"Andelin, Kyle;",
	"Anderson, Taylor;",
	"Baker, Nathan;",
	"Barton, Zachary;",
	"Berry, Solomon;",
	"Bullock, Taylor;",
	"Busco, Brian;",
	"Davis, Michael;",
	"Egbert, Seth;",
	"Glazier, Tanner;",
	"Goulding, Matt;",
	"Jackson, Spencer;",
	"Jensen, Emily;",
	"Karras, Caden;",
	"King, Spencer;",
	"Lund, Thomas;",
	"Luper, Abbie;",
	"Maxfield, Chase;",
	"Miller, Brinley;",
	"Moody, Josh;",
	"Moulton, McKay;",
	"Nabrotzky, Keanna;",
	"Nelson, Hunter;",
	"Nielsen, Dustin;",
	"Prock, Kamryn;",
	"Sanderson, Ian;",
	"Schow, Jackson;",
	"Scorse, Brett;",
	"Smith, Ali;",
	"Sorensen, Stephen;",
	"Souter, Kaden;",
	"Spencer, Jessie;",
	"Taylor, Chandler;",
	"Washburn, Jackson;",
	"Williams, Brennan;"}
	var numInTeam int

	// Prompt the number of members for each team
	fmt.Println("Enter the number for each team:")
	fmt.Scan(&numInTeam)
	fmt.Println("You entered: ",numInTeam)

	// Randomize the array
	rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(studentArray), func(i, j int) {
        studentArray[i], studentArray[j] = studentArray[j], studentArray[i]
    })
	
	// Calculate num of teams and remainders
	res := len(studentArray) % numInTeam
	numOfTeams := (len(studentArray) - res) / numInTeam
	if res != 0 {
		numOfTeams += 1
	} 
	fmt.Println(res)
	fmt.Println(numOfTeams)
	
	// Output the teams
	for i:= 0; i < len(studentArray)-res; i+=numInTeam{
		j := 0
		j += i + numInTeam
		fmt.Println(studentArray[i:j])
	}
	if res != 0{
		fmt.Println(studentArray[len(studentArray)-1-res:len(studentArray)-1])
	}
}