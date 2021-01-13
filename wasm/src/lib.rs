// std

// crates
use lanner::{evaluate as lanner_evaluate, LannerError};
use wasm_bindgen::prelude::*;

// local

#[cfg(feature = "wee_alloc")]
#[global_allocator]
static ALLOC: wee_alloc::WeeAlloc = wee_alloc::WeeAlloc::INIT;

#[wasm_bindgen]
pub fn evaluate(inp: &str) -> String {
    match lanner_evaluate(inp) {
        Ok(result) => result.to_string(),
        Err(error) => error.to_string(),
    }
}

#[cfg(test)]
mod tests {
    use super::*;
}
