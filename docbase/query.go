package docbase

import (
	"fmt"
	"strings"
)

// PostQuerySort specifies sort order to get posts.
func PostQuerySort(property PostQuerySortName, asc bool) string {
	prefix := "desc:"
	if asc {
		prefix = "asc:"
	}
	return prefix + property.String()
}

// PostQuerySortName specifies a property to sort posts.
type PostQuerySortName string

// Concrete properties to sort posts.
const (
	PostQuerySortNameScore     = PostQuerySortName("score")
	PostQuerySortNameChangedAt = PostQuerySortName("changed_at")
	PostQuerySortNameCreatedAt = PostQuerySortName("created_at")
	PostQuerySortNameStars     = PostQuerySortName("stars")
	PostQuerySortNameComments  = PostQuerySortName("comments")
	PostQuerySortNameLikes     = PostQuerySortName("likes")
)

func (s PostQuerySortName) String() string {
	return string(s)
}

// PostQueryProperty specifies a property to search the keyword from posts
func PostQueryProperty(property PostQueryPropertyName, keyword string) string {
	if strings.ContainsRune(keyword, ' ') {
		keyword = "\"" + strings.Trim(keyword, "\"") + "\""
	}
	return property.String() + ":" + keyword
}

// PostQueryPropertyName specifies a property to search the keyword from posts
type PostQueryPropertyName string

// Concrete properties to search the keyword from posts
const (
	PostQueryPropertyNameTitle       = PostQueryPropertyName("title")
	PostQueryPropertyNameBody        = PostQueryPropertyName("body")
	PostQueryPropertyNameComments    = PostQueryPropertyName("comments")
	PostQueryPropertyNameAttachments = PostQueryPropertyName("attachments")
	PostQueryPropertyNameAuthor      = PostQueryPropertyName("author")
	PostQueryPropertyNameCommentedBy = PostQueryPropertyName("commented_by")
	PostQueryPropertyNameLikedBy     = PostQueryPropertyName("liked_by")
	PostQueryPropertyNameTag         = PostQueryPropertyName("tag")
	PostQueryPropertyNameGroup       = PostQueryPropertyName("group")
)

func (s PostQueryPropertyName) String() string {
	return string(s)
}

// PostQueryDateName specifies a property to search the keyword from posts
type PostQueryDateName string

// Concrete properties to search the keyword from posts
const (
	PostQueryDateNameCreatedAt = PostQueryDateName("created_at")
	PostQueryDateNameChangedAt = PostQueryDateName("changed_at")
)

func (s PostQueryDateName) String() string {
	return string(s)
}

// PostQueryDateTo specifies the end date with property name to search posts
func PostQueryDateTo(property PostQueryDateName, year, month, day int) string {
	return fmt.Sprintf("%s:*~%02d-%02d-%02d", property, year, month, day)
}

// PostQueryDateFrom specifies the start date with property name to search posts
func PostQueryDateFrom(property PostQueryDateName, year, month, day int) string {
	return fmt.Sprintf("%s:%02d-%02d-%02d~*", property, year, month, day)
}

// PostQueryDate specifies the date with property name to search posts
func PostQueryDate(property PostQueryDateName, year, month, day int) string {
	return fmt.Sprintf("%s:%02d-%02d-%02d", property, year, month, day)
}

// PostQueryDateRange specifies the date range with property name to search posts
func PostQueryDateRange(property PostQueryDateName, year1, month1, day1, year2, month2, day2 int) string {
	return fmt.Sprintf("%s:%02d-%02d-%02d~%%2d-%02d-%02d", property, year1, month1, day1, year2, month2, day2)
}

// PostQueryMissing specifies a property to search posts which is not filled
func PostQueryMissing(property PostQueryMissingName) string {
	return "missing:" + property.String()
}

// PostQueryMissingName specifies a property to search the keyword from posts
type PostQueryMissingName string

// Concrete properties to search the keyword from posts
const (
	PostQueryMissingNameTag = PostQueryMissingName("tag")
)

func (s PostQueryMissingName) String() string {
	return string(s)
}

// PostQueryOr :
func PostQueryOr() string {
	return "OR"
}

// PostQueryHasStar :
func PostQueryHasStar() string {
	return "has:star"
}

// PostQueryIsDraft :
func PostQueryIsDraft() string {
	return "is:draft"
}

// PostQueryIsUnread :
func PostQueryIsUnread() string {
	return "is:unread"
}

// PostQueryIsShared :
func PostQueryIsShared() string {
	return "is:shared"
}
