### Root span benthos issue 

There are  two benthos stream defined in code 
1.  [Input is not defined](src/stream.yaml). Use ProductionFunc for start stream 
2. [With input](src/stream-input.yaml). Defined input http_server on 8081 

Run containers 
```shell
docker compose up -d 
```


execute stream 
```shell
curl -X POST http://localhost:8081
curl -X POST http://localhost:8080
```


Open [jaeger dashboard](http://localhost:16686)  

You should see result like [this](screenshot.png)

Root span is not created for stream without declared input. 
