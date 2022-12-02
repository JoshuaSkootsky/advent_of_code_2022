## Notes for Day 2

- Elves like snacks this year.
- Note the problem refers to the plural as elves, not elfs.

- The word 'encrypted' in the question threw me off, I thought I'd need to decrypt something before I finished reading

- I also thought I might need to randomly simulate the strategy guide against random inputs.

- The first question is how to quickly add up one round, then repeat that process.

- Had a great moment with Embed where the compiler caught `//go:embed input_day2.txt` instead of `//go:embed input_day02.txt`. That's really cool.

### Reflections

- I ended up using a map to make a table to store results
- I ended up needing to write tests because I was getting mixed up between when to win and when to lose. Even with "descriptive" constant names, I still got mixed up and needed tests, but the clear code and descriptive constant names made it easy to see how to fix when the tests pointed out that something was wrong.
- I thought for a moment that having "X", "Y", and "Z" encoded in two different places to mean two different things would result in some problems, but this did not occur.
