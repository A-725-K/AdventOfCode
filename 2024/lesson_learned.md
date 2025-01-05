# Lesson Learned
This file is a reminder for myself to understand what I did not know
when solving the AoC challenges, and how can I improve myself next
time I will face a similar problem. As well as some noticeable facts that
happened during the AoC.

## Day 25
Super simple problem, just to let all the players enjoy the Christmas day with
their loved ones. Once again has been very fun to get to the end! I will reserve
some time as soon as I can to complete all the remaining problems early next
year. A good resolution for 2025 would be to improve a bit in competitive
programming, and try to be a better player next year. Until then, merry Christmas
to all the players! The end!

## Day 24
Quite a weird one! The first part has been extremely easy, like too much easy.
It took me probably less than 5 minutes to code the solution, but oh boy, part
two has been a nightmare! I couldn't even imagine an algorithm to approach the
problem, let alone solve it. I also got quite scared after noticing the length
of video of the solution proposed by one of the best player of AoC on YouTube
(i.e. [hyperneutrino](https://www.youtube.com/watch?v=SU6lp6wyd3I)). Don't want
to make excuses, of course I am not a great player, but I did not have enough
time this year to spend trying to visualize and scribble something reasonable...
I had to resort to the walkthrough once again. The idea of reverse-engineering
the binary addition is very brilliant! Also the only idea I had was to try every
possible 4-way-combinations of wires swap to try to get a solution, but this
would have took way too long to compute the solution. Eventually there's at least
some brute-force but for every pair, after a failure is found. Way smarter then
what I thought. Almost!

## Day 23
Another graph problem! I am not too good in working with graphs, and that has
been clearly highlighted by the solution I came up with for part 1. I also
found a way to generalize to solve part 2 but it would have taken another 4.5
billions of years to compute the solution :( I had to look up how to optimize
the algorithm to find the `connected component` faster. In today's problem Rust
has not been really helpful since to implement simple operations required more
lines than expected, and of course that is also due to the fact that I am no
expert in such language. At least I learned something new, and that's the whole
purpose of playing AoC. So I am overall satisfied :) Not bad!

## Day 22
This solution is not really satisfactory, but at least got me the second star :)
```bash
Best seq: [(0, 0), (2, 2), (2, 0), (2, 0)]
The most bananas I can get is ***

real	3m55.868s
user	3m55.548s
sys	0m0.034
```
It has been quite fun to solve today's problem! I would like to spend some more
time to figure out a better and faster algorithm to get to the solution, mine
seems to be an order of magnitude slower than needed. Unfortunately time is
running out, and it is every day harder to find few hours to spare for AoC,
again life is happening in the background! I also have to solve day 21, which I
haven't tackled yet. Anyway I am pretty happy with how it went! Great!

## Day 21
What a problem! This day has been extremely challenging and rewarding at the
same time. Even though I couldn't figure out a proper solution I enjoyed trying.
I spent quite some time before understanding that the Cartesian product was the
operation I was looking for. Moreover, I managed to visualize the difference
between DFS and BFS and how they solve different kind of problems sometimes. For
part 2 I resorted once again to one of the top-player YouTube videos to get some
insights. Turned out this time I needed more than just a peek, but it has been
very worth it. The monster has been overcome also for this year. Eventually,
the first part computes in ~30s while the second only takes milliseconds, and it
is quite amazing to see the difference. Late!

## Day 20
When you think you are comfortable solving maze problems the reality arrives and
hits you very hard! Again, today is one of these days where I don't have enough
time to spend on AoC, so I had to look up a solution and once again learn from
better players. I had no idea how to approach the problem, and the attempt I
drafted was not even close to the real solution. Interesting the "diamond"
approach to get to the second star, now I know how to explore at a certain
distance an area on a grid. Come on!

## Day 19
Today has been a bit of up and down. To be fair I hadn't got much time to spend
on AoC as much as I wanted :( Life is happening really fast, and it's hard to
reserve time for everything. But mumbling aside, I was quite happy to figure out
the first part on my own. I managed to find an optimal algorithm to compute the
solution. Unfortunately, when it came the second part, the code I produced was
not simple enough to refactor to count also the number of possible solutions.
I need to improve in these kind of problem, it is a bit sad that I am not able
to reason on the algorithm I just wrote to adapt it in such a way. On the other
hand, it is always super-interesting to see how top-notch players can so easily
find their way. I have still lots to learn and so little time to do it :(
So and so!

## Day 18
After a hectic couple of days here some more easy-going puzzle :) After realizing
that `memoization` was the best today's best friend, I had no issue in figuring
out all by myself! It has been relieving after giving up for at least 2 days.
It also took very short time to adapt the solution to part 1 for getting the
second star, and the code is clean too. Relaxing!

## Day 17
Cpus has always been my favorite kind of problem to implement for AoC. The first
star has been very fun to reach. Unfortunately, I had no idea how to approach
the second part whatsoever. After few attempts, and inspecting (or more like
dissecting with a careful reverse engineering approach) the input I got, I found
some bits of information but that were not enough to make me think of a solution.
Thanks to a top-player mindset reveal I eventually understood what was the way
to go, but, hey, I would have never-ever-ever get to that by myself. Here's a
[link](https://www.youtube.com/watch?v=OjFGKL54yJQ) that inspired my solution
after hearing and understanding the reasoning. It has been overall very fun, but
too difficult for my current possibilities. The idea behind this puzzle is
incredibly smart, and there could be a general solution if and only if all the
players' inputs share a common pattern (which I believe they do). Kudos to the
people who thought of it, they have my complete admiration! Surprised!

## Day 16
Not every day is bright and shining... Today I have been not able to figure out
the solution by myself. Or actually I found an algorithm that solved the two
small inputs in few seconds that was a BFS visit with a sparkle of backtracking.
Unfortunately it miserably failed on the real input. After several hours of
runtime it was not finished yet. It is probably correct, but it would take
forever to get an answer, and then it is when I realized my solution was not the
intended one. After reading a bit on the internet I understood that today's
problem required `Dijkstra's` algorithm to win the stars. I also learned a new
interesting way to express the direction while walking a maze as a pair of values
among -1|0|1. Here is how:
 - (1,  0): move to the right (_East_)
 - (-1, 0): move to the left (_West_)
 - (0,  1): move down (_South_)
 - (0, -1): move up (_North_)
where the positive values is toward the positive direction in a Cartesian plane
oriented like this:
-|--------> X
 |
 |
 V
 Y

Then after looking up a proper solution I managed to overcome the first part of
the problem. When I read the second half, I literally gave up. I knew it had to
do with storing the full path somehow while walking through the maze, but how to
save only the best one that lead to a minimum score was completely out of my
ability. I am happy I learned something new today, even if I did not managed to
solve anything. I wish I could do better with these kind of problems. I also
learned some new tricks with Rust, therefore the day has not been entirely
wasted. But hey, life happens. Positive!

## Day 15
I am pretty amazed by how today I managed to solve both parts. For the first one
I had no problem at all, the algorithm I imagined worked perfectly once I managed
to write it down as intended. For the second it took me a while to understand
why my idea didn't work, but I have been enlightened somehow and realized that 
`backtrack` technique was the way to proceed. Probably not the most efficient
solution, but it served the purpose. Bonus points, both solutions are recursive!
Incredible!

## Day 14
What a day! I love these kind of puzzles where the output is also visual :)
Luckily for me, today has been quite easy to get both stars, and I couldn't have
asked for a better day for a simple challenge. It has been a very long day for
personal reasons, but eventually I managed to play for some time and I am very
delighted and pleased with the outcome. Grateful!

## Day 13
It is always fun to deal with math problems. It has also been fun to solve it
together with a loved one. The first part has been really straightforward, and
it took me really short time to get the star. When I started the second half, I
spent too much time trying to find a Rust library to solve linear algebra systems
without any luck. Turned out it was not really needed, just a piece of paper and
some scribbling was the perfect recipe to find the solution. I spent too much
time debugging some stupid errors due to operand orders, missing factors, and
variable poorly named but eventually I was able to reach also the second star.
Careful!

## Day 12
The game starts to become complicated as expected. I was surprised to figure out
at least the first part in a very short time. I also had a hunch about the
second one, and I was already using the half-coordinates thanks to some memories
from past years and to a lucky Google search. Unfortunately my algorithm
was not able to solve the corner case presented in one of the examples where
2 regions were touching an angle but were not part of the same segment. But it
has been fun to spend the day on AoC after a rough couple of weeks, so I am not
disappointed after all! I am starting to appreciate Rust and I will try to stick
to it for the rest of the competition. Cool!

## Day 11
Very interesting riddle! It was one of these days that while reading part one,
you might already guess what part two will be asking you. After solving quickly
the first half of the day, I figured out that the only way to tackle the next one
would have been caching some data. I only had too little time to spend to figure
out on my own how. Unfortunately today is one of those days. I had to look up
how to implement correctly the memoization and only then I understood that I
could have written it myself. Only the divide-et-impera approach didn't occur to
me. On the other hand I now know a new Rust crate, namely `cached`, that can be
used almost in the same way as `@cache` and `@lru_cache` decorators in Python,
which I was already familiar with. Ok!

## Day 10
After yesterday debacle my morale was very low. Today has been a breath of fresh
air! I revamped my maze-walking skills from previous AoC, and I managed to
complete both parts pretty quickly. Fun fact: at some point I got to the solution
of part 2 before solving part 1 :) Switching between the 2 has been quite smooth.
Glad!

## Day 9
The first terrible defeat, as expected :( It was all going too well, until today.
The problem was fun, and in this moment in my life I feel really anxious for a
number of things. This doesn't want to be an excuse, but just underline how I
was not in the right mindset to face this challenge. First of all, I didn't
understand properly part 1. I just misread the instructions and developed an
overcomplex algorithm that did not even solved the problem. After quite a lot
of struggling and juggling with edge cases, eventually, I managed to get to the
point. For part 2 I had a lot of ideas, but had hard time implementing them
because of the aforementioned mood. Thanks to one of the best AoC players and
YouTubers, I managed to understand where I was failing. Due to time constraints
I had not enough time to develop a solution of my own, so I ended up almost
copying it. Today, also Rust came in the way few times. I understand now the
power of this language and its compiler. It really prevents you from writing bad
code, but not always is easy to deal with it. Sad!

## Day 8
The author of AoC wanted to be merciful on the first weekend of the competition,
therefore a nice and easy challenge kept us busing this Sunday. Some basic math
from high school was enough to solve today's puzzle. Thank you!

## Day 7
I feel it! The problem are becoming more complex. Today struggle, on the other
hand, it is not about the algorithm. It is about me being stupid using wrong
data structure to store the input. I spent an excessive amount of time trying to
understand why my solution was working on the test input and miserably failing
by a small amount on the real one. Turned out I had duplicates in my input, and
using HashMaps in this case was not the smartest choices since some values were
overridden. After understanding this bit, the solution has been pretty
straightforward. In the beginning I copied-and-pasted some code from
[Rosetta repository](https://rosettacode.org/wiki/Permutations_with_repetitions#Rust)
to generate all possible combinations of N elements in K slots with repetitions.
For the first time in my life I made use of some probability theory I learned
at university :) I was quite happy even though it is probably (pun intended) not
the fastest solution. Anyway it has been very easy to adapt the first solution
to part 2. Today I also had my first argument with Rust, it did not helped me
in writing what I actually had in mind. I also struggled a bit with the
`itertools` package to find a better implementation for the Rosetta part.
Semi-happy!

## Day 6
Today I put into practice something I learned on day 2. Start with the dirty
solution and then refine your way up! I got the second star in \~1 minute, my
script was quite slow even though I was very much convinced of my solution.
I also learned that the Rust compiler makes quite a lot of difference when
running in the `--release` mode. From 30+ seconds it got to \~16. After reviewing
other solutions, I also found a nice optimization that saved a lot of time, and
eventually I managed to run my code in \~3sec. Satisfied!

## Day 5
I am starting to enjoy Rust even if sometimes I still fight with its compiler :)
It has a quite extensive standard library with lots of basic operations. I liked
the fact it has a `.contains()` method on `Vec` type, it simplified my life.
With a bit of struggle and mental gymnastic to figure out the first part, the
second one was surprisingly smooth. Let's go!

## Day 4
Today I propose a nice and clean solution! I was a little bit worried in the
beginning because I know Rust is known to complicate iterating on a string
due to its extensive UTF-8 support, but eventually turned out to be not too
difficult to manipulate a `Vec<Vec<char>>`. I enjoyed refining my solution and
exploring some Rust features (e.g. `std::slice::Iter` crate) to make it nicer!
Moreover, crosswords are one of my favorites hobby, so today I also appreciated
the lore :) Git-push!

