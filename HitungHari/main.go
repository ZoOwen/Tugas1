package main

import (
	"fmt"
	"time"
)

func main(){
	
		time1 := time.Date(2019, time.December,30, 23, 0, 0, 0, time.UTC)
		time2 := time.Date(2019, time.January,1, 23, 0, 0, 0, time.UTC)
        diff := time1.Sub(time2)
        hari := int(diff.Hours() / 24)
        fmt.Println("jumlah hari antara 3 januari sampai 10 Februari adalah",hari,"hari")
}