package models

type Insta struct {
	Config                Config       `json:"config"`
	CountryCode           string       `json:"country_code"`
	LanguageCode          string       `json:"language_code"`
	Locale                string       `json:"locale"`
	EntryData             EntryData    `json:"entry_data"`
	Hostname              string       `json:"hostname"`
	IsWhitelistedCrawlBot bool         `json:"is_whitelisted_crawl_bot"` // (!)
	DeploymentStage       string       `json:"deployment_stage"`
	Platform              string       `json:"platform"`
	Nonce                 string       `json:"nonce"`
	MidPct                float64      `json:"mid_pct"`
	ZeroData              *interface{} `json:"zero_data"`
	CacheSchemaVersion    int64        `json:"cache_schema_version"`
	ServerChecks          *interface{} `json:"server_checks"`
	Knobx                 *interface{} `json:"knobx"`
	ToCache               ToCache      `json:"to_cache"`
	DeviceID              string       `json:"device_id"` // (!)
	Encryption            Encryption   `json:"encryption"`
	RolloutHash           string       `json:"rollout_hash"`
	BundleVariant         string       `json:"bundle_variant"`
	IsCanary              bool         `json:"is_canary"`
}

type EntryData struct {
	ProfilePage []ProfilePage `json:"ProfilePage"`
	TagPage     []GraphqlPage `json:"TagPage"`
	PostPage    []GraphqlPage `json:"PostPage"`
}

type GraphqlPage struct {
	Graphql Graphql `json:"graphql"`
}

type Graphql struct {
	User           User           `json:"user"`
	Hashtag        Hashtag        `json:"hashtag"`
	ShortcodeMedia ShortcodeMedia `json:"shortcode_media"`
}

type ProfilePage struct {
	LoggingPageID         string      `json:"logging_page_id"`
	ShowSuggestedProfiles bool        `json:"show_suggested_profiles"`
	ShowFollowDialog      bool        `json:"show_follow_dialog"`
	Graphql               Graphql     `json:"graphql"`
	ToastContentOnLoad    interface{} `json:"toast_content_on_load"`
}

// ===================
type InstaNext struct {
	Data   Data   `json:"data"`
	Status string `json:"status"`
}
type Data struct {
	User    User    `json:"user"`
	Hashtag Hashtag `json:"hashtag"`
}
type Config struct {
	CSRFToken string  `json:"csrf_token"`
	Viewer    *string `json:"viewer"`
	ViewerID  *string `json:"viewerId"`
}

type Encryption struct {
	KeyID     string `json:"key_id"`
	PublicKey string `json:"public_key"`
}

type User struct {
	Biography                string                      `json:"biography"`
	BlockedByViewer          bool                        `json:"blocked_by_viewer"`
	CountryBlock             bool                        `json:"country_block"`
	ExternalURL              string                      `json:"external_url"`
	ExternalURLLinkshimmed   string                      `json:"external_url_linkshimmed"`
	EdgeFollowedBy           EdgeFollowClass             `json:"edge_followed_by"`
	FollowedByViewer         bool                        `json:"followed_by_viewer"`
	EdgeFollow               EdgeFollowClass             `json:"edge_follow"`
	FollowsViewer            bool                        `json:"follows_viewer"`
	FullName                 string                      `json:"full_name"`
	HasChannel               bool                        `json:"has_channel"`
	HasBlockedViewer         bool                        `json:"has_blocked_viewer"`
	HighlightReelCount       int64                       `json:"highlight_reel_count"`
	HasRequestedViewer       bool                        `json:"has_requested_viewer"`
	ID                       string                      `json:"id"`
	IsBusinessAccount        bool                        `json:"is_business_account"`
	IsJoinedRecently         bool                        `json:"is_joined_recently"`
	BusinessCategoryName     interface{}                 `json:"business_category_name"`
	IsPrivate                bool                        `json:"is_private"`
	IsVerified               bool                        `json:"is_verified"`
	EdgeMutualFollowedBy     EdgeMutualFollowedBy        `json:"edge_mutual_followed_by"`
	ProfilePicURL            string                      `json:"profile_pic_url"`
	ProfilePicURLHD          string                      `json:"profile_pic_url_hd"`
	RequestedByViewer        bool                        `json:"requested_by_viewer"`
	Username                 string                      `json:"username"`
	ConnectedFbPage          interface{}                 `json:"connected_fb_page"`
	EdgeFelixVideoTimeline   EdgeFelixVideoTimelineClass `json:"edge_felix_video_timeline"`
	EdgeOwnerToTimelineMedia EdgeFelixVideoTimelineClass `json:"edge_owner_to_timeline_media"`
	EdgeSavedMedia           EdgeFelixVideoTimelineClass `json:"edge_saved_media"`
	EdgeMediaCollections     EdgeFelixVideoTimelineClass `json:"edge_media_collections"`
}

