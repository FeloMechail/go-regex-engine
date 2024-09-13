# Todo List: Implement a Regex Engine with NFA in Go

1. [x] Set up the project structure
   - [x] Setup Frontend using svelte
2. [x] Implement literal string matching

3. [x] Implement basic character matching
   - [x] Single character matching (e.g., 'a', 'b', 'c')
   - [x] Any character matching (e.g., '.')
   - [ ] implement multiple lines

4. [ ] Design and implement NFA structure
   - [x] Design Stack data structure and RPN parser using Shunting yard Algo
   - [ ] Define State and NFA structs
   - [ ] Implement state transitions

5. [ ] Implement Thompson's construction algorithm
   - [ ] Convert regex to NFA

6. [ ] Implement character classes
   - [ ] Character set matching (e.g., [abc])
   - [ ] Negated character set matching (e.g., [^abc])

7. [ ] Implement quantifiers
   - [ ] Zero or more (*)
   - [ ] One or more (+)
   - [ ] Zero or one (?)

8. [ ] Implement alternation (|)

9.  [ ] Implement grouping with parentheses ()

10. [ ] Implement anchors
    - [ ] Start of string (^)
    - [ ] End of string ($)

11. [ ] Implement escape characters (\)

12. [ ] Implement NFA simulation for pattern matching

13. [ ] Optimize NFA simulation (e.g., lazy evaluation)

14. [ ] Refactor and optimize the code

15. [ ] Add more advanced features (optional)
    - [ ] Lazy quantifiers
    - [ ] Lookahead and lookbehind assertions
    - [ ] Named capture groups

16. [ ] Comprehensive testing and benchmarking
