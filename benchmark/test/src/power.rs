extern crate num;
use num::pow;

#[no_mangle]
pub extern fn power(x:i64, y:isize) -> i64 {
   pow(x, y)
}

fn main(){}