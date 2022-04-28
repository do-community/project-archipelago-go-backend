package entities

type Profile struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Bio      string `json:"bio"`
}

type Links struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	Icon  string `json:"icon"`
}

type QR struct {
	Data string `json"data"`
}

// Data Model
// {
// 	“Username” : “chris”,
// 	“Avatar” : “URL”
// 	“Bio” : “I’m chris and I’m awesome”,
// 	“Links” : [
// 	  {
// 	  Title: “my awesome site”,
// 			  URL: “https://chris.com”,
// 			  Icon: “https://kasjdfksjdfksa”
//   },
//   {
// 	  Title: “my awesome site”,
// 			  URL: “https://chris.com”,
// 			  Icon: “https://kasjdfksjdfksa”
//   },

//   ],
//    “QR” : {
// 	  Data: “kasdjkaf”,
// 	 }
//   }
