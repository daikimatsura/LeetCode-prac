package main

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
 

Example 1:

Input: s = "()"

Output: true

Example 2:

Input: s = "()[]{}"

Output: true

Example 3:

Input: s = "(]"

Output: false

Example 4:

Input: s = "([])"

Output: true

 

Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

func isValid(s string) bool {
    var stockOpenS []string
    for _, str := range s {
        targetS := string(str)
        if !isOpenS(targetS) && len(stockOpenS) == 0 {
            return false
        }
        if !isOpenS(targetS) && !isCloseValidS(stockOpenS[len(stockOpenS) - 1], targetS) {
            return false   
        }
        if isOpenS(targetS) {
            stockOpenS = append(stockOpenS, targetS)
        }else{
            stockOpenS = stockOpenS[:len(stockOpenS) - 1]
        }
    }
    return len(stockOpenS) == 0
}

func isOpenS(s string) bool {
    return s == "(" || s == "{" || s == "["
}

func isCloseValidS(openS string, closeS string) bool {
    return openS == "(" && closeS == ")" || openS == "{" && closeS == "}" || openS == "[" && closeS == "]"
}