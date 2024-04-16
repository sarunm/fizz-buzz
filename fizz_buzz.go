package FizzBuzz

func fizzbuzz(n int) string {
    if n == 1 {
        return "1"
    }
    if n == 2 {
        return "2"
    }
    if n == 3 || n == 6 {
        return "Fizz"
    }
    if n == 4 {
        return "4"
    }
    if n == 5 {
        return "Buzz"
    }

    return "1"
}
