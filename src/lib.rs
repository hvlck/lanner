#[macro_use]
// std

// crates
use pest_derive::*;
use pest::{error::Error, Parser};

// local
mod eval;
use eval::evaluate_simple;

#[derive(Debug, Clone)]
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

#[derive(Debug, Clone)]
pub enum Function {
    Cos,
    Sin,
    Tan,
}

#[derive(Debug, Clone)]
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

pub fn parse(src: &str) -> Result<Vec<AstNode>, Error<Rule>> {
    let mut ast: Vec<AstNode> = Vec::new();

    let exp = LannerParser::parse(Rule::expression, src)?;
    //    println!("{:#?}", exp); // debugging
    for pair in exp {
        match pair.as_rule() {
            Rule::expression => {
                ast.push(parse_to_ast(pair));
            }
            _ => ast.push(AstNode::Error("Something went wrong.")),
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
        }
        Rule::function => {
            let mut pair = pair.into_inner();
            let function = pair.next().unwrap();
            let value = pair.next().unwrap().as_str().parse::<f64>().unwrap();

            let function = parse_function(function);

            match function {
                Ok(function) => AstNode::FunctionOp {
                    expr: value,
                    function,
                },
                Err(()) => AstNode::Error("Failed to parse function."),
            }
        }
        _ => AstNode::Error("Invalid input."),
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
    match operator.as_str() {
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

fn parse_function(function: pest::iterators::Pair<Rule>) -> Result<Function, ()> {
    match function.as_str() {
        "sin" => Ok(Function::Sin),
        "cos" => Ok(Function::Cos),
        "tan" => Ok(Function::Tan),
        _ => Err(()),
    }
}

pub fn evaluate(src: &str) -> Result<f64, String> {
    let ast = parse(src);
    match ast {
        Ok(ast) => {
            let pair = ast.get(0).unwrap();
            match pair {
                //                AstNode::FunctionOp { function, expr } => {}
                AstNode::Expression { lhs, rhs, operator } => {
                    match evaluate_simple(pair.to_owned()) {
                        Ok(result) => Ok(result),
                        Err(()) => Err(String::from("Somethign went wrong.")),
                    }
                }
                AstNode::Error(_) => Err(String::from("Something went wrong.")),
                _ => Err(String::from("Something went wrong.")),
            }
        }
        Err(error) => Err(String::from("Something went wrong.")),
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn parse_fn_test() {
        parse("20 + 20");
        parse("sin(20)");
    }

    #[test]
    fn addition_eval() {
        assert_eq!(evaluate("20 + 30").unwrap(), 50.0);
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
