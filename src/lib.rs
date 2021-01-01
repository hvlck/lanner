#[macro_use]
// std

// crates
use pest::Parser;
use pest_derive::*;

// local

#[derive(Parser)]
#[grammar = "lib.pest"]
pub struct LannerParser;

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn check_number() {
        let whole = LannerParser::parse(Rule::number, "10");
        assert!(whole.is_ok());
        let integer = LannerParser::parse(Rule::number, "10.15");
        assert!(integer.is_ok());
    }

    #[test]
    fn check_operators() {}

    #[test]
    fn check_expression() {
        let exp = LannerParser::parse(Rule::expression, "10+10");
        assert!(exp.is_ok());

        let exp_with_spaces = LannerParser::parse(Rule::expression, "10 + 10");
        assert!(exp_with_spaces.is_ok());
    }
}
