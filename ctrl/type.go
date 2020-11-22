package ctrl

type (
	StrList struct {
		List []string `json:"list"`
	}

	ObjList struct {
		List []interface{} `json:"list"`
	}

	AuthData struct {
		Exp int
		Iat int
		Uid string
		Tk  string // 验证admin
		Gid string
		Wid string
	}

	ResData struct {
		Data string `json:"data"`
		Path string `json:"path,omitempty"`
	}

	// Vcode struct {
	// 	Value string
	// 	Ip    string
	// 	Exp   int
	// }

	WSData struct {
		Path string `json:"path"`
		Sign string `json:"sign"`
		Data string `json:"data"`
	}

	SearchData struct {
		Topic []interface{} `json:"topic"`
		Group []interface{} `json:"group"`
	}
)
