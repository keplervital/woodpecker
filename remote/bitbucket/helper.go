package bitbucket

import (
	"strings"

	"github.com/drone/drone/model"
)

// convertRepo is a helper function used to convert a Bitbucket
// repository structure to the common Drone repository structure.
func convertRepo(from *Repo) *model.Repo {
	repo := model.Repo{
		Owner:     from.Owner.Login,
		Name:      from.Name,
		FullName:  from.FullName,
		Link:      from.Links.Html.Href,
		IsPrivate: from.IsPrivate,
		Avatar:    from.Owner.Links.Avatar.Href,
		Branch:    "master",
	}

	// in some cases, the owner of the repository is not
	// provided, however, we do have the full name.
	if len(repo.Owner) == 0 {
		repo.Owner = strings.Split(repo.FullName, "/")[0]
	}

	// above we manually constructed the repository clone url.
	// below we will iterate through the list of clone links and
	// attempt to instead use the clone url provided by bitbucket.
	for _, link := range from.Links.Clone {
		if link.Name == "https" {
			repo.Clone = link.Href
			break
		}
	}

	// if no repository name is provided, we use the Html link.
	// this excludes the .git suffix, but will still clone the repo.
	if len(repo.Clone) == 0 {
		repo.Clone = repo.Link
	}

	return &repo
}

// convertRepoLite is a helper function used to convert a Bitbucket
// repository structure to the simplified Drone repository structure.
func convertRepoLite(from *Repo) *model.RepoLite {
	return &model.RepoLite{
		Owner:    from.Owner.Login,
		Name:     from.Name,
		FullName: from.FullName,
		Avatar:   from.Owner.Links.Avatar.Href,
	}
}