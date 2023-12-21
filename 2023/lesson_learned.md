# Lesson Learned
This file is a reminder for myself to understand what I did not know
when solving the AoC challenges, and how can I improve myself next
time I will face a similar problem. As well as some noticeable facts that
happened during the AoC.

# Day 22
Very intricated problem, that stretched my imagination skills. It is a lot
harder trying to visualize data in 3D rather than in 2D. It has not been very
easy to deal with this assignment, especially the first part where I spent most
of the time without really being able to figure out a satisfying solution. I had
to peak a bit from the top-notch players of this year AoC Youtube channels to
get a reasonable explanation on how to deal with this kind of data without
wasting all my available resources in terms of computational power. It has been
quite smart to divide the blocks in "those which support" from "those which are
supported by", it simplified a lot the reasoning for part 2. Eventually, I was
able to get the golden star all by myself, after a long shower used to figure
out the algorithm :) Solid!

# Day 21
Another defeat, but this time more lighthearted :) Part 1 has not been too hard
to solve, but for part 2, even though I found a nice optimization of the
algorithm I previously used, I had to wave the white flag. My solution was
consuming too much resources, I even wrote a function to check that the RAM of
my poor laptop (given that is 32GB!!!) and stop if the usage was growing too
much. Since I am coding the solutions in Golang it has been interesting to see
how to use the *runtime* library and the power of *goroutines* to achieve so.
When I watched few Youtube videos from another outstanding AoC player named
![Anshuman Dash](https://www.youtube.com/@anshumandash1304), I realized that I
would have never reached that solution. This time, I considered *hyper-neutrino*
's one too convoluted to try to give it a shot, even though it is incredibly
smart and coincise. To wrap up today's experience I would like to say kudos to
all the players able to solve the problem themselves! For me, instead, it is
still a long way to the (metaphorical) top :) Climbing up!

# Day 20
Very cool problem! For a moment I though an eletrician would have been needed to
solve the problem, but eventually was very fun to reason about and solve. No
major concerns today to be honest, the only remark is to be careful to possible
rabbit-holes that you end up in, like TRYING TO RUN THE SAME ALGORITHM OF PART 1
TO SOLVE PART 2 ;) Shocking!

# Day 19
Today I had a lot of fun! The problem was not incredibly difficult but few trap
were hidden between the lines. The silver star was almost immediate, after
figuring out how to store in proper data structures the input. For the second
one instead I had an intuition that turned out to be the correct way to go.
Converting the input to a *graph* has been quite simple, and the algorithm to
travers it to compute the solution unravel itself along the way. The only issue
that took longer time to investigate was the fact that I should have used the
outer range when I was not applying a rule in a step. And also an off-by-one
error when computing the possible combinationation within a range. Incredible!

# Day 18
What a blast! Today I learnt A LOT of new algorithms and techniques to add to my
toolbox! The first part has been not too hard, it has been fun to spend some
time to get to visualize the output. For the second one instead I had to spend
some time on Google before finding the best way to solve it. Once understood the
formula, the algorithm has been quite straightforward to implement, the only
hurdle has been some float/integer rounding errors. Once I sorted that out, the
second golden star was behind the corner. Things I have learnt today:
- *Ray casting* algorithm
- *Shoelace* formula
- *Pick*'s theorem

Blazing!

# Day 17
All this problems on a grid are testing my nerves \o.O/ They have always been my
Achille's tendon, but this AoC I am finally improving my problem solving skills
in this area. I spent some time trying to figure out a solution using a BFS,
without any luck and also without really understanding why it was not entirely
working. After a while I realized that there are few cases where this search
will fail because of a wrong assumption regarding the tiles already visited.
After a while, I realized that, since the assigment was to find the **minimum**
value of a path in a *weighted* grid, the optimal and more straightforward
solution would have been the one and only ***Djikstra***. Of course a bit of
additional care in implementing the constraints (i.e. maximum 3 steps in a
single direction in part 1, and at least 4 but a maximum of 10 in part 2) was
necessary. Notice that the real difficulty in these kind of problem is debugging
the paths that are generated, since it might be cumbersome to analyze every and
each state pushed in the priority queue at each step, they are TOO many! Some
smart tricks are the key to find where the code is wrong, but that comes also
with experience. Wonderful!

# Day 16
Partecipatin in AoC while you have a flight to take, it is not a good idea,
especially if you haven't sleep enoguh (and by enough I mean "at all") :D
That said today's problem was insanely FUN to solve! I had clear from the start
what my algorithm should have looked like, and it has been very straightforward
to implement it. Also it works flawlessly at the first attempt on both inputs
for both parts, and this is something that does not happen very often, but is
VERY rewarding. The experience I have accumulated playing AoC so far helped me
in the thinking process, I saw some progress at least. Terrific! :)

# Day 15
Today's problem was both a relief and also a very fun to solve! Beside some
indexes gymnastics no major hurdles to get to the solution. Super!

# Day 14
Today has been a blast! The problem was really fun to solve and finally I was
able to put into practice some tricks learnt last year in AoC. The first part
has been really straightforward and it was also announcing where it would have
landed with the second one. I found it so funny that I even spent time to create
a small animation to visualize the rolling process between every cycle. Needless
to say I had a lot of fun today, given that I was able to figure out the
solution all by myself. Neat!

# Day 13
This day marks my biggest defeat so far. The algorithm I tried to implement is
giving me an incorrect result on 2 test cases, therefore I was not able even to
achieve the silver star without peeking at the real solution. This time I would
like to blame my comprehension of English language, I am still not 100% sure I
undestood what the real assignment was. Of course, admitting the incapacity it
is the first step to improve, therefore I accept the fate and I look forward
next challenges. Sad :(
OBS: 13 is an unlucky number, I will blame the misfortune brought by it ;)

# Day 12
The level of difficulty of AoC has increased. For today's problem an optimal
solution was necessary in order to get to the result in a time that was in
range of a human life. For the first part I was proud of myself since I used
some notions of *combinatorics* and a smart bit-wise algorithm to generate all
possible permutations of two elements over a certain number of variables.
Unfortunately, the solution I come up with was clearly not good enough to
achieve the gold star. I tried to come up with a smarter line-solving algorithm
that would have been applicable to the classic nonogram enigmistics game, but
I was not able to provide a proper solution for it. I, once again, decided to
rely on the explanation of
![hyper-neutrino](https://www.youtube.com/@hyper-neutrino), and this would be
the last time I will mention this AMAZING player since I have decided that after
spending a reasonable amount of time tackling the problem without succeeding I
will get to the solution learning from his YT tutorials. That said, when I
peeked at his solution for the first part of the problem, which was incredibly
smart and coincise, I was able to understand that the key for the second one was
a long-time friend of AoC: **memoization**. It was less then trivial to apply
a caching mechanism to the previous code, and in very short time I was able to
get to the end of the day. Voila'!

# Day 11
Very cool riddle! I was able to tackle the first part with a very bloated and
too-long algorithm that was performing a BFS for each pair of points. Obviously
it was taking forever also for the small input to compute the solution so, as
soon as I read the second part of the assignment, I realized that there should
have been a smarter and more optimal way to compute the distance between two
galaxies (i.e. points on a 2D matrix). After googling for a bit, I got to know
the so called "**Manhattan Distance**", and immediately some reminiscence from
Graph Theory studied at university became vivid in my mind. The difficult bit
of this problem, or at least the part that took me the majority of the time to
implement it, was to scale properly the coordinates of the input. In contrast,
the formula is VERY simple and elegant and it computes immediately. Noice!

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
