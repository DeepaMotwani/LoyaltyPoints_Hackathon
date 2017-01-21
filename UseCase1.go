/*
Loyalty_Points_Use_Case
*/

package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// LoyaltyPointsChaincode example simple Chaincode implementation
type LoyaltyPointsChaincode struct {
}

func (t *LoyaltyPointsChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Printf("Init called, initializing chaincode")
	
	//var A, B string    // Entities
	//var Aval, Bval,Apts,Bpts int // Asset holdings
	var err error


       if len(args) == 2 {

          customer_name := args[0]
          merchant_name := args[0]+"#"+args[1]

        fmt.Printf("Customer Name= %s, Merchant Name with Customer Name= %s\n", customer_name, merchant_name)

        err = stub.PutState(customer_name, []byte(merchant_name))
	if err != nil {
		return nil, err
	}

       if(args[1]=="K")
       {
          
      fmt.Printf("Enrolling loyalty points=%d\n",100 )

        err = stub.PutState(merchant_name, []byte(strconv.Itoa(100))))
	if err != nil {
		return nil, err
	}
       }
    else if(args[1]=="S")
      {
       fmt.Printf("Enrolling loyalty points=%d\n",150 )

        err = stub.PutState(merchant_name, []byte(strconv.Itoa(150))))
	if err != nil {
		return nil, err
	}

      } 
    
		
	}


      if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}


==================================================================================================

	
==================================================================================================

	return nil, nil
}

// Transaction for Customer for earning and transferring Loyalty points
func (t *LoyaltyPointsChaincode) invoke(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	fmt.Printf("Running invoke")
	
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var X int          // Transaction value
	var err error


        if len(args) == 2 {
		

         customer_name, err := stub.GetState(args[0])


          s := make([]string,len(customer_name))
          for i := range b {
         s[i] = strconv.Itoa(int(b[i]))
          }
         customer_name_temp:=strings.Join(s,",")

         loyalty_points, err := stub.GetState(customer_name_temp)

       
       s1 := make([]string,len(loyalty_points))
          for j := range c {
         s1[i] = strconv.Itoa(int(c[i]))
          }
         loyalty_points_temp:=strings.Join(s1,",")

        loyalty_points_value := strconv.Atoi(loyalty_points_temp)




        transaction_value := strconv.Atoi(args[1])


          cust_merchnatname := strings.Split(customer_name_temp, "#")
             customer_name_new, merchant_name := cust_merchnatname[0], cust_merchnatname[1]




        if(merchant_name=="K")
        {
var transaction_value_temp int=transaction_value/25

loyalty_points_value+=transaction_value_temp

        }
else if(merchant_name=="S")
{
var transaction_value_temp int=transaction_value/20
loyalty_points_value+=transaction_value_temp
}


 err = stub.PutState(customer_name_temp, []byte(strconv.Itoa(loyalty_points_value))))
if err != nil {
		return nil, err
	}



	}

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}



	return nil, nil
}

==================================================================================================

// Invoke callback representing the invocation of a chaincode
// This chaincode will manage two accounts A and B and will transfer X units from A to B upon invoke
func (t *LoyaltyPointsChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Printf("Invoke called, determining function")
	
	// Handle different functions
	if function == "invoke" {
		// Transaction makes payment of X units from A to B
		fmt.Printf("Function is invoke")
		return t.invoke(stub, args)
	} else if function == "init" {
		fmt.Printf("Function is init")
		return t.Init(stub, function, args)
	} else if function == "delete" {
		// Deletes an entity from its state
		fmt.Printf("Function is delete")
		return t.delete(stub, args)
	}

	return nil, errors.New("Received unknown function invocation")
}

func (t* LoyaltyPointsChaincode) Run(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Printf("Run called, passing through to Invoke (same function)")
	
	// Handle different functions
	if function == "invoke" {
		// Transaction makes payment of X units from A to B
		fmt.Printf("Function is invoke")
		return t.invoke(stub, args)
	} else if function == "init" {
		fmt.Printf("Function is init")
		return t.Init(stub, function, args)
	} else if function == "delete" {
		// Deletes an entity from its state
		fmt.Printf("Function is delete")
		return t.delete(stub, args)
	}

	return nil, errors.New("Received unknown function invocation")
}

// Query callback representing the query of a chaincode
func (t *LoyaltyPointsChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Printf("Query called, determining function")
	
	if function != "query" {
		fmt.Printf("Function is query")
		return nil, errors.New("Invalid query function name. Expecting \"query\"")
	}
	var A string // Entities
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return nil, errors.New(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return nil, errors.New(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return Avalbytes, nil
}

func main() {
	err := shim.Start(new(LoyaltyPointsChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}