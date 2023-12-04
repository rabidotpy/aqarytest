package questions

import (
	"sort"
)


func Questions2(s string) string {
    // here we will store the frequency of each character in a key: value pair
    characterExistenceMap := make(map[rune]int)

    // Counting characters just like I'd do in JavaScript
    for _, character := range s {
        characterExistenceMap[character]++
    }

    // Figuring out the maximum frequency
    maxExistence := 0
    for _, existence := range characterExistenceMap {
        if existence > maxExistence {
            maxExistence = existence
        }
    }

    // Aha! If a character appears too much, can't rearrange
    if maxExistence > (len(s)+1)/2 {
        return ""
    }

    // Sort characters by frequency (like sort() in JS but feels different)
    sortedCharacters := make([]rune, 0, len(characterExistenceMap))
    for character := range characterExistenceMap {
        sortedCharacters = append(sortedCharacters, character)
    }
    sort.Slice(sortedCharacters, func(i, j int) bool {
        return characterExistenceMap[sortedCharacters[i]] > characterExistenceMap[sortedCharacters[j]]
    })

    // rearrange chars
    rearranged := make([]rune, len(s))
    index := 0
    for _, character := range sortedCharacters {
        count := characterExistenceMap[character]
        for count > 0 {
            rearranged[index] = character
            index += 2
            if index >= len(s) {
                index = 1
            }
            count--
        }
    }

    // return the rearranged string if we got here
    return string(rearranged)
}



