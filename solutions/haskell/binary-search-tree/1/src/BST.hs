module BST
    ( BST
    , bstLeft
    , bstRight
    , bstValue
    , empty
    , fromList
    , insert
    , singleton
    , toList
    ) where

data BST a = Empty | Node a (BST a) (BST a)
    deriving (Eq, Show)

bstLeft :: BST a -> Maybe (BST a)
bstLeft Empty = Just Empty
bstLeft (Node x Empty _) = Just Empty
bstLeft (Node x left _) = Just left

bstRight :: BST a -> Maybe (BST a)
bstRight Empty = Just Empty
bstRight (Node x _ Empty) = Just Empty
bstRight (Node x _ right) = Just right

bstValue :: BST a -> Maybe a
bstValue Empty = Nothing
bstValue (Node x _ _) = Just x

empty :: BST a
empty = Empty

fromList :: Ord a => [a] -> BST a
fromList xs = fromListHelper xs Empty

fromListHelper :: Ord a => [a] -> BST a -> BST a
fromListHelper [] tree = tree
fromListHelper (x:xs) tree = fromListHelper xs (insert x tree)




insert :: Ord a => a -> BST a -> BST a
insert x Empty = Node x Empty Empty
insert x (Node y left right)
    | x <= y    = Node y (insert x left) right
    | otherwise = Node y left (insert x right)

singleton :: a -> BST a
singleton x = Node x Empty Empty

toList :: BST a -> [a]
toList tree = toListHelper tree []


toListHelper :: BST a -> [a] -> [a]
toListHelper Empty xs = xs
toListHelper (Node x Empty Empty) xs = x:xs 
toListHelper (Node x left right) xs = (toListHelper left xs) ++ (x : toListHelper right xs)

