use std::f64::consts::*;

/// Euler's number
pub const E: f64 = std::f64::consts::E;
/// Pi
pub const Pi: f64 = PI;
/// Tau
pub const Tau: f64 = TAU;

/// Imaginary unit
pub fn i() -> f64 {
    let num: f64 = -1.0;
    num.sqrt()
}
