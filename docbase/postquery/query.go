package postquery

import (
	"fmt"
	"strings"
)

func Join(queries ...string) string {
	return strings.Join(queries, " ")
}

// Sort specifies sort order to get posts.
func Sort(property SortName, asc bool) string {
	prefix := "desc:"
	if asc {
		prefix = "asc:"
	}
	return prefix + property.String()
}

// SortName specifies a property to sort posts.
type SortName string

// Concrete properties to sort posts.
const (
	SortNameScore     = SortName("score")
	SortNameChangedAt = SortName("changed_at")
	SortNameCreatedAt = SortName("created_at")
	SortNameStars     = SortName("stars")
	SortNameComments  = SortName("comments")
	SortNameLikes     = SortName("likes")
)

func (s SortName) String() string {
	return string(s)
}

// Property specifies a property to search the keyword from posts
func Property(property PropertyName, keyword string) string {
	if strings.ContainsRune(keyword, ' ') {
		keyword = "\"" + strings.Trim(keyword, "\"") + "\""
	}
	return property.String() + ":" + keyword
}

// PropertyName specifies a property to search the keyword from posts
type PropertyName string

// Concrete properties to search the keyword from posts
const (
	PropertyNameTitle       = PropertyName("title")
	PropertyNameBody        = PropertyName("body")
	PropertyNameComments    = PropertyName("comments")
	PropertyNameAttachments = PropertyName("attachments")
	PropertyNameAuthor      = PropertyName("author")
	PropertyNameCommentedBy = PropertyName("commented_by")
	PropertyNameLikedBy     = PropertyName("liked_by")
	PropertyNameTag         = PropertyName("tag")
	PropertyNameGroup       = PropertyName("group")
)

func (s PropertyName) String() string {
	return string(s)
}

// DateName specifies a property to search the keyword from posts
type DateName string

// Concrete properties to search the keyword from posts
const (
	DateNameCreatedAt = DateName("created_at")
	DateNameChangedAt = DateName("changed_at")
)

func (s DateName) String() string {
	return string(s)
}

// DateTo specifies the end date with property name to search posts
func DateTo(property DateName, year, month, day int) string {
	return fmt.Sprintf("%s:*~%02d-%02d-%02d", property, year, month, day)
}

// DateFrom specifies the start date with property name to search posts
func DateFrom(property DateName, year, month, day int) string {
	return fmt.Sprintf("%s:%02d-%02d-%02d~*", property, year, month, day)
}

// Date specifies the date with property name to search posts
func Date(property DateName, year, month, day int) string {
	return fmt.Sprintf("%s:%02d-%02d-%02d", property, year, month, day)
}

// DateRange specifies the date range with property name to search posts
func DateRange(property DateName, year1, month1, day1, year2, month2, day2 int) string {
	return fmt.Sprintf("%s:%02d-%02d-%02d~%2d-%02d-%02d", property, year1, month1, day1, year2, month2, day2)
}

// Missing specifies a property to search posts which is not filled
func Missing(property MissingName) string {
	return "missing:" + property.String()
}

// MissingName specifies a property to search the keyword from posts
type MissingName string

// Concrete properties to search the keyword from posts
const (
	MissingNameTag = MissingName("tag")
)

func (s MissingName) String() string {
	return string(s)
}

// Or :
func Or() string {
	return "OR"
}

// HasStar :
func HasStar() string {
	return "has:star"
}

// IsDraft :
func IsDraft() string {
	return "is:draft"
}

// IsUnread :
func IsUnread() string {
	return "is:unread"
}

// IsShared :
func IsShared() string {
	return "is:shared"
}
