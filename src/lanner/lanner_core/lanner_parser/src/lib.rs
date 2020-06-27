/// Represents a mathematical expression.
/// ```markdown
/// The `original` field represents the original expression, in a `String`.
/// The `pieces` field represents a `vector` of `chars`, containing the pieces of `original` broken into chunks to iterate over.
/// The `total` field represents the value of the expression.
/// ```
pub struct Expression {
    pub original: String,
    pieces: Vec<char>,
    total: f64
}

impl Expression {
    /// Returns a new `Expression` struct.  Accepts only a `String`, representing the original expression.
    pub fn new(exp: String) -> Expression {
        let mut pieces = Vec::new();

        for character in exp.chars() {
            pieces.push(character);
        }

        Expression {
            original: exp,
            pieces: pieces,
            total: 0.0
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn it_works() {
        let example = Expression {
            original: String::from("1+1"),
            pieces: vec!['1', '+', '1'],
            total: 0.0
        };
        let fromNew = Expression::new(String::from("1+1"));
        assert_eq!(fromNew.original, example.original);
        assert_eq!(fromNew.total, example.total);
        assert_eq!(fromNew.pieces, example.pieces);
    }
}
