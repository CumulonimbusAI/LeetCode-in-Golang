//@author - vidya ranganathan

func isValid(s string) bool {
    // initialize the stack
    stack := ""

    // iterate over the string 
    for _, c := range s {
        // if character is open brace, push into the stack
        if c == '(' || c == '{' || c == '[' {
            // push
            stack += string(c)
        }else if c == ')' && len(stack)>0 && stack[len(stack)-1] == '(' {
            // pop out of stack if matching brace is found
            stack = stack[:len(stack)-1]
        }else if c == '}' && len(stack)>0 && stack[len(stack)-1] == '{' {
            // pop out of stack if matching brace is found
            stack = stack[:len(stack)-1]
        }else if c == ']' && len(stack)>0 && stack[len(stack)-1] == '[' {
            // pop out of stack if matching brace is found
            stack = stack[:len(stack)-1]
        }else {
            return false
        }
    }
    // verify if stack is empty
    return len(stack) == 0
}
