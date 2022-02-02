# run go

go run go/*.go gingerbreadtemplate.uk ns-1401.awsdns-47.org ns-1722.awsdns-23.co.uk ns-402.awsdns-50.com ns-982.awsdns-58.com

# build and run
go build -o remap go/*.go    
remap gingerbreadtemplate.uk ns-1401.awsdns-47.org ns-1722.awsdns-23.co.uk ns-402.awsdns-50.com ns-982.awsdns-58.com

# run docker
docker build -t tester .
docker run -it tester   