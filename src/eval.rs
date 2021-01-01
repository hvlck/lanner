use crate::{AstNode, LannerError, Operation};

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
        _ => Err(LannerError::Other(
            "Expression contains something other than simple arithmetic.",
        )),
    }
}
