use time::{Duration, PrimitiveDateTime as DateTime};

const one_billion: i64 = 1000000000;

// Returns a DateTime one billion seconds after start.
pub fn after(start: DateTime) -> DateTime {
    start.saturating_add(Duration::new(one_billion, 0))
}
