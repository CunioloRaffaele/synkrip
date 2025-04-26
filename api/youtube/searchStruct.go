package youtube

// Generated with https://transform.tools/json-to-go
type searchStruct struct {
	ResponseContext struct {
		VisitorData           string `json:"visitorData"`
		ServiceTrackingParams []struct {
			Service string `json:"service"`
			Params  []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"params"`
		} `json:"serviceTrackingParams"`
		MaxAgeSeconds int `json:"maxAgeSeconds"`
	} `json:"responseContext"`
	Contents struct {
		TabbedSearchResultsRenderer struct {
			Tabs []struct {
				TabRenderer struct {
					Title    string `json:"title"`
					Selected bool   `json:"selected"`
					Content  struct {
						SectionListRenderer struct {
							Contents []struct {
								ItemSectionRenderer struct {
									Contents []struct {
										MessageRenderer struct {
											TrackingParams string `json:"trackingParams"`
											Button         struct {
												ButtonRenderer struct {
													Style      string `json:"style"`
													IsDisabled bool   `json:"isDisabled"`
													Text       struct {
														SimpleText string `json:"simpleText"`
													} `json:"text"`
													Icon struct {
														IconType string `json:"iconType"`
													} `json:"icon"`
													NavigationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														URLEndpoint         struct {
															URL    string `json:"url"`
															Target string `json:"target"`
														} `json:"urlEndpoint"`
													} `json:"navigationEndpoint"`
													TrackingParams string `json:"trackingParams"`
													IconPosition   string `json:"iconPosition"`
												} `json:"buttonRenderer"`
											} `json:"button"`
											Style struct {
												Value string `json:"value"`
											} `json:"style"`
										} `json:"messageRenderer"`
									} `json:"contents"`
									TrackingParams string `json:"trackingParams"`
								} `json:"itemSectionRenderer,omitempty"`
								MusicCardShelfRenderer* struct {
									TrackingParams string `json:"trackingParams"`
									Thumbnail      struct {
										MusicThumbnailRenderer struct {
											Thumbnail struct {
												Thumbnails []struct {
													URL    string `json:"url"`
													Width  int    `json:"width"`
													Height int    `json:"height"`
												} `json:"thumbnails"`
											} `json:"thumbnail"`
											ThumbnailCrop  string `json:"thumbnailCrop"`
											ThumbnailScale string `json:"thumbnailScale"`
											TrackingParams string `json:"trackingParams"`
										} `json:"musicThumbnailRenderer"`
									} `json:"thumbnail"`
									Title struct {
										Runs []struct {
											Text               string `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												WatchEndpoint       struct {
													VideoID                            string `json:"videoId"`
													WatchEndpointMusicSupportedConfigs struct {
														WatchEndpointMusicConfig struct {
															MusicVideoType string `json:"musicVideoType"`
														} `json:"watchEndpointMusicConfig"`
													} `json:"watchEndpointMusicSupportedConfigs"`
												} `json:"watchEndpoint"`
											} `json:"navigationEndpoint"`
										} `json:"runs"`
									} `json:"title"`
									Subtitle struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
										Accessibility struct {
											AccessibilityData struct {
												Label string `json:"label"`
											} `json:"accessibilityData"`
										} `json:"accessibility"`
									} `json:"subtitle"`
									Buttons []struct {
										ButtonRenderer struct {
											Style      string `json:"style"`
											Size       string `json:"size"`
											IsDisabled bool   `json:"isDisabled"`
											Text       struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											Icon struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
											Accessibility struct {
												Label string `json:"label"`
											} `json:"accessibility"`
											TrackingParams    string `json:"trackingParams"`
											AccessibilityData struct {
												AccessibilityData struct {
													Label string `json:"label"`
												} `json:"accessibilityData"`
											} `json:"accessibilityData"`
											Command struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												WatchEndpoint       struct {
													VideoID                            string `json:"videoId"`
													Params                             string `json:"params"`
													WatchEndpointMusicSupportedConfigs struct {
														WatchEndpointMusicConfig struct {
															MusicVideoType string `json:"musicVideoType"`
														} `json:"watchEndpointMusicConfig"`
													} `json:"watchEndpointMusicSupportedConfigs"`
												} `json:"watchEndpoint"`
											} `json:"command"`
										} `json:"buttonRenderer"`
									} `json:"buttons"`
									Menu struct {
										MenuRenderer struct {
											Items []struct {
												MenuNavigationItemRenderer struct {
													Text struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"text"`
													Icon struct {
														IconType string `json:"iconType"`
													} `json:"icon"`
													NavigationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														WatchEndpoint       struct {
															VideoID        string `json:"videoId"`
															PlaylistID     string `json:"playlistId"`
															Params         string `json:"params"`
															LoggingContext struct {
																VssLoggingContext struct {
																	SerializedContextData string `json:"serializedContextData"`
																} `json:"vssLoggingContext"`
															} `json:"loggingContext"`
															WatchEndpointMusicSupportedConfigs struct {
																WatchEndpointMusicConfig struct {
																	MusicVideoType string `json:"musicVideoType"`
																} `json:"watchEndpointMusicConfig"`
															} `json:"watchEndpointMusicSupportedConfigs"`
														} `json:"watchEndpoint"`
													} `json:"navigationEndpoint"`
													TrackingParams string `json:"trackingParams"`
												} `json:"menuNavigationItemRenderer,omitempty"`
												MenuServiceItemRenderer struct {
													Text struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"text"`
													Icon struct {
														IconType string `json:"iconType"`
													} `json:"icon"`
													ServiceEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														QueueAddEndpoint    struct {
															QueueTarget struct {
																VideoID      string `json:"videoId"`
																OnEmptyQueue struct {
																	ClickTrackingParams string `json:"clickTrackingParams"`
																	WatchEndpoint       struct {
																		VideoID string `json:"videoId"`
																	} `json:"watchEndpoint"`
																} `json:"onEmptyQueue"`
															} `json:"queueTarget"`
															QueueInsertPosition string `json:"queueInsertPosition"`
															Commands            []struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																AddToToastAction    struct {
																	Item struct {
																		NotificationTextRenderer struct {
																			SuccessResponseText struct {
																				Runs []struct {
																					Text string `json:"text"`
																				} `json:"runs"`
																			} `json:"successResponseText"`
																			TrackingParams string `json:"trackingParams"`
																		} `json:"notificationTextRenderer"`
																	} `json:"item"`
																} `json:"addToToastAction"`
															} `json:"commands"`
														} `json:"queueAddEndpoint"`
													} `json:"serviceEndpoint"`
													TrackingParams string `json:"trackingParams"`
												} `json:"menuServiceItemRenderer,omitempty"`
												ToggleMenuServiceItemRenderer struct {
													DefaultText struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"defaultText"`
													DefaultIcon struct {
														IconType string `json:"iconType"`
													} `json:"defaultIcon"`
													DefaultServiceEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														ModalEndpoint       struct {
															Modal struct {
																ModalWithTitleAndButtonRenderer struct {
																	Title struct {
																		Runs []struct {
																			Text string `json:"text"`
																		} `json:"runs"`
																	} `json:"title"`
																	Content struct {
																		Runs []struct {
																			Text string `json:"text"`
																		} `json:"runs"`
																	} `json:"content"`
																	Button struct {
																		ButtonRenderer struct {
																			Style      string `json:"style"`
																			IsDisabled bool   `json:"isDisabled"`
																			Text       struct {
																				Runs []struct {
																					Text string `json:"text"`
																				} `json:"runs"`
																			} `json:"text"`
																			NavigationEndpoint struct {
																				ClickTrackingParams string `json:"clickTrackingParams"`
																				SignInEndpoint      struct {
																					Hack bool `json:"hack"`
																				} `json:"signInEndpoint"`
																			} `json:"navigationEndpoint"`
																			TrackingParams string `json:"trackingParams"`
																		} `json:"buttonRenderer"`
																	} `json:"button"`
																} `json:"modalWithTitleAndButtonRenderer"`
															} `json:"modal"`
														} `json:"modalEndpoint"`
													} `json:"defaultServiceEndpoint"`
													ToggledText struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"toggledText"`
													ToggledIcon struct {
														IconType string `json:"iconType"`
													} `json:"toggledIcon"`
													TrackingParams string `json:"trackingParams"`
												} `json:"toggleMenuServiceItemRenderer,omitempty"`
											} `json:"items"`
											TrackingParams string `json:"trackingParams"`
											Accessibility  struct {
												AccessibilityData struct {
													Label string `json:"label"`
												} `json:"accessibilityData"`
											} `json:"accessibility"`
										} `json:"menuRenderer"`
									} `json:"menu"`
									OnTap struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										WatchEndpoint       struct {
											VideoID                            string `json:"videoId"`
											WatchEndpointMusicSupportedConfigs struct {
												WatchEndpointMusicConfig struct {
													MusicVideoType string `json:"musicVideoType"`
												} `json:"watchEndpointMusicConfig"`
											} `json:"watchEndpointMusicSupportedConfigs"`
										} `json:"watchEndpoint"`
									} `json:"onTap"`
									Header struct {
										MusicCardShelfHeaderBasicRenderer struct {
											Title struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"title"`
											TrackingParams string `json:"trackingParams"`
											Strapline      struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"strapline"`
										} `json:"musicCardShelfHeaderBasicRenderer"`
									} `json:"header"`
									ThumbnailOverlay struct {
										MusicItemThumbnailOverlayRenderer struct {
											Background struct {
												VerticalGradient struct {
													GradientLayerColors []string `json:"gradientLayerColors"`
												} `json:"verticalGradient"`
											} `json:"background"`
											Content struct {
												MusicPlayButtonRenderer struct {
													PlayNavigationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														WatchEndpoint       struct {
															VideoID                            string `json:"videoId"`
															WatchEndpointMusicSupportedConfigs struct {
																WatchEndpointMusicConfig struct {
																	MusicVideoType string `json:"musicVideoType"`
																} `json:"watchEndpointMusicConfig"`
															} `json:"watchEndpointMusicSupportedConfigs"`
														} `json:"watchEndpoint"`
													} `json:"playNavigationEndpoint"`
													TrackingParams string `json:"trackingParams"`
													PlayIcon       struct {
														IconType string `json:"iconType"`
													} `json:"playIcon"`
													PauseIcon struct {
														IconType string `json:"iconType"`
													} `json:"pauseIcon"`
													IconColor             int64 `json:"iconColor"`
													BackgroundColor       int   `json:"backgroundColor"`
													ActiveBackgroundColor int   `json:"activeBackgroundColor"`
													LoadingIndicatorColor int   `json:"loadingIndicatorColor"`
													PlayingIcon           struct {
														IconType string `json:"iconType"`
													} `json:"playingIcon"`
													IconLoadingColor      int    `json:"iconLoadingColor"`
													ActiveScaleFactor     int    `json:"activeScaleFactor"`
													ButtonSize            string `json:"buttonSize"`
													RippleTarget          string `json:"rippleTarget"`
													AccessibilityPlayData struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"accessibilityPlayData"`
													AccessibilityPauseData struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"accessibilityPauseData"`
												} `json:"musicPlayButtonRenderer"`
											} `json:"content"`
											ContentPosition string `json:"contentPosition"`
											DisplayStyle    string `json:"displayStyle"`
										} `json:"musicItemThumbnailOverlayRenderer"`
									} `json:"thumbnailOverlay"`
								} `json:"musicCardShelfRenderer,omitempty"`
								MusicShelfRenderer* struct {
									Title struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"title"`
									Contents []struct {
										MusicResponsiveListItemRenderer struct {
											TrackingParams string `json:"trackingParams"`
											Thumbnail      struct {
												MusicThumbnailRenderer struct {
													Thumbnail struct {
														Thumbnails []struct {
															URL    string `json:"url"`
															Width  int    `json:"width"`
															Height int    `json:"height"`
														} `json:"thumbnails"`
													} `json:"thumbnail"`
													ThumbnailCrop  string `json:"thumbnailCrop"`
													ThumbnailScale string `json:"thumbnailScale"`
													TrackingParams string `json:"trackingParams"`
												} `json:"musicThumbnailRenderer"`
											} `json:"thumbnail"`
											Overlay struct {
												MusicItemThumbnailOverlayRenderer struct {
													Background struct {
														VerticalGradient struct {
															GradientLayerColors []string `json:"gradientLayerColors"`
														} `json:"verticalGradient"`
													} `json:"background"`
													Content struct {
														MusicPlayButtonRenderer struct {
															PlayNavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																WatchEndpoint       struct {
																	VideoID                            string `json:"videoId"`
																	WatchEndpointMusicSupportedConfigs struct {
																		WatchEndpointMusicConfig struct {
																			MusicVideoType string `json:"musicVideoType"`
																		} `json:"watchEndpointMusicConfig"`
																	} `json:"watchEndpointMusicSupportedConfigs"`
																} `json:"watchEndpoint"`
															} `json:"playNavigationEndpoint"`
															TrackingParams string `json:"trackingParams"`
															PlayIcon       struct {
																IconType string `json:"iconType"`
															} `json:"playIcon"`
															PauseIcon struct {
																IconType string `json:"iconType"`
															} `json:"pauseIcon"`
															IconColor             int64 `json:"iconColor"`
															BackgroundColor       int   `json:"backgroundColor"`
															ActiveBackgroundColor int   `json:"activeBackgroundColor"`
															LoadingIndicatorColor int   `json:"loadingIndicatorColor"`
															PlayingIcon           struct {
																IconType string `json:"iconType"`
															} `json:"playingIcon"`
															IconLoadingColor      int    `json:"iconLoadingColor"`
															ActiveScaleFactor     int    `json:"activeScaleFactor"`
															ButtonSize            string `json:"buttonSize"`
															RippleTarget          string `json:"rippleTarget"`
															AccessibilityPlayData struct {
																AccessibilityData struct {
																	Label string `json:"label"`
																} `json:"accessibilityData"`
															} `json:"accessibilityPlayData"`
															AccessibilityPauseData struct {
																AccessibilityData struct {
																	Label string `json:"label"`
																} `json:"accessibilityData"`
															} `json:"accessibilityPauseData"`
														} `json:"musicPlayButtonRenderer"`
													} `json:"content"`
													ContentPosition string `json:"contentPosition"`
													DisplayStyle    string `json:"displayStyle"`
												} `json:"musicItemThumbnailOverlayRenderer"`
											} `json:"overlay"`
											FlexColumns []struct {
												MusicResponsiveListItemFlexColumnRenderer struct {
													Text struct {
														Runs []struct {
															Text               string `json:"text"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																WatchEndpoint       struct {
																	VideoID                            string `json:"videoId"`
																	WatchEndpointMusicSupportedConfigs struct {
																		WatchEndpointMusicConfig struct {
																			MusicVideoType string `json:"musicVideoType"`
																		} `json:"watchEndpointMusicConfig"`
																	} `json:"watchEndpointMusicSupportedConfigs"`
																} `json:"watchEndpoint"`
															} `json:"navigationEndpoint"`
														} `json:"runs"`
													} `json:"text"`
													DisplayPriority string `json:"displayPriority"`
												} `json:"musicResponsiveListItemFlexColumnRenderer"`
											} `json:"flexColumns"`
											Menu struct {
												MenuRenderer struct {
													Items []struct {
														MenuNavigationItemRenderer struct {
															Text struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"text"`
															Icon struct {
																IconType string `json:"iconType"`
															} `json:"icon"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																WatchEndpoint       struct {
																	VideoID        string `json:"videoId"`
																	PlaylistID     string `json:"playlistId"`
																	Params         string `json:"params"`
																	LoggingContext struct {
																		VssLoggingContext struct {
																			SerializedContextData string `json:"serializedContextData"`
																		} `json:"vssLoggingContext"`
																	} `json:"loggingContext"`
																	WatchEndpointMusicSupportedConfigs struct {
																		WatchEndpointMusicConfig struct {
																			MusicVideoType string `json:"musicVideoType"`
																		} `json:"watchEndpointMusicConfig"`
																	} `json:"watchEndpointMusicSupportedConfigs"`
																} `json:"watchEndpoint"`
															} `json:"navigationEndpoint"`
															TrackingParams string `json:"trackingParams"`
														} `json:"menuNavigationItemRenderer,omitempty"`
														MenuServiceItemRenderer struct {
															Text struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"text"`
															Icon struct {
																IconType string `json:"iconType"`
															} `json:"icon"`
															ServiceEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																QueueAddEndpoint    struct {
																	QueueTarget struct {
																		VideoID      string `json:"videoId"`
																		OnEmptyQueue struct {
																			ClickTrackingParams string `json:"clickTrackingParams"`
																			WatchEndpoint       struct {
																				VideoID string `json:"videoId"`
																			} `json:"watchEndpoint"`
																		} `json:"onEmptyQueue"`
																	} `json:"queueTarget"`
																	QueueInsertPosition string `json:"queueInsertPosition"`
																	Commands            []struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		AddToToastAction    struct {
																			Item struct {
																				NotificationTextRenderer struct {
																					SuccessResponseText struct {
																						Runs []struct {
																							Text string `json:"text"`
																						} `json:"runs"`
																					} `json:"successResponseText"`
																					TrackingParams string `json:"trackingParams"`
																				} `json:"notificationTextRenderer"`
																			} `json:"item"`
																		} `json:"addToToastAction"`
																	} `json:"commands"`
																} `json:"queueAddEndpoint"`
															} `json:"serviceEndpoint"`
															TrackingParams string `json:"trackingParams"`
														} `json:"menuServiceItemRenderer,omitempty"`
														ToggleMenuServiceItemRenderer struct {
															DefaultText struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"defaultText"`
															DefaultIcon struct {
																IconType string `json:"iconType"`
															} `json:"defaultIcon"`
															DefaultServiceEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																ModalEndpoint       struct {
																	Modal struct {
																		ModalWithTitleAndButtonRenderer struct {
																			Title struct {
																				Runs []struct {
																					Text string `json:"text"`
																				} `json:"runs"`
																			} `json:"title"`
																			Content struct {
																				Runs []struct {
																					Text string `json:"text"`
																				} `json:"runs"`
																			} `json:"content"`
																			Button struct {
																				ButtonRenderer struct {
																					Style      string `json:"style"`
																					IsDisabled bool   `json:"isDisabled"`
																					Text       struct {
																						Runs []struct {
																							Text string `json:"text"`
																						} `json:"runs"`
																					} `json:"text"`
																					NavigationEndpoint struct {
																						ClickTrackingParams string `json:"clickTrackingParams"`
																						SignInEndpoint      struct {
																							Hack bool `json:"hack"`
																						} `json:"signInEndpoint"`
																					} `json:"navigationEndpoint"`
																					TrackingParams string `json:"trackingParams"`
																				} `json:"buttonRenderer"`
																			} `json:"button"`
																		} `json:"modalWithTitleAndButtonRenderer"`
																	} `json:"modal"`
																} `json:"modalEndpoint"`
															} `json:"defaultServiceEndpoint"`
															ToggledText struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"toggledText"`
															ToggledIcon struct {
																IconType string `json:"iconType"`
															} `json:"toggledIcon"`
															TrackingParams string `json:"trackingParams"`
														} `json:"toggleMenuServiceItemRenderer,omitempty"`
													} `json:"items"`
													TrackingParams string `json:"trackingParams"`
													Accessibility  struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"accessibility"`
												} `json:"menuRenderer"`
											} `json:"menu"`
											Badges []struct {
												MusicInlineBadgeRenderer struct {
													TrackingParams string `json:"trackingParams"`
													Icon           struct {
														IconType string `json:"iconType"`
													} `json:"icon"`
													AccessibilityData struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"accessibilityData"`
												} `json:"musicInlineBadgeRenderer"`
											} `json:"badges"`
											PlaylistItemData struct {
												VideoID string `json:"videoId"`
											} `json:"playlistItemData"`
											FlexColumnDisplayStyle string `json:"flexColumnDisplayStyle"`
											ItemHeight             string `json:"itemHeight"`
										} `json:"musicResponsiveListItemRenderer"`
									} `json:"contents"`
									TrackingParams string `json:"trackingParams"`
									BottomText     struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"bottomText"`
									BottomEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										SearchEndpoint      struct {
											Query  string `json:"query"`
											Params string `json:"params"`
										} `json:"searchEndpoint"`
									} `json:"bottomEndpoint"`
									ShelfDivider struct {
										MusicShelfDividerRenderer struct {
											Hidden bool `json:"hidden"`
										} `json:"musicShelfDividerRenderer"`
									} `json:"shelfDivider"`
								} `json:"musicShelfRenderer,omitempty"`
							} `json:"contents"`
							TrackingParams string `json:"trackingParams"`
							Header         struct {
								ChipCloudRenderer struct {
									Chips []struct {
										ChipCloudChipRenderer struct {
											Style struct {
												StyleType string `json:"styleType"`
											} `json:"style"`
											Text struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												SearchEndpoint      struct {
													Query  string `json:"query"`
													Params string `json:"params"`
												} `json:"searchEndpoint"`
											} `json:"navigationEndpoint"`
											TrackingParams    string `json:"trackingParams"`
											AccessibilityData struct {
												AccessibilityData struct {
													Label string `json:"label"`
												} `json:"accessibilityData"`
											} `json:"accessibilityData"`
											IsSelected bool   `json:"isSelected"`
											UniqueID   string `json:"uniqueId"`
										} `json:"chipCloudChipRenderer"`
									} `json:"chips"`
									CollapsedRowCount    int    `json:"collapsedRowCount"`
									TrackingParams       string `json:"trackingParams"`
									HorizontalScrollable bool   `json:"horizontalScrollable"`
								} `json:"chipCloudRenderer"`
							} `json:"header"`
						} `json:"sectionListRenderer"`
					} `json:"content"`
					TabIdentifier  string `json:"tabIdentifier"`
					TrackingParams string `json:"trackingParams"`
				} `json:"tabRenderer"`
			} `json:"tabs"`
		} `json:"tabbedSearchResultsRenderer"`
	} `json:"contents"`
	TrackingParams string `json:"trackingParams"`
}