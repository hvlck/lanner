//! Mathematical constants.

/// Euler's number
pub const E: f64 = std::f64::consts::E;
/// Pi
pub const PI: f64 = std::f64::consts::PI;
/// Tau
pub const TAU: f64 = std::f64::consts::TAU;

/// Imaginary unit
pub fn i() -> f64 {
    let num: f64 = -1.0;
    num.sqrt()
}
