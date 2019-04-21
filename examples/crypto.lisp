(defun check (...mtp, i)
    (if (= i 0)
        (throw "Bad function...")
        (sha256 (cat
                    (head mtp)
                    (check rest mtp 32))))

(main (loc ...mtp)
    (and
        ; Ensure that the transaction id is the same
        (= (get mtp loc) 0x0000)

        ; Iterate through each element of the set, max depth 32
        ; 32 = 4 billion transactions
        (for (( 0) (< i 32) (+ i 1))
            (if (eq (length loc) 0)
                ()
                (sha256 (head mtp) (head mtp))))))