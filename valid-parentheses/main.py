class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        for c in s:
            if c in ['(', '[', '{']:
                stack.append(c)
            else:
                try:
                    popped = stack.pop()
                except IndexError:
                    return False
                if popped == "(" and c == ")":
                    continue
                if popped == "[" and c == "]":
                    continue
                if popped == "{" and c == "}":
                    continue
                return False
        return len(stack) == 0

s = Solution()
print(s.isValid("()"))
print(s.isValid("()[]{}"))
print(s.isValid("(]"))
print(s.isValid("["))
print(s.isValid("]"))