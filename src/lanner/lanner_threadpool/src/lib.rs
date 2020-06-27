use std::thread;

pub struct SolverController {
    solvers: Vec<Solver>
}

impl SolverController {
    pub fn new(max: usize) -> SolverController {
        SolverController { solvers: vec![] }
    }
}

struct Solver {
    expression: Vec<char>,
    thread: thread::JoinHandle<()>
}

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        
    }
}
