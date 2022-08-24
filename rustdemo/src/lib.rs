extern crate libc;

use sha256::digest_bytes;
use std::ffi::{CStr, CString};

#[no_mangle] 
pub extern "C" fn rustdemo(name: *const libc::c_char) -> *const libc::c_char {
    let cstr_name = unsafe { CStr::from_ptr(name) };
    let bytes = cstr_name.to_bytes();
    CString::new(digest_bytes(bytes)).unwrap().into_raw()
}
