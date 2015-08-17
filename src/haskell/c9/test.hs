double x = x + x

quadruple = double . double
-- quadruple x = double (double x)
-- quadruple x = x + x + x + x
-- quadruple x = x * 4

n = a `div` length xs
    where
      a  = 10
      xs = [1..5]

newlast xs = drop ((length xs) - 1) xs !! 0
anotherlast xs = xs !! ((length xs) - 1)


qs [] = []
qs (x:xs) = qs [a | a <- xs, a < x]
            ++ [x] ++
            qs [a | a <- xs, a >= x]


factorial n = product [1..n]
-- factorial n = product (take n [1..])
