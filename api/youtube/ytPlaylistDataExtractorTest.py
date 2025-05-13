import requests # type: ignore
import re
import json
from urllib.parse import urljoin

def get_yt_api_key_and_version():
    """Get INNERTUBE API key and client version from YouTube homepage."""
    html = requests.get("https://www.youtube.com", headers={"User-Agent": "Mozilla/5.0"}).text
    api_key = re.search(r'"INNERTUBE_API_KEY":"([^"]+)"', html).group(1)
    client_version = re.search(r'"INNERTUBE_CLIENT_VERSION":"([^"]+)"', html).group(1)
    return api_key, client_version



def get_playlist_videos_internal_api(playlist_id):
    api_key, client_version = get_yt_api_key_and_version()

    headers = {
        "User-Agent": "Mozilla/5.0",
        "Content-Type": "application/json"
    }

    payload = {
        "context": {
            "client": {
                "clientName": "WEB",
                "clientVersion": client_version
            }
        },
        "browseId": f"VL{playlist_id}"
    }

    url = f"https://www.youtube.com/youtubei/v1/browse"
    response = requests.post(url, json=payload)
    data = response.json()

    # Extract video items from the response
    try:
        contents = (
            data["contents"]["twoColumnBrowseResultsRenderer"]["tabs"][0]["tabRenderer"]
            ["content"]["sectionListRenderer"]["contents"][0]["itemSectionRenderer"]
            ["contents"][0]["playlistVideoListRenderer"]["contents"]
        )

        videos = []
        for item in contents:
            video = item.get("playlistVideoRenderer")
            if video:
                title = video["title"]["runs"][0]["text"]
                video_id = video["videoId"]
                videos.append({
                    "title": title,
                    "video_id": video_id,
                    "url": f"https://www.youtube.com/watch?v={video_id}"
                })
        return videos

    except KeyError:
        raise Exception("Failed to parse internal playlist API response")




playlist_id = "PLMmqTuUsDkRJs4S89GKN3hnmW_88Hin6R" 

# api_key, client_version = get_yt_api_key_and_version()
# print(f"API Key: {api_key}")
# print(f"Client Version: {client_version}")

get_playlist_videos_internal_api(playlist_id)
print(f"Videos in playlist {playlist_id}:")
videos = get_playlist_videos_internal_api(playlist_id)
for video in videos:
    print(f"Title: {video['title']}, URL: {video['url']}")
