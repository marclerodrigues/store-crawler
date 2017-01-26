# store-crawler

This is a simples store scraper build with Go(Golang).
The main goal for this project is to be a simple, fast and run in parallel.

# Running

After cloning this repo, run `go get` to install the dependencies.
Then run the `go run store-crawler/main.go` to see the program in action.

# Next Steps

1. Add tests, and test with different stores
2. Fetch stores in parallel
3. Create a rest end-point to submit/return data as json
4. Create a rest end-point to check the health of the application
5. Implement functionality of price checker, which will keep looking for a product until the price change and will make a post to some url passed.
