import sys
S = input()

while len(S) > 0:
    if S[-len("dream"):] == "dream":
        S = S[:-len("dream")]
    elif S[-len("erase"):] == "erase":
        S = S[:-len("dream")]
    elif S[-len("dreamer"):] == "dreamer":
        S = S[:-len("dreamer")]
    elif S[-len("eraser"):] == "eraser":
        S = S[:-len("eraser")]
    else:
        print("NO")
        sys.exit(0)
print("YES")
