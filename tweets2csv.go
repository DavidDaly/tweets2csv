package main

import "strconv"
import "github.com/ChimeraCoder/anaconda"
import "encoding/csv"
import "os"
import "net/url"

func main() {

	var lastId string

	anaconda.SetConsumerKey("Consumer Key")
	anaconda.SetConsumerSecret("Consumer Secret")
	api := anaconda.NewTwitterApi("Access Token", "Access Token Secret")

	var records = [][]string{ 
				{"Date and Time",
				"Tweet Text",
				"User",
				"Number of Followers", 
				"Is Retweet",
				"Number of Retweets",
				"Number of Likes"} }	

	for i := 0; i < 20; i++ {
		
		v := url.Values{}
		v.Set("count", "100")
		v.Set("until", "2016-09-02")
		if i > 0 {
			v.Set("max_id", lastId)
		}
		
		searchResponse, err := api.GetSearch("#DevOpsChat", v)

		if err != nil {
    			panic(err)
		}
		
		for _, tweet := range searchResponse.Statuses {
    	
			if ( tweet.IdStr != lastId ) {
				records = append(records, []string{
						tweet.CreatedAt,
						tweet.Text,
						tweet.User.ScreenName,
						strconv.Itoa(tweet.User.FollowersCount),
						strconv.FormatBool(tweet.RetweetedStatus != nil),
						strconv.Itoa(tweet.RetweetCount),
						strconv.Itoa(tweet.FavoriteCount) } )
			}

			lastId = tweet.IdStr
		}

		
	}

	w := csv.NewWriter(os.Stdout)
	w.WriteAll(records)

}
