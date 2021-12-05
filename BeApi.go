package BEAPI

import (
  "os"
  "io"
  "fmt"
  "bytes"
  "strings"
  "mime/multipart"
  neturl "net/url"
  
  "github.com/valyala/fasthttp"
)

type BeAPIClient struct{
  Host string
  Url string
}

var DefaultClient = &BeAPIClient{Host: "https://beta.beapi.me", Url: "https://beta.beapi.me"}

func (self *BeAPIClient) AddParams(endpoint string, param map[string]string) {
  p := neturl.Values{}
  for k, v := range param{
    p.Add(k, v)
  }
  self.Url = self.Host + endpoint + "?" + p.Encode()
  return
}

func (self *BeAPIClient) returnErr(err error) string{
  self.Url = self.Host
  return err.Error()
}

func (self *BeAPIClient) Request(method string, args map[string]interface{}) string{
  req := fasthttp.AcquireRequest()
  defer fasthttp.ReleaseRequest(req)
  req.Header.SetMethod(method)
  req.SetRequestURI(self.Url)
  if method == "POST"{
    var b bytes.Buffer
    w := multipart.NewWriter(&b)
    for k, v := range args{
      var fw io.Writer
      var err error
      var r io.Reader
      switch v.(type){
        case string:
          if fw, err = w.CreateFormField(k); err != nil{
            return self.returnErr(err)
          }
          r = strings.NewReader(v.(string))
        case *os.File:
          if x, ok := v.(io.Closer); ok {
            defer x.Close()
          }
          x, _ := v.(*os.File)
          if fw, err = w.CreateFormFile(k, x.Name()); err != nil{
            return self.returnErr(err)
          }
          r = v.(io.Reader)
      }
      if _, err = io.Copy(fw, r); err != nil {
        return self.returnErr(err)
      }
    }
    w.Close()
    req.SetBodyStream(&b, -1)
    req.Header.SetContentType(w.FormDataContentType())
  }
  resp := fasthttp.AcquireResponse()
  if err := fasthttp.Do(req, resp); err != nil {
    return self.returnErr(err)
  }
  defer fasthttp.ReleaseResponse(resp)
  return string(resp.Body())
}

func (self *BeAPIClient) Get() string{
  return self.Request("GET", nil)
}

func (self *BeAPIClient) Post(args map[string]interface{}) string{
  return self.Request("POST", args)
}

func (self *BeAPIClient) AlphaCoders(args ...string) string{
  var search, page string
  la := len(args)
  if la != 0{
    search, page = args[0], "1"
    if la > 1{ page = args[1] }
  } else { return "Args = search(must), page(optional)" }
  self.AddParams("/alphacoders", map[string]string{
    "search": search,
    "page": page,
  })
  return self.Get()
}

func (self *BeAPIClient) OnGoingAnime() string{
  self.Url += "/animeongoing"
  return self.Get()
}

func (self *BeAPIClient) AnimeXin() string{
  self.Url += "/animexin"
  return self.Get()
}

func (self *BeAPIClient) AuthKey2Primary(authKey string) string{
  self.AddParams("/authkey2primary", map[string]string{"authkey":authKey})
  return self.Get()
}

func (self *BeAPIClient) BrainlySearch(search string) string{
  self.AddParams("/brainly", map[string]string{"search":search})
  return self.Get()
}

func (self *BeAPIClient) GIFSearch(search string) string{
  self.AddParams("/gifsearch", map[string]string{"search":search})
  return self.Get()
}

func (self *BeAPIClient) GoogleImg(search string) string{
  self.AddParams("/googleimg", map[string]string{"search":search})
  return self.Get()
}

func (self *BeAPIClient) GoogleSearch(search string) string{
  self.AddParams("/googlesearch", map[string]string{"search":search})
  return self.Get()
}

func (self *BeAPIClient) GoogleTranslate(lang, text string) string{
  self.AddParams("/googletrans", map[string]string{"lang":lang, "text": text})
  return self.Get()
}

func (self *BeAPIClient) GoogleImgReverse(url string) string{
  self.AddParams("/imgreverse", map[string]string{"url":url})
  return self.Get()
}

func (self *BeAPIClient) LanguageList() string{
  self.Url += "/language"
  return self.Get()
}

