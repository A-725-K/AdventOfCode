# Lesson Learned
This file is a reminder for myself to understand what I did not know
when solving the AoC challenges, and how can I improve myself next
time I will face a similar problem. As well as some noticeable facts that
happened during the AoC.

## Day 24
After all those days of solving riddles, I was able to figure out by myself
one of the gimmick that helped finding a better solution to this problem.
I realized that since the maze is finite, there should have been only a
certain amount of blizzerd "configurations". If I pre-compute them, before
trying to walk the maze, I would have had an easier life to see where at
each step my character could have moved. Moreover I also realized that
the amount of possible "configurations" were equal to the Least Common
Multiplier (LCM) of the number of rows and the number of columns. After
that amount of steps, the patter repeats itself over and over again.

## Day 22
Today was extremely fun to solve, and not so difficult eventually! Sometimes
thinking out of the box helps to find creative solutions and today I did
some handwork in order to figure out the solution. I draw on a piece of
paper the exact same shape given as input, and only when I actually saw the
cube in real life I realized the simplicity and the elegance of the
algorithm. Cool!

## Day 21
For the first time, I appreciated how making small steps every day
eventually repayed! Today I am very proud for being able to find a solution
quickly to the problem that was also very efficient (~0.1s for both parts).
Day after day, my reasoning skills have improved. Besides designing a good
data structure to deal with the problem, after reading the second part I was
immediately aware that I could not have solved it with an exhaustive search
of the index since the quantities involved were too big. I realized only by
myself that a ***merge-sort*** like approach would have been the winning
technique to reduce the space of possible solutions rapidly. In fact, trying
to find a solution in the range [0, 2^64] took only about 1/10 of a second,
or ~100ms, which is absolutely amazing, IMHO! I do not want to spend too
many words on the problem, since probably it was medium-easy, but still I
felt really great by solving it in a reasonable amount of time (2/3h).
I would not consider myself not even a shadow of a good programmer, but
sometimes these kind of small victories help the morale and the self esteem.

## Day 20
The problem per se was not super-difficult. In fact I was able to come up
with a working algorithm for the first part in quite short time. The real
issue was dealing with the next phase: shifting the numbers of a large
offset pointed out one of the most subtle but, at the same time, dangerous
errors in the history of programming: **off-by-1**. In fact, the underlying
logic was sound, what I was doing wrong were the boundary checks. Of course,
those kind of errors are super-difficult to spot, especially if some algebra
modulo N is involved, and in this case it was indeed. Eventually, I am
very satisfied with a solution I produced since it is relying on *2 mutual
recursive functions*. It is always surpring when you see some crazy theory
you studied on the desk of the university coming in hand to solve an actual
task. Nice!

Mistakes I did:
- Wrong approach to debugging: sometimes it is way easiear to take paper and
  pen, and try to visualize the problem instead of just conjecturing in your
  head very abstract reasonings.

