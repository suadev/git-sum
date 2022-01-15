### git-sum cli tool

See open issue and pull request counts for each repository of the user. 

### Installation

```
go install github.com/suadev/git-sum@latest
```

### Usage

```bash
git-sum             Lists all repositories that have at least one open issue or pull request.
git-sum set-user    Sets GitHub Username. # Required. Can be changed later.
git-sum set-token   Sets GitHub Personel Access Token. # Optional. But, you have 60 requests per hour without token, so you better set your token.
```
<img src="https://github.com/suadev/git-sum/blob/main/how-to.gif" />

### Todo
 
* Add more useful information that is worth tracking for each repository.

