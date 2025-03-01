package helmutil

type RepoEntry interface {
	// Name returns repo name.
	// Example: "my-charts".
	Name() string

	// URL returns repo URL.
	// Examples:
	// - https://kubernetes-charts.storage.googleapis.com/
	// - s3://my-charts
	URL() string

	// IndexURL returns repo index file URL.
	// Examples:
	// - https://kubernetes-charts.storage.googleapis.com/index.yaml
	// - s3://my-charts/index.yaml
	IndexURL() string

	// CacheFile returns repo local cache file path.
	// Examples:
	// - /Users/foo/Library/Caches/helm/repository/my-charts-index.yaml (on macOS)
	// - /home/foo/.cache/helm/repository/my-charts-index.yaml (on Linux)
	CacheFile() string
}

// LookupRepoEntry returns an entry from helm's repositories.yaml file by name.
func LookupRepoEntry(name string) (RepoEntry, error) {
	if IsHelm3() {
		return lookupV3(name)
	}
	return lookupV2(name)
}

// LookupRepoEntryByURL returns an entry from helm's repositories.yaml file by
// repo URL. If not found, returns false and <nil> error.
func LookupRepoEntryByURL(url string) (RepoEntry, bool, error) {
	if IsHelm3() {
		return lookupByURLV3(url)
	}
	return lookupByURLV2(url)
}
