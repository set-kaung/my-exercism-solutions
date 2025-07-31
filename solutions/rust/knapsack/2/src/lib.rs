#[derive(Debug)]
pub struct Item {
    pub weight: u32,
    pub value: u32,
}

pub fn maximum_value(max_weight: u32, items: &[Item]) -> u32 {
    let mut dp: Vec<Vec<u32>> = Vec::new();
    for i in 0..=max_weight {
        let mut v = Vec::new();
        for j in 0..=items.len() {
            v.push(0);
        }
        dp.push(v)
    }
    for w in 1..=max_weight {
        for i in 1..=items.len() {
            let take = if w >= items[i - 1].weight {
                items[i - 1].value + dp[(w - items[i - 1].weight) as usize][i - 1]
            } else {
                0
            };
            let skip = dp[w as usize][i - 1];
            dp[w as usize][i] = take.max(skip);
        }
    }
    return dp[max_weight as usize][items.len()];
}
