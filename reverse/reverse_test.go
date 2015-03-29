package main

import (
    "testing"
    "strings"
)
func TestReverse(t *testing.T) {
    testCases := []struct {
        Input []string
        WantWords []string
    }{
        {
            Input: []string{ "dog", "god", "tea" },
            WantWords: []string{ "dog", "god" },
        },
        {
            Input: []string{ "aet", "god", "tea" },
            WantWords: []string{ "aet", "tea" },
        },
        {
            Input: []string{ "2dc", "god", "tea" },
            WantWords: []string{},
        },
        {
            Input: []string{ "tsang", "god", "tea" },
            WantWords: []string{},
        },
        {
            Input: []string{ "tsang", "gnast", "tea" },
            WantWords: []string{ "tsang", "gnast" },
        },
        {
            Input: []string{ "tsang", "gnast", "tea", "aet", "god", "dog" },
            WantWords: []string{ "tsang", "gnast", "tea", "aet", "god", "dog" },
        },
    }
    
    for _, testcase := range testCases {
        result := ReverseMatch(testcase.Input)
        t.Logf("Got result: %s", strings.Join(result, ","))
        t.Logf("Want result: %s", strings.Join(testcase.WantWords, ","))
        if strings.Join(result, ",") != strings.Join(testcase.WantWords, ",") {
            t.Errorf("testCase mismatched for '%s'", strings.Join(testcase.Input, " ")) 
        }
    }
    
}
