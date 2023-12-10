# Lesson Learned
This file is a reminder for myself to understand what I did not know
when solving the AoC challenges, and how can I improve myself next
time I will face a similar problem. As well as some noticeable facts that
happened during the AoC.

# Day 10
Today mark my first defeat! :( I was not able at all to figure out a solution
for part 2 (and partly also part one). I had to peek to one of the top player
well known from previous year, namely
![hyper-neutrino](https://www.youtube.com/@hyper-neutrino). It is very smart
the solution and it does not take too many lines of complicated code to present
the correct numbers. The idea of trying to count how many pipes I needed to
traverse, and the reasoning behind the shape of the corner pipes in contrast
with the direction we were scanning the maze, was not something that I could
have figure out by myself alone. Thanks once again AoC to making me learn new
techniques to add to the arsenal! To summarize today's experience in some sort
of coding quotes: knowledge++.

# Day 9
Very fun problem to solve, lightweighted compared to other to let people enjoy
the weekend :) Some slow down due to the debugging of part 2, but writing down
the computations on paper gave its results! After writing down the examples on
the small input, I understood where I was mistaking. Cool!

# Day 8
That sensation when you start to see patterns \*w*/ For the first time in this
AoC I had a flash of inspiration! Realizing that computing the LCM instead of
brute-forcing the solution would had led me to an optimal solution was a boost
of morale. I am very excited to acknowledge how fast I was to solve today's
riddle. Very nice!

# Day 7
This problem really required my utmost attention. After a first attempt a good
refactoring of the code was needed. I am pretty satisfied with how I shaped the
code eventually. Today, debugging has been more difficult than the usual since
the main logic had to be wrapped in the "lambda" function that determined the
sorting, and that function is called a lot of times, especially with the real
input ($O(n * log(n))$). One last note is regarding the very special corner
case in the second problem that got me for a moment, when all the cards in a
round were all jokers. Fun to solve it!

# Day 6
Very fun and easy challenge, with a little bit of physics involved. It has been
a race against time, like the assigment topic :)

# Day 5
Not satisfied with the implementation that got me the 2nd star :( It took about
18 minutes to run on my super-powerful machine, it is not an optimal solution I
would say. Time to refactor and see where it goes... After reviewing my Golang
knowledge I tried to scale horizontally using Go-routines. That lead to a
better runtime, ~4 minutes, still not satisfying though. I am clearly missing
the point of the problem. There should be a way to optimize the algorithm,
either by pruning A LOT of cases to test, or by using some smart optimization
to avoid recomputing the same values over and over.

# Day 4
Nothing really exciting to say about today. Just acknowledging that reading
very carefully the text of the assignment enabled me to get to a solution in a
reasonably short time. As always some parsing-gymnastics has to be done.

# Day 3
The problems start to get tricky very earlier this year \O.O/!! Very important
to do not overcomplicate and overthink the solution, sometimes it is way easier
than it looks like. One general consideration, even if it seems harder while
parsing, try to give sane direction to the coordinate axes when you store the
2D matrix. Sometimes a different orientation while ease the reasoning! Nice
exercise to test off-by-1 or end-of-line corner cases.

# Day 2
Parsing gymnastics. Very fun and still very easy, once you figure out how to
organize properly the input. READ CAREFULLY THE EXAMPLE! That said, let's go
to the next!

# Day 1
A bit of stretching and... let's start! AoC once again, the most wonderful time
of the year is colored with ton of fun given by the joy of solving the riddles!
Can't wait for tommorow :)
