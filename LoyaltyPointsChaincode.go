/*
Loyalty_Points_Use_Case
*/

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// LoyaltyPointsChaincode example simple Chaincode implementation
type LoyaltyPointsChaincode struct {
}

func (t *LoyaltyPointsChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Init called, initializing chaincode")
	
	//var A, B string    // Entities
	//var Aval, Bval,Apts,Bpts int // Asset holdings
	var err error
	var customerNamePrint string


       if len(args) == 2 {

          customer_name := args[0]
          merchant_name := args[0]+"#"+args[1]

		  customerNamePrint=customer_name

        fmt.Printf("Customer Name= %s, Merchant Name with Customer Name= %s\n", customer_name, merchant_name)

        err = stub.PutState(customer_name, []byte(merchant_name))
	if err != nil {
		return nil, err
	}

       if args[1]=="K" {
          
      fmt.Printf("Enrolling loyalty points=%d\n",100 )

        err = stub.PutState(merchant_name, []byte(strconv.Itoa(100)))
	if err != nil {
		return nil, err
	}
       }
     if args[1]=="S" {
       fmt.Printf("Enrolling loyalty points=%d\n",150 )

        err = stub.PutState(merchant_name, []byte(strconv.Itoa(150)))
	      if err != nil {
		      return nil, err
	         }

      } 
	  
    
		
	}


      if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}


//==================================================================================================
 
    

         gettingData, err := stub.GetState(customerNamePrint)
	
//==================================================================================================

	return gettingData, nil
}

// Transaction for Customer for earning and transferring Loyalty points
func (t *LoyaltyPointsChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Running invoke")
	
	//var A, B string    // Entities
	//var Aval, Bval int // Asset holdings
	var loyalty_points_value int          // Transaction value
	var err1 error


        if len(args) == 2 {
		
         customer_name, err := stub.GetState(args[0])
          if err != nil {
		    return nil, err
	          }

          s := make([]string,len(customer_name))
          for i := range customer_name {
         s[i] = strconv.Itoa(int(customer_name[i]))
          }
         customer_name_temp:=strings.Join(s,"")

         loyalty_points, err := stub.GetState(customer_name_temp)
          if err != nil {
		return nil, err
	        }
       
  


       //s1 := make([]string,len(loyalty_points))
        //  for j := range loyalty_points {
        // s1[j] = strconv.Itoa(int(loyalty_points[j]))
        //  }
       //  loyalty_points_temp:=strings.Join(s1,"")

		loyalty_points_value, _ = strconv.Atoi(string(loyalty_points))
		




        transaction_value,_:= strconv.Atoi(args[1])


          cust_merchnatname := strings.Split(customer_name_temp, "#")
              merchant_name := cust_merchnatname[1]




        if merchant_name=="K" {
var transaction_value_temp int=transaction_value/25

loyalty_points_value+=transaction_value_temp

        }
 if merchant_name=="S"{
var transaction_value_temp int=transaction_value/20
loyalty_points_value+=transaction_value_temp
}

//fmt.Printf("Total loyalty Points, %d",loyalty_points_value)

 err1 = stub.PutState(customer_name_temp, []byte(strconv.Itoa(loyalty_points_value)))
if err1 != nil {
		return nil, err1
	}






	}

	if len(args) == 3{



  customer_name, err := stub.GetState(args[0])
          if err != nil {
		    return nil, err
	          }

          s := make([]string,len(customer_name))
          for i := range customer_name {
         s[i] = strconv.Itoa(int(customer_name[i]))
          }
         customer_name_temp:=strings.Join(s,"")

         loyalty_points, err := stub.GetState(customer_name_temp)
          if err != nil {
		return nil, err
	        }

fmt.Printf("Enrolling loyalty points=%d\n",loyalty_points )

cust_merchnatname := strings.Split(customer_name_temp, "#")
              merchant_name := cust_merchnatname[1]

			  fmt.Printf("Enrolling loyalty points=%s\n",merchant_name )
//if



	}



	if len(args) != 2  && len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}



	return nil, nil
}

//==================================================================================================



// Query callback representing the query of a chaincode
func (t *LoyaltyPointsChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Query called, determining function")
	
	
//var key, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

    customer_name, err := stub.GetState(args[0])
          if err != nil {
		    return nil, err
	          }

          s := make([]string,len(customer_name))
          for i := range customer_name {
         s[i] = strconv.Itoa(int(customer_name[i]))
          }
         customer_name_temp:=strings.Join(s,"")

         loyalty_points, err := stub.GetState(customer_name_temp)

		// if err != nil {
		//jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
	//	return nil, errors.New(jsonResp)
	//}

		// loyalty_points_value, _ = strconv.Atoi(string(loyalty_points))

	
	

	return loyalty_points, nil

	
}

func main() {
	err := shim.Start(new(LoyaltyPointsChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}