# go-exercise

## Usage

### Fibonacci
To calculate the Fibonacci sequence of a number enter the following command:
```shell
go run main.go fibonacciSequence -number=10
```
```-n``` is the nth number of sequence to be calculated.

### Palindrome API
To run the Palindrome API run the following command:
```shell
go run main.go palindromeApi -port=8080
```
```-port``` is the port where the resource to be exposed. By default, it's 80.  

To check if a sentence is a palindrome run curl below:
```shell
curl --location --request GET 'http://localhost:80/palindrome/validate?sentence=racecar'
```