## Day 19
["O tempora o mores!"](https://en.wikipedia.org/wiki/O_tempora,_o_mores!),
also today I have to stop at the silver star! Since I have never had the
chance to study operational research, this optimization problem was way
harder than expected. Luckily I recognized almost immediately the class of
algorithm I had to deal with, so I did not spend too much time investigating
how to implement it, and TBH, the algorithm itself it is fairly simple to
understand and quite well written in my solution. It is surprising to see,
even in the simple case, how long it takes before being able to observe the
outcome of the program. This fact is a constant in the last few days of AoC.
The biggest issue is that we are asked to find a ***global maximum*** of a
cost function, therefore a *greedy* approach would not suffice since it
would only find the best value locally, at each step, disregarding other
possibly better states. In today's problem, sometimes it is more convenient
continue to produce even though we have enough resource to build some robot.
Needless to say that the difficulty increases even more when we are
approaching the golden star. In this case we have to be sure that not only
the algorithm is efficient **time**-wise, but also **space**-wise. It was
crazy seeing my laptop, a very recent and powerful one (32GB RAM, Intel i7
7th generation 16 cores) to almost fulfill the memory (~20GB before stopping
the program) keeping a reasonable amount of CPU running (~10%). Eventually I
did not found a way to store intelligently the data of the problem, knowing
that it would have take a while to get to a number. On the bright side, I
consider very elegant the data structure I have adopted to modelize the
problem, making a good use of pointers and structure offered by Go language.

### EDIT:
After few days it was time to finally solve this challenge. I seeked for
hints on the faboulous Reddit page of Advent of Code. There I found some
suggestions on how to improve the algorithm, without necessarily finding
code snippets. This was very helpful, since then I realized how to optimize
my solution, which was already kinda good. Eventually I was able to run my
code in ~40s using "only" ~5GB of RAM. Even though those are not small
numbers in IT terms, I would consider them good enough for the purpose of
solving this challenge. Eventually I learned how to reason in terms of DFS
and what `pruning` means in this context. The optimizations I applied are
the following:
- use `uint8` instead of `int` to store state since each number costs only
  8 bits instead of 32
- in a turn, if I decide to build a Geode robot, that's the best that I can
  do for that turn, so do not even try to explore other solutions
- save the DP keys in `base64` encoding and only the first 23 characters
  (this number comes after some tuning experiments, $\lt23$ leads to wrong
  results, while $\gt23$ takes more time and space to get to the solution).
- if I already own enough robot to cover the max cost for each resource,
  there is no point in exploring a path where I try to build an additional
  one

## Day 18
Today I learnt about a new type of algorithm called ***Flood-fill***. It
was clear that this was the way to go in order to achieve the second star,
but unfortunately I needed some input. Thanks to this amazing YT channel of
a top-player of AoC, namely
[William Y. Feng](https://www.youtube.com/channel/UCMhk437GeN8069t7y0lJQbw),
I was able to understand the description of the second part. This time I had
hard time understanding what the author wanted to know (English is not my
first language). After watching the explanation and taking a glance at his
solution, I was able to implement my own version. The real golden star was
the discovery of new techniques to add to my knowledge base!

## Day 17
Today mark my very first failure :( After solving smoothly the first star,
the second part completely destroyed me. Not a good day for my morale,
especially because I figure it out by myself a possible algorithm to get to
the solution, but I was not able to actually implement it. After looking at
the problem for several time, I realized that in the outcome there would
have been a recurring pattern, so the overall height could have been easily
computed as soon as the repetition is detected. Unfortunately my abilities
were not sufficient at this point to code this algorithm. But admitting the
defeat is also part of the growing and learning processes, so let's deal
with it, and move on to the other problems! I am sure that later I will be,
soon or later, good enough to solve it! ;)

### EDIT:
After Christmas I decided to try to solve the second part. Again, the
algorithm was pretty clear in my head but I couldn't figure out how to
"find the beginning of a pattern". After reading around some solutions I
found a very interesting approach in the one proposed by
[Carolina Sol Fernandez](https://github.com/carolinasolfernandez). I
finally understood how to detect a pattern: if a specific rock fall down
at the same jet of gas, that means that the structure will repeat itself
over and over. At this point the only values to keep track of are the height
of the tower so far (which I was already doing with the variable `maxY`),
and the number of rocks fallen so far. At this point with some quick math
was fairly easy to find the well after an arbitrary number of rounds. Thank
you so much for letting me learn how to modelize in a very clever way an
incredible hard concept. Eventually the algorithm did not change that much
from part 1, besides the addition of an if-else statement.

## Day 16
Today's problem was very tough! Graphs have always been not my strong
suit, and moreover I was not able to recognize the optimizaion problem
lying between the lines. After taking a look at some of the MVP coders
solution, I realized that Dynamic Programming was the correct technique
to adopt. Honestly I couldn't have figured it out by myself. Even though
from the challenge point of view this was kinda of a failure, it was an
incredible learning experience. As soon as I saw a variable called "DP" I
tried to implement my own version of memoization, and it eventually worked,
leading me to the correct number.

Mistakes I did:
- Trying to implement again a backtracking algorithm even though it would
have clearly been a wrong algorithm since the request was only the result,
not the path to reach it.

Things I did good:
- Dealing with graph traversals, I was surprised to see how my knowledge has
advanced in these techniques.

## Day 13
Man, what a journey! I learned a lot by solving this riddle. First of
all, *RECURSION 101*. Mistakes I did:
- the input of this problem was clearly suited for a dynamic language
  not strongly typed, persevering with Go wouldn't have brought me
  anywhere, luckily I decided to switch to Python fairly soon
- read, *Read*, **READ**, ***R-E-A-D-!*** the f*ing description of
  the problem! The explanation of the solving algorithm was clearly
  stated in the example section of the first part of the problem.
  With that in mind, debugging the work in progress code was
  ***A LOT*** easier.
- realizing that a 3-values solution was a better idea than just a
  boolean one took me a lot of time, especially because I did not
  fully understood how the decision should have been taken
Luckily, my implementation design helped incredibly in solving the
second part of today's problem. Once the *comparison* callback was in
place, finding out the solution was just a simple matter of Python
gymnastics.

## Day 12
Not always the first solution that comes to mind is the best for the
problem. In this case I was fairly sure about the *backtracking*
technique, not taking into account that can explode pretty fast, even
with reasonably small inputs. A small research, sometimes, lead to a
better approach. It was the case as soon as I realized that I could
have used BFS to implement my solution.

## Day 11
First part went smoothly, but to get the second star I struggle a lot
before realizing how easy and elegant the solution would have been.
I did 2 mistakes:
- trying to use Go library `math/big` to implement the solution
  under the assumption that the issue I had was due to an integer
  overflow
- rewrite my solution in Python3 thinking that the issue with the
  Go program was a bad memory management since `math/big` library
  operations are done with pointers, and Python3 handles arbitrary
  length integers by design

Then I peeked at the solution of some of the top-notch coders in the
competition to realize that, in order to save a lot of space and
speed up the computations, I had to use some basic math knowledge.
To carry smaller number and not upset my CPU and RAM, I had to save
numbers up to LCM (Least Common Multiple) of the dividend of all
monkeys. Turns out that this number can be stored in a simple `int`
variable.

## Day 9
This problem will remain a mystery for me. I thought I have
understood the algorithm, but, unfortunately, I did not get entirely
how the tail should have reached the head in case of multiple nodes.
The issue I had was really weird, since the algorithm I developed
worked on the small input provided by the challenge, as well as on
the "Try to match those exact steps" section in the description of
the second part of the day. It was a little bit frustrating to not
figure out by myself, but then I decided to look at the solution of
my colleague [Jesper Högström](https://github.com/jhogstrom) in order
to find where I was at fault. Eventually, I realized that the issue
with my algorithm was in the movement of the tail once the head was
more than 2 steps ahed in diagonal. It is still unknown to me where
is the pitfall in my algorithm, since it seems, at a first glance,
even more optimized than the expected solution. Nevertheless, it was
wrong. So, thank you for the implicit help!

## Day 7
The only issue that lead me to a huge loss of time while solving this
challenge was the same as day 4: **READ CAREFULLY THE DESCRIPTION!**.
In the directory-search algorithm instead of looking up starting
from the root of the file system, I had to read from the current
directory (and this fact was exactly described in the specs of the
`cd` command).

## Day 5
Only small note: next time be less lazy and parse the input instead
of doing the setup programmatically ;)

## Day 4
While solving this problem I had encountered the biggest issue of all
times:
> READ CAREFULLY THE DESCRIPTION OF THE CHALLENGE!!!

In fact I spent an unreasonable amount of time trying to implement a
difficult algorithm that checked how many intervals were not
overlapping at all in the list, instead of simply looking for those
in the same line that were not independent.

# How to improve
- READ CAREFULLY THE DESCRIPTION OF THE CHALLENGE!!!
- Test always the code with the small input and save it!
- Think about possible optimizations when the input grows: sometimes also
  the space is important, even when dealing with a time-optimal algorithm!
- Use the language that suits best the reasoning. Although solutions
  written in languages like `C`, `C++`, `Golang`, `Rust`, etc. which are
  statically typed are usually fast to run, they may require extra time to
  code. On the other hand languages like `Python` allow a quick prototype
  that can be refined later on.

# Results:
|Day|Stars||Day|Stars|
|---|---|---|---|---|
|1|$\textcolor{gold}{\textsf{**}}$||14|$\textcolor{gold}{\textsf{**}}$|
|2|$\textcolor{gold}{\textsf{**}}$||15|$\textcolor{gold}{\textsf{**}}$|
|3|$\textcolor{gold}{\textsf{**}}$||16|$\textcolor{gold}{\textsf{**}}$|
|4|$\textcolor{gold}{\textsf{**}}$||17|$\textcolor{silver}{\textsf{*}}$|
|5|$\textcolor{gold}{\textsf{**}}$||18|$\textcolor{gold}{\textsf{**}}$|
|6|$\textcolor{gold}{\textsf{**}}$||19|$\textcolor{silver}{\textsf{*}}$|
|7|$\textcolor{gold}{\textsf{**}}$||20|$\textcolor{gold}{\textsf{**}}$|
|8|$\textcolor{gold}{\textsf{**}}$||21|$\textcolor{gold}{\textsf{**}}$|
|9|$\textcolor{gold}{\textsf{**}}$||22|$\textcolor{gold}{\textsf{**}}$|
|10|$\textcolor{gold}{\textsf{**}}$||23|$\textcolor{gold}{\textsf{**}}$|
|11|$\textcolor{gold}{\textsf{**}}$||24|$\textcolor{gold}{\textsf{**}}$|
|12|$\textcolor{gold}{\textsf{**}}$||25|$\textcolor{gold}{\textsf{*}}$|
|13|$\textcolor{gold}{\textsf{**}}$|

Eventually solved also 17 and 19 part 2, after the competition ended.

# TODO(s):
- create script to grab input
- use a command line argument to run either small or big input
