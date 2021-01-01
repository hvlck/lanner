#[macro_use]
// std

// crates
use pest_derive::*;
use pest::{error::Error, Parser};

// local

#[derive(Debug)]
pub enum AstNode {
    FunctionOp {
        function: Function,
        expr: f64,
    },
    Expression {
        lhs: f64,
        rhs: f64,
        operator: Operation,
    },
}

#[derive(Debug)]
pub enum Function {
    Cos,
    Sin,
    Tan,
}

#[derive(Debug)]
pub enum Operation {
    Add,
    Subtract,
    Divide,
    Multiply,
    RmdDivsiion,
    Exponent,
    Radical,
}

#[derive(Parser)]
#[grammar = "lib.pest"]
pub struct LannerParser;

pub fn parse(src: &str) -> Result<Vec<AstNode>, Error<Rule>> {
    let mut ast = Vec::new();

    let exp = LannerParser::parse(Rule::expression, src)?;
    for pair in exp {
        match pair.as_rule() {
            Rule::number => {}
            Rule::function => {}
            _ => {}
        }
    }
    //    match exp.as_rule() {}
    println!("{:#?}", ast);

    Ok(ast)
}

pub fn evaluate() -> Result<f64, Error<Rule>> {
    Ok(10.0)
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test() {
        parse("20 + 20");
    }
}

#[cfg(test)]
mod parser_tests {
    use super::*;

    #[test]
    fn check_number() {
        let whole = LannerParser::parse(Rule::number, "10");
        assert!(whole.is_ok());
        let integer = LannerParser::parse(Rule::number, "10.15");
        assert!(integer.is_ok());
    }

    #[test]
    fn check_fn() {
        let function = LannerParser::parse(Rule::function, "cos(20)");
        assert!(function.is_ok());
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
