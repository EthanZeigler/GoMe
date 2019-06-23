package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Groups []Group

type Group struct {
	// Group Unique ID
	ID string `json:"id"`
	// Group name
	Name string `json:"name"`
	// Group type
	// Private or public (documentation unclear)?
	Type string `json:"type"`
	// Group members defined description
	Description string `json:"description"`
	// Group icon URL
	ImageURL string `json:"image_url"`
	// User ID of the group creator
	CreatorUserID string `json:"creator_user_id"`
	// UNIX time (s) the group was made at
	CreatedAt int64 `json:"created_at"`
	// UNIX time (s) the group last changed at
	UpdatedAt int64 `json:"updated_at"`
	// Group Members
	Members []Member `json:"members"`
	// URL that can be given to other users
	// to join this group
	ShareURL string `json:"share_url"`
	// Generic data about the messages of this chat.
	// Does not provide specifics, however provides the ID needed
	// to do so.
	Messages Messages `json:"messages"`
}

type Member struct {
	// Unique user ID
	UserID string `json:"user_id"`
	// Display name of the member within the relevant group
	Nickname string `json:"nickname"`
	// Whether the member receives notifications for the relevent group
	// Mentioning the member will override this
	Muted bool `json:"muted"`
	// User's image icon
	ImageURL string `json:"image_url"`
}

type Messages struct {
	// Number of messages within the relevent group
	Count int64 `json:"count"`
	// Most recent message's ID
	LastMessageID string `json:"last_message_id"`
	// UNIX time (s) of the last message within the relevent group
	LastMessageCreatedAt int64 `json:"last_message_created_at"`
	// Preview messafe
	Preview MessagePreview `json:"preview"`
}

type MessagePreview struct {
	Nickname    string       `json:"nickname"`
	Text        string       `json:"text"`
	ImageURL    string       `json:"image_url"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Type        string    `json:"type"`
	URL         *string   `json:"url,omitempty"`
	Lat         *string   `json:"lat,omitempty"`
	Lng         *string   `json:"lng,omitempty"`
	Name        *string   `json:"name,omitempty"`
	Token       *string   `json:"token,omitempty"`
	Placeholder *string   `json:"placeholder,omitempty"`
	Charmap     [][]int64 `json:"charmap"`
}

//  #############################################
//  #                                           #
//  #       Restful Adapter Functionality       #
//  #                                           #
//  #############################################

func UserGroups(
	apiKey          string,
	page            int64,
	perPage         int64,
	omitMemberships bool) (response Groups, err error) {
	// convert from the omitMembership bool to what the API expects
	// Man why on earth did they use a string for a bool value?
	var omitStr string
	if omitMemberships {
		omitStr = "memberships"
	} else {
		omitStr = ""
	}

	request, _ := http.NewRequest(
		"GET",
		"https://api.groupme.com/v3/groups",
		nil,
	)

	// Add in parameters for the request
	query := request.URL.Query()
	query.Add("token", apiKey)
	query.Add("page", string(page))
	query.Add("per_page", string(perPage))
	query.Add("omit", omitStr)

	// Build final query string, make request
	request.URL.RawQuery = query.Encode()
	rawResponse, err := httpClient.Do(request)
	if err != nil {
		// request to the groupme API failed
		return
	}

	// Decode the response
	jsonResponse, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		// response data is invalid
		return
	}

	// Unmarshal expected data
	// On failure, see if it's a groupme error response
	if err = json.Unmarshal(jsonResponse, &response); err != nil {
		// maybe it's a error response?
		var errResponse GMError
		if err = json.Unmarshal(jsonResponse, &errResponse); err != nil {
			// okay its just bad data
			err = GMError{}
		}
	}
	// no return necessary because of predefined return values
}