func (self *BeAPIClient) IgPost(url string) string{
  self.AddParams("/igpost", map[string]string{"url":url})
  return self.Get()
}

func (self *BeAPIClient) IgUser(username string) string{
  self.AddParams("/igprofile", map[string]string{"user":username})
  return self.Get()
}

func (self *BeAPIClient) JooxSearch(search string) string{
  self.AddParams("/joox", map[string]string{"search":search})
  return self.Get()
}

func (self *BeAPIClient) JooxId(id string) string{
  self.AddParams("/joox", map[string]string{"id":id})
  return self.Get()
}

func (self *BeAPIClient) KBBI(search string) string{
  self.AddParams("/kbbi", map[string]string{"search":search})
  return self.Get()
}

func (self *BeAPIClient) LineAppName() string{
  self.Url += "/lineappname"
  return self.Get()
}

var OsNameList = []string{"android","ios","androidlite","chromeos","desktopmac","desktopwin","iosipad"}
func (self *BeAPIClient) RandomLineAppName(osname string) string{
  self.AddParams("/lineappname_random", map[string]string{"osname":osname})
  return self.Get()
}

func (self *BeAPIClient) PrimaryToSecondary(appName, authToken string) string{
  self.AddParams("/lineprimary2secondary", map[string]string{"appname":appName, "authtoken":authToken})
  return self.Get()
}

func (self *BeAPIClient) LineGetQR(appName string, certs ...string) string{
  params := make(map[string]string)
  params["appname"] = appName
  if len(certs) == 1{
    params["cert"] = certs[0]
  }
  self.AddParams("/lineqr", params)
  return self.Get()
}

func (self *BeAPIClient) LineGetQRPincode(sess string) string{
  self.Url += "/lineqr/pincode/"+sess
  return self.Get()
}

func (self *BeAPIClient) LineGetQRAuth(sess string) string{
  self.Url += "/lineqr/auth/"+sess
  return self.Get()
}

var NineGagCategoryList = []string{"funny", "among-us", "animals", "anime-manga", "animewaifu", "animewallpaper", "apexlegends", "ask9gag", "awesome", "car", "comic-webtoon", "coronavirus", "cosplay", "countryballs", "home-living", "crappydesign", "cyberpunk2077", "drawing-diy-crafts", "rate-my-outfit", "food-drinks", "football", "fortnite", "got", "gaming", "gif", "girl", "girlcelebrity", "guy", "history", "horror", "kpop", "timely", "leagueoflegends", "lego", "superhero", "meme", "movie-tv", "music", "basketball", "nsfw", "overwatch", "pcmr", "pokemon", "politics", "pubg", "random", "relationship", "savage", "satisfying", "science-tech", "sport", "starwars", "school", "travel-photography", "video", "wallpaper", "warhammer", "wholesome", "wtf", "darkhumor", "funny", "nsfw", "girl", "wtf", "anime-manga", "random", "animals", "animewaifu", "awesome", "car", "comic-webtoon", "cosplay", "cyberpunk2077", "gaming", "gif", "girlcelebrity", "leagueoflegends", "meme", "politics", "relationship", "savage", "video", "algeria", "argentina", "australia", "austria", "bosniaherzegovina", "bahrain", "belgium", "bolivia", "brazil", "bulgaria", "canada", "chile", "colombia", "costarica", "croatia", "cyprus", "czechia", "denmark", "dominicanrepublic", "ecuador", "egypt", "estonia", "finland", "france", "georgia", "germany", "ghana", "greece", "guatemala", "hongkong", "hungary", "iceland", "india", "indonesia", "iraq", "ireland", "israel", "italy", "japan", "jordan", "kenya", "kuwait", "latvia", "lebanon", "lithuania", "luxembourg", "malaysia", "mexico", "montenegro", "morocco", "nepal", "netherlands", "newzealand", "nigeria", "norway", "oman", "pakistan", "peru", "philippines", "poland", "portugal", "puertorico", "qatar", "romania", "russia", "saudiarabia", "senegal", "serbia", "singapore", "slovakia", "slovenia", "southafrica", "southkorea", "spain", "srilanka", "sweden", "switzerland", "taiwan", "tanzania", "thailand", "tunisia", "turkey", "uae", "usa", "ukraine", "uk", "uruguay", "vietnam", "yemen", "zimbabwe"}
func (self *BeAPIClient) NineGagFresh(ctg string) string{
  self.AddParams("/9gag-fresh", map[string]string{"category":ctg})
  return self.Get()
}

