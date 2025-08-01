module SumOfMultiples (sumOfMultiples) where

sumOfMultiples :: [Integer] -> Integer -> Integer
sumOfMultiples factors limit = sum [y | y <- [1 .. limit - 1], any (\f -> f /= 0 && mod y f == 0) factors]
