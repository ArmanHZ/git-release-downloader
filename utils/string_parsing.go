package utils

import (
	"errors"
	"regexp"
)

var githubRepoRegex = regexp.MustCompile(
	// Capturing group will extract `onwer/repo-name`
	`^(?:https?:\/\/)?(?:www\.)?github\.com\/([A-Za-z0-9_.-]+\/[A-Za-z0-9_.-]+)(?:\.git)?\/?$`,
)

func ExtractOwnerAndRepoNames(repoURL string) (repo string, err error) {
	matches := githubRepoRegex.FindStringSubmatch(repoURL)
	if len(matches) < 2 {
		return "", errors.New("Error parsing the url.")
	}

	return matches[1], nil
}

func GetAssetsFromRelease(release Release) []Asset {
	return release.Assets
}

func AssetDigestSpaceCalc(assets []Asset) int {
	maxNameLen := 0
	for _, asset := range assets {
		if len(asset.Name) > maxNameLen {
			maxNameLen = len(asset.Name)
		}
	}

	maxNameLen += 2
	return maxNameLen
}