func (self *BeAPIClient) NineGagHot(ctg string) string{
  self.AddParams("/9gag-hot", map[string]string{"category":ctg})
  return self.Get()
}

func (self *BeAPIClient) OneCakRandom() string{
  self.Url += "/onecak"
  return self.Get()
}

func (self *BeAPIClient) PhotoFunia(params map[string]string) string{
  self.AddParams("/photofunia", params)
  return self.Get()
}

func (self *BeAPIClient) Reface(params map[string]string) string{
  self.AddParams("/reface", params)
  return self.Get()
}

var SimSimiLang = []string{"af", "al*", "ar", "hy", "az", "eu", "be", "bn", "bs", "bg", "ca", "cx*", "ch*", "hr", "cs", "da", "nl", "en", "et", "ph*", "fi", "fr", "fy", "gl", "ka", "de", "el", "gu", "he", "hi", "hu", "is", "id", "it", "ja", "kn", "kk", "kh*", "ko", "ku", "lv", "lt", "mk", "ms", "ml", "mr", "mn", "my", "ne", "nb", "as", "br", "gn", "jv", "or", "rw", "zh*", "ps", "fa", "pl", "pt", "pa", "ro", "ru", "rs*", "si", "sk", "sl", "es", "sw", "sv", "tg", "ta", "te", "th", "tr", "uk", "ur", "uz", "vn*", "cy"}
func (self *BeAPIClient) SimSimi(lang, text string) string{
  self.AddParams("/simisimi", map[string]string{"lang":lang, "text": text})
  return self.Get()
}

func (self *BeAPIClient) SmulePost(url string) string{
  self.AddParams("/smule/post", map[string]string{"url":url})
  return self.Get()
}

func (self *BeAPIClient) SmuleUser(u string) string{
  self.AddParams("/smule/user", map[string]string{"user":u})
  return self.Get()
}

func (self *BeAPIClient) SmulePerformance(u string) string{
  self.AddParams("/smule/performance", map[string]string{"user":u})
  return self.Get()
}

func (self *BeAPIClient) FileUpload(path string) string{
  file, err := os.Open(path)
  if err != nil{
    return err.Error()
  }
  m := map[string]interface{}{
    "file": file,
  }
  self.Url += "/storage"
  fmt.Println(self.Url)
  return self.Post(m)
}

func (self *BeAPIClient) ShortLink(url string) string{
  self.Url += "/short-link"
  return self.Post(map[string]interface{}{"url": url})
}

func (self *BeAPIClient) WebScreenshot(url string) string{
  self.Url += "/ss-web"
  return self.Post(map[string]interface{}{"url": url})
}

func (self *BeAPIClient) TextPro(params map[string]string) string{
  self.AddParams("/textpro", params)
  return self.Get()
}

var CourierList = []string{"pos", "wahana", "jnt", "sap", "sicepat", "jet", "dse", "first", "ninja", "lion", "idl", "rex", "ide", "sentral"}
func (self *BeAPIClient) ResiTracking(resi, courier string) string{
  self.AddParams("/track-resi", map[string]string{"resi":resi, "courier": courier})
  return self.Get()
}

func (self *BeAPIClient) TiktokPost(url string) string{
  self.AddParams("/tiktok", map[string]string{"url":url})
  return self.Get()
}

func (self *BeAPIClient) TiktokPostV2(url string) string{
  self.AddParams("/musicallydown", map[string]string{"url":url})
  return self.Get()
}

func (self *BeAPIClient) TiktokUser(u string) string{
  self.AddParams("/tiktok", map[string]string{"user":u})
  return self.Get()
}

func (self *BeAPIClient) YoutubeSearch(search string) string{
  self.AddParams("/youtube", map[string]string{"search":search})
  return self.Get()
}

func (self *BeAPIClient) YoutubeDownload(url string) string{
  self.AddParams("/youtube", map[string]string{"url":url})
  return self.Get()
}