type Hashtag struct {
	ID                           string                       `json:"id"`
	Name                         string                       `json:"name"`
	AllowFollowing               bool                         `json:"allow_following"`
	IsFollowing                  bool                         `json:"is_following"`
	IsTopMediaOnly               bool                         `json:"is_top_media_only"`
	ProfilePicURL                string                       `json:"profile_pic_url"`
	EdgeHashtagToMedia           EdgeHashtagToMedia           `json:"edge_hashtag_to_media"`
	EdgeHashtagToTopPosts        EdgeHashtagToTopPosts        `json:"edge_hashtag_to_top_posts"`
	EdgeHashtagToContentAdvisory EdgeHashtagToContentAdvisory `json:"edge_hashtag_to_content_advisory"`
	EdgeHashtagToRelatedTags     EdgeHashtagToRelatedTags     `json:"edge_hashtag_to_related_tags"`
	EdgeHashtagToNullState       EdgeHashtagToNullStateClass  `json:"edge_hashtag_to_null_state"`
}

type ShortcodeMedia struct {
	Typename                    string                        `json:"__typename"`
	ID                          string                        `json:"id"`
	Shortcode                   string                        `json:"shortcode"`
	Dimensions                  Dimensions                    `json:"dimensions"`
	GatingInfo                  interface{}                   `json:"gating_info"`
	FactCheckOverallRating      interface{}                   `json:"fact_check_overall_rating"`
	FactCheckInformation        interface{}                   `json:"fact_check_information"`
	MediaPreview                interface{}                   `json:"media_preview"`
	DisplayURL                  string                        `json:"display_url"`
	DisplayResources            []DisplayResource             `json:"display_resources"`
	IsVideo                     bool                          `json:"is_video"`
	TrackingToken               string                        `json:"tracking_token"`
	EdgeMediaToTaggedUser       EdgeMediaToTaggedUser         `json:"edge_media_to_tagged_user"`
	EdgeMediaToCaption          EdgeMediaToCaptionClass       `json:"edge_media_to_caption"`
	CaptionIsEdited             bool                          `json:"caption_is_edited"`
	HasRankedComments           bool                          `json:"has_ranked_comments"`
	EdgeMediaToParentComment    EdgeMediaToParentCommentClass `json:"edge_media_to_parent_comment"`
	EdgeMediaToHoistedComment   EdgeMediaToCaptionClass       `json:"edge_media_to_hoisted_comment"`
	EdgeMediaPreviewComment     EdgeMediaPreview              `json:"edge_media_preview_comment"`
	CommentsDisabled            bool                          `json:"comments_disabled"`
	CommentingDisabledForViewer bool                          `json:"commenting_disabled_for_viewer"`
	TakenAtTimestamp            int64                         `json:"taken_at_timestamp"`
	EdgeMediaPreviewLike        EdgeMediaPreview              `json:"edge_media_preview_like"`
	EdgeMediaToSponsorUser      EdgeMediaToCaptionClass       `json:"edge_media_to_sponsor_user"`
	Location                    *Location                     `json:"location"`
	ViewerHasLiked              bool                          `json:"viewer_has_liked"`
	ViewerHasSaved              bool                          `json:"viewer_has_saved"`
	ViewerHasSavedToCollection  bool                          `json:"viewer_has_saved_to_collection"`
	ViewerInPhotoOfYou          bool                          `json:"viewer_in_photo_of_you"`
	ViewerCanReshare            bool                          `json:"viewer_can_reshare"`
	Owner                       ShortcodeMediaOwner           `json:"owner"`
	IsAd                        bool                          `json:"is_ad"`
	EdgeWebMediaToRelatedMedia  EdgeMediaToCaptionClass       `json:"edge_web_media_to_related_media"`
	EdgeSidecarToChildren       EdgeSidecarToChildren         `json:"edge_sidecar_to_children"`
}

