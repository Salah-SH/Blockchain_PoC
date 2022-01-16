# PoCBlockchain

# CagnotteApp




## Build, initialise and start the blockchain application

  **Build the application :**
  
  Install the binaries that will be served to create and initialise and communicate with the blockchain application.
   
    make build
  **Initialise the blockchain application :** 
 
 
  The blockchain node:
  
  1- Create a blockchain node and name it nodename, the chain-id parameter refers to the name of the chain of blocks that will be stored on the node:
  
     svc init nodename --chain-id=namechain
     
  2- Configure the client ( the user executing transactions ):
  
   Specify the output format:
    
     cli config output json
     cli config indent true
   
   Trust connected node : don't verify proofs for responses. If the node is distrusted, various checks are done on the response received from the node on the        executed queries by the client.
     
     cli config trust-node true
     
   Set the name of the chain of blocks where the transaction will be saved, this has to be the same as the name specified when creating the node:
   
     cli config chain-id namechain
     
   The keyring holds the private/public keypairs. Backend are where the private keys are stored, the default is os,
   here we store the account keys in the app's configuration directory unencrypted. The keys can then be found in "$HOME/.cli/keyring-test-cagnotte".
   
     cli config keyring-backend test
     
  3- For our application, we need an admin account. At this level, we are going to create the admin keys, then allocate the admin account 
  with the tokens. We will also save the admin adress in the initial state. The address will be needed later by the application.
   
   Create admin keys
   
    cli keys add admin
   Add Admin account in the initial state of the application and allocate it with two types of tokens: nametoken and stake.
   
    svc add-genesis-account $(cli keys show admin -a) 100000nametoken,100000000stake
    
   Register the admin address in the genesis state.
   
    svc add-admin $(cli keys show admin -a)
    
   For the reasons of test, we will initial a user in the genesis state:
   
    cli keys add user
    svc add-genesis-account $(cli keys show user -a) 100nametoken
    
    
   4- Create a validator node (in our case it will be the same node we have created).
   
   This command is used to add a validator node to the chain. In this case the admin has become a validator operator ( the account that controls the validator node ).
   This command also links the operator account ( in our case "admin" ) with a Tendermint node pubkey that will be used for signing blocks.
   
    svc gentx --name admin --keyring-backend test
    
   Register the transaction of the creation of the validator in the genesis state:
  
    svc collect-gentxs
    
   5- Validate the genesis state:
   
    svc validate-genesis
  
  Ps: all these commands can be executed by executing the 'svc-init' script:
  
     cd /bin
    ./svc-init
    
  **Start the blockchain node :**
   
    svc start
    


## Test the application 

First check the accounts to ensure they have funds 
 
    cli query account $(cli keys show admin -a)
    cli query account $(cli keys show user -a)

Create a cagnotte, the parameter gas-prices indicates how much token the user is willing to pay per gas unit. 

    cli tx cagnotte create-cagnotte cagnotte_test  --from user --gas-prices=0.0000025nametoken

Contribute to a cagnotte: 

    cli tx cagnotte add-cagnotte cagnotte_test 10 --from user --gas-prices=0.000003nametoken
    
Get cagnotte "cagnotte_test" informations:

    cli tx cagnotte get-cagnotte cagnotte_test

The admin confirms the tx:

    cli tx cagnotte confirm-tx [name of cagnotte] [amount] [user address] [success (true/false)] --from admin --gas-prices=0.00001nametoken
   
   Example
   
    cli tx cagnotte confirm-tx cagnotte_test 10 $(cli keys show user -a) true --from admin  --gas-prices=0.00001nametoken
    
 

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos Tutorials](https://tutorials.cosmos.network)
- [Channel on Discord](https://discord.gg/W8trcGV)
