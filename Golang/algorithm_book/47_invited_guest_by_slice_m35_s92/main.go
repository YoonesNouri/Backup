package main

import (
	"fmt"
)

func main() {
	var name string
	invited := []string{}

	for {
		fmt.Print("Enter an invited name (press enter to stop): ")
		fmt.Scanln(&name)
		if name == "" {
			break
		}
		invited = append(invited, name)
	}

	for {
		fmt.Print("Insert a name: ")
		fmt.Scanln(&name)

		found := false
		for i := 0; i < len(invited); i++ {
			if invited[i] == name {
				fmt.Printf("%v is invited\n", name)
				found = true
				if len(invited) == 1 {
					return //خروج از فانک مِین چون برنامه تمام شده است
				}
				copy(invited[i:], invited[i+1:])
				invited = invited[:len(invited)-1]
				break
			}
		}

		if !found {
			fmt.Printf("%v is not invited\n", name)
		}
	}

	fmt.Println("Invited names:")
	for _, v := range invited {
		fmt.Println(v)
	}
}
