use crate::{AstNode, Function, LannerError, LannerParser, Operation};

pub fn evaluate_simple_expression(node: AstNode) -> Result<f64, LannerError> {
    match node {
        AstNode::Expression { lhs, rhs, operator } => match operator {
            Operation::Add => Ok(lhs + rhs),
            Operation::Subtract => Ok(lhs - rhs),
            Operation::Divide => Ok(lhs / rhs),
            Operation::Multiply => Ok(lhs * rhs),
            Operation::Exponent => Ok(lhs.powf(rhs)),
            _ => Err(LannerError::InvalidExpression),
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
            _ => Err(LannerError::InvalidFunction),
        },
        _ => Err(LannerError::InvalidInput),
    }
}
