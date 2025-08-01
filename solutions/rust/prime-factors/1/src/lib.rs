pub fn factors(n: u64) -> Vec<u64> {
    let mut f = 2;
    let mut p = n;
    let mut result = Vec::new();
    while f * f <= n {
        if p % f == 0 && is_prime(f) {
            result.push(f);
            p /= f;
        } else {
            f += 1
        }
    }
    if p != 1 {
        result.push(p);
    }
    result
}

pub fn is_prime(n: u64) -> bool {
    let mut divisor = 2;
    while divisor * divisor < n {
        if n % 2 == 0 {
            return false;
        }
        divisor += 1;
    }
    true
}
