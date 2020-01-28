#[no_mangle]
pub extern fn sum(x: i32, y: i32) -> i32 {
    x + y 
}
#[no_mangle]
pub extern fn diff(x: i32, y: i32) -> i32 {
    x - y 
}
#[no_mangle]
pub extern fn mul(x: i32, y: i32) -> i32 {
    x * y 
}
#[no_mangle]
pub extern fn div(x: i32, y: i32) -> i32 {
    x / y 
}

fn main(){
}
