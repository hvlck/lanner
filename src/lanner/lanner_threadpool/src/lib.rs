// std
use std::thread;

// external
use rayon::prelude::*;

// local
use lanner_parser::*;
use lanner_core::*;

pub struct SolverController {
    pub original: String,
    expression: Expression,
    solvers: Vec<Solver>,
    total: f64
}

impl SolverController {
    pub fn new(exp: String) -> SolverController {
        let mut value: f64 = 0.0;
        let new_exp = Expression::new(exp.clone());

        let mut solvers = Vec::new();

        for piece in new_exp.pieces.iter() {
            let solver = Solver::new(piece.pieces.clone(), piece.operation);
            value += solver.value;

            solvers.push(solver);
        }

        SolverController {
            original: exp,
            expression: new_exp,
            solvers: solvers,
            total: value
        }
    }
}

struct Solver {
    expression: Vec<f64>,
    pub value: f64,
    operation: char
}

impl Solver {
    fn new(exp: Vec<f64>, operation: char) -> Solver {
        let result = match operation {
            '+' => exp.par_iter().sum(),
            _ => 0.0
        };

        Solver {
            expression: exp.clone(),
            value: result,
            operation: operation.to_owned()
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn solve() {
        let solver = SolverController::new(String::from("1+1"));

        assert_eq!(solver.total, 2.0);
    }
}
