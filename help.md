# Help

### Test Apache Bench (Windows)

-> start apache + open terminal
-> run: .\apache\bin\ab.exe -k -c 60 -n 90 -m POST http://localhost:4000/game

### Run Test

-> go test ./... -v -cover
