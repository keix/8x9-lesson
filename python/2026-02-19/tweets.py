import json
from datetime import datetime, timezone

TWEETS_FILE = "tweets.json"


# ---------- File I/O ----------
def load_tweets():
    """Load tweets from JSON file"""
    with open(TWEETS_FILE) as f:
        return json.load(f)


def save_tweets(tweets):
    """Save tweets to JSON file"""
    with open(TWEETS_FILE, "w") as f:
        json.dump(tweets, f, indent=2, ensure_ascii=False)


# ---------- Search ----------
def find_tweet_by_id(tweets, tweet_id):
    """Find a tweet by ID"""
    for tweet in tweets:
        if tweet["id"] == tweet_id:
            return tweet
    return None


def find_tweets_by_account(tweets, account_id):
    """Find all tweets by account ID"""
    return [tweet for tweet in tweets if tweet["account_id"] == account_id]


# ---------- API Functions ----------
def add_tweet(account_id, content):
    """Add a new tweet"""
    tweets = load_tweets()

    if tweets:
        new_id = max(tweet["id"] for tweet in tweets) + 1
    else:
        new_id = 1

    new_tweet = {
        "id": new_id,
        "account_id": account_id,
        "content": content,
        "created_at": datetime.now(timezone.utc).isoformat(),
    }

    tweets.append(new_tweet)
    save_tweets(tweets)

    return new_tweet


def delete_tweet(tweet_id):
    """Delete a tweet"""
    tweets = load_tweets()
    tweet = find_tweet_by_id(tweets, tweet_id)

    if tweet is None:
        return {"ERROR": "tweet not found"}

    tweets.remove(tweet)
    save_tweets(tweets)

    return tweet


def list_tweets():
    """List all tweets"""
    return load_tweets()


def list_tweets_by_account(account_id):
    """List all tweets by account"""
    tweets = load_tweets()
    return find_tweets_by_account(tweets, account_id)
