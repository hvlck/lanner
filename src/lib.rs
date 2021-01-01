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
    Error(&'static str),
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
    RmdDivision,
    Exponent,
    Radical,
}

#[derive(Parser)]
#[grammar = "lib.pest"]
pub struct LannerParser;

pub fn parse(src: &str) -> Result<AstNode, Error<Rule>> {
    let mut ast: AstNode;

    let exp = LannerParser::parse(Rule::expression, src)?;
    println!("{:#?}", exp); // debugging
    for pair in exp {
        match pair.as_rule() {
            Rule::expression => {
                ast = parse_to_ast(pair);
            }
            _ => ast = AstNode::Error("Something went wrong."),
        }
    }

    Ok(ast)
}

// todo: rename
fn parse_to_ast(pair: pest::iterators::Pair<Rule>) -> AstNode {
    match pair.as_rule() {
        Rule::expression => parse_to_ast(pair.into_inner().next().unwrap()),
        Rule::simple_expression => {
            let mut pair = pair.into_inner();
            let lhs = pair.next().unwrap();
            let op = pair.next().unwrap();
            let rhs = pair.next().unwrap();
            parse_expression(op, lhs, rhs)
        },
        _ => {
            AstNode::Error("Something went wrong.")
        }
    }
}

fn parse_expression(
    operator: pest::iterators::Pair<Rule>,
    lhs: pest::iterators::Pair<Rule>,
    rhs: pest::iterators::Pair<Rule>,
) -> AstNode {
    AstNode::Expression {
        lhs: lhs.as_str().parse::<f64>().unwrap(),
        rhs: rhs.as_str().parse::<f64>().unwrap(),
        operator: parse_operator(operator),
    }
}

fn parse_operator(operator: pest::iterators::Pair<Rule>) -> Operation {
    let operator = operator.as_str();
    match operator {
        "+" => Operation::Add,
        "-" => Operation::Subtract,
        "*" => Operation::Multiply,
        "/" => Operation::Divide,
        "%" => Operation::RmdDivision,
        "**" => Operation::Exponent,
        "*/*" => Operation::Radical,
        _ => Operation::Add, // todo: fix
    }
}

pub fn evaluate() -> Result<f64, Error<Rule>> {
    todo!()
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test() {
        parse("20 + 20");
        parse("sin(20)");
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
