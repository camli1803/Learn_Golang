package comparemapsfunction

import "fmt"

func Comparetwomaps(map1, map2 map[int]int) bool {
	check := true
	if len(map1) == len(map2) { //check chiều dài đảm bảo mọi key trong map1 đều được kiểm tra trong map2 và ngược lại
		for key, value := range map1 {
			valuemap2, ok := map2[key] //check xem map2 có key đó giống map1 không, ok là biến check.
			if ok == true {
				if valuemap2 != value {
					check = false
				}
			} else {
				check = false
			}
		}
	} else {
		check = false
	}
	return check
}
func PrintNotice(map1, map2 map[int]int) {
	check := Comparetwomaps(map1, map2)
	if check == true {
		fmt.Println("Equal!")
	} else {
		fmt.Println("Not Equal!")
	}

}
