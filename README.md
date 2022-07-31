# stream-service
    This is a demo for testing streaming rpc.
## How to use this demo 
    1 go build -o stream-srv
    2 ./stream-srv
    3 cd bin && go build -o client
    4 echo '{"count": 5}' > config
    5 ./client stream-demo ping-pong -f config --address 127.0.0.1:port --insecure
 # Note
    The stream-srv will listen on a random port.
    
