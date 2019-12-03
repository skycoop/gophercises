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
complicated and feature complete, so I think I'll loop back to it after doing some more exercises.
TODOs:
- [ ] Investigate building a full webapp with [Buffalo](https://gobuffalo.io/en/)

## Exercise 3: [Choose Your Own Adventure](https://github.com/gophercises/cyoa) ([solution](cyoa))
Fairly simple one. The templating library is a bit odd, but not terrible to work with. 
TODOs:
- [ ] Investigate using [pkger](https://github.com/markbates/pkger) for embedding the template files in the binary
- [ ] Make it prettier
- [ ] CLI version

## Exercise 4: [HTML Link Extractor](https://github.com/gophercises/link) ([solution](link))
XML is the worst :confounded: I got the href extraction working, don't really feel like doing the text extraction. The
`testing` module is cool, but I can see why there are frameworks built on top of it. No one wants to write the standard
set of equality operations required.
TODOs:
- [ ] Refactor the recursive algorithms to use goroutines
- [ ] Do the text extraction
- [ ] Try a testing framework like [Testify](https://github.com/stretchr/testify) or maybe a BDD framework if I'm 
feeling fancy

## Exercise 5: [Sitemap Builder](https://github.com/gophercises/sitemap)
Well if the point of #4 was just to build an href extractor, I think I might be ok with where I left it 
:stuck_out_tongue_closed_eyes:. I'm going to skip this one for the moment and come back when I want to wrangle with XML.

## Exercise 6: [HackerRank](https://github.com/gophercises/hr1)
Huh well I guess this counts as an exercise.
Solutions:
- [CamelCase](https://www.hackerrank.com/challenges/camelcase/submissions/code/132837633)
- [Caesar Cipher](https://www.hackerrank.com/challenges/caesar-cipher-1/submissions/code/132839737)

## Excercise 7: [TODO List](https://github.com/gophercises/task) ([solution](todo))
Oh sweet this exercise (despite it's non-obvious name) is exactly what I wanted. I wanted to fiddle with 
[cobra](https://github.com/spf13/cobra) and [BoltDB](https://github.com/etcd-io/bbolt) (I'm using the etcd fork since
the original isn't being updated and I like etcd).