// Package platforms 参考 https://github.com/beautiful-you/anniversary/wechat/tree/master/util
package platforms

import (
	"encoding/json"
	"fmt"

	"github.com/beautiful-you/anniversary/wechat/util"
)

// 接口信息
const (
	ComponentTokenURL = "https://api.weixin.qq.com/cgi-bin/component/api_component_token"
	PreAuthCodeURL    = "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token=%s"
	// AuthURL           = "https://mp.weixin.qq.com/safe/bindcomponent?action=bindcomponent&auth_type=3&no_scan=1&component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=xxx&biz_appid=xxxx#wechat_redirect"
	AuthURL = "https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=3"
)

// ResComponentAccessToken ComponentAccessToken
type ResComponentAccessToken struct {
	util.CommonError
	ComponentAccessToken string `json:"component_access_token"`
	ExpiresIn            int64  `json:"expires_in"`
}

// ResPreAuthCode 预授权码
type ResPreAuthCode struct {
	util.CommonError
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int64  `json:"expires_in"`
}

// ComponentAccessToken // 获取第三方平台 component_access_token
func ComponentAccessToken(AppID, AppSecret, cvt string) (*ResComponentAccessToken, error) {
	// 获取第三方平台 component_access_token
	body, err := util.PostJSON(ComponentTokenURL, map[string]string{"component_appid": AppID, "component_appsecret": AppSecret, "component_verify_ticket": cvt})
	if err != nil {
		return nil, err
	}
	resComponentAccessToken := new(ResComponentAccessToken)
	err = json.Unmarshal(body, &resComponentAccessToken)
	if err != nil {
		return nil, err
	}
	if resComponentAccessToken.ErrMsg != "" {
		err = fmt.Errorf("get access_token error : errcode=%v , errormsg=%v", resComponentAccessToken.ErrCode, resComponentAccessToken.ErrMsg)
		return nil, err
	}
	return resComponentAccessToken, nil
}

// PreAuthCode 获取预授权码 pre_auth_code
func PreAuthCode(AppID, ComponentAccessToken string) (*ResPreAuthCode, error) {
	// 获取预授权码 pre_auth_code
	url := fmt.Sprintf(PreAuthCodeURL, ComponentAccessToken)
	body, err := util.PostJSON(url, map[string]string{"component_appid": AppID})
	if err != nil {
		return nil, err
	}
	resPreAuthCode := new(ResPreAuthCode)
	err = json.Unmarshal(body, &resPreAuthCode)
	if err != nil {
		return nil, err
	}
	if resPreAuthCode.ErrMsg != "" {
		err = fmt.Errorf("get auth_code error : errcode=%v , errormsg=%v", resPreAuthCode.ErrCode, resPreAuthCode.ErrMsg)
		return nil, err
	}
	return resPreAuthCode, nil
}
