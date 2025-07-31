pub fn secret_add(secret: Int) -> fn(Int) -> Int {
  fn(y) { secret + y }
}

pub fn secret_subtract(secret: Int) -> fn(Int) -> Int {
  fn(y) { y - secret }
}

pub fn secret_multiply(secret: Int) -> fn(Int) -> Int {
  fn(y) { secret * y }
}

pub fn secret_divide(secret: Int) -> fn(Int) -> Int {
  fn(y) { y / secret }
}

pub fn secret_combine(
  secret_function1: fn(Int) -> Int,
  secret_function2: fn(Int) -> Int,
) -> fn(Int) -> Int {
  fn(y) { secret_function1(y) |> secret_function2 }
}
