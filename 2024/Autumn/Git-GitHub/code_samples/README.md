# Contributing guidelines

The above sample files contain a bug as mentioned in the issue. Do fix a bug in the files which you can and send a PR to this repository. Please make sure you write proper commits and commit messages.

A good practice would be to have one commit for each bug fix, with an appropriate message.

## Instructions to send a PR

1. Before starting to work, you need to have a copy of this repo on your github, `Fork` it
2. After forking the repository, you have a copy of it in your repositories list on github. To bring that repo on your machine - `Clone` it
3. Once the codebase is one your machine, move to the directory. It is always a good idea to not experiment on `main` branch. Make a new branch - `git branch BRANCH_NAME`
4. Hop on to the branch you just created - `git checkout BRANCH_NAME`
5. Fix the bugs, add the files to the staging area, simplest way - `git add -A` (adds all the files in which changes are made)
6. Commit it along with a proper message - `git commit -m  "A proper commit message which explains what you did in simple and short way"`
7. Push to the github - `git push origin BRANCH_NAME` (Note that you are pushing to the forked repository)
8. Open the forked repository on github, it shows you an option for sending a pull request.
9. Click on the button.
10. While sending a PR add a proper title and description so that the repo maintainer gets an idea of what the PR is about. It is always a good idea to mention the issue in the description of the PR. For example, if you are fixing issue number - 5, you can include `Fixes #5`

Check out [MetaKGP](https://github.com/metakgp) repositories and [good first issues](https://goodfirstissue.dev/) to start open-source contributions by creating PRs!