# Boggle (Find all possible words in a board of characters) 

Problem description: [GeeksforGeeks](https://www.geeksforgeeks.org/boggle-find-possible-words-board-characters/amp/)

Given a dictionary, a method to do lookup in dictionary and a M x N board where every cell has one character. Find all possible words that can be formed by a sequence of adjacent characters. Note that we can move to any of 8 adjacent characters, but a word should not have multiple instances of same cell.

```bash
Input: dictionary[] = {"GEEKS", "FOR", "QUIZ", "GO"};
       boggle[][]   = {{'G', 'I', 'Z'},
                       {'U', 'E', 'K'},
                       {'Q', 'S', 'E'}};
      isWord(str): returns true if str is present in dictionary
                   else false.

Output:  Following words of dictionary are present
         GEEKS
         QUIZ
```

## Algorithm

* create the graph with the links between the nodes
* iterate over the list of Nodes until we found one with the Value == first letter of the word
* apply a modified breadth for search algorithm on the graph that add to `visited` only the nodes which satisfy the condition that next letter of the word == Node.Value
* run the for loop until either the word was found, by comparing the index with the length of the word, or no `Next` suitable Node was found

## Result

```bash
$ go run main.go 
This word was found: GEEKS
This word was found: QUIZ
```
