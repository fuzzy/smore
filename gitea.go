package main

type GiteaPush struct {
	After   string `json:"after"`
	Before  string `json:"before"`
	Commits []struct {
		Added  interface{} `json:"added"`
		Author struct {
			Email    string `json:"email"`
			Name     string `json:"name"`
			Username string `json:"username"`
		} `json:"author"`
		Committer struct {
			Email    string `json:"email"`
			Name     string `json:"name"`
			Username string `json:"username"`
		} `json:"committer"`
		ID           string      `json:"id"`
		Message      string      `json:"message"`
		Modified     interface{} `json:"modified"`
		Removed      interface{} `json:"removed"`
		Timestamp    string      `json:"timestamp"`
		URL          string      `json:"url"`
		Verification interface{} `json:"verification"`
	} `json:"commits"`
	CompareURL string      `json:"compare_url"`
	HeadCommit interface{} `json:"head_commit"`
	Pusher     struct {
		AvatarURL string `json:"avatar_url"`
		Created   string `json:"created"`
		Email     string `json:"email"`
		FullName  string `json:"full_name"`
		ID        int    `json:"id"`
		IsAdmin   bool   `json:"is_admin"`
		Language  string `json:"language"`
		LastLogin string `json:"last_login"`
		Login     string `json:"login"`
		Username  string `json:"username"`
	} `json:"pusher"`
	Ref        string `json:"ref"`
	Repository struct {
		AllowMergeCommits         bool   `json:"allow_merge_commits"`
		AllowRebase               bool   `json:"allow_rebase"`
		AllowRebaseExplicit       bool   `json:"allow_rebase_explicit"`
		AllowSquashMerge          bool   `json:"allow_squash_merge"`
		Archived                  bool   `json:"archived"`
		AvatarURL                 string `json:"avatar_url"`
		CloneURL                  string `json:"clone_url"`
		CreatedAt                 string `json:"created_at"`
		DefaultBranch             string `json:"default_branch"`
		Description               string `json:"description"`
		Empty                     bool   `json:"empty"`
		Fork                      bool   `json:"fork"`
		ForksCount                int    `json:"forks_count"`
		FullName                  string `json:"full_name"`
		HasIssues                 bool   `json:"has_issues"`
		HasPullRequests           bool   `json:"has_pull_requests"`
		HasWiki                   bool   `json:"has_wiki"`
		HTMLURL                   string `json:"html_url"`
		ID                        int    `json:"id"`
		IgnoreWhitespaceConflicts bool   `json:"ignore_whitespace_conflicts"`
		Mirror                    bool   `json:"mirror"`
		Name                      string `json:"name"`
		OpenIssuesCount           int    `json:"open_issues_count"`
		OriginalURL               string `json:"original_url"`
		Owner                     struct {
			AvatarURL string `json:"avatar_url"`
			Created   string `json:"created"`
			Email     string `json:"email"`
			FullName  string `json:"full_name"`
			ID        int    `json:"id"`
			IsAdmin   bool   `json:"is_admin"`
			Language  string `json:"language"`
			LastLogin string `json:"last_login"`
			Login     string `json:"login"`
			Username  string `json:"username"`
		} `json:"owner"`
		Parent      interface{} `json:"parent"`
		Permissions struct {
			Admin bool `json:"admin"`
			Pull  bool `json:"pull"`
			Push  bool `json:"push"`
		} `json:"permissions"`
		Private       bool   `json:"private"`
		Size          int    `json:"size"`
		SSHURL        string `json:"ssh_url"`
		StarsCount    int    `json:"stars_count"`
		UpdatedAt     string `json:"updated_at"`
		WatchersCount int    `json:"watchers_count"`
		Website       string `json:"website"`
	} `json:"repository"`
	Secret string `json:"secret"`
	Sender struct {
		AvatarURL string `json:"avatar_url"`
		Created   string `json:"created"`
		Email     string `json:"email"`
		FullName  string `json:"full_name"`
		ID        int    `json:"id"`
		IsAdmin   bool   `json:"is_admin"`
		Language  string `json:"language"`
		LastLogin string `json:"last_login"`
		Login     string `json:"login"`
		Username  string `json:"username"`
	} `json:"sender"`
}
