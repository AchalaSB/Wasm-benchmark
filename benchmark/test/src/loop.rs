#[no_mangle]
pub extern fn loop(x:u64, n: u64) -> u64 {
    let mut count = 0;
    for j in 0 .. n {
        count = x + j;
    }
    count
}
fn main(){}