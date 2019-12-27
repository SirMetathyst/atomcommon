prog = atom component add -p atomcommon -n

test:
	go test -race -coverprofile=coverage.txt -covermode=atomic

testreport: test
	go tool cover -html=coverage.txt -o ./coverage.html

bench:
	go test -bench . -benchmem -benchtime 15s -cpu 1,4 -count 2

clean:
	rm -f ./coverage.html
	rm -f ./coverage.txt

generate:
	$(prog) LocalPosition2 -d x:float32 -d y:float32 
	$(prog) LocalRotation2 -d x:float32 -d y:float32
	$(prog) LocalScale2    -d x:float32 -d y:float32
	$(prog) Velocity2 -d x:float32 -d y:float32 