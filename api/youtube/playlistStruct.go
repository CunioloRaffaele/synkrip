package youtube

// Generated with https://transform.tools/json-to-go
type playlistStruct struct {
	Contents struct {
		TwoColumnBrowseResultsRenderer struct {
			Tabs []struct {
				TabRenderer struct {
					Selected bool `json:"selected"`
					Content  struct {
						SectionListRenderer struct {
							Contents []struct {
								ItemSectionRenderer struct {
									Contents []struct {
										PlaylistVideoListRenderer struct {
											Contents []struct {
												PlaylistVideoRenderer struct {
													VideoID   string `json:"videoId"`
													Thumbnail struct {
														Thumbnails []struct {
															URL    string `json:"url"`
															Width  int    `json:"width"`
															Height int    `json:"height"`
														} `json:"thumbnails"`
													} `json:"thumbnail"`
													Title struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
														Accessibility struct {
															AccessibilityData struct {
																Label string `json:"label"`
															} `json:"accessibilityData"`
														} `json:"accessibility"`
													} `json:"title"`
													Index struct {
														SimpleText string `json:"simpleText"`
													} `json:"index"`
													ShortBylineText struct {
														Runs []struct {
															Text               string `json:"text"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		URL         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																		APIURL      string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																BrowseEndpoint struct {
																	BrowseID         string `json:"browseId"`
																	CanonicalBaseURL string `json:"canonicalBaseUrl"`
																} `json:"browseEndpoint"`
															} `json:"navigationEndpoint"`
														} `json:"runs"`
													} `json:"shortBylineText"`
													LengthText struct {
														Accessibility struct {
															AccessibilityData struct {
																Label string `json:"label"`
															} `json:"accessibilityData"`
														} `json:"accessibility"`
														SimpleText string `json:"simpleText"`
													} `json:"lengthText"`
													NavigationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																URL         string `json:"url"`
																WebPageType string `json:"webPageType"`
																RootVe      int    `json:"rootVe"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														WatchEndpoint struct {
															VideoID        string `json:"videoId"`
															PlaylistID     string `json:"playlistId"`
															Index          int    `json:"index"`
															Params         string `json:"params"`
															PlayerParams   string `json:"playerParams"`
															LoggingContext struct {
																VssLoggingContext struct {
																	SerializedContextData string `json:"serializedContextData"`
																} `json:"vssLoggingContext"`
															} `json:"loggingContext"`
															WatchEndpointSupportedOnesieConfig struct {
																HTML5PlaybackOnesieConfig struct {
																	CommonConfig struct {
																		URL string `json:"url"`
																	} `json:"commonConfig"`
																} `json:"html5PlaybackOnesieConfig"`
															} `json:"watchEndpointSupportedOnesieConfig"`
														} `json:"watchEndpoint"`
													} `json:"navigationEndpoint"`
													LengthSeconds  string `json:"lengthSeconds"`
													TrackingParams string `json:"trackingParams"`
													IsPlayable     bool   `json:"isPlayable"`
													Menu           struct {
														MenuRenderer struct {
															Items []struct {
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
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				SendPost bool `json:"sendPost"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		SignalServiceEndpoint struct {
																			Signal  string `json:"signal"`
																			Actions []struct {
																				ClickTrackingParams  string `json:"clickTrackingParams"`
																				AddToPlaylistCommand struct {
																					OpenMiniplayer      bool   `json:"openMiniplayer"`
																					VideoID             string `json:"videoId"`
																					ListType            string `json:"listType"`
																					OnCreateListCommand struct {
																						ClickTrackingParams string `json:"clickTrackingParams"`
																						CommandMetadata     struct {
																							WebCommandMetadata struct {
																								SendPost bool   `json:"sendPost"`
																								APIURL   string `json:"apiUrl"`
																							} `json:"webCommandMetadata"`
																						} `json:"commandMetadata"`
																						CreatePlaylistServiceEndpoint struct {
																							VideoIds []string `json:"videoIds"`
																							Params   string   `json:"params"`
																						} `json:"createPlaylistServiceEndpoint"`
																					} `json:"onCreateListCommand"`
																					VideoIds     []string `json:"videoIds"`
																					VideoCommand struct {
																						ClickTrackingParams string `json:"clickTrackingParams"`
																						CommandMetadata     struct {
																							WebCommandMetadata struct {
																								URL         string `json:"url"`
																								WebPageType string `json:"webPageType"`
																								RootVe      int    `json:"rootVe"`
																							} `json:"webCommandMetadata"`
																						} `json:"commandMetadata"`
																						WatchEndpoint struct {
																							VideoID                            string `json:"videoId"`
																							WatchEndpointSupportedOnesieConfig struct {
																								HTML5PlaybackOnesieConfig struct {
																									CommonConfig struct {
																										URL string `json:"url"`
																									} `json:"commonConfig"`
																								} `json:"html5PlaybackOnesieConfig"`
																							} `json:"watchEndpointSupportedOnesieConfig"`
																						} `json:"watchEndpoint"`
																					} `json:"videoCommand"`
																				} `json:"addToPlaylistCommand"`
																			} `json:"actions"`
																		} `json:"signalServiceEndpoint"`
																	} `json:"serviceEndpoint"`
																	TrackingParams string `json:"trackingParams"`
																} `json:"menuServiceItemRenderer"`
															} `json:"items"`
															TrackingParams string `json:"trackingParams"`
															Accessibility  struct {
																AccessibilityData struct {
																	Label string `json:"label"`
																} `json:"accessibilityData"`
															} `json:"accessibility"`
														} `json:"menuRenderer"`
													} `json:"menu"`
													ThumbnailOverlays []struct {
														ThumbnailOverlayTimeStatusRenderer struct {
															Text struct {
																Accessibility struct {
																	AccessibilityData struct {
																		Label string `json:"label"`
																	} `json:"accessibilityData"`
																} `json:"accessibility"`
																SimpleText string `json:"simpleText"`
															} `json:"text"`
															Style string `json:"style"`
														} `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`
														ThumbnailOverlayNowPlayingRenderer struct {
															Text struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"text"`
														} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
													} `json:"thumbnailOverlays"`
													VideoInfo struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"videoInfo"`
												} `json:"playlistVideoRenderer"`
											} `json:"contents"`
											PlaylistID     string `json:"playlistId"`
											IsEditable     bool   `json:"isEditable"`
											CanReorder     bool   `json:"canReorder"`
											TrackingParams string `json:"trackingParams"`
											TargetID       string `json:"targetId"`
										} `json:"playlistVideoListRenderer"`
									} `json:"contents"`
									TrackingParams string `json:"trackingParams"`
								} `json:"itemSectionRenderer,omitempty"`
								ContinuationItemRenderer struct {
									Trigger              string `json:"trigger"`
									ContinuationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												SendPost bool   `json:"sendPost"`
												APIURL   string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										ContinuationCommand struct {
											Token   string `json:"token"`
											Request string `json:"request"`
										} `json:"continuationCommand"`
									} `json:"continuationEndpoint"`
								} `json:"continuationItemRenderer,omitempty"`
							} `json:"contents"`
							TrackingParams string `json:"trackingParams"`
							TargetID       string `json:"targetId"`
						} `json:"sectionListRenderer"`
					} `json:"content"`
					TrackingParams string `json:"trackingParams"`
				} `json:"tabRenderer"`
			} `json:"tabs"`
		} `json:"twoColumnBrowseResultsRenderer"`
	} `json:"contents"`
	Header struct {
		PageHeaderRenderer struct {
			PageTitle string `json:"pageTitle"`
			Content   struct {
				PageHeaderViewModel struct {
					Title struct {
						DynamicTextViewModel struct {
							Text struct {
								Content string `json:"content"`
							} `json:"text"`
							RendererContext struct {
								LoggingContext struct {
									LoggingDirectives struct {
										TrackingParams string `json:"trackingParams"`
										Visibility     struct {
											Types string `json:"types"`
										} `json:"visibility"`
										ClientVeSpec struct {
											UIType    int `json:"uiType"`
											VeCounter int `json:"veCounter"`
										} `json:"clientVeSpec"`
									} `json:"loggingDirectives"`
								} `json:"loggingContext"`
							} `json:"rendererContext"`
						} `json:"dynamicTextViewModel"`
					} `json:"title"`
					Metadata struct {
						ContentMetadataViewModel struct {
							MetadataRows []struct {
								MetadataParts []struct {
									AvatarStack struct {
										AvatarStackViewModel struct {
											Avatars []struct {
												AvatarViewModel struct {
													Image struct {
														Sources []struct {
															URL    string `json:"url"`
															Width  int    `json:"width"`
															Height int    `json:"height"`
														} `json:"sources"`
														Processor struct {
															BorderImageProcessor struct {
																Circular bool `json:"circular"`
															} `json:"borderImageProcessor"`
														} `json:"processor"`
													} `json:"image"`
													AvatarImageSize string `json:"avatarImageSize"`
												} `json:"avatarViewModel"`
											} `json:"avatars"`
											Text struct {
												Content     string `json:"content"`
												CommandRuns []struct {
													StartIndex int `json:"startIndex"`
													Length     int `json:"length"`
													OnTap      struct {
														InnertubeCommand struct {
															ClickTrackingParams string `json:"clickTrackingParams"`
															CommandMetadata     struct {
																WebCommandMetadata struct {
																	URL         string `json:"url"`
																	WebPageType string `json:"webPageType"`
																	RootVe      int    `json:"rootVe"`
																	APIURL      string `json:"apiUrl"`
																} `json:"webCommandMetadata"`
															} `json:"commandMetadata"`
															BrowseEndpoint struct {
																BrowseID         string `json:"browseId"`
																CanonicalBaseURL string `json:"canonicalBaseUrl"`
															} `json:"browseEndpoint"`
														} `json:"innertubeCommand"`
													} `json:"onTap"`
												} `json:"commandRuns"`
												StyleRuns []struct {
													StartIndex  int    `json:"startIndex"`
													Length      int    `json:"length"`
													FontColor   int64  `json:"fontColor"`
													WeightLabel string `json:"weightLabel"`
												} `json:"styleRuns"`
											} `json:"text"`
											RendererContext struct {
												LoggingContext struct {
													LoggingDirectives struct {
														TrackingParams string `json:"trackingParams"`
														Visibility     struct {
															Types string `json:"types"`
														} `json:"visibility"`
														ClientVeSpec struct {
															UIType    int `json:"uiType"`
															VeCounter int `json:"veCounter"`
														} `json:"clientVeSpec"`
													} `json:"loggingDirectives"`
												} `json:"loggingContext"`
												AccessibilityContext struct {
													Label string `json:"label"`
												} `json:"accessibilityContext"`
												CommandContext struct {
													OnTap struct {
														InnertubeCommand struct {
															ClickTrackingParams string `json:"clickTrackingParams"`
															CommandMetadata     struct {
																WebCommandMetadata struct {
																	URL         string `json:"url"`
																	WebPageType string `json:"webPageType"`
																	RootVe      int    `json:"rootVe"`
																	APIURL      string `json:"apiUrl"`
																} `json:"webCommandMetadata"`
															} `json:"commandMetadata"`
															BrowseEndpoint struct {
																BrowseID         string `json:"browseId"`
																CanonicalBaseURL string `json:"canonicalBaseUrl"`
															} `json:"browseEndpoint"`
														} `json:"innertubeCommand"`
													} `json:"onTap"`
												} `json:"commandContext"`
											} `json:"rendererContext"`
										} `json:"avatarStackViewModel"`
									} `json:"avatarStack"`
								} `json:"metadataParts"`
							} `json:"metadataRows"`
							Delimiter       string `json:"delimiter"`
							RendererContext struct {
								LoggingContext struct {
									LoggingDirectives struct {
										TrackingParams string `json:"trackingParams"`
										Visibility     struct {
											Types string `json:"types"`
										} `json:"visibility"`
										ClientVeSpec struct {
											UIType    int `json:"uiType"`
											VeCounter int `json:"veCounter"`
										} `json:"clientVeSpec"`
									} `json:"loggingDirectives"`
								} `json:"loggingContext"`
							} `json:"rendererContext"`
						} `json:"contentMetadataViewModel"`
					} `json:"metadata"`
					Actions struct {
						FlexibleActionsViewModel struct {
							ActionsRows []struct {
								Actions []struct {
									ButtonViewModel struct {
										IconName string `json:"iconName"`
										Title    string `json:"title"`
										OnTap    struct {
											InnertubeCommand struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														URL         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												WatchEndpoint struct {
													VideoID        string `json:"videoId"`
													PlaylistID     string `json:"playlistId"`
													PlayerParams   string `json:"playerParams"`
													LoggingContext struct {
														VssLoggingContext struct {
															SerializedContextData string `json:"serializedContextData"`
														} `json:"vssLoggingContext"`
													} `json:"loggingContext"`
													WatchEndpointSupportedOnesieConfig struct {
														HTML5PlaybackOnesieConfig struct {
															CommonConfig struct {
																URL string `json:"url"`
															} `json:"commonConfig"`
														} `json:"html5PlaybackOnesieConfig"`
													} `json:"watchEndpointSupportedOnesieConfig"`
												} `json:"watchEndpoint"`
											} `json:"innertubeCommand"`
										} `json:"onTap"`
										AccessibilityText string `json:"accessibilityText"`
										Style             string `json:"style"`
										TrackingParams    string `json:"trackingParams"`
										IsFullWidth       bool   `json:"isFullWidth"`
										Type              string `json:"type"`
										ButtonSize        string `json:"buttonSize"`
									} `json:"buttonViewModel,omitempty"`
									ToggleButtonViewModel struct {
										DefaultButtonViewModel struct {
											ButtonViewModel struct {
												IconName string `json:"iconName"`
												OnTap    struct {
													InnertubeCommand struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																IgnoreNavigation bool `json:"ignoreNavigation"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														ModalEndpoint struct {
															Modal struct {
																ModalWithTitleAndButtonRenderer struct {
																	Title struct {
																		SimpleText string `json:"simpleText"`
																	} `json:"title"`
																	Content struct {
																		SimpleText string `json:"simpleText"`
																	} `json:"content"`
																	Button struct {
																		ButtonRenderer struct {
																			Style      string `json:"style"`
																			Size       string `json:"size"`
																			IsDisabled bool   `json:"isDisabled"`
																			Text       struct {
																				SimpleText string `json:"simpleText"`
																			} `json:"text"`
																			NavigationEndpoint struct {
																				ClickTrackingParams string `json:"clickTrackingParams"`
																				CommandMetadata     struct {
																					WebCommandMetadata struct {
																						URL         string `json:"url"`
																						WebPageType string `json:"webPageType"`
																						RootVe      int    `json:"rootVe"`
																					} `json:"webCommandMetadata"`
																				} `json:"commandMetadata"`
																				SignInEndpoint struct {
																					NextEndpoint struct {
																						ClickTrackingParams string `json:"clickTrackingParams"`
																						CommandMetadata     struct {
																							WebCommandMetadata struct {
																								URL         string `json:"url"`
																								WebPageType string `json:"webPageType"`
																								RootVe      int    `json:"rootVe"`
																								APIURL      string `json:"apiUrl"`
																							} `json:"webCommandMetadata"`
																						} `json:"commandMetadata"`
																						BrowseEndpoint struct {
																							BrowseID string `json:"browseId"`
																						} `json:"browseEndpoint"`
																					} `json:"nextEndpoint"`
																					IdamTag string `json:"idamTag"`
																				} `json:"signInEndpoint"`
																			} `json:"navigationEndpoint"`
																			TrackingParams string `json:"trackingParams"`
																		} `json:"buttonRenderer"`
																	} `json:"button"`
																} `json:"modalWithTitleAndButtonRenderer"`
															} `json:"modal"`
														} `json:"modalEndpoint"`
													} `json:"innertubeCommand"`
												} `json:"onTap"`
												AccessibilityText string `json:"accessibilityText"`
												Style             string `json:"style"`
												TrackingParams    string `json:"trackingParams"`
												IsFullWidth       bool   `json:"isFullWidth"`
												Type              string `json:"type"`
												ButtonSize        string `json:"buttonSize"`
												Tooltip           string `json:"tooltip"`
											} `json:"buttonViewModel"`
										} `json:"defaultButtonViewModel"`
										ToggledButtonViewModel struct {
											ButtonViewModel struct {
												IconName          string `json:"iconName"`
												AccessibilityText string `json:"accessibilityText"`
												Style             string `json:"style"`
												TrackingParams    string `json:"trackingParams"`
												IsFullWidth       bool   `json:"isFullWidth"`
												Type              string `json:"type"`
												ButtonSize        string `json:"buttonSize"`
												Tooltip           string `json:"tooltip"`
											} `json:"buttonViewModel"`
										} `json:"toggledButtonViewModel"`
										IsToggled      bool   `json:"isToggled"`
										Identifier     string `json:"identifier"`
										TrackingParams string `json:"trackingParams"`
									} `json:"toggleButtonViewModel,omitempty"`
								} `json:"actions"`
							} `json:"actionsRows"`
							MinimumRowHeight int `json:"minimumRowHeight"`
							RendererContext  struct {
								LoggingContext struct {
									LoggingDirectives struct {
										TrackingParams string `json:"trackingParams"`
										Visibility     struct {
											Types string `json:"types"`
										} `json:"visibility"`
										ClientVeSpec struct {
											UIType    int `json:"uiType"`
											VeCounter int `json:"veCounter"`
										} `json:"clientVeSpec"`
									} `json:"loggingDirectives"`
								} `json:"loggingContext"`
							} `json:"rendererContext"`
						} `json:"flexibleActionsViewModel"`
					} `json:"actions"`
					Description struct {
						DescriptionPreviewViewModel struct {
							Description struct {
								Content string `json:"content"`
							} `json:"description"`
							TruncationText struct {
								Content   string `json:"content"`
								StyleRuns []struct {
									StartIndex int `json:"startIndex"`
									Length     int `json:"length"`
									Weight     int `json:"weight"`
								} `json:"styleRuns"`
							} `json:"truncationText"`
							RendererContext struct {
								LoggingContext struct {
									LoggingDirectives struct {
										TrackingParams string `json:"trackingParams"`
										Visibility     struct {
											Types string `json:"types"`
										} `json:"visibility"`
										ClientVeSpec struct {
											UIType    int `json:"uiType"`
											VeCounter int `json:"veCounter"`
										} `json:"clientVeSpec"`
									} `json:"loggingDirectives"`
								} `json:"loggingContext"`
								AccessibilityContext struct {
									Label string `json:"label"`
								} `json:"accessibilityContext"`
								CommandContext struct {
									OnTap struct {
										InnertubeCommand struct {
											ClickTrackingParams         string `json:"clickTrackingParams"`
											ShowEngagementPanelEndpoint struct {
												EngagementPanel struct {
													EngagementPanelSectionListRenderer struct {
														Header struct {
															EngagementPanelTitleHeaderRenderer struct {
																Title struct {
																	Runs []struct {
																		Text string `json:"text"`
																	} `json:"runs"`
																} `json:"title"`
																VisibilityButton struct {
																	ButtonRenderer struct {
																		Icon struct {
																			IconType string `json:"iconType"`
																		} `json:"icon"`
																		TrackingParams    string `json:"trackingParams"`
																		AccessibilityData struct {
																			AccessibilityData struct {
																				Label string `json:"label"`
																			} `json:"accessibilityData"`
																		} `json:"accessibilityData"`
																		Command struct {
																			ClickTrackingParams                   string `json:"clickTrackingParams"`
																			ChangeEngagementPanelVisibilityAction struct {
																				TargetID   string `json:"targetId"`
																				Visibility string `json:"visibility"`
																			} `json:"changeEngagementPanelVisibilityAction"`
																		} `json:"command"`
																	} `json:"buttonRenderer"`
																} `json:"visibilityButton"`
																TrackingParams string `json:"trackingParams"`
															} `json:"engagementPanelTitleHeaderRenderer"`
														} `json:"header"`
														Content struct {
															SectionListRenderer struct {
																Contents []struct {
																	ItemSectionRenderer struct {
																		Contents []struct {
																			MessageRenderer struct {
																				Text struct {
																					Runs []struct {
																						Text string `json:"text"`
																					} `json:"runs"`
																				} `json:"text"`
																				TrackingParams string `json:"trackingParams"`
																				Style          struct {
																					Value string `json:"value"`
																				} `json:"style"`
																			} `json:"messageRenderer"`
																		} `json:"contents"`
																		TrackingParams string `json:"trackingParams"`
																	} `json:"itemSectionRenderer"`
																} `json:"contents"`
																TrackingParams string `json:"trackingParams"`
															} `json:"sectionListRenderer"`
														} `json:"content"`
														DisablePullRefresh bool   `json:"disablePullRefresh"`
														TargetID           string `json:"targetId"`
														Identifier         struct {
															Surface string `json:"surface"`
															Tag     string `json:"tag"`
														} `json:"identifier"`
														OutsideScrimTapBehavior string `json:"outsideScrimTapBehavior"`
													} `json:"engagementPanelSectionListRenderer"`
												} `json:"engagementPanel"`
												Identifier struct {
													Surface string `json:"surface"`
													Tag     string `json:"tag"`
												} `json:"identifier"`
												EngagementPanelPresentationConfigs struct {
													EngagementPanelPopupPresentationConfig struct {
														PopupType string `json:"popupType"`
													} `json:"engagementPanelPopupPresentationConfig"`
												} `json:"engagementPanelPresentationConfigs"`
											} `json:"showEngagementPanelEndpoint"`
										} `json:"innertubeCommand"`
									} `json:"onTap"`
								} `json:"commandContext"`
							} `json:"rendererContext"`
						} `json:"descriptionPreviewViewModel"`
					} `json:"description"`
					HeroImage struct {
						ContentPreviewImageViewModel struct {
							Image struct {
								Sources []struct {
									URL    string `json:"url"`
									Width  int    `json:"width"`
									Height int    `json:"height"`
								} `json:"sources"`
							} `json:"image"`
							Style      string `json:"style"`
							LayoutMode string `json:"layoutMode"`
							Overlays   []struct {
								ThumbnailHoverOverlayViewModel struct {
									Icon struct {
										Sources []struct {
											ClientResource struct {
												ImageName string `json:"imageName"`
											} `json:"clientResource"`
										} `json:"sources"`
									} `json:"icon"`
									Text struct {
										Content   string `json:"content"`
										StyleRuns []struct {
											StartIndex int `json:"startIndex"`
											Length     int `json:"length"`
										} `json:"styleRuns"`
									} `json:"text"`
									Style           string `json:"style"`
									RendererContext struct {
										CommandContext struct {
											OnTap struct {
												InnertubeCommand struct {
													ClickTrackingParams string `json:"clickTrackingParams"`
													CommandMetadata     struct {
														WebCommandMetadata struct {
															URL         string `json:"url"`
															WebPageType string `json:"webPageType"`
															RootVe      int    `json:"rootVe"`
														} `json:"webCommandMetadata"`
													} `json:"commandMetadata"`
													WatchEndpoint struct {
														VideoID        string `json:"videoId"`
														PlaylistID     string `json:"playlistId"`
														PlayerParams   string `json:"playerParams"`
														LoggingContext struct {
															VssLoggingContext struct {
																SerializedContextData string `json:"serializedContextData"`
															} `json:"vssLoggingContext"`
														} `json:"loggingContext"`
														WatchEndpointSupportedOnesieConfig struct {
															HTML5PlaybackOnesieConfig struct {
																CommonConfig struct {
																	URL string `json:"url"`
																} `json:"commonConfig"`
															} `json:"html5PlaybackOnesieConfig"`
														} `json:"watchEndpointSupportedOnesieConfig"`
													} `json:"watchEndpoint"`
												} `json:"innertubeCommand"`
											} `json:"onTap"`
										} `json:"commandContext"`
									} `json:"rendererContext"`
								} `json:"thumbnailHoverOverlayViewModel"`
							} `json:"overlays"`
							RendererContext struct {
								LoggingContext struct {
									LoggingDirectives struct {
										TrackingParams string `json:"trackingParams"`
										Visibility     struct {
											Types string `json:"types"`
										} `json:"visibility"`
										ClientVeSpec struct {
											UIType    int `json:"uiType"`
											VeCounter int `json:"veCounter"`
										} `json:"clientVeSpec"`
									} `json:"loggingDirectives"`
								} `json:"loggingContext"`
								AccessibilityContext struct {
									Label string `json:"label"`
								} `json:"accessibilityContext"`
							} `json:"rendererContext"`
						} `json:"contentPreviewImageViewModel"`
					} `json:"heroImage"`
					Background struct {
						CinematicContainerViewModel struct {
							BackgroundImageConfig struct {
								Image struct {
									Sources []struct {
										URL    string `json:"url"`
										Width  int    `json:"width"`
										Height int    `json:"height"`
									} `json:"sources"`
								} `json:"image"`
							} `json:"backgroundImageConfig"`
							GradientColorConfig []struct {
								LightThemeColor int64 `json:"lightThemeColor"`
								DarkThemeColor  int64 `json:"darkThemeColor"`
							} `json:"gradientColorConfig"`
							Config struct {
								LightThemeBackgroundColor int64 `json:"lightThemeBackgroundColor"`
								DarkThemeBackgroundColor  int64 `json:"darkThemeBackgroundColor"`
								ColorSourceSizeMultiplier int   `json:"colorSourceSizeMultiplier"`
								ApplyClientImageBlur      bool  `json:"applyClientImageBlur"`
							} `json:"config"`
						} `json:"cinematicContainerViewModel"`
					} `json:"background"`
					HasTopbarAnimation                  bool `json:"hasTopbarAnimation"`
					EnableFlexibleActionsButtonsWrapper bool `json:"enableFlexibleActionsButtonsWrapper"`
					RendererContext                     struct {
						LoggingContext struct {
							LoggingDirectives struct {
								TrackingParams string `json:"trackingParams"`
								Visibility     struct {
									Types string `json:"types"`
								} `json:"visibility"`
								ClientVeSpec struct {
									UIType    int `json:"uiType"`
									VeCounter int `json:"veCounter"`
								} `json:"clientVeSpec"`
							} `json:"loggingDirectives"`
						} `json:"loggingContext"`
					} `json:"rendererContext"`
				} `json:"pageHeaderViewModel"`
			} `json:"content"`
			EnableSidebarView bool `json:"enableSidebarView"`
		} `json:"pageHeaderRenderer"`
	} `json:"header"`
	OnResponseReceivedActions* []struct {
		ClickTrackingParams           string `json:"clickTrackingParams"`
		AppendContinuationItemsAction struct {
			ContinuationItems []struct {
				PlaylistVideoRenderer* struct {
					VideoID   string `json:"videoId"`
					Thumbnail struct {
						Thumbnails []struct {
							URL    string `json:"url"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
						} `json:"thumbnails"`
					} `json:"thumbnail"`
					Title struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
						Accessibility struct {
							AccessibilityData struct {
								Label string `json:"label"`
							} `json:"accessibilityData"`
						} `json:"accessibility"`
					} `json:"title"`
					Index struct {
						SimpleText string `json:"simpleText"`
					} `json:"index"`
					ShortBylineText struct {
						Runs []struct {
							Text               string `json:"text"`
							NavigationEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										URL         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
										APIURL      string `json:"apiUrl"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								BrowseEndpoint struct {
									BrowseID         string `json:"browseId"`
									CanonicalBaseURL string `json:"canonicalBaseUrl"`
								} `json:"browseEndpoint"`
							} `json:"navigationEndpoint"`
						} `json:"runs"`
					} `json:"shortBylineText"`
					LengthText struct {
						Accessibility struct {
							AccessibilityData struct {
								Label string `json:"label"`
							} `json:"accessibilityData"`
						} `json:"accessibility"`
						SimpleText string `json:"simpleText"`
					} `json:"lengthText"`
					NavigationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								URL         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						WatchEndpoint struct {
							VideoID        string `json:"videoId"`
							PlaylistID     string `json:"playlistId"`
							Index          int    `json:"index"`
							Params         string `json:"params"`
							PlayerParams   string `json:"playerParams"`
							LoggingContext struct {
								VssLoggingContext struct {
									SerializedContextData string `json:"serializedContextData"`
								} `json:"vssLoggingContext"`
							} `json:"loggingContext"`
							WatchEndpointSupportedOnesieConfig struct {
								HTML5PlaybackOnesieConfig struct {
									CommonConfig struct {
										URL string `json:"url"`
									} `json:"commonConfig"`
								} `json:"html5PlaybackOnesieConfig"`
							} `json:"watchEndpointSupportedOnesieConfig"`
						} `json:"watchEndpoint"`
					} `json:"navigationEndpoint"`
					LengthSeconds  string `json:"lengthSeconds"`
					TrackingParams string `json:"trackingParams"`
					IsPlayable     bool   `json:"isPlayable"`
					Menu           struct {
						MenuRenderer struct {
							Items []struct {
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
										CommandMetadata     struct {
											WebCommandMetadata struct {
												SendPost bool `json:"sendPost"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										SignalServiceEndpoint struct {
											Signal  string `json:"signal"`
											Actions []struct {
												ClickTrackingParams  string `json:"clickTrackingParams"`
												AddToPlaylistCommand struct {
													OpenMiniplayer      bool   `json:"openMiniplayer"`
													VideoID             string `json:"videoId"`
													ListType            string `json:"listType"`
													OnCreateListCommand struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																SendPost bool   `json:"sendPost"`
																APIURL   string `json:"apiUrl"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														CreatePlaylistServiceEndpoint struct {
															VideoIds []string `json:"videoIds"`
															Params   string   `json:"params"`
														} `json:"createPlaylistServiceEndpoint"`
													} `json:"onCreateListCommand"`
													VideoIds []string `json:"videoIds"`
												} `json:"addToPlaylistCommand"`
											} `json:"actions"`
										} `json:"signalServiceEndpoint"`
									} `json:"serviceEndpoint"`
									TrackingParams string `json:"trackingParams"`
								} `json:"menuServiceItemRenderer"`
							} `json:"items"`
							TrackingParams string `json:"trackingParams"`
							Accessibility  struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibility"`
						} `json:"menuRenderer"`
					} `json:"menu"`
					ThumbnailOverlays []struct {
						ThumbnailOverlayTimeStatusRenderer struct {
							Text struct {
								Accessibility struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
								SimpleText string `json:"simpleText"`
							} `json:"text"`
							Style string `json:"style"`
						} `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`
						ThumbnailOverlayNowPlayingRenderer struct {
							Text struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"text"`
						} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
					} `json:"thumbnailOverlays"`
					VideoInfo struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"videoInfo"`
				} `json:"playlistVideoRenderer"`
			} `json:"continuationItems"`
			TargetID string `json:"targetId"`
		} `json:"appendContinuationItemsAction"`
	} `json:"onResponseReceivedActions"`
}
