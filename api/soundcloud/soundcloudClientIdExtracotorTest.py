import re
import requests # type: ignore

# Test in python per estrarre il client_id da SoundCloud -> TODO: RIMUOVERE
# Questo script estrae il client_id da SoundCloud analizzando i file JS
# presenti nella homepage di SoundCloud.

def get_soundcloud_client_id():
    # Scarica la homepage di SoundCloud
    homepage = requests.get('https://soundcloud.com/').text

    # Trova tutti i link ai file JS
    js_files = re.findall(r'src="(https://a-v2\.sndcdn\.com/assets/[^"]+\.js)"', homepage)

    print(f"Trovati {len(js_files)} file JS, inizio la ricerca...")

    # Cerca il client_id in ogni file JS
    for js_url in js_files:
        print(f"Controllo {js_url}...")
        js_content = requests.get(js_url).text
        match = re.search(r'client_id\s*:\s*"(?P<client_id>[a-zA-Z0-9]{32})"', js_content)
        if match:
            client_id = match.group('client_id')
            print(f"Trovato client_id: {client_id}")
            return client_id

    print("Nessun client_id trovato.")
    return None

if __name__ == "__main__":
    client_id = get_soundcloud_client_id()
    if client_id:
        print(f"\n✅ Client ID: {client_id}")
    else:
        print("\n❌ Client ID non trovato.")
