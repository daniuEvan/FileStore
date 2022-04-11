/**
 * @date: 2022/3/20
 * @desc: 网易云信短信服务
 */

package config

//
// SmConfig
// @Description: 网易云信
//
type SmConfig struct {
	SendSmBaseUrl   string `mapstructure:"sendSmBaseUrl" json:"sendSmBaseUrl"`
	AppSecret       string `mapstructure:"appSecret" json:"appSecret"`
	AppKey          string `mapstructure:"appKey" json:"appKey"`
	SMTemplateCode  int    `mapstructure:"SMTemplateCode" json:"SMTemplateCode"`
	CodeLen         int    `mapstructure:"codeLen" json:"codeLen"`
	VerifySmBaseUrl string `mapstructure:"verifySmBaseUrl" json:"verifySmBaseUrl"`
}
