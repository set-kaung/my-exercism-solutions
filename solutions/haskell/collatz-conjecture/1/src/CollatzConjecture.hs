module CollatzConjecture (collatz) where

collatz :: Integer -> Maybe Integer
collatz n = if n <= 0 then Nothing else helper n 0

helper :: Integer -> Integer -> Maybe Integer
helper 1 acc = Just acc
helper n acc = if even n then helper (n `div` 2) (acc + 1) else helper (3 * n + 1) (acc + 1)