## Day 3
Very easy and fun problem to solve! Rust FTW once again, even though the fact
that regex are not part of the core language was a little weird :) Let's go!

## Day 2
First star easy-peasy... and then the challenge hit me! It took way longer than
expected to figure out the second one. Trying to remove each and every item it
looked like a not-so-smart move, but turned out to be an accepted solution. I
am not sure at this point if there's a better way to achieve the result, but at
least I got my second gold star! Bonus: I still managed to use Rust :) Go on!

## Day 1
Here we go again! This year will be different for me. A lot in my real life is
happening, and I hope I will find the time to participate at least a bit in AoC.
One thing I would like to try is Rust language even though I am not sure how
far I can get with this. Let's see!

# Results:
|Day|Stars||Day|Stars|
|---|---|---|---|---|
|1|$\textcolor{gold}{\textsf{**}}$||14|$\textcolor{gold}{\textsf{**}}$|
|2|$\textcolor{gold}{\textsf{**}}$||15|$\textcolor{gold}{\textsf{**}}$|
|3|$\textcolor{gold}{\textsf{**}}$||16|$\textcolor{gold}{\textsf{**}}$|
|4|$\textcolor{gold}{\textsf{**}}$||17|$\textcolor{gold}{\textsf{**}}$|
|5|$\textcolor{gold}{\textsf{**}}$||18|$\textcolor{gold}{\textsf{**}}$|
|6|$\textcolor{gold}{\textsf{**}}$||19|$\textcolor{gold}{\textsf{**}}$|
|7|$\textcolor{gold}{\textsf{**}}$||20|$\textcolor{gold}{\textsf{**}}$|
|8|$\textcolor{gold}{\textsf{**}}$||21||
|9|$\textcolor{gold}{\textsf{**}}$||22|$\textcolor{gold}{\textsf{**}}$|
|10|$\textcolor{gold}{\textsf{**}}$||23|$\textcolor{gold}{\textsf{**}}$|
|11|$\textcolor{gold}{\textsf{**}}$||24|$\textcolor{silver}{\textsf{**}}$|
|12|$\textcolor{gold}{\textsf{**}}$||25|$\textcolor{silver}{\textsf{**}}$|
|13|$\textcolor{gold}{\textsf{**}}$|

Eventually solved also 21 and 24 part 2 after the competition ended.