type DisplayResource struct {
	Src          string `json:"src"`
	ConfigWidth  int64  `json:"config_width"`
	ConfigHeight int64  `json:"config_height"`
}

type EdgeMediaPreview struct {
	Count int64                         `json:"count"`
	Edges []EdgeMediaPreviewCommentEdge `json:"edges"`
}

type EdgeHashtagToContentAdvisory struct {
	Count int64         `json:"count"`
	Edges []interface{} `json:"edges"`
}

type EdgeHashtagToMedia struct {
	Count    int64                    `json:"count"`
	PageInfo PageInfo                 `json:"page_info"`
	Edges    []EdgeHashtagToMediaEdge `json:"edges"`
}

type EdgeHashtagToMediaEdge struct {
	Node PurpleNode `json:"node"`
}
type EdgeHashtagToTopPosts struct {
	Edges []EdgeHashtagToMediaEdge `json:"edges"`
}
type EdgeHashtagToRelatedTags struct {
	Edges []EdgeHashtagToRelatedTagsEdge `json:"edges"`
}

type EdgeHashtagToRelatedTagsEdge struct {
	Node TentacledNode `json:"node"`
}

type TentacledNode struct {
	Name string `json:"name"`
}
type EdgeHashtagToNullStateClass struct {
	Edges []EdgeHashtagToNullStateEdge `json:"edges"`
}
type EdgeHashtagToNullStateEdge struct {
	Node FluffyNode `json:"node"`
}

type EdgeFelixVideoTimelineClass struct {
	Count    int64                        `json:"count"`
	PageInfo PageInfo                     `json:"page_info"`
	Edges    []EdgeFelixVideoTimelineEdge `json:"edges"`
}

type EdgeFelixVideoTimelineEdge struct {
	Node PurpleNode `json:"node"`
}

type PurpleNode struct {
	Typename               string                         `json:"__typename"`
	ID                     string                         `json:"id"`
	EdgeMediaToCaption     EdgeMediaToCaption             `json:"edge_media_to_caption"`
	Shortcode              string                         `json:"shortcode"`
	EdgeMediaToComment     EdgeFollowClass                `json:"edge_media_to_comment"`
	CommentsDisabled       bool                           `json:"comments_disabled"`
	TakenAtTimestamp       int64                          `json:"taken_at_timestamp"`
	Dimensions             Dimensions                     `json:"dimensions"`
	DisplayURL             string                         `json:"display_url"`
	EdgeLikedBy            EdgeFollowClass                `json:"edge_liked_by"`
	EdgeMediaPreviewLike   EdgeFollowClass                `json:"edge_media_preview_like"`
	Location               *Location                      `json:"location"`
	GatingInfo             interface{}                    `json:"gating_info"`
	FactCheckOverallRating interface{}                    `json:"fact_check_overall_rating"`
	FactCheckInformation   interface{}                    `json:"fact_check_information"`
	MediaPreview           string                         `json:"media_preview"`
	Owner                  Owner                          `json:"owner"`
	ThumbnailSrc           string                         `json:"thumbnail_src"`
	ThumbnailResources     []ThumbnailResource            `json:"thumbnail_resources"`
	IsVideo                bool                           `json:"is_video"`
	AccessibilityCaption   string                         `json:"accessibility_caption"`
	Text                   string                         `json:"text"`
	CreatedAt              int64                          `json:"created_at"`
	DidReportAsSpam        bool                           `json:"did_report_as_spam"`
	ViewerHasLiked         bool                           `json:"viewer_has_liked"`
	IsRestrictedPending    bool                           `json:"is_restricted_pending"`
	EdgeThreadedComments   *EdgeMediaToParentCommentClass `json:"edge_threaded_comments,omitempty"`
}

type Dimensions struct {
	Height int64 `json:"height"`
	Width  int64 `json:"width"`
}

type EdgeFollowClass struct {
	Count int64 `json:"count"`
}

type EdgeMediaToCaption struct {
	Edges []EdgeMediaToCaptionEdge `json:"edges"`
}

