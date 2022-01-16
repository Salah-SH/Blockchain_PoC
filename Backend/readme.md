# CagnotteApp
The blockchain node has to be started


## Backend
We have a set of env variables that have to be set:

BASE_URL: A REST interface for state queries, transaction generation and broadcasting.

DATABASE_URL: The url of the database used for user data storage. mongodb server has to be started

ADMIN_ADDR: the blockchain address of the admin user.

ROOT_DIR: the root of the directory of the keyring.


   Set the environment variables for the backend: 
   
   PS: the mongodb server has to be already started
   
   
    ROOT_DIR="$ROOT/data/cli"
    BASE_URL="http://localhost:1317"
    DATABASE_URL="mongodb://localhost:27017/Blockchain"
    
    
   The admin address is also need to be set as an environment variable:
   
   First 
   
    cli keys list
    
   Copy the address of the admin, then: 
   
    ADMIN_ADDR=the address copied
    
   For example
   
    ADMIN_ADDR=cosmos18433huwdwpqq7f507lzsykhdkjgzyalv0qwl3q
   
  Start the blockchain-node rest server:
  
    cli rest-server --chain-id namechain --trust-node

  You have to pass the paiement account parameters as an environment viriables : the piment account number an the authorization token which identify the account of craft foundry receiving the paiements on paymee.
  
  For example 
  
      PAIEMENT_ACCOUNT=1796 AUTHORIZATION_TOKEN=d647198a53cb86379615ace5f2af1c302ee21c08

  the environment variable "REPEAT" indicates the time to wait before the token expires: Time = REPEAT * 5
  
  then, in another terminal: 
 
      ADMIN_ADDR=cosmos18433huwdwpqq7f507lzsykhdkjgzyalv0qwl3q  ROOT_DIR="$ROOT/data/cli" PAIEMENT_ACCOUNT=1711  AUTHORIZATION_TOKEN=45b82b861967508cb279d08c5e7ad4618075ee7d REPEAT=50 go run main.go  
    

  
  
  
