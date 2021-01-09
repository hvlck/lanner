// std

// crates

// local
use crate::{evaluate, LannerError};

trait Evaluate {
    fn evaluate(&self) -> Result<f64, LannerError>;
}

impl Evaluate for str {
    fn evaluate(&self) -> Result<f64, LannerError> {
        evaluate(self.as_ref())
    }
}

impl Evaluate for String {
    fn evaluate(&self) -> Result<f64, LannerError> {
        evaluate(self.as_str())
    }
}

#[cfg(test)]
mod prelude_tests {
    use super::*;

    #[test]
    fn check_str_evaluate_impl() {
        let s = "1+1";
        assert_eq!(s.evaluate().unwrap(), 2.);
    }

    #[test]
    fn check_string_evaluate_impl() {
        let s = String::from("1+1");
        assert_eq!(s.evaluate().unwrap(), 2.);
    }
}
