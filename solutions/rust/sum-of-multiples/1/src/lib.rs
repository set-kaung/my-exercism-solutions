pub fn sum_of_multiples(limit: u32, factors: &[u32]) -> u32 {
    (1..limit)
        .filter(|k| factors.into_iter().any(|x| *x != 0 && k % x == 0))
        .sum()
}
