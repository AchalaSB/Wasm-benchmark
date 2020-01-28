use std::ffi::{CStr, CString};
use std::os::raw::{c_char};


#[no_mangle]
pub extern fn greet(subject: *mut c_char) -> *mut c_char {
    let subject = unsafe { CStr::from_ptr(subject).to_bytes().to_vec() };
    let mut output = b"Hello ".to_vec();
    output.extend(&subject);
    output.extend(&[b'!']);

    unsafe { CString::from_vec_unchecked(output) }.into_raw()
}

fn main(){}