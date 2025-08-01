pub fn square(s: u32) -> u64 {
    if s < 1 || (s > 64) {
        panic!("Square must be between 1 and 64")
    }
    let mut res: u64 = 1;
    for _ in 1..s {
        res *= 2;
    }
    return res;
}

pub fn total() -> u64 {
    u64::MAX
}
