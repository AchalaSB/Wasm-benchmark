#[no_mangle]
pub extern fn eval(x: i64, y: i64, op:i64) -> i64 {
  match op {
    1 => x + y,
    2 => x - y,
    3 => x * y,
    4 => x / y,
    5 => x % y,
    _ => 0
  }
}

// #[no_mangle]
// pub extern fn eval_main() -> i64 {
//     eval(4, 5, 1)
// }

fn main() {

}
