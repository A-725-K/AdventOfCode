# Lesson Learned
This file is a reminder for myself to understand what I did not know
when solving the AoC challenges, and how can I improve myself next
time I will face a similar problem. As well as some noticeable facts that
happened during the AoC.

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