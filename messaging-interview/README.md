# Exercise for Messaging Team at The New York Times

To better assess a candidate's programming skills, we ask that they solve the following challenge. You have as much time as you'd like, though candidates normally take between 3 and 7 business days to complete the exercise.

## Instructions

1. First, create a branch. This is a Git repository, so you should show your familiarity with Git by creating a branch and committing all your changes to that branch. The final submission will be the zipped Git repository with your branch.
2. Write the program that solves the challenge described below.
3. Finally, send your submission to the HR contact at The New York Times via email. Please make a note about the name of your Git branch in your email. They will pass it along to the team for review, which normally takes 1-3 business days.

## Technical Constraints

* You must use Go 1.8. The team is currently using this version of Go. This will help prevent issues with running your program.
* Be aware that the reviewers will be using machines with MacOS installed. If you pre-build your program and ask that the built program be run, make sure you built it in MacOS. Otherwise, make it clear you built it in a different OS so reviewers know to re-build it before running it.

## Project Description

This problem consists of finding a valid way to separate a string consisting of concatenated words from a dictionary. This input string can be found in a file named `concatenated.txt`. The words in the dictionary can be found separated by newlines in a file named `words.txt`. Both files are part of the project and you can find them in this directory.

For this problem, use the Go language to create an API that reads in the input string from `concatenated.txt` and dictionary words from `words.txt`, separating the entire concatenated string into a set of valid dictionary words. For example, if the input string is "applepie" and the dictionary contains a standard set of English words, "apple pie" would be a valid separation. You should provide endpoints to output the separated words in a clear format of your choice; printing one solution word per line is fine. 

You can assume that the input string consists solely of a concatenation of words randomly selected from the dictionary. In other words, there will be a way to space-separate it into dictionary words. There may be multiple valid ways to separate the string but all valid separations are equally valid and no specific one is preferred. You can assume that the entire dictionary can be loaded into memory, but do not make assumptions about the maximum length of words in the dictionary.

### Here's what your API must do:

1. It must be written in Go.
2. It must have an endpoint `GET /healthcheck` that returns a json response such as `{ 'status':'OK' }`.
3. It must include an endpoint `POST /dictionary` that will upload a file containing words and initialize the API’s dictionary.
4. It must include an endpoint `POST /split` that will upload a file containing concatenated words and reply with the separated words, one per line.
5. It must include a `SOLUTION.md` file with an explanation of your approach and analysis of the worst case runtime in terms of the length of the concatenated string. Please also include instructions for how to run your solution. This is very crucial. Reviewers must be able to easily run your program. The easier you make it to run, the better.
5. All undefined routes should return a 404.


### Your application does not need to:

1. Handle authentication or authorization (bonus points if it does, extra bonus points if it's via Google Cloud Endpoints). If you do handle authentication/authorization, make sure it's clear that it is a requirement for making requests, explain how to be authenticated/authorized, and make it easy for reviewers to actually be authenticated/authorized.
2. Be written with any particular framework (although bonus points for using our Gizmo framework https://github.com/nytimes/gizmo).

Your application should be easy to set up and should run on either Linux or Mac OS X. It should not require any for-pay software.

## Evaluation

Evaluation of your submission will be based on the following criteria. Additionally, reviewers will attempt to assess your familiarity with standard libraries and how you've structured your submission.

1. Did your application fulfill the basic requirements?
2. Did you document the method for setting up and running your application?
3. Did you follow the instructions for submission?
4. Did you follow best practices, methodologies like 12factor app?
5. Is the program well-tested?
6. Would the program easily pass through a code review cycle?
