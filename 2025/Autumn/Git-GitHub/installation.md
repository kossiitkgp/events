# Git Installation Guide

1. Install Git
2. Make a GitHub account
3. Authenticating to GitHub in the terminal

## Install Git

### Linux

Here are a few links to help you install git on your Linux system:

- https://git-scm.com/download/linux
- https://www.atlassian.com/git/tutorials/install-git#linux

### Windows

Here are a few links to help you install git on your Windows system:

- https://phoenixnap.com/kb/how-to-install-git-windows
- https://www.stanleyulili.com/git/how-to-install-git-bash-on-windows/

### Mac OS X

Here are a few links to help you install git on your Mac OS X system:

- https://git-scm.com/download/mac
- https://www.atlassian.com/git/tutorials/install-git#mac-os-x

------

After Git is installed in your system, run the following command in a newly opened shell window:

```bash
git --version
```
If you see a version number, then this confirms that git has been properly installed!

## Create a GitHub Account

If you don't have a GitHub account already, go to https://github.com, click on Sign Up, and create your account.

## Authenticating to GitHub in the terminal

> [!NOTE]
> We will be using `gh` CLI, a command line utility by GitHub, for setting up authentication in Git.
> For other ways to set this up, see https://docs.github.com/en/get-started/git-basics/caching-your-github-credentials-in-git.

1. Install GH CLI (GitHub CLI): Follow the download instructions from https://cli.github.com/. Once it has been installed, run the following command in a newly opened shell window:
   ```bash
   gh --version
   ```
   If you see a version number, this confirms that the `gh` utility has been installed.

2. Run the command
   ```bash
   gh auth login
   ```
   Go through the wizard, answering the prompts. When a browser window is opened, login to GitHub, and paste the 'one-time code' in your browser.
   ```
   ? What account do you want to log into? GitHub.com
   ? What is your preferred protocol for Git operations on this host? HTTPS
   ? Authenticate Git with your GitHub credentials? Yes
   ? How would you like to authenticate GitHub CLI? Login with a web browser

   ! First copy your one-time code: XXXX-XXXX
   Press Enter to open github.com in your browser...
   ✓ Authentication complete.
   ✓ Configured git protocol
   ✓ Logged in as <username>
   ```

3. Finally, set your identity in git (name and email), this should match the email connected to your GitHub account.
    ```
    git config --global user.name "Your Name"
    git config --global user.email "your@email.com"
    ```
   
 You are now ready to master the ways of git!
