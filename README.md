**This project setups a dummy tax calculator gRPC Service.**
#### Client implementation has been done for 
#####    1. Rest endpoint using gRPC-gateway
#####    2. Using protobufs
#
#### For setup 
##### Install go (version > 1.12) and setup the PATH and ENV variables
#
#### Inorder to execute below commands move to the taxCalculator directory
#####    cd taxCalculator/
##
#### To build:
#####    *make build*
##
#### To run grpc-server:
#####    *make grpc-server*
##
#### To run rest-server:
#####    *make rest-server*
##
#### To run client:
#####    *make grpc-client*
##
#### To run tests:
#####    *make test*
##
#### To run load tests(sending 200 requests at a time):
#####    *make load-test*
###### Used ghz project https://github.com/bojand/ghz - needs installation on local

#
## NOTE
#### *Client implementation for protobufs have 2 types of dummy requests - Buying a house and Renting a house, and currently client code calls with a request returning buying a house*
#### *For Rest, Postman collections are available in the folder*