// Convenience type to allow Result return types to only specify the type
// for the true case; failures are specified as static strings.
pub type Result<T> = std::result::Result<T, &'static str>;

// The function under test.
//
// Accepts a string and an integer and returns the
// result of sticking them together separated by a dash as a string.
pub fn join_params_with_dash(str: &str, num: i32) -> Result<String> {
    if str == "" {
        return Err("string cannot be blank");
    }

    if num <= 0 {
        return Err("number must be positive");
    }

    let result = format!("{}-{}", str, num);

    Ok(result)
}



#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_join_params_with_dash() {
        // This is a type used to record all details of the inputs
        // and outputs of the function under test.
        #[derive(Debug)]
        struct TestData<'a> {
            str: &'a str,
            num: i32,
            result: Result<String>,
        }
    
        // The tests can now be specified as a set of inputs and outputs
        let tests = &[
            // Failure scenarios
            TestData {
                str: "",
                num: 0,
                result: Err("string cannot be blank"),
            },
            TestData {
                str: "foo",
                num: -1,
                result: Err("number must be positive"),
            },

            // Success scenarios
            TestData {
                str: "foo",
                num: 42,
                result: Ok("foo-42".to_string()),
            },
            TestData {
                str: "-",
                num: 1,
                result: Ok("--1".to_string()),
            },
        ];
    
        // Run the tests
        for (i, d) in tests.iter().enumerate() {
            // Create a string containing details of the test
            let msg = format!("test[{}]: {:?}", i, d);
    
            // Call the function under test
            let result = join_params_with_dash(d.str, d.num);
    
            // Update the test details string with the results of the call
            let msg = format!("{}, result: {:?}", msg, result);
    
            // Perform the checks
            if d.result.is_ok() {
                assert!(result == d.result, "{}", msg);
                continue;
            }
    
            let expected_error = format!("{}", d.result.as_ref().unwrap_err());
            let actual_error = format!("{}", result.unwrap_err());
            assert!(actual_error == expected_error, "{}", msg);
        }
    }
}
