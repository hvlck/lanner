// std
use std::thread;

// external
use rayon::prelude::*;

// local
use lanner_parser::*;

pub struct SolverController {
    pub original: String,
    expression: Expression,
    solvers: Vec<Solver>
}

impl SolverController {
    pub fn new(exp: String) -> SolverController {
        let new_exp = Expression::new(exp.clone());

        let mut solvers = Vec::new();

        for piece in new_exp.pieces.iter() {
            solvers.push(Solver::new(piece.pieces.clone(), piece.operation));
        }

        SolverController {
            original: exp,
            expression: new_exp,
            solvers: solvers
        }
    }
}

struct Solver {
    expression: Vec<f64>,
    operation: char
}

impl Solver {
    fn new(exp: Vec<f64>, operation: char) -> Solver {
        Solver {
            expression: exp,
            operation: operation.to_owned()
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn solve() {
        let solver = SolverController::new(String::from("1+1"), 20);
    }
}
