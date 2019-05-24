# Zest Go Security

## Install 

```go get -u github.com/abhisheklalzest/security```

A go package to check the validity of JWT token.

### Example

```
serr := security.Validate("<Token>")
if serr != nil {
		fmt.Println("Invalid")
} else {
		fmt.Println("Valid")
}
```
