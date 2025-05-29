# <release-version>

**Use this template to draft changelog and submit PR to review by the team**

## Release Manager

github handle of release manager

## Highlights

🚀 New Features – Highlight major new functionality
- ...
- ...

⛔ Breaking Changes – Call out backward-incompatible changes.
- ...
- ...

📌 Deprecations – Mention features that are being phased out.
- ...
- ...

🛠️ Security Fixes – Specify patched vulnerabilities.
- ...
- ...

🔧 Improvements – Enhancements to existing features.
- ...
- ...

🐛 Bug Fixes – List resolved issues.
- ...
- ...


## Changelog

To generate a changelog of commits added since the last release using Git on 
the command line, follow these steps:

1. Identify the last release tag

First, list your tags (assuming you use Git tags for releases):

```
git tag --sort=-creatordate
```

This shows your most recent tags at the top. Let's say the last release tag is `v1.4.2`


2. Generate the changelog

Now, use the following command to list the commits since that tag, and auto generate github PR link if there's any

```
git log v1.4.2..HEAD --pretty=format:"%s" --no-merges | \
sed -E 's/^(.*)\(#([0-9]+)\)$/- \1[PR #\2](https:\/\/github.com\/layr-labs\/eigenlayer-contracts\/pull\/\2)/' | \
sed -E '/\[PR #[0-9]+\]/! s/^(.*)$/- \1/'
```

This will show:

- Only commits since v1.4.2 up to the current HEAD
- One-line commit messages (%s) with the author name (%an)


An example output is:

```
- ci: add explicit permissions to workflows to  mitigate security concerns [PR #1392](https://github.com/layr-labs/eigenlayer-contracts/pull/1392)
- ci: remove branch constraint for foundry coverage job
- docs: add release managers to changelogs
- docs: add templates for changelog and release notes [PR #1382](https://github.com/layr-labs/eigenlayer-contracts/pull/1382)
- docs: add doc for steps to write deploy scripts [PR #1380](https://github.com/layr-labs/eigenlayer-contracts/pull/1380)
- ci: add testnet envs sepolia and hoodi to validate-deployment-scripts [PR #1378](https://github.com/layr-labs/eigenlayer-contracts/pull/1378)
- docs: update MAINTENANCE to include practices of merging multiple release-dev branches
- docs: updating readme for dead links, readability, new language, and more [PR #1377](https://github.com/layr-labs/eigenlayer-contracts/pull/1377)
...
```

3. Commit the Changelog

Copy the output and add here with a commit, then proceed to cut the release from the commit.
