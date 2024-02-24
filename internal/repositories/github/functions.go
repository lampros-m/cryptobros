package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"repositories/cryptobros/internal/networking"
)

var (
	GitHubValidPattern       = `^https:\/\/github\.com\/([a-zA-Z0-9-]+)\/?([a-zA-Z0-9-]+)?$`
	GitHubOrganizaiton       = `^https:\/\/github\.com\/([^\/]+)`
	GitHubAPIRepositoriesURL = "https://api.github.com/orgs/%s/repos"
)

func GetRepositoriesPerOrganization(url string) (Repositories, error) {
	organization, err := extractOrganizationFromURL(url)
	if err != nil {
		return Repositories{}, err
	}

	repositoriesURL := fmt.Sprintf(GitHubAPIRepositoriesURL, organization)
	bytes, statusCode, err := networking.GetRequest(repositoriesURL)
	if err != nil {
		return Repositories{}, errors.New("error on GetRepositoriesPerOrganization: networking.GetRequest: " + url + err.Error())
	}

	if statusCode == networking.StatusNotFound {
		return Repositories{}, nil
	}

	var repositories []Repository
	err = json.Unmarshal(bytes, &repositories)
	if err != nil {
		return Repositories{}, errors.New("error on GetRepositoriesPerOrganization: json.Unmarshal: url: " + url + fmt.Sprint(statusCode) + err.Error())
	}

	return repositories, nil
}

func extractOrganizationFromURL(url string) (string, error) {
	if !regexp.MustCompile(GitHubValidPattern).MatchString(url) {
		return "", ErrInvalidGithubURL
	}

	organizationSlice := regexp.MustCompile(GitHubOrganizaiton).FindStringSubmatch(url)
	if len(organizationSlice) != 2 {
		return "", ErrNoOrganizationFound
	}

	organization := organizationSlice[1]

	return organization, nil
}
