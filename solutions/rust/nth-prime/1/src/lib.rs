fn is_prime(v: &Vec<u32>, n: u32) -> bool {
    for i in v {
        if n % i == 0 {
            return false;
        }
    }
    true
}

pub fn nth(n: u32) -> u32 {
    let mut v: Vec<u32> = vec![2, 3, 5, 7, 11, 13, 17];
    let idx = n as usize;
    while idx >= v.len() {
        let mut l = *v.last().unwrap();
        l += 1;
        while !is_prime(&v, l) {
            l += 1;
        }
        v.push(l);
    }
    v[idx]
}
