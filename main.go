package main

import (
	"encoding/json"
	"fmt"
	srv "main/httpServerNew"
	kafka "main/kafka"
	routine "main/routine"
	"os"
	"sync"
)

type XYZ struct {
	Name int
	Data map[string]Data
}

type Data struct {
	Value map[string]*Value
}

type Value struct {
	ID int
	ip *int
}

func main() {
	ip := 2000
	val := Value{ID: 100000, ip: &ip}
	dta := Data{}
	dta.Value = make(map[string]*Value)
	dta.Value["abcd"] = &val

	x := XYZ{}
	x.Data = make(map[string]Data)
	x.Data["xyz"] = dta
	x.Name = 100
	b, err := json.Marshal(x.Data)
	if err != nil {
		fmt.Println("error marshalling : ", err)
		os.Exit(1)
	}
	fmt.Println("marshalling Data : ", string(b))

	//os.Exit(1)
	var argMent string
	if len(os.Args) > 1 {
		argMent = os.Args[1]
	}

	fmt.Println("argMent : ", argMent)

	if argMent == "Route" {
		srv.InitPProfServer()
		server := srv.GetHttpRouterServer()
		if err := server.ListenAndServe(); err != nil {
			fmt.Println("error while route server : ", err)
			os.Exit(1)
		}
	} else if argMent == "HTTP" {
		srv.GetAndStartNormalHttpServer()
	} else if argMent == "NORMAL_GO" {
		fmt.Println("calling normalGoRoutine")
		go routine.NormalGoRoutine()
		fmt.Println("after normalGoRoutine")
	} else if argMent == "WG_GO" {
		fmt.Println("calling WaitGroupGoRoutine")
		var wg sync.WaitGroup
		wg.Add(1)
		go routine.WaitGroupGoRoutine(&wg)
		wg.Wait()
		fmt.Println("after WaitGroupGoRoutine")
	} else if argMent == "STRING_CHAN_GO" {
		ch := make(chan string)
		go routine.NormalChanGoRoutineString(ch)
		select {
		case resp, ok := <-ch:
			if ok {
				fmt.Println("after NormalChanGoRoutineString resp : ", resp)
			} else {
				break
			}
		}

	} else if argMent == "SR_CHAN_GO_Async" {
		fmt.Println("SR_CHAN_GO")
		ch := make(chan int, 10)
		var wg sync.WaitGroup
		wg.Add(1)
		go routine.SendChanGoRoutine(ch, &wg)
		wg.Wait()
		wg.Add(1)
		go routine.ReceivedChanGoRoutine(ch, &wg)
		wg.Wait()
	} else if argMent == "SR_CHAN_GO_sync" {
		fmt.Println("SR_CHAN_GO")
		ch := make(chan int)
		var wg sync.WaitGroup
		wg.Add(2)
		go routine.SendChanGoRoutine(ch, &wg)
		go routine.ReceivedChanGoRoutine(ch, &wg)
		wg.Wait()
	} else if argMent == "kafka" {

		var wg sync.WaitGroup
		wg.Add(1)
		go kafka.ConsumeAgain(&wg)
		wg.Wait()
	}
}
