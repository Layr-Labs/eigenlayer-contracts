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

Now, use the following command to list the commits since that tag:

```
git log v1.4.2..HEAD --pretty=format:"- %s (%an)"
```

This will show:

- Only commits since v1.2.3 up to the current HEAD
- One-line commit messages (%s) with the author name (%an)


An example output is:

```
% git log v1.4.2..HEAD --pretty=format:"- %s (%an)" --no-merges
- ci: add explicit permissions to workflows to  mitigate security concerns (#1392) (bowenli86)
- ci: remove branch constraint for foundry coverage job (Bowen Li)
- docs: add release managers to changelogs (Bowen Li)
- docs: add templates for changelog and release notes (#1382) (bowenli86)
- docs: add doc for steps to write deploy scripts (#1380) (bowenli86)
- ci: add testnet envs sepolia and hoodi to validate-deployment-scripts (#1378) (bowenli86)
...
```

3. Commit the Changelog, Before Cutting Release

Copy the output and add here with a commit, then proceed to cut the release from the commit.
