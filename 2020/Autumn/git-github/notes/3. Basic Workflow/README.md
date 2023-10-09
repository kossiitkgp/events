# Let's start your journey with Git.

### 1. Creating a local repository. (Git command: git init)
- Simply cd into the directory you want to track and type `git init`. This will create a `.git` folder in the directory with a basic skeleton without any commits.

### 2. Checking files which can be commited. (Git command: git status)
>`$ git status`
- This will list the files that have been changed in red and those which can be commited in green

- To ignore files that you dont want to track create a `.gitignore` file using `touch .gitignore`
- Open this using text editor and add names of files you want to ignore. 

### 3. Staging Area.
- Moving files to staging area:
- For adding files individually use 
> `$ git add <filename>`
- For adding all files at once use 
> `$ git add -A`
- To remove files from staging area use 
> `$ git reset <filename>`
- and to remove all simply type 
> `$ git reset`

### 4. To track the changes you made:
> `$ git diff`

### 5. To commit the files:
> `$ git commit`
- but it is necessary to add messages with the commit to make sure what we did ; so for that we use `-m` extension like `git commit -m “message”`
- You can also compare the changes you made w.r.t a commit using `git diff <commit>`.If `<commit>` field is kept empty, git will compare it with the last commit.

### 6. To check the commit history:
>` $ git log`

>` $ git log --graph` presents the change history in the tree format. 

### 7. Creating a repository on GitHub. (Git command: git remote)
>`$ git remote add origin <URL>`
- and to update it use `git push origin master`(explained later).

### 8. Uploading your changes to Github:
- We use `git push origin <branchname>` to upload the changes to Github.
- But before pushing it back we need to pull and check whether any other person made a change in the branch since the last time type in `git pull origin master` ( If on master branch else use branch name in place of master).

##### Congrats you made your first commit!

### 9. Fetching changes without merging
- `pull` fetches the changes and merges it in directly to the branch you are in. That may cause some problems if there are conflicts. Instead, it's better to use `fetch`. 
- `fetch` does exactly what it says. It fetches commits from a repo/branch and you can either merge it or inspect it by checking out to that branch.

### 10. Branching:
- To create a new branch use `git branch <branchname>`
- To check all the branches present use `git branch`
- To switch over branch use `git checkout <branchname>`
- To do the above processes of creating and switching together use `git checkout -b <branchname>`
- To push a commited changes type in `git push -u origin <branchname>`
- To merge the current branch you are working on to master use `git merge <branchname>`
- To push the changes we simply use `git push origin master`
- To check it they are merged type in `git branch --merged`
- If the branch is successfully merged we can now delete the branch–`git branch -d <branchname>`

### 11. Switching to previous commit version (Extreme Situations):

- We do `git reset --hard <SOME-COMMIT>` . Remember all code changes will be lost after this commit number.
- If we do `git reset <SOME-COMMIT>` . It will reach to a state where the changes made after the commit number were still to added to the staging area.

#### Pull Request:
- Send a pull request to sandbox repository.
