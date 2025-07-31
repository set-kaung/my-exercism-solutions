pub fn factors(n: u64) -> Vec<u64> {
    let mut f = 2;
    let mut p = n;
    let mut result = Vec::new();
    while p > 1{
        while p % f == 0 {
            result.push(f);
            p /= f;
        }
        f += 1;
    }

    result
}

