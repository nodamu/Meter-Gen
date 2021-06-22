## MeterGen App

* Requirements
    - Docker
    - Golang >= 1.15
    - MongoDb >= 4.x
    - Python >=3.6

* Then run the command below to start mongodb, rabbitmq and the api service
```
docker-compose up
```
* Access the fast api documention at localhost:8008/docs

* If you want to start meter data generator manually run  
```  
./run.sh  
```
* ./run.sh is a wrapper around 
```
./main --startdate=1231233423 --meterid=hello --mqaddress=amqp://guest:guest@localhost:5672/
```
* NB: startdate is in unix time format