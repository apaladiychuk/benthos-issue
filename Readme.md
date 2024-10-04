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

***input_http_server_post*** it is span with child spans for each processor and output (from stream with input)

each others are spans for each processor and output but without root span and looks like independent (from stream without input) 
Root span is not created for stream without declared input. 


### My goal 

How I can define root span for stream without declared input 
