package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

type Data struct {
	MasterChannelID string  `json:"master_channel_id"`
	VideoURL        string  `json:"video_url"`
	BadgeURL        string  `json:"badgeUrl"`
	AuthorType      string  `json:"authorType"`
	IsVerified      bool    `json:"isVerified"`
	IsChatOwner     bool    `json:"isChatOwner"`
	IsChatSponsor   bool    `json:"isChatSponsor"`
	IsChatModerator bool    `json:"isChatModerator"`
	ChannelID       string  `json:"channelId"`
	Name            string  `json:"name"`
	ImageURL        string  `json:"imageUrl"`
	Type            string  `json:"type"`
	SID             string  `json:"id"`
	Timestamp       int64   `gorm:"type:BIGINT" json:"timestamp"`
	ElapsedTime     string  `json:"elapsedTime"`
	Datetime        string  `json:"datetime"`
	Message         string  `gorm:"type:text" json:"message"`
	AmountValue     float64 `json:"amountValue"`
	AmountString    string  `json:"amountString"`
	Currency        string  `json:"currency"`
	BgColor         float64 `json:"bgColor"`
}

var BaseURL string = "https://needmov239087.df.r.appspot.com"
var TatiID string = "UCvUc0m317LWTTPZoBQV479A"

func decodingMsg(enc string) (string, error) {
	dec, err := base64.StdEncoding.DecodeString(enc)
	if err != nil {
		return "", err
	}
	return string(dec), nil
}
func myResp(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode != http.StatusOK {
		return resp, err
	}
	return resp, nil
}
func myUnmarshal(r *http.Response) ([]Data, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	var d []Data
	if err := json.Unmarshal(body, &d); err != nil {
		return nil, err
	}
	return d, nil
}
func reqCommeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "req",
		Short: "req culc",
		Args:  cobra.RangeArgs(2, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return nil
			}
			id, msg := args[0], args[1]
			url := BaseURL + "/api/comme/chvimsg_simi?chid=UCvUc0m317LWTTPZoBQV479A&id=" + id + "&msg=" + msg
			resp, err := myResp(url)
			if err != nil {
				return nil
			}
			defer resp.Body.Close()
			d, err := myUnmarshal(resp)
			if err != nil {
				return nil
			}
			for _, v := range d {
				v.Name, _ = decodingMsg(v.Name)
				v.Message, _ = decodingMsg(v.Message)
				fmt.Println(v.MasterChannelID, v.Name, "videoID = "+v.VideoURL, v.ElapsedTime, v.Message)
			}
			return nil
		},
	}
	return cmd
}

// z_4QQvZJJf4 草

func reqSuperChatCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "superchat",
		Short: "superchat culc",
		RunE: func(cmd *cobra.Command, args []string) error {
			url := BaseURL + "/api/comme/all_sc?chid=" + TatiID
			resp, err := myResp(url)
			if err != nil {
				return nil
			}
			defer resp.Body.Close()
			d, err := myUnmarshal(resp)
			if err != nil {
				return nil
			}
			for _, v := range d {
				v.Name, _ = decodingMsg(v.Name)
				v.Message, _ = decodingMsg(v.Message)
				fmt.Println(v.MasterChannelID, v.AmountString, v.Name, "videoID = "+v.VideoURL, v.ElapsedTime, v.Message)
			}
			return nil
		},
	}
	return cmd
}

// video_sc?chid=xxx&id=xxx
func reqSuperChatVideoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "superchat-video",
		Short: "superchat culc",
		Args:  cobra.RangeArgs(1, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			url := BaseURL + "/api/comme/video_sc?chid=" + TatiID + "&id=" + id
			resp, err := myResp(url)
			if err != nil {
				return nil
			}
			defer resp.Body.Close()
			d, err := myUnmarshal(resp)
			if err != nil {
				return nil
			}
			for _, v := range d {
				v.Name, _ = decodingMsg(v.Name)
				v.Message, _ = decodingMsg(v.Message)
				fmt.Println(v.MasterChannelID, v.AmountString, v.Name, "videoID = "+v.VideoURL, v.ElapsedTime, v.Message)
			}
			return nil
		},
	}
	return cmd
}
func myPri(d []Data, di Data, i interface{}) {
	for _, v := range d {
		v.Name, _ = decodingMsg(v.Name)
		v.Message, _ = decodingMsg(v.Message)
		fmt.Println(v.MasterChannelID, v.AmountString, v.Name, "videoID = "+v.VideoURL, v.ElapsedTime, v.Message)
	}
	fmt.Println(di.AmountString)
}
