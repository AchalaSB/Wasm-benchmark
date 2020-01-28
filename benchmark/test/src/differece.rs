#[no_mangle]
pub extern fn test(x: i32, y: i32) -> i32 {
    x - y 
}

#[no_mangle]
pub extern fn sub_main() -> i32 {
   test(5, 1)
}

fn main() {}