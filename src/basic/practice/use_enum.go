package practice
import "fmt"


type State int

const (
    Deleted = -1
    Default State = iota
    Enable    
    Verified
    Disabled 
)

const (
    A = iota
    B
    C
    D = "disabled"
    E 
    F = 100
    G = iota
    H
)

func(this State) String() string {
    switch this {
        case Deleted:
            return "deleted"
        case Default:
            return "initialed"
        case Enable:
            return "enabled"
        case Verified:
            return "verified"
        case Disabled:
            return "disabled"
        default:
            return "Unknown"
    }
}

func UseEnum() {
    state := Disabled
    fmt.Println("state", state)

    fmt.Println(A,B,C,D,E,F,G,H)
}