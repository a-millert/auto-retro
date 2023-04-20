package core

type Interaction int

// `IssueState` from the `Contributor`'s perspective.
const (
	OPEN Interaction = iota // also commented
	APPROVED
	MERGED
	TO_REVIEW
	SELF_COMMENTED
	SELF_APPROVED
)

type Issue struct {
	URL   string
	Title string
	Event Interaction
}

type Contribution struct {
	Open          []Issue
	Approved      []Issue
	Merged        []Issue
	ToReview      []Issue
	SelfCommented []Issue
	SelfApproved  []Issue
}

type Contributor struct {
	User
	Contribution
}