type EdgeMediaToCaptionEdge struct {
	Node FluffyNode `json:"node"`
}

type FluffyNode struct {
	Text string `json:"text"`
}

type Location struct {
	ID            string `json:"id"`
	HasPublicPage bool   `json:"has_public_page"`
	Name          string `json:"name"`
	Slug          string `json:"slug"`
}

type Owner struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type ThumbnailResource struct {
	Src          string `json:"src"`
	ConfigWidth  int64  `json:"config_width"`
	ConfigHeight int64  `json:"config_height"`
}

type PageInfo struct {
	HasNextPage bool    `json:"has_next_page"`
	EndCursor   *string `json:"end_cursor"`
}

type EdgeMutualFollowedBy struct {
	Count int64         `json:"count"`
	Edges []interface{} `json:"edges"`
}

type ToCache struct {
	Gatekeepers    map[string]bool `json:"gatekeepers"`
	Qe             *interface{}    `json:"qe"`
	ProbablyHasApp bool            `json:"probably_has_app"`
	Cb             bool            `json:"cb"`
}

type EdgeMediaToParentCommentClass struct {
	Count    int64                         `json:"count"`
	PageInfo PageInfo                      `json:"page_info"`
	Edges    []EdgeMediaPreviewCommentEdge `json:"edges"`
}

type EdgeMediaPreviewCommentEdge struct {
	Node PurpleNode `json:"node"`
}

type EdgeOwnerToTimelineMediaClass struct {
	Count int64 `json:"count"`
}

type UserClass struct {
	ID            string  `json:"id"`
	IsVerified    bool    `json:"is_verified"`
	ProfilePicURL string  `json:"profile_pic_url"`
	Username      string  `json:"username"`
	FullName      *string `json:"full_name,omitempty"`
}

type EdgeMediaToCaptionClass struct {
	Edges []EdgeMediaToCaptionEdge `json:"edges"`
}

type EdgeMediaToTaggedUser struct {
	Edges []EdgeMediaToTaggedUserEdge `json:"edges"`
}

type EdgeMediaToTaggedUserEdge struct {
	Node TentacledNode `json:"node"`
}

type EdgeSidecarToChildren struct {
	Edges []EdgeSidecarToChildrenEdge `json:"edges"`
}

type EdgeSidecarToChildrenEdge struct {
	Node StickyNode `json:"node"`
}

type StickyNode struct {
	Typename               string                `json:"__typename"`
	ID                     string                `json:"id"`
	Shortcode              string                `json:"shortcode"`
	Dimensions             Dimensions            `json:"dimensions"`
	GatingInfo             interface{}           `json:"gating_info"`
	FactCheckOverallRating interface{}           `json:"fact_check_overall_rating"`
	FactCheckInformation   interface{}           `json:"fact_check_information"`
	MediaPreview           string                `json:"media_preview"`
	DisplayURL             string                `json:"display_url"`
	DisplayResources       []DisplayResource     `json:"display_resources"`
	AccessibilityCaption   string                `json:"accessibility_caption"`
	IsVideo                bool                  `json:"is_video"`
	TrackingToken          string                `json:"tracking_token"`
	EdgeMediaToTaggedUser  EdgeMediaToTaggedUser `json:"edge_media_to_tagged_user"`
}

type ShortcodeMediaOwner struct {
	ID                       string                        `json:"id"`
	IsVerified               bool                          `json:"is_verified"`
	ProfilePicURL            string                        `json:"profile_pic_url"`
	Username                 string                        `json:"username"`
	BlockedByViewer          bool                          `json:"blocked_by_viewer"`
	RestrictedByViewer       bool                          `json:"restricted_by_viewer"`
	FollowedByViewer         bool                          `json:"followed_by_viewer"`
	FullName                 string                        `json:"full_name"`
	HasBlockedViewer         bool                          `json:"has_blocked_viewer"`
	IsPrivate                bool                          `json:"is_private"`
	IsUnpublished            bool                          `json:"is_unpublished"`
	RequestedByViewer        bool                          `json:"requested_by_viewer"`
	EdgeOwnerToTimelineMedia EdgeOwnerToTimelineMediaClass `json:"edge_owner_to_timeline_media"`
}
