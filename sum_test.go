package main

import "testing"

func TestSumFunc(t *testing.T) {
    result := sum(5, 5)
    if result != 10 {
       t.Errorf("Func sum() error: expected %d, got: %d.", 10, result)
    }
}