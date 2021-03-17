package model

type Data struct {
	Config struct {
		CsrfToken string      `json:"csrf_token"`
		Viewer    interface{} `json:"viewer"`
		ViewerID  interface{} `json:"viewerId"`
	} `json:"config"`
	CountryCode  string `json:"country_code"`
	LanguageCode string `json:"language_code"`
	Locale       string `json:"locale"`
	EntryData    struct {
		ProfilePage []struct {
			LoggingPageID         string `json:"logging_page_id"`
			ShowSuggestedProfiles bool   `json:"show_suggested_profiles"`
			ShowFollowDialog      bool   `json:"show_follow_dialog"`
			Graphql               struct {
				User struct {
					Biography              string      `json:"biography"`
					BlockedByViewer        bool        `json:"blocked_by_viewer"`
					RestrictedByViewer     interface{} `json:"restricted_by_viewer"`
					CountryBlock           bool        `json:"country_block"`
					ExternalURL            string      `json:"external_url"`
					ExternalURLLinkshimmed string      `json:"external_url_linkshimmed"`
					EdgeFollowedBy         struct {
						Count int `json:"count"`
					} `json:"edge_followed_by"`
					Fbid             string `json:"fbid"`
					FollowedByViewer bool   `json:"followed_by_viewer"`
					EdgeFollow       struct {
						Count int `json:"count"`
					} `json:"edge_follow"`
					FollowsViewer        bool        `json:"follows_viewer"`
					FullName             string      `json:"full_name"`
					HasArEffects         bool        `json:"has_ar_effects"`
					HasClips             bool        `json:"has_clips"`
					HasGuides            bool        `json:"has_guides"`
					HasChannel           bool        `json:"has_channel"`
					HasBlockedViewer     bool        `json:"has_blocked_viewer"`
					HighlightReelCount   int         `json:"highlight_reel_count"`
					HasRequestedViewer   bool        `json:"has_requested_viewer"`
					ID                   string      `json:"id"`
					IsBusinessAccount    bool        `json:"is_business_account"`
					IsJoinedRecently     bool        `json:"is_joined_recently"`
					BusinessCategoryName interface{} `json:"business_category_name"`
					OverallCategoryName  interface{} `json:"overall_category_name"`
					CategoryEnum         interface{} `json:"category_enum"`
					CategoryName         string      `json:"category_name"`
					IsPrivate            bool        `json:"is_private"`
					IsVerified           bool        `json:"is_verified"`
					EdgeMutualFollowedBy struct {
						Count int           `json:"count"`
						Edges []interface{} `json:"edges"`
					} `json:"edge_mutual_followed_by"`
					ProfilePicURL          string      `json:"profile_pic_url"`
					ProfilePicURLHd        string      `json:"profile_pic_url_hd"`
					RequestedByViewer      bool        `json:"requested_by_viewer"`
					ShouldShowCategory     bool        `json:"should_show_category"`
					Username               string      `json:"username"`
					ConnectedFbPage        interface{} `json:"connected_fb_page"`
					EdgeFelixVideoTimeline struct {
						Count    int `json:"count"`
						PageInfo struct {
							HasNextPage bool   `json:"has_next_page"`
							EndCursor   string `json:"end_cursor"`
						} `json:"page_info"`
						Edges []struct {
							Node struct {
								Typename   string `json:"__typename"`
								ID         string `json:"id"`
								Shortcode  string `json:"shortcode"`
								Dimensions struct {
									Height int `json:"height"`
									Width  int `json:"width"`
								} `json:"dimensions"`
								DisplayURL            string `json:"display_url"`
								EdgeMediaToTaggedUser struct {
									Edges []interface{} `json:"edges"`
								} `json:"edge_media_to_tagged_user"`
								FactCheckOverallRating interface{} `json:"fact_check_overall_rating"`
								FactCheckInformation   interface{} `json:"fact_check_information"`
								GatingInfo             interface{} `json:"gating_info"`
								SharingFrictionInfo    struct {
									ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
									BloksAppURL               interface{} `json:"bloks_app_url"`
								} `json:"sharing_friction_info"`
								MediaOverlayInfo interface{} `json:"media_overlay_info"`
								MediaPreview     interface{} `json:"media_preview"`
								Owner            struct {
									ID       string `json:"id"`
									Username string `json:"username"`
								} `json:"owner"`
								IsVideo              bool        `json:"is_video"`
								AccessibilityCaption interface{} `json:"accessibility_caption"`
								DashInfo             struct {
									IsDashEligible    bool        `json:"is_dash_eligible"`
									VideoDashManifest interface{} `json:"video_dash_manifest"`
									NumberOfQualities int         `json:"number_of_qualities"`
								} `json:"dash_info"`
								HasAudio           bool   `json:"has_audio"`
								TrackingToken      string `json:"tracking_token"`
								VideoURL           string `json:"video_url"`
								VideoViewCount     int    `json:"video_view_count"`
								EdgeMediaToCaption struct {
									Edges []struct {
										Node struct {
											Text string `json:"text"`
										} `json:"node"`
									} `json:"edges"`
								} `json:"edge_media_to_caption"`
								EdgeMediaToComment struct {
									Count int `json:"count"`
								} `json:"edge_media_to_comment"`
								CommentsDisabled bool `json:"comments_disabled"`
								TakenAtTimestamp int  `json:"taken_at_timestamp"`
								EdgeLikedBy      struct {
									Count int `json:"count"`
								} `json:"edge_liked_by"`
								EdgeMediaPreviewLike struct {
									Count int `json:"count"`
								} `json:"edge_media_preview_like"`
								Location           interface{} `json:"location"`
								ThumbnailSrc       string      `json:"thumbnail_src"`
								ThumbnailResources []struct {
									Src          string `json:"src"`
									ConfigWidth  int    `json:"config_width"`
									ConfigHeight int    `json:"config_height"`
								} `json:"thumbnail_resources"`
								FelixProfileGridCrop struct {
									CropLeft   float64 `json:"crop_left"`
									CropRight  float64 `json:"crop_right"`
									CropTop    float64 `json:"crop_top"`
									CropBottom float64 `json:"crop_bottom"`
								} `json:"felix_profile_grid_crop"`
								EncodingStatus interface{} `json:"encoding_status"`
								IsPublished    bool        `json:"is_published"`
								ProductType    string      `json:"product_type"`
								Title          string      `json:"title"`
								VideoDuration  float64     `json:"video_duration"`
							} `json:"node"`
						} `json:"edges"`
					} `json:"edge_felix_video_timeline"`
					EdgeOwnerToTimelineMedia struct {
						Count    int `json:"count"`
						PageInfo struct {
							HasNextPage bool   `json:"has_next_page"`
							EndCursor   string `json:"end_cursor"`
						} `json:"page_info"`
						Edges []struct {
							Node struct {
								Typename   string `json:"__typename"`
								ID         string `json:"id"`
								Shortcode  string `json:"shortcode"`
								Dimensions struct {
									Height int `json:"height"`
									Width  int `json:"width"`
								} `json:"dimensions"`
								DisplayURL            string `json:"display_url"`
								EdgeMediaToTaggedUser struct {
									Edges []struct {
										Node struct {
											User struct {
												FullName      string `json:"full_name"`
												ID            string `json:"id"`
												IsVerified    bool   `json:"is_verified"`
												ProfilePicURL string `json:"profile_pic_url"`
												Username      string `json:"username"`
											} `json:"user"`
											X float64 `json:"x"`
											Y float64 `json:"y"`
										} `json:"node"`
									} `json:"edges"`
								} `json:"edge_media_to_tagged_user"`
								FactCheckOverallRating interface{} `json:"fact_check_overall_rating"`
								FactCheckInformation   interface{} `json:"fact_check_information"`
								GatingInfo             interface{} `json:"gating_info"`
								SharingFrictionInfo    struct {
									ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
									BloksAppURL               interface{} `json:"bloks_app_url"`
								} `json:"sharing_friction_info"`
								MediaOverlayInfo interface{} `json:"media_overlay_info"`
								MediaPreview     interface{} `json:"media_preview"`
								Owner            struct {
									ID       string `json:"id"`
									Username string `json:"username"`
								} `json:"owner"`
								IsVideo              bool   `json:"is_video"`
								AccessibilityCaption string `json:"accessibility_caption"`
								EdgeMediaToCaption   struct {
									Edges []struct {
										Node struct {
											Text string `json:"text"`
										} `json:"node"`
									} `json:"edges"`
								} `json:"edge_media_to_caption"`
								EdgeMediaToComment struct {
									Count int `json:"count"`
								} `json:"edge_media_to_comment"`
								CommentsDisabled bool `json:"comments_disabled"`
								TakenAtTimestamp int  `json:"taken_at_timestamp"`
								EdgeLikedBy      struct {
									Count int `json:"count"`
								} `json:"edge_liked_by"`
								EdgeMediaPreviewLike struct {
									Count int `json:"count"`
								} `json:"edge_media_preview_like"`
								Location struct {
									ID            string `json:"id"`
									HasPublicPage bool   `json:"has_public_page"`
									Name          string `json:"name"`
									Slug          string `json:"slug"`
								} `json:"location"`
								ThumbnailSrc       string `json:"thumbnail_src"`
								ThumbnailResources []struct {
									Src          string `json:"src"`
									ConfigWidth  int    `json:"config_width"`
									ConfigHeight int    `json:"config_height"`
								} `json:"thumbnail_resources"`
								EdgeSidecarToChildren struct {
									Edges []struct {
										Node struct {
											Typename   string `json:"__typename"`
											ID         string `json:"id"`
											Shortcode  string `json:"shortcode"`
											Dimensions struct {
												Height int `json:"height"`
												Width  int `json:"width"`
											} `json:"dimensions"`
											DisplayURL            string `json:"display_url"`
											EdgeMediaToTaggedUser struct {
												Edges []struct {
													Node struct {
														User struct {
															FullName      string `json:"full_name"`
															ID            string `json:"id"`
															IsVerified    bool   `json:"is_verified"`
															ProfilePicURL string `json:"profile_pic_url"`
															Username      string `json:"username"`
														} `json:"user"`
														X float64 `json:"x"`
														Y float64 `json:"y"`
													} `json:"node"`
												} `json:"edges"`
											} `json:"edge_media_to_tagged_user"`
											FactCheckOverallRating interface{} `json:"fact_check_overall_rating"`
											FactCheckInformation   interface{} `json:"fact_check_information"`
											GatingInfo             interface{} `json:"gating_info"`
											SharingFrictionInfo    struct {
												ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
												BloksAppURL               interface{} `json:"bloks_app_url"`
											} `json:"sharing_friction_info"`
											MediaOverlayInfo interface{} `json:"media_overlay_info"`
											MediaPreview     string      `json:"media_preview"`
											Owner            struct {
												ID       string `json:"id"`
												Username string `json:"username"`
											} `json:"owner"`
											IsVideo              bool   `json:"is_video"`
											AccessibilityCaption string `json:"accessibility_caption"`
										} `json:"node"`
									} `json:"edges"`
								} `json:"edge_sidecar_to_children"`
							} `json:"node"`
						} `json:"edges"`
					} `json:"edge_owner_to_timeline_media"`
					EdgeSavedMedia struct {
						Count    int `json:"count"`
						PageInfo struct {
							HasNextPage bool        `json:"has_next_page"`
							EndCursor   interface{} `json:"end_cursor"`
						} `json:"page_info"`
						Edges []interface{} `json:"edges"`
					} `json:"edge_saved_media"`
					EdgeMediaCollections struct {
						Count    int `json:"count"`
						PageInfo struct {
							HasNextPage bool        `json:"has_next_page"`
							EndCursor   interface{} `json:"end_cursor"`
						} `json:"page_info"`
						Edges []interface{} `json:"edges"`
					} `json:"edge_media_collections"`
					EdgeRelatedProfiles struct {
						Edges []struct {
							Node struct {
								ID            string `json:"id"`
								FullName      string `json:"full_name"`
								IsPrivate     bool   `json:"is_private"`
								IsVerified    bool   `json:"is_verified"`
								ProfilePicURL string `json:"profile_pic_url"`
								Username      string `json:"username"`
							} `json:"node"`
						} `json:"edges"`
					} `json:"edge_related_profiles"`
				} `json:"user"`
			} `json:"graphql"`
			ToastContentOnLoad      interface{} `json:"toast_content_on_load"`
			ShowViewShop            bool        `json:"show_view_shop"`
			ProfilePicEditSyncProps interface{} `json:"profile_pic_edit_sync_props"`
		} `json:"ProfilePage"`
	} `json:"entry_data"`
	Hostname                string  `json:"hostname"`
	IsWhitelistedCrawlBot   bool    `json:"is_whitelisted_crawl_bot"`
	ConnectionQualityRating string  `json:"connection_quality_rating"`
	DeploymentStage         string  `json:"deployment_stage"`
	Platform                string  `json:"platform"`
	Nonce                   string  `json:"nonce"`
	MidPct                  float64 `json:"mid_pct"`
	ZeroData                struct {
	} `json:"zero_data"`
	CacheSchemaVersion int `json:"cache_schema_version"`
	ServerChecks       struct {
	} `json:"server_checks"`
	Knobx struct {
		Num4  bool `json:"4"`
		Num17 bool `json:"17"`
		Num22 bool `json:"22"`
		Num23 bool `json:"23"`
		Num24 bool `json:"24"`
		Num25 bool `json:"25"`
		Num26 bool `json:"26"`
		Num27 bool `json:"27"`
		Num29 bool `json:"29"`
		Num32 bool `json:"32"`
		Num34 bool `json:"34"`
		Num35 bool `json:"35"`
		Num38 int  `json:"38"`
		Num39 bool `json:"39"`
		Num40 bool `json:"40"`
		Num41 bool `json:"41"`
		Num42 bool `json:"42"`
		Num44 bool `json:"44"`
		Num45 bool `json:"45"`
		Num46 bool `json:"46"`
	} `json:"knobx"`
	ToCache struct {
		Gatekeepers struct {
			Num4   bool `json:"4"`
			Num5   bool `json:"5"`
			Num6   bool `json:"6"`
			Num7   bool `json:"7"`
			Num8   bool `json:"8"`
			Num9   bool `json:"9"`
			Num10  bool `json:"10"`
			Num11  bool `json:"11"`
			Num12  bool `json:"12"`
			Num13  bool `json:"13"`
			Num14  bool `json:"14"`
			Num15  bool `json:"15"`
			Num16  bool `json:"16"`
			Num26  bool `json:"26"`
			Num27  bool `json:"27"`
			Num28  bool `json:"28"`
			Num29  bool `json:"29"`
			Num31  bool `json:"31"`
			Num32  bool `json:"32"`
			Num34  bool `json:"34"`
			Num38  bool `json:"38"`
			Num40  bool `json:"40"`
			Num41  bool `json:"41"`
			Num43  bool `json:"43"`
			Num59  bool `json:"59"`
			Num61  bool `json:"61"`
			Num62  bool `json:"62"`
			Num63  bool `json:"63"`
			Num64  bool `json:"64"`
			Num65  bool `json:"65"`
			Num67  bool `json:"67"`
			Num68  bool `json:"68"`
			Num69  bool `json:"69"`
			Num71  bool `json:"71"`
			Num73  bool `json:"73"`
			Num74  bool `json:"74"`
			Num75  bool `json:"75"`
			Num78  bool `json:"78"`
			Num79  bool `json:"79"`
			Num81  bool `json:"81"`
			Num82  bool `json:"82"`
			Num84  bool `json:"84"`
			Num86  bool `json:"86"`
			Num91  bool `json:"91"`
			Num95  bool `json:"95"`
			Num97  bool `json:"97"`
			Num100 bool `json:"100"`
			Num101 bool `json:"101"`
			Num102 bool `json:"102"`
			Num103 bool `json:"103"`
			Num104 bool `json:"104"`
			Num105 bool `json:"105"`
			Num106 bool `json:"106"`
			Num107 bool `json:"107"`
			Num108 bool `json:"108"`
			Num112 bool `json:"112"`
			Num113 bool `json:"113"`
			Num114 bool `json:"114"`
			Num116 bool `json:"116"`
			Num120 bool `json:"120"`
			Num123 bool `json:"123"`
			Num128 bool `json:"128"`
			Num131 bool `json:"131"`
			Num132 bool `json:"132"`
			Num137 bool `json:"137"`
			Num140 bool `json:"140"`
			Num142 bool `json:"142"`
			Num146 bool `json:"146"`
			Num147 bool `json:"147"`
			Num149 bool `json:"149"`
			Num150 bool `json:"150"`
			Num151 bool `json:"151"`
			Num153 bool `json:"153"`
			Num156 bool `json:"156"`
			Num157 bool `json:"157"`
			Num158 bool `json:"158"`
			Num159 bool `json:"159"`
			Num160 bool `json:"160"`
			Num162 bool `json:"162"`
			Num166 bool `json:"166"`
			Num167 bool `json:"167"`
			Num168 bool `json:"168"`
			Num169 bool `json:"169"`
			Num170 bool `json:"170"`
			Num173 bool `json:"173"`
			Num174 bool `json:"174"`
			Num175 bool `json:"175"`
			Num178 bool `json:"178"`
			Num179 bool `json:"179"`
			Num181 bool `json:"181"`
			Num185 bool `json:"185"`
			Num186 bool `json:"186"`
			Num187 bool `json:"187"`
			Num188 bool `json:"188"`
			Num189 bool `json:"189"`
			Num190 bool `json:"190"`
			Num191 bool `json:"191"`
			Num192 bool `json:"192"`
			Num193 bool `json:"193"`
			Num195 bool `json:"195"`
			Num196 bool `json:"196"`
			Num197 bool `json:"197"`
			Num198 bool `json:"198"`
			Num199 bool `json:"199"`
			Num200 bool `json:"200"`
			Num201 bool `json:"201"`
			Num202 bool `json:"202"`
			Num203 bool `json:"203"`
			Num204 bool `json:"204"`
			Num205 bool `json:"205"`
			Num208 bool `json:"208"`
			Num209 bool `json:"209"`
			Num211 bool `json:"211"`
			Num212 bool `json:"212"`
			Num213 bool `json:"213"`
			Num215 bool `json:"215"`
			Num218 bool `json:"218"`
			Num219 bool `json:"219"`
			Num221 bool `json:"221"`
			Num222 bool `json:"222"`
			Num223 bool `json:"223"`
			Num224 bool `json:"224"`
			Num226 bool `json:"226"`
			Num229 bool `json:"229"`
			Num230 bool `json:"230"`
			Num231 bool `json:"231"`
			Num234 bool `json:"234"`
			Num237 bool `json:"237"`
			Num238 bool `json:"238"`
			Num239 bool `json:"239"`
			Num240 bool `json:"240"`
		} `json:"gatekeepers"`
		Qe struct {
			Num0 struct {
				P struct {
					Num9 bool `json:"9"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"0"`
			Num12 struct {
				P struct {
					Num0 int `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"12"`
			Num13 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"13"`
			Num16 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"16"`
			Num22 struct {
				P struct {
					Num2  float64 `json:"2"`
					Num3  float64 `json:"3"`
					Num4  float64 `json:"4"`
					Num10 float64 `json:"10"`
					Num11 int     `json:"11"`
					Num12 int     `json:"12"`
					Num13 bool    `json:"13"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"22"`
			Num23 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"23"`
			Num25 struct {
				P struct {
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"25"`
			Num26 struct {
				P struct {
					Num0 string `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"26"`
			Num28 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"28"`
			Num31 struct {
				P struct {
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"31"`
			Num33 struct {
				P struct {
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"33"`
			Num34 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"34"`
			Num36 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"36"`
			Num37 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"37"`
			Num39 struct {
				P struct {
					Num0  bool `json:"0"`
					Num8  bool `json:"8"`
					Num14 bool `json:"14"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"39"`
			Num41 struct {
				P struct {
					Num3 bool `json:"3"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"41"`
			Num42 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"42"`
			Num43 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"43"`
			Num44 struct {
				P struct {
					Num1 string  `json:"1"`
					Num2 float64 `json:"2"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"44"`
			Num45 struct {
				P struct {
					Num13 bool   `json:"13"`
					Num17 int    `json:"17"`
					Num32 bool   `json:"32"`
					Num35 bool   `json:"35"`
					Num36 string `json:"36"`
					Num37 bool   `json:"37"`
					Num38 bool   `json:"38"`
					Num40 string `json:"40"`
					Num46 bool   `json:"46"`
					Num47 bool   `json:"47"`
					Num52 bool   `json:"52"`
					Num53 bool   `json:"53"`
					Num55 bool   `json:"55"`
					Num56 string `json:"56"`
					Num57 bool   `json:"57"`
					Num58 bool   `json:"58"`
					Num59 bool   `json:"59"`
					Num60 string `json:"60"`
					Num62 string `json:"62"`
					Num64 bool   `json:"64"`
					Num65 bool   `json:"65"`
					Num66 int    `json:"66"`
					Num67 bool   `json:"67"`
					Num68 bool   `json:"68"`
					Num69 string `json:"69"`
					Num71 bool   `json:"71"`
					Num72 bool   `json:"72"`
					Num74 bool   `json:"74"`
					Num76 bool   `json:"76"`
					Num77 bool   `json:"77"`
					Num78 bool   `json:"78"`
				} `json:"p"`
				L struct {
					Num76 bool `json:"76"`
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"45"`
			Num46 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"46"`
			Num47 struct {
				P struct {
					Num0  bool `json:"0"`
					Num9  bool `json:"9"`
					Num10 bool `json:"10"`
					Num11 bool `json:"11"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"47"`
			Num49 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"49"`
			Num50 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"50"`
			Num54 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"54"`
			Num58 struct {
				P struct {
					Num0 float64 `json:"0"`
					Num1 bool    `json:"1"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"58"`
			Num59 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"59"`
			Num62 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"62"`
			Num67 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
					Num3 bool `json:"3"`
					Num4 bool `json:"4"`
					Num5 bool `json:"5"`
					Num7 bool `json:"7"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"67"`
			Num69 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"69"`
			Num71 struct {
				P struct {
					Num1 string `json:"1"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"71"`
			Num72 struct {
				P struct {
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
				} `json:"p"`
				L struct {
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"72"`
			Num73 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"73"`
			Num74 struct {
				P struct {
					Num1  bool `json:"1"`
					Num2  bool `json:"2"`
					Num3  bool `json:"3"`
					Num4  bool `json:"4"`
					Num7  bool `json:"7"`
					Num9  bool `json:"9"`
					Num13 bool `json:"13"`
					Num15 bool `json:"15"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"74"`
			Num75 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"75"`
			Num77 struct {
				P struct {
					Num1 bool `json:"1"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"77"`
			Num80 struct {
				P struct {
					Num3 bool `json:"3"`
					Num4 bool `json:"4"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"80"`
			Num84 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
					Num3 bool `json:"3"`
					Num4 bool `json:"4"`
					Num5 bool `json:"5"`
					Num6 bool `json:"6"`
					Num8 bool `json:"8"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"84"`
			Num85 struct {
				P struct {
					Num0 bool   `json:"0"`
					Num1 string `json:"1"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"85"`
			Num87 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"87"`
			Num93 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"93"`
			Num95 struct {
				P struct {
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"95"`
			Num98 struct {
				P struct {
					Num1 bool `json:"1"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"98"`
			Num100 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"100"`
			Num101 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"101"`
			Num108 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"108"`
			Num109 struct {
				P struct {
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"109"`
			Num111 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"111"`
			Num113 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
					Num4 bool `json:"4"`
					Num5 bool `json:"5"`
					Num7 bool `json:"7"`
					Num8 bool `json:"8"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"113"`
			Num118 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"118"`
			Num119 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"119"`
			Num121 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"121"`
			Num122 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"122"`
			Num123 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"123"`
			Num124 struct {
				P struct {
					Num0 bool   `json:"0"`
					Num1 bool   `json:"1"`
					Num2 bool   `json:"2"`
					Num4 bool   `json:"4"`
					Num5 string `json:"5"`
					Num6 string `json:"6"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"124"`
			Num125 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"125"`
			Num127 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"127"`
			Num128 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"128"`
			Num129 struct {
				P struct {
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"129"`
			Num131 struct {
				P struct {
					Num2 bool `json:"2"`
					Num3 bool `json:"3"`
					Num4 bool `json:"4"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"131"`
			Num132 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"132"`
			Num135 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
					Num3 bool `json:"3"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"135"`
			Num137 struct {
				P struct {
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"137"`
			Num141 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
					Num3 bool `json:"3"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"141"`
			Num142 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"142"`
			Num143 struct {
				P struct {
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
					Num3 bool `json:"3"`
					Num4 bool `json:"4"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"143"`
			Num146 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"146"`
			Num148 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"148"`
			Num149 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"149"`
			Num152 struct {
				P struct {
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"152"`
			Num154 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"154"`
			Num155 struct {
				P struct {
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"155"`
			Num156 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"156"`
			Num158 struct {
				P struct {
					Num2 bool `json:"2"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"158"`
			Num159 struct {
				P struct {
					Num1 bool `json:"1"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"159"`
			Num160 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"160"`
			Num162 struct {
				P struct {
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"162"`
			Num163 struct {
				P struct {
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"163"`
			Num164 struct {
				P struct {
					Num0 bool `json:"0"`
					Num2 bool `json:"2"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"164"`
			Num165 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"165"`
			Num166 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"166"`
			Num167 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"167"`
			Num168 struct {
				P struct {
					Num0 bool `json:"0"`
					Num3 bool `json:"3"`
					Num4 bool `json:"4"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"168"`
			Num169 struct {
				P struct {
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"169"`
			Num170 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"170"`
			Num171 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				L struct {
				} `json:"l"`
				Qex bool `json:"qex"`
			} `json:"171"`
			AppUpsell struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"app_upsell"`
			IglAppUpsell struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"igl_app_upsell"`
			Notif struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"notif"`
			Onetaplogin struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"onetaplogin"`
			FelixClearFbCookie struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"felix_clear_fb_cookie"`
			FelixCreationDurationLimits struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"felix_creation_duration_limits"`
			FelixCreationFbCrossposting struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"felix_creation_fb_crossposting"`
			FelixCreationFbCrosspostingV2 struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"felix_creation_fb_crossposting_v2"`
			FelixCreationValidation struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"felix_creation_validation"`
			PostOptions struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"post_options"`
			StickerTray struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"sticker_tray"`
			WebSentry struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"web_sentry"`
		} `json:"qe"`
		ProbablyHasApp bool `json:"probably_has_app"`
		Cb             bool `json:"cb"`
	} `json:"to_cache"`
	DeviceID          string `json:"device_id"`
	BrowserPushPubKey string `json:"browser_push_pub_key"`
	Encryption        struct {
		KeyID     string `json:"key_id"`
		PublicKey string `json:"public_key"`
		Version   string `json:"version"`
	} `json:"encryption"`
	IsDev                   bool        `json:"is_dev"`
	SignalCollectionConfig  interface{} `json:"signal_collection_config"`
	ShouldShowConsentDialog interface{} `json:"should_show_consent_dialog"`
	RolloutHash             string      `json:"rollout_hash"`
	BundleVariant           string      `json:"bundle_variant"`
	FrontendEnv             string      `json:"frontend_env"`
}
