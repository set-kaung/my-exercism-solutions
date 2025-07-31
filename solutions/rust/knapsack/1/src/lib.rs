#[derive(Debug)]
pub struct Item {
    pub weight: u32,
    pub value: u32,
}

pub fn maximum_value(max_weight: u32, items: &[Item]) -> u32 {
   solve(0,0,max_weight,items)
}

pub fn solve(i:usize,total:u32,max_weight:u32,items:&[Item]) -> u32{
    if i == items.len(){
        total
    }else{
        let mut take = 0;
        if items[i].weight <= max_weight{
             take = solve(i+1,total+items[i].value,max_weight-items[i].weight,items);
        }
        let skip = solve(i+1,total,max_weight,items);
        take.max(skip)
    }
}
