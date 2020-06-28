/// Represents a mathematical expression.
/// 
/// The `original` field represents the original expression, in a `String`.
/// 
/// The `pieces` field represents a `vector` of `chars`, containing the pieces of `original` broken into chunks to iterate over.
/// 
/// The `total` field is an Option<T> enum, representing the value of the expression (Some<f64>), or the lack of being solved (None).
pub struct Expression {
    pub original: String,
    pub pieces: Vec<Vec<char>>,
    total: Option<f64>
}

impl Expression {
    /// Returns a new `Expression` struct.  Accepts only a `String`, representing the original expression.
    pub fn new(exp: String) -> Expression {
        let mut pieces = Vec::new();

        let mut current: Vec<char> = Vec::new();
        for character in exp.chars() {
            current.push(character);
            match exp.chars().next() {
                Some(val) if !character.is_numeric() => {
                    current.push(val);
                    pieces.push(current.clone());
                    current.clear();
                },
                Some(_) => {
                    continue
                },
                None => break
            }
        }

        Expression {
            original: exp,
            pieces: pieces,
            total: None
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
            pieces: vec![vec!['1', '+', '1']],
            total: None
        };
        let from_new = Expression::new(String::from("1+1"));
        assert_eq!(from_new.original, example.original);
        assert_eq!(from_new.total, example.total);
        assert_eq!(from_new.pieces, example.pieces);
    }
}
