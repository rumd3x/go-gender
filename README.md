# go-gender

Find person gender by name. Works better with Brazilian or International Names.

[![GoDoc](https://godoc.org/github.com/rumd3x/go-gender?status.png)](https://godoc.org/github.com/rumd3x/go-gender)
![GitHub release](https://img.shields.io/github/release/rumd3x/go-gender.svg)
[![Build Status](https://travis-ci.org/rumd3x/go-gender.svg?branch=master)](https://travis-ci.org/rumd3x/go-gender)
[![codecov](https://codecov.io/gh/rumd3x/go-gender/branch/master/graph/badge.svg)](https://codecov.io/gh/rumd3x/go-gender)
![License](https://img.shields.io/github/license/rumd3x/go-gender.svg)

## How it works

This project uses the API from Brazilian Institute of Geography and Statistics (IBGE) to gather data and predict gender statiscally

## Install

```sh
go get github.com/rumd3x/go-gender
```

## Usage

Read the [GoDoc](https://godoc.org/github.com/rumd3x/go-gender) for all available methods and their description.

## Examples

```go
brianGender, _ := gender.GetGender("Brian")
jenGender, err := gender.GetGender("Jennifer")

if err != nil {
    panic(err)
}

if !jenGender.Equals(brianGender) {
    fmt.Println("Names Jennifer and Brian are not of the same gender")
}

if jenGender.Female() {
    fmt.Println("Jennifer is a female")
}

if brianGender.Male() {
    fmt.Println("Brian is a Male")
}

if !brianGender.Unisex() {
    fmt.Println("Brian is not a Unisex name")
}

fmt.Printf("Name Brian is of gender %s", brianGender.String()) // Outputs: Name Brian is of gender Male
```
