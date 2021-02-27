use crate::{constants::*, AstNode, Constant, Function, LannerError, LannerParser, Operation};

pub fn evaluate_constant(constant: Constant) -> Result<f64, LannerError> {
    match constant {
        Constant::E => Ok(E),
        Constant::Pi => Ok(PI),
        Constant::Tau => Ok(TAU),
        Constant::I => Ok(i()),
        _ => Err(LannerError::InvalidConstant),
    }
}

pub fn evaluate_simple_expression(node: AstNode) -> Result<f64, LannerError> {
    match node {
        AstNode::Expression { lhs, rhs, operator } => match operator {
            Operation::Add => Ok(lhs + rhs),
            Operation::Subtract => Ok(lhs - rhs),
            Operation::Multiply => Ok(lhs * rhs),
            Operation::Divide => Ok(lhs / rhs),
            Operation::Modulo => Ok(lhs % rhs),
            Operation::Exponent => Ok(lhs.powf(rhs)),
            _ => Err(LannerError::InvalidExpression),
        },
        AstNode::Constant(constant) => match evaluate_constant(constant) {
            Ok(result) => Ok(result),
            Err(err) => Err(err),
        },
        AstNode::Value(value) => Ok(value),
        _ => Err(LannerError::Other(
            "Expression contains something other than simple arithmetic.",
        )),
    }
}

pub fn evaluate_function(node: AstNode) -> Result<f64, LannerError> {
    match node {
        AstNode::FunctionOp { function, value } => match function {
            Function::Cos => Ok(value.cos()),
            Function::Sin => Ok(value.sin()),
            Function::Tan => Ok(value.tan()),
            Function::Sqrt => Ok(value.sqrt()),
            Function::Acos => Ok(value.acos()),
            Function::Asin => Ok(value.asin()),
            Function::Atan => Ok(value.atan()),
            Function::Abs => Ok(value.abs()),
            _ => Err(LannerError::InvalidFunction),
        },
        _ => Err(LannerError::InvalidInput),
    }
}
