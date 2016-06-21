# SoundCloud API
SoundCloudのAPIクライアント

# Usage
```
func main() {
    clientId := <your app client-id>
    clientSecret := <your app client-secret>
    api, err := soundcloud.NewClient(clientId, clientSecret)
    if err != nil {
    	log.Fatal(err.Error())
    }
    
    userId := 12345
    for _, v := range api.GetFavorites(userId) {
    	log.Printf("title: %s", v.Title)
    }
}
```
