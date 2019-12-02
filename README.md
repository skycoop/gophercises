# gophercises
These are my solutions for Jon Calhoun's [gophercises](https://gophercises.com/)

## Exercise 1: [Quiz](https://github.com/gophercises/quiz) ([solution](quiz/main.go))
Completed parts 1 and 2, as well as both bonus tasks. I was already sold on Go's concurrency abstraction but writing 
this drove home how powerful it could be. I'm very pleased with the simplicity, although the stupid issue with the 
random seed is driving me crazy. For some reason I keep getting the same value out of `time.Now()` on my Windows 
10 system. I'm still a little unsure on how best to handle command line flags. The `flags` module is great for handling 
the parsing, but I feel like I'm using it slightly wrong...

## Exercise 2: [URL Shortener](https://github.com/gophercises/urlshort) ([solution](urlshort))
I mostly solved the problem. YAML is easy so I ignored that. Overall, this solution is on the way to something more
complicated and feature complete, so I think I'll loop back to it after doing some more excercises.

## Exercise 3: [Choose Your Own Adventure](https://github.com/gophercises/cyoa) ([solution](cyoa))
Solution pending
