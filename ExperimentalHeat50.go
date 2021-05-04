/*
Copyright 2020 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"strconv"

    //MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)
var newSensorReading string
var wg sync.WaitGroup
func workerOne(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 1 done"
}

func workerTwo(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 2 done"
}

func workerThree(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 3 done"
}

func workerFour(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 4 done"
}

func workerFive(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 5 done"
}

func workerSix(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
     ch <- "worker 6 done"
}

func workerSeven(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 7 done"
}

func workerEight(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 8 done"
}

func workerNine(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 9 done"
}

func workerTen(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 10 done"
}

func workerEleven(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 11 done"
}

func workerTwelve(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 12 done"
}

func workerThirteen(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 13 done"
}

func workerFourteen(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 14 done"
}

func workerFifteen(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 15 done"
}

func workerSixteen(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 16 done"
}

func workerSeventeen(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 17 done"
}

func workerEighteen(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 18 done"
}

func workerNineteen(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 19 done"
}

func workerTwenty(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 20 done"
}

func workerTwentyone(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 21 done"
}

func workerTwentytwo(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 22 done"
}

func workerTwentythree(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 23 done"
}

func workerTwentyfour(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 24 done"
}

func workerTwentyfive(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 25 done"
}

func workerTwentysix(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 26 done"
}

func workerTwentyseven(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 27 done"
}

func workerTwentyeight(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 28 done"
}

func workerTwentynine(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 29 done"
}

func workerThirty(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 30 done"
}

func workerThirtyone(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 31 done"
}

func workerThirtytwo(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 32 done"
}

func workerThirtythree(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 33 done"
}

func workerThirtyfour(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 34 done"
}

func workerThirtyfive(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 35 done"
}

func workerThirtysix(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 36 done"
}

func workerThirtyseven(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 37 done"
}

func workerThirtyeight(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 38 done"
}

func workerThirtynine(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 39 done"
}

func workerForty(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 40 done"
}

func workerFortyone(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 41 done"
}

func workerFortytwo(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 42 done"
}

func workerFortythree(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 43 done"
}

func workerFortyfour(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 44 done"
}

func workerFortyfive(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 45 done"
}

func workerFortysix(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 46 done"
}

func workerFortyseven(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 47 done"
}

func workerFortyeight(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 48 done"
}

func workerFortynine(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 49 done"
}

func workerFifty(id int, ch chan<- string) {

    fmt.Printf("Worker %d starting\n", id)
    ch <- "worker 50 done"
}

func main() {
	os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		fmt.Printf("Failed to create wallet: %s\n", err)
		os.Exit(1)
	}

	if !wallet.Exists("appUser") {
		err = populateWallet(wallet)
		if err != nil {
			fmt.Printf("Failed to populate wallet contents: %s\n", err)
			os.Exit(1)
		}
	}

	ccpPath := filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		fmt.Printf("Failed to connect to gateway: %s\n", err)
		os.Exit(1)
	}
	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		fmt.Printf("Failed to get network: %s\n", err)
		os.Exit(1)
	}
	
	contract := network.GetContract("fabcar")

	var newSensorReading string
	    firstchan := make(chan string)
   	    secondchan := make(chan string)
    	thirdchan := make(chan string)
    	fourthchan := make(chan string)
   	    fifthchan := make(chan string)
    	sixthchan := make(chan string)
    	seventhchan := make(chan string)
    	eightchan := make(chan string)
    	ninthchan := make(chan string)
    	tenthchan := make(chan string)
    	eleventhchan := make(chan string)
    	twelvethchan := make(chan string)
    	thirteenthchan := make(chan string)
    	fourteenthchan := make(chan string)
    	fifteenthchan := make(chan string)
    	sixteenthchan := make(chan string)
    	seventeenthchan := make(chan string)
    	eighteenthchan := make(chan string)
    	nineteenthchan := make(chan string)
    	twenteithchan := make(chan string)
        twentyfirstchan := make(chan string)
        twentysecondchan := make(chan string)
        twentythirdchan := make(chan string)
        twentyfourthchan := make(chan string)
        twentyfifthchan := make(chan string)
        twentysixthchan := make(chan string)
        twentyseventhchan := make(chan string)
        twentyeighthchan := make(chan string)
        twentyninthchan := make(chan string)
        thirteithchan := make(chan string)
        thirtyfirstchan := make(chan string)
        thirtysecondchan := make(chan string)
        thirtythirdchan := make(chan string)
        thirtyfourthchan := make(chan string)
        thirtyfifthchan := make(chan string)
        thirtysixthchan := make(chan string)
        thirtyseventhchan := make(chan string)
        thirtyeighthchan := make(chan string)
        thirtyninthchan := make(chan string)
        fourteithchan := make(chan string)
        fortyfirstchan := make(chan string)
        fortysecondchan := make(chan string)
        fortythirdchan := make(chan string)
        fortyfourthchan := make(chan string)
        fortyfifthchan := make(chan string)
        fortysixthchan := make(chan string)
        fortyseventhchan := make(chan string)
        fortyeighthchan := make(chan string)
        fortyninthchan := make(chan string)
        fifteithchan := make(chan string)
	for i := 1; i <= 1000; i++ {
		
		wg.Add(1)
		
		j:=i%50
	
        	switch j {
        	case 1:
            	go func(){
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
			         go workerOne(i, firstchan)
			         return
		             }()
		     fmt.Println(<-firstchan)
		case 2:
            	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerTwo(i, secondchan)
				     return
		        }()
		     fmt.Println(<-secondchan)
	    	case 3:
            	go func(){
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
                     go workerThree(i, thirdchan)
                     return
		        }()
		     fmt.Println(<-thirdchan)
	    	case 4:
            	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerFour(i,fourthchan)
				     return
		        }()
		     fmt.Println(<-fourthchan)
	    	case 5:
            	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerFive(i,fifthchan)
				     return
		        }()
		     fmt.Println(<-fifthchan)
	    	case 6:
            	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerSix(i,sixthchan)
				     return
				     }()
		     fmt.Println(<-sixthchan)
	    	case 7:
            	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerSeven(i,seventhchan)
				     return
		        }()
		    fmt.Println(<-seventhchan)
	    	case 8:
            	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerEight(i,eightchan)
				     return
		        }()
		     fmt.Println(<-eightchan)
	    	case 9:
            	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerNine(i,ninthchan)
				     return
		        }()
		     fmt.Println(<-ninthchan)
		     case 10:
            	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerTen(i,tenthchan)
				     return
			}()
			fmt.Println(<-tenthchan)
		     case 11:
            	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerEleven(i,eleventhchan)
				     return
			}()
			fmt.Println(<-eleventhchan)
		     case 12:
            	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerTwelve(i,twelvethchan)
				     return
			}()
			fmt.Println(<-twelvethchan)
		     case 13:
            	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerThirteen(i,thirteenthchan)
				     return
			}()
			fmt.Println(<-thirteenthchan)
		     case 14:
            	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerFourteen(i,fourteenthchan)
				     return
			}()
			fmt.Println(<-fourteenthchan)
		
		     case 15:
            	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerFifteen(i,fifteenthchan)
				     return
			}()
			fmt.Println(<-fifteenthchan)
		
		     case 16:
            	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerSixteen(i,sixteenthchan)
				     return
			}()
			fmt.Println(<-sixteenthchan)
		
		     case 17:
            	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerSeventeen(i,seventeenthchan)
				     return
			}()
			fmt.Println(<-seventeenthchan)

		     case 18:
            	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerEighteen(i,eighteenthchan)
				     return
			}()
			fmt.Println(<-eighteenthchan)		     
		     case 19:
            	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerNineteen(i,nineteenthchan)
				     return
			}()
			fmt.Println(<-nineteenthchan)
            case 20:
            go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerTwenty(i,twenteithchan)
				     return
			}()
			fmt.Println(<-twenteithchan)
            case 21:
            go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerTwenty(i, twentyfirstchan)
				     return
			}()
			fmt.Println(<- twentyfirstchan)
            case 22:
            go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerTwentytwo(i, twentysecondchan)
				     return
			}()
			fmt.Println(<- twentysecondchan)      
            case 23:
            go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerTwentythree(i, twentythirdchan)
				     return
			}()
			fmt.Println(<- twentythirdchan)
            case 24:
             go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerTwentyfour(i, twentyfourthchan)
				     return
			}()
			fmt.Println(<- twentyfourthchan)
	    	case 25:
	    	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerTwentyfive(i, twentyfifthchan)
				     return
			}()
			fmt.Println(<- twentyfifthchan)
	    	case 26:
	    	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerTwentysix(i, twentysixthchan)
				     return
			}()
			fmt.Println(<- twentysixthchan)
	    	case 27:
	    	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerTwentyseven(i, twentyseventhchan)
				     return
			}()
			fmt.Println(<- twentyseventhchan)
	    	case 28:
	    	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerTwentyeight(i, twentyeighthchan)
				     return
			}()
			fmt.Println(<- twentyeighthchan)
	    	case 29:
	    	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerTwentynine(i, twentyninthchan)
				     return
			}()
			fmt.Println(<- twentyninthchan)
	    	case 30:
	    	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerThirty(i, thirteithchan)
				     return
			}()
			fmt.Println(<- thirteithchan)
	    	case 31:
	    	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerThirtyone(i, thirtyfirstchan)
				     return
			}()
			fmt.Println(<- thirtyfirstchan)
	    	case 32:
	    	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerThirtytwo(i, thirtysecondchan)
				     return
			}()
			fmt.Println(<- thirtysecondchan)
	    	case 33:
	    	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerThirtythree(i, thirtythirdchan)
				     return
			}()
			fmt.Println(<- thirtythirdchan)
	    	case 34:
	    	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerThirtyfour(i, thirtyfourthchan)
				     return
			}()
			fmt.Println(<- thirtyfourthchan)
	    	case 35:
	    	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerThirtyfive(i, thirtyfifthchan)
				     return
			}()
			fmt.Println(<- thirtyfifthchan)
	    	case 36:
	    	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerThirtysix(i, thirtysixthchan)
				     return
			}()
			fmt.Println(<- thirtysixthchan)
	    	case 37:
	    	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerThirtyseven(i, thirtyseventhchan)
				     return
			}()
			fmt.Println(<- thirtyseventhchan)
	    	case 38:
	    	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerThirtyeight(i, thirtyeighthchan)
				     return
			}()
			fmt.Println(<- thirtyeighthchan)
	    	case 39:
	    	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerThirtynine(i, thirtyninthchan)
				     return
			}()
			fmt.Println(<- thirtyninthchan)
	    	case 40:
	    	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerForty(i, fourteithchan)
				     return
			}()
			fmt.Println(<- fourteithchan)
	    	case 41:
	    	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerFortyone(i, fortyfirstchan)
				     return
			}()
			fmt.Println(<- fortyfirstchan)
	    	case 42:
	    	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerFortytwo(i, fortysecondchan)
				     return
			}()
			fmt.Println(<- fortysecondchan)
	    	case 43:
	    	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerFortythree(i, fortythirdchan)
				     return
			}()
			fmt.Println(<- fortythirdchan)
	    	case 44:
	    	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerFortyfour(i, fortyfourthchan)
				     return
			}()
			fmt.Println(<- fortyfourthchan)
	    	case 45:
	    	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerFortyfive(i, fortyfifthchan)
				     return
			}()
			fmt.Println(<- fortyfifthchan)
	    	case 46:
	    	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerFortysix(i, fortysixthchan)
				     return
			}()
			fmt.Println(<- fortysixthchan)
	    	case 47:
	    	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerFortyseven(i, fortyseventhchan)
				     return
			}()
			fmt.Println(<- fortyseventhchan)
	    	case 48:
	    	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerFortyeight(i, fortyeighthchan)
				     return
			}()
			fmt.Println(<- fortyeighthchan)
	    	case 49:
	    	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         defer contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerFortynine(i, fortyninthchan)
				     return
			}()
			fmt.Println(<- fortyninthchan)
	    	case 0:
            	go func() {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
				     go workerFifty(i,fifteithchan)
				     return
		        }()
		    fmt.Println(<-fifteithchan)
	   }
	wg.Done()		
	}

	wg.Wait()

}

func populateWallet(wallet *gateway.Wallet) error {
	credPath := filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return errors.New("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	err = wallet.Put("appUser", identity)
	if err != nil {
		return err
	}
	return nil
}
