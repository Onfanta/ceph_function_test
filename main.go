package main

import (
	"fmt"
)


func main() {
	chos := [] string {"Pool Test","Ec Test","User Test", "Rgw Test", "Test All", "exit"}
	var cho int
	fmt.Println("function:")
	for i:=0;i< len(chos);i++  {
		fmt.Println(i+1,chos[i])
	}
	fmt.Printf("plz input option u want[1-6] :")
	fmt.Scanln(cho)
	if cho == 5  {
		fmt.Printf("processing...")
		PoolTest()
		EcTest()
		UserTest()
		RgwTest()
	}else if cho < 5 {
		switch cho {
		case 1:
			fmt.Println("PoolTest is started")
			PoolTest()
		case 2:
			fmt.Println("EcTest is started")
			EcTest()
		case 3:
			fmt.Println("UserTest is started")
			UserTest()
		case 4:
			fmt.Println("RGWTest is started")
			RgwTest()
		}
	}else if cho == 6 {
		fmt.Println("Exiting...")

	}else {
		fmt.Println("type err,plz input again...")
	}
}
