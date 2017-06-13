package p355

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	obj := Constructor()
	obj.PostTweet(1, 1)
	obj.PostTweet(1, 2)
	obj.PostTweet(1, 3)
	obj.PostTweet(1, 4)
	obj.PostTweet(1, 5)
	obj.PostTweet(1, 6)
	obj.PostTweet(1, 7)
	obj.PostTweet(1, 8)
	obj.PostTweet(1, 9)
	obj.PostTweet(1, 10)
	assert.Equal(t, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, obj.GetNewsFeed(1))
	obj.Follow(2, 1)
	obj.Unfollow(2, 3)
	obj.Unfollow(2, 2)
}
