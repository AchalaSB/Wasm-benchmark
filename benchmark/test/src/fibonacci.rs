#[no_mangle]
pub extern fn fib(n: i32) -> i32 {
	if n == 1 || n == 2 {
		1
	} else {
		fib(n - 1) + fib(n - 2)
	}
}

#[no_mangle]
pub extern fn fib_main() -> i32 {
	fib(4)
}

fn main() {}
