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

	for i := 1; i <= 1000; i++ {
		
		wg.Add(1)
		
		j:=i%10
	
        	switch j {
        	case 1:
            	go func() ([]byte, error) {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         result, err := contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
			         if err != nil {
				        return result, fmt.Errorf("failed to evaluate transaction: %v", err)
			         }
			         go workerOne(i, firstchan)
			         return result, nil
		             }()
		     fmt.Println(<-firstchan)
		case 2:
            	go func() ([]byte, error) {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         result, err := contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
			         if err != nil {
				        return result, fmt.Errorf("failed to evaluate transaction: %v", err)
			         }
				 go workerTwo(i, secondchan)
			         return result, nil
		        }()
		     fmt.Println(<-secondchan)
	    	case 3:
            	go func() ([]byte, error) {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         result, err := contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
			         if err != nil {
				        return result, fmt.Errorf("failed to evaluate transaction: %v", err)
			         }
                     go workerThree(i, thirdchan)
			         return result, nil
		        }()
		     fmt.Println(<-thirdchan)
	    	case 4:
            	go func() ([]byte, error) {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         result, err := contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
			         if err != nil {
				        return result, fmt.Errorf("failed to evaluate transaction: %v", err)
			         }
				 go workerFour(i,fourthchan)
			         return result, nil
		        }()
		     fmt.Println(<-fourthchan)
	    	case 5:
            	go func() ([]byte, error) {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         result, err := contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
			         if err != nil {
				        return result, fmt.Errorf("failed to evaluate transaction: %v", err)
			         }
				 go workerFive(i,fifthchan)
			         return result, nil
		        }()
		     fmt.Println(<-fifthchan)
	    	case 6:
            	go func() ([]byte, error) {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         result, err := contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
			         if err != nil {
				        return result, fmt.Errorf("failed to evaluate transaction: %v", err)
			         }
				 go workerSix(i,sixthchan)
			         return result, nil
		        }()
		     fmt.Println(<-sixthchan)
	    	case 7:
            	go func() ([]byte, error) {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         result, err := contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
			         if err != nil {
				        return result, fmt.Errorf("failed to evaluate transaction: %v", err)
			         }
				 go workerSeven(i,seventhchan)
			         return result, nil
		        }()
		    fmt.Println(<-seventhchan)
	    	case 8:
            	go func() ([]byte, error) {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         result, err := contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
			         if err != nil {
				        return result, fmt.Errorf("failed to evaluate transaction: %v", err)
			         }
				 go workerEight(i,eightchan)
			         return result, nil
		        }()
		     fmt.Println(<-eightchan)
	    	case 9:
            	go func() ([]byte, error) {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         result, err := contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
			         if err != nil {
				        return result, fmt.Errorf("failed to evaluate transaction: %v", err)
			         }
				 go workerNine(i,ninthchan)
			         return result, nil
		        }()
		     fmt.Println(<-ninthchan)
	    	case 0:
            	go func() ([]byte, error) {
            	     newSensorReading = "SensorRead"+strconv.Itoa(i)
			         result, err := contract.SubmitTransaction("createSensorRead", newSensorReading, "Buddy")
			         if err != nil {
				        return result, fmt.Errorf("failed to evaluate transaction: %v", err)
			         }
				 go workerTen(i,tenthchan)
			         return result, nil
		        }()
		    fmt.Println(<-tenthchan)
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
