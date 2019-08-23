# go-gender

Find person gender by name

## Install

```sh
go get github.com/rumd3x/go-gender
```

## Examples

```go
brianGender, _ := gender.GetGender("Brian")
jenGender, err := gender.GetGender("Jennifer")

if err != nil {
    panic(err)
}

if !jenGender.Equals(brianGender) {
    fmt.Println("Names Jennifer and Brian are of opposite gender")
}

fmt.Printf("Name Brian is of gender %s", brianGender.String())
```
