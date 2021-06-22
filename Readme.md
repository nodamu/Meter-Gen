## MeterGen App

* Requirements
    - Docker
    - Golang >= 1.15
    - MongoDb >= 4.x
    - Python >=3.6


* To to start meter data generator run  
```  
./run.sh  
```

* Then run the command below to start mongodb, rabbitmq and the api service
```
docker-compose up
```

* Access the fast api documention at localhost:8008/docs