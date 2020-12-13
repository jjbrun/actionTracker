# ActionTracker

ActionTracker is a Golang library for keeping statistics of performed actions.

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

go get -u github.com/jjbrun/actionTracker

## Examples

```
func main() {
	jump100 := `{"action":"jump", "time":100}`
	jump200 := `{"action":"jump", "time":200}`
	run75 := `{"action":"run", "time":75}`

	at := actionTracker.NewTracker()

	at.AddAction(jump100)
	at.AddAction(jump200)
	at.AddAction(run75)

	fmt.Println(at.GetStats())
}
```

## Output

```
adding act:ion: {jump 100}
adding act:ion: {jump 200}
adding act:ion: {run 75}
[{"action":"jump","avg":150},{"action":"run","avg":75}]
```

## Running Unit Test

- Download the code
- Navigate to the root of the project
- Run all unit tests:
```
go test ./...
```


## Future Considerations

- Only supporting a collection of predefined actions 
- Add a database for persistence of the action statistics
