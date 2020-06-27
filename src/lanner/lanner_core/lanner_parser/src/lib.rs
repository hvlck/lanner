pub struct Expression {
    pub original: String,
    pieces: Vec<char>,
    total: f64
}

impl Expression {
    pub fn new(exp: String) -> Expression {
        let mut pieces = Vec::new();

        for character in exp.chars() {
            pieces.push(character);
        }

        Expression {
            original: exp,
            pieces: pieces
            total: 0.0
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
}
