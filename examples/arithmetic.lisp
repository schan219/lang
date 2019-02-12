
# You get to spend bitcoin if you can provide a 3 numbers (a,b,c,d)
# such that 10 = (10+a)*(c+d)
#

(main (a b c d)
    (eq 10
        (*
            (+ a 10)
            (+ c d)))