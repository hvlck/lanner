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

/// Phi, or the Golden Ratio
/// Phi is approximately equal to `1.6180339`.
pub fn phi() -> f64 {
    0.5 * ((1.0 + ((5.0 as f32).sqrt())) as f64)
}
