package a

func _()                {}              // OK
func _() int            { return 1 }    // OK
func _() (int, int)     { return 1, 1 } // OK
func _() (a int)        { return 1 }    // want "a is named return value"
func _() (a int, b int) { return 1, 1 } // want "a is named return value" "b is named return value"
func _() (a, b int)     { return 1, 1 } // want "a is named return value" "b is named return value"
