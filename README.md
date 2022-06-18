![GoCoverage](https://img.shields.io/badge/Coverage-99.99%25-brightgreen)
# go-badges
Badges generation

# Usage
- Coverage Badge
```yaml
...
    - name: Run Test
      # Run your test with coverage and generate output file
      run: |
        go test -v ./... -covermode=count -coverprofile=coverage.out
        go tool cover -func=coverage.out -o=coverage.out
    - name: Go Coverage Badge
      # Run action to create coverage badge and update README.md file
      uses: dannysy/go-badges@v1.13
      with:
        # Path to coverage report file
        coveragePath: coverage.out
        # Path to README.md
        # Default: README.md
        readmePath: README.md
# Commit and Push changes to README.md file    
    - name: Commit changes
      run: |
        git config --local user.email "action@github.com"
        git config --local user.name "GitHub Action"
        git add README.md
        git commit -m "chore: Updated coverage badge."

    - name: Push changes
      uses: ad-m/github-push-action@master
      with:
        github_token: ${{ github.token }}
        branch: ${{ github.head_ref }}
...
```        