package p355

/**
Design a simplified version of Twitter where users can post tweets,
follow/unfollow another user and is able to see the 10 most recent
tweets in the user's news feed. Your design should support the following methods:

postTweet(userId, tweetId): Compose a new tweet.
getNewsFeed(userId): Retrieve the 10 most recent tweet ids in the user's news feed.
Each item in the news feed must be posted by users who the user followed or by the user herself.
Tweets must be ordered from most recent to least recent.
follow(followerId, followeeId): Follower follows a followee.
unfollow(followerId, followeeId): Follower unfollows a followee.
Example:

Twitter twitter = new Twitter();

// User 1 posts a new tweet (id = 5).
twitter.postTweet(1, 5);

// User 1's news feed should return a list with 1 tweet id -> [5].
twitter.getNewsFeed(1);

// User 1 follows user 2.
twitter.follow(1, 2);

// User 2 posts a new tweet (id = 6).
twitter.postTweet(2, 6);

// User 1's news feed should return a list with 2 tweet ids -> [6, 5].
// Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
twitter.getNewsFeed(1);

// User 1 unfollows user 2.
twitter.unfollow(1, 2);

// User 1's news feed should return a list with 1 tweet id -> [5],
// since user 1 is no longer following user 2.
twitter.getNewsFeed(1);
*/

type Tweet struct {
	uid int
	id  int
}

type Twitter struct {
	follows map[int]map[int]bool
	tweets  []Tweet
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		follows: make(map[int]map[int]bool),
		tweets:  make([]Tweet, 0),
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.tweets = append(this.tweets, Tweet{userId, tweetId})
}

/** Retrieve the 10 most recent tweet ids in the user's news feed.
Each item in the news feed must be posted by users who the user followed or by the user herself.
Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	res := make([]int, 0, 10)
	follows := this.follows[userId]
	if follows == nil {
		follows = make(map[int]bool)
	}
	follows[userId] = true
	for i := len(this.tweets) - 1; i >= 0; i-- {
		t := this.tweets[i]
		if follows[t.uid] {
			res = append(res, t.id)
		}
		if len(res) >= 10 {
			break
		}
	}
	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.follows[followerId] == nil {
		this.follows[followerId] = make(map[int]bool, 0)
	}
	this.follows[followerId][followeeId] = true
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if this.follows[followerId] != nil && followerId != followeeId {
		delete(this.follows[followerId], followeeId)
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
