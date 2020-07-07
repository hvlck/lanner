const OPERATORS: [char; 4] = ['+', '-', '/', '*'];

/// Represents a mathematical expression.
/// 
/// The `original` field represents the original expression, in a `String`.
/// 
/// The `pieces` field represents a `vector` of `ExpressionFragment`, containing the pieces of `original` broken into chunks to iterate over.
/// 
/// The `total` field is an Option<T> enum, representing the value of the expression (Some<f64>), or the lack of being solved (None).
pub struct Expression {
    pub original: String,
    pub pieces: Vec<ExpressionFragment>,
    total: Option<f64>
}

impl Expression {
    /// Returns a new `Expression` struct.  Accepts only a `String`, representing the original expression.
    pub fn new(exp: String) -> Expression {
        let mut pieces: Vec<ExpressionFragment> = Vec::new();

        let mut nums: Vec<f64> = Vec::new();
        
        let mut iter = exp.chars().fuse();

        while let Some(item) = iter.next() {
            if item.is_numeric() {
                match &iter.next() {
                    Some(val) => {
                        match item.to_owned().to_digit(10) {
                            Some(val) => nums.push(val.into()),
                            None => continue
                        }
                    },
                    None => {
                        for operator in OPERATORS.iter() {
                            if exp.contains(*operator) {
                                pieces.push(ExpressionFragment::new(&nums, *operator));
                            }
                        }
                        
                        nums.clear();
                    }
                }
            } else if !item.is_numeric() {
                for operator in OPERATORS.iter() {
                    if exp.contains(*operator) {
                        pieces.push(ExpressionFragment::new(&nums, *operator));
                    }
                }

                nums.clear();
            }
        }

        Expression {
            original: exp,
            pieces: pieces,
            total: None
        }
    }
}

pub struct ExpressionFragment {
    pub operation: char,
    pub pieces: Vec<f64>
}

impl ExpressionFragment {
    pub fn new(original: &Vec<f64>, operation: char) -> ExpressionFragment {
        println!("{:?} {}", original.clone(), operation);
        ExpressionFragment {
            operation: operation,
            pieces: original.clone()
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn it_works() {
        let example = Expression {
            original: String::from("1+1/2"),
            pieces: vec![ExpressionFragment::new(vec![1.0, 1.0], '+'), ExpressionFragment::new(vec![2.0], '/')],
            total: None
        };

        let from_new = Expression::new(String::from("1+1/2"));
        assert_eq!(from_new.original, example.original);
    }
}
