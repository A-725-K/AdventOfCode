# Lesson Learned
This file is a reminder for myself to understand what I did not know
when solving the AoC challenges, and how can I improve myself next
time I will face a similar problem. As well as some noticeable facts that
happened during the AoC.

# Day 9
What a day! I reached the first star basically immediately, and then a big wall was presented
in front of me. Boy, that was hard! I have tried (and you can see by the amount of commented
code I will commit anyway...) few different approaches, mostly based on `ray-casting`
algorithm, but none of them where enabling me to reach an answer in a limited amount of time.
The input space grew enormously, such that an exhaustive search wouldn't led me anywhere. I
had to resort to some walkthrough this time, and I realized I couldn't resolved it with my
current knowledge. I got to know 2-3 very interesting algorithm (one of which I already
explored at some point a couple of years ago, in AoC):
- **coordinate compression**: a very interesting approach to reduce of some magnitude the size
                  of the input space. By saving the deltas between two Xs or Ys
                  the number of coordinates to consider is much lower
- **flood-fill**: the idea to find a set of points that fall inside a perimeter is to start
                  a flood from the outside and see when they reach a perimeter, and then
                  exclude them.
- **prefix sum**: an efficient algorithm to get the sum of some values within a range, here
                  in its 2D version, where you keep a matrix of pre-compuited values instead
                  of a single list.
The mix of these 3 algorithms lead to the solution in a very short time. A bit hard to imagine
and visualize, but printing out intermediate results helped quite a lot in understanding the
whole process. Interesting!

# Day 8
After a defeat, this was absolutely needed! It was a boost for my journey. I was able to
identify the correct algorithm, and also implement it in a reasonable time (for my standards,
of course). *Union-find*s are not the most common data structures used in everyday's job.
Luckily there were stuck in the back of my head. First I've tried a connected-component
approach, but the limits were clear after examining the output even on the samll input.
Good that existed an already available implementation in a Rust crate called `union_find`,
that simplified the solution. The code was also very elegant. I have also learned that `f64`
data type in Rust is not totally ordered since it can hold the value NaN. Astonishing!

# Day 7
Today marked my first defeat :( The first part has been fairly easy to solve, and I
manage to come up with a sort of iterative algorithm that worked (and graphically
showed the result, too). Second part was easy to solve on the small input with a
*backtracking* algorithm, but when it get to the real input it failed miserably. Fun
fact: today I was at a Christmas party at some friends, and I decided to let run the
solution until I was back home... well, after ~11h the program was still running, while
no solution in sight. After poking a bit at the sub-reddit, I've finally found what I
was not understanding. I've completely missed 2 aspects: caching and simplicity of the
algorithm design. Turned out I could've crafted a very trivial solution in ~10 lines, but
I had to change completely the way I was thinking. It is also something I should've learned
from past Advents of Code, but still it got me. It happens!

# Day 6
Very fun challenge! It's a bit sad to see that we are already halfway through this
year, time is flying too fast! Anyway, a bit of parsing gymnastics was all that was
needed to solve today's riddle. I was a bit surprised when after solving the
second part for the small input, my code panicked on the real one, but then I have
learned that the `char::to_digit()` function in Rust returns an `i32` and not a
`usize`. Fixing the cast, seamlessly led me to the correct result since the
algorithm was sound. Super!

# Day 5
Also today's problem was a walk in the park. It was clear from the beginning that
there would have been some range-merging algorithm involved, but for the first
part I decided to go with the simplest and inefficient solution that took ~1ms
to run anyway. Nothing compared to day 9 of 2024 :) Simple!

# Day 4
Today was a relaxing problem. Simple matrix structure and easy requirements, the
easiest so far for me. Not much to add. Happy :)

# Day 3
After 4 years of AoC I start to see some fruits of my work :) The first problem
has been fairly trivial for me. For the second part, instead I was surprised how
fast I could come up with an optimal solution. I tried to implement 3 different
strategies, but the first 2 were not correct. Nonetheless, they helped me to
understand what I was doing wrong. All in all was a successful day. Enjoyable!

# Day 2
Given the limited amount of days, this year the difficulty slope will be steep.
While the first part was quite straightforward, adapting the algorithm for the
second part took me a bit of time. Eventually I kept it simple, after yesterday,
I've learned that sometimes simplicity pays off. While it took ~1s to solve the
problem on the real input, it is still a fairly fast algorithm, given that I was
running the debug version (`cargo run` instead of `cargo run --release`). There
was an edge case I had to treat it separately which I believe could be integrated
in the actual solving algorithm, but I am happy with how the code looked like
eventually. Satisfied!

# Day 1
Another year, another Aoc! Sad to see that this year it will be shorter,
but the author has my full support given his reasons! And in this AI world,
it is still good to experience such good quality in this game, I am very thankful!
From the beginning, part 2 hit me hard: I always forget to keep it as simple
as possible, and reason about the constraints, it would have saved me some time.
Nevertheless the start is smooth, looking forward to tomorrow's problem already!
Let's begin!

# Results:
|Day|Stars|
|---|---|
|1|$\textcolor{gold}{\textsf{**}}$|
|2|$\textcolor{gold}{\textsf{**}}$|
|3|$\textcolor{gold}{\textsf{**}}$|
|4|$\textcolor{gold}{\textsf{**}}$|
|5|$\textcolor{gold}{\textsf{**}}$|
|6|$\textcolor{gold}{\textsf{**}}$|
|7|$\textcolor{gold}{\textsf{**}}$|
|8|$\textcolor{gold}{\textsf{**}}$|
|9|$\textcolor{silver}{\textsf{*}}$|
|10||
|11||
|12||
