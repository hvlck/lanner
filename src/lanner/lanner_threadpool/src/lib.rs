use std::thread;

use lanner_parser::*;

pub struct SolverController {
    pub original: String,
    expression: Expression,
    solvers: Vec<Solver>
}

impl SolverController {
    pub fn new(exp: String) -> SolverController {
        let new_exp = Expression::new(exp.clone());
        
        let mut threads = Vec::new();

        for piece in new_exp.pieces.iter() {
            threads.push(Solver::new(piece))
        }

        SolverController {
            original: exp,
            expression: new_exp,
            solvers: threads
        }
    }
}

struct Solver {
    expression: Vec<char>,
    thread: thread::JoinHandle<()>
}

impl Solver {
    fn new(exp: Vec<char>) -> Solver {
        let clone = exp.clone();
        let thread = thread::spawn(move || {
            
        });

        Solver {
            expression: clone,
            thread: thread